package wrapper

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/w-woong/common/dto"
	"go.elastic.co/apm/module/apmgrpc/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/resolver"
)

type GrpcServer struct {
	Svr *grpc.Server
}

// func NewGrpcServer(svr *grpc.Server) *grpcServer {
// 	return &grpcServer{
// 		Svr: svr,
// 	}
// }
func (s *GrpcServer) Serve(lis net.Listener) error {
	// lis, err := net.Listen("tcp", addr)
	// if err != nil {
	// 	return nil, err
	// }
	return s.Svr.Serve(lis)
}
func (s *GrpcServer) Stop() error {
	s.Svr.GracefulStop()
	return nil
}

// NewGrpcServer
// pb.RegisterEventServer(svr, eventGrpcServer)
//
func NewGrpcServer(conf dto.ConfigGrpc, certPem, certKey string, apmActive bool, opt ...grpc.ServerOption) (*grpc.Server, error) {

	opts := []grpc.ServerOption{}
	if certPem != "" && certKey != "" {
		cert, err := tls.LoadX509KeyPair(certPem, certKey)
		if err != nil {
			return nil, err
		}
		opts = append(opts, grpc.Creds(credentials.NewServerTLSFromCert(&cert)))
	}
	if conf.EnforcementPolicy.Use {
		kaep := keepalive.EnforcementPolicy{
			MinTime:             time.Duration(conf.EnforcementPolicy.MinTime) * time.Second,
			PermitWithoutStream: conf.EnforcementPolicy.PermitWithoutStream,
		}
		opts = append(opts, grpc.KeepaliveEnforcementPolicy(kaep))
	}
	kasp := keepalive.ServerParameters{
		MaxConnectionIdle:     time.Duration(conf.KeepAlive.MaxConnIdle) * time.Second,
		MaxConnectionAge:      time.Duration(conf.KeepAlive.MaxConnAge) * time.Second,
		MaxConnectionAgeGrace: time.Duration(conf.KeepAlive.MaxConnAgeGrace) * time.Second,
		Time:                  time.Duration(conf.KeepAlive.Time) * time.Second,
		Timeout:               time.Duration(conf.KeepAlive.Timeout) * time.Second,
	}
	opts = append(opts, grpc.KeepaliveParams(kasp))
	if apmActive {
		opts = append(opts, grpc.UnaryInterceptor(apmgrpc.NewUnaryServerInterceptor()))
		opts = append(opts, grpc.StreamInterceptor(apmgrpc.NewStreamServerInterceptor()))
	}
	opts = append(opts, opt...)
	svr := grpc.NewServer(opts...)
	if conf.HealthCheck {
		healthCheck := health.NewServer()
		healthpb.RegisterHealthServer(svr, healthCheck)
	}

	return svr, nil
}

// NewGrpcClient
// defer conn.Close()
// TODO: Dial timeout
func NewGrpcClient(conf dto.ConfigGrpcClient, apmActive bool) (*grpc.ClientConn, error) {

	resolver.Register(&grpcResolverBuilder{
		scheme:      conf.ResolverScheme,
		serviceName: conf.ResolverServiceName,
		addrs:       strings.Split(conf.Addr, ","),
	})
	kacp := keepalive.ClientParameters{
		Time:                time.Duration(conf.KeepAlive.Time) * time.Second,
		Timeout:             time.Duration(conf.KeepAlive.Timeout) * time.Second,
		PermitWithoutStream: conf.KeepAlive.PermitWithoutStream,
	}
	opts := []grpc.DialOption{
		grpc.WithKeepaliveParams(kacp),
	}
	if conf.CaCertPem != "" && conf.CertServerName != "" {
		creds, err := credentials.NewClientTLSFromFile(conf.CaCertPem, conf.CertServerName)
		if err != nil {
			return nil, err
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	if conf.DefaultServiceConfig != "" {
		opts = append(opts, grpc.WithDefaultServiceConfig(conf.DefaultServiceConfig))
	}
	if conf.DialBlock {
		opts = append(opts, grpc.WithBlock())
	}
	if apmActive {
		opts = append(opts, grpc.WithUnaryInterceptor(apmgrpc.NewUnaryClientInterceptor()))
		opts = append(opts, grpc.WithStreamInterceptor(apmgrpc.NewStreamClientInterceptor()))
	}
	var dialTimeout time.Duration
	if conf.DialTimeout == 0 {
		dialTimeout = 12 * time.Second
	} else {
		dialTimeout = time.Duration(conf.DialTimeout) * time.Second
	}
	ctx, cancel := context.WithTimeout(context.Background(), dialTimeout)
	defer cancel()
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:///%s", conf.ResolverScheme, conf.ResolverServiceName),
		opts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
