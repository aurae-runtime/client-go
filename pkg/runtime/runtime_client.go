package runtime

import (
	"context"
	"github.com/aurae-runtime/client-go/pkg/api/v0/runtime"
	"google.golang.org/grpc"
)

type Client struct {
	grpc runtime.RuntimeClient
}

func NewRuntimeClient(conn *grpc.ClientConn) Client {
	return Client{
		grpc: runtime.NewRuntimeClient(conn),
	}
}

func (rc Client) Exec(ctx context.Context, in *runtime.Executable) (*runtime.ExecutableStatus, error) {
	return rc.grpc.Exec(ctx, in)
}
