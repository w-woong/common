package wrapper_test

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"os"
	"strconv"
	"testing"
	"time"

	"gorm.io/gorm"

	"github.com/w-woong/common/dto"
	"github.com/w-woong/common/wrapper"
	pb "github.com/w-woong/common/wrapper/tests/protos"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	onlinetest, _ = strconv.ParseBool(os.Getenv("ONLINE_TEST"))
	benchtest, _  = strconv.ParseBool(os.Getenv("BENCH_TEST"))
	manualtest, _ = strconv.ParseBool(os.Getenv("MANUAL_TEST"))
	db            *sql.DB
	gdb           *gorm.DB

	grpcSvr  *grpc.Server
	grpcConn *grpc.ClientConn
)

func setup() error {
	var err error
	grpcConf := dto.ConfigGrpc{
		Timeout:     15,
		HealthCheck: false,
		EnforcementPolicy: dto.ConfigEnforcementPolicy{
			Use:                 true,
			MinTime:             15,
			PermitWithoutStream: true,
		},
		KeepAlive: dto.ConfigKeepAlive{
			MaxConnIdle:         300,
			MaxConnAge:          300,
			MaxConnAgeGrace:     6,
			Time:                30,
			Timeout:             1,
			PermitWithoutStream: true,
		},
	}
	grpcSvr, err = wrapper.NewGrpcServer(grpcConf, "./certs/server.crt", "./certs/server.key", false)
	if err != nil {
		return err
	}
	pb.RegisterStudentServer(grpcSvr, &studentGrpcServer{})
	lis, err := net.Listen("tcp", ":18080")
	if err != nil {
		return err
	}
	go func() {
		grpcSvr.Serve(lis)
	}()

	grpcClientConf := dto.ConfigGrpcClient{
		Addr: ":18080",
		KeepAlive: dto.ConfigKeepAlive{
			Time:                60,
			Timeout:             1,
			PermitWithoutStream: true,
		},
		ResolverScheme:      "student",
		ResolverServiceName: "student-svc",
		DefaultServiceConfig: `{
			"loadBalancingConfig": [{"round_robin":{}}],
			"methodConfig": [{
				"name": [{}],
				"waitForReady": true,
				"retryPolicy": {
					"MaxAttempts": 4,
					"InitialBackoff": ".01s",
					"MaxBackoff": ".01s",
					"BackoffMultiplier": 1.0,
					"RetryableStatusCodes": [ "UNAVAILABLE" ]
				}
			}]
		}
		`,
		// DefaultServiceConfig: `
		// {"loadBalancingConfig": [{"round_robin":{}}],
		//   "methodConfig": [{
		// 	"name": [{"service": "grpc.event"}],
		// 	"waitForReady": true,
		// 	"retryPolicy": {
		// 	  "MaxAttempts": 4,
		// 	  "InitialBackoff": ".01s",
		// 	  "MaxBackoff": ".01s",
		// 	  "BackoffMultiplier": 1.0,
		// 	  "RetryableStatusCodes": [ "UNAVAILABLE" ]
		// 	}
		//   }]
		// }
		// `,
		CaCertPem:      "./certs/server.crt",
		CertServerName: "localhost",
	}
	grpcConn, err = wrapper.NewGrpcClient(grpcClientConf, false)
	if err != nil {
		return err
	}
	// if onlinetest {
	// 	//
	// }
	return err
}

func shutdown() {
	if grpcConn != nil {
		grpcConn.Close()
	}
	if grpcSvr != nil {
		grpcSvr.GracefulStop()
	}
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Println(err)
		shutdown()
		os.Exit(1)
	}

	exitCode := m.Run()

	shutdown()
	os.Exit(exitCode)
}

type studentGrpcServer struct {
	pb.StudentServer
}

func (d *studentGrpcServer) Read(ctx context.Context, in *pb.StudentRequest) (*pb.StudentReply, error) {
	docs := make([]*pb.StudentEntity, 0)
	docs = append(docs, &pb.StudentEntity{
		Name:        "wonk",
		Age:         10,
		DateTime:    timestamppb.New(time.Now()),
		DoubleValue: 10.1,
	})

	var count int64 = 1
	rep := pb.StudentReply{
		Status:    200,
		Documents: docs,
		Count:     &count,
	}

	return &rep, nil
}
