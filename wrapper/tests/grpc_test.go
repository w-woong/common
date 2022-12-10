package wrapper_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	pb "github.com/w-woong/common/wrapper/tests/protos"
)

func TestGrpcWrapper(t *testing.T) {
	c := pb.NewStudentClient(grpcConn)
	rep, err := c.Read(context.Background(), &pb.StudentRequest{
		Name: "wonka",
	})
	assert.Nil(t, err)

	fmt.Println(rep.String())
}

func BenchmarkGrpcWrapper(b *testing.B) {

	c := pb.NewStudentClient(grpcConn)
	b.ReportAllocs()
	// b.ReportMetric(0, "/op")
	for i := 0; i < b.N; i++ {
		c.Read(context.Background(), &pb.StudentRequest{
			Name: "wonka",
		})
	}
}
