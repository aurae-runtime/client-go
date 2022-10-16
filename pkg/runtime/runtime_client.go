package runtime

import (
	"context"
	"github.com/aurae-runtime/client-go/pkg/stdlib/v0/runtime"
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

func (rc Client) ExecutableStart(ctx context.Context, in *runtime.Executable) (*runtime.ExecutableStatus, error) {
	return rc.grpc.ExecutableStart(ctx, in)
}

func (rc Client) ExecutableStop(ctx context.Context, in *runtime.Executable) (*runtime.ExecutableStatus, error) {
	return rc.grpc.ExecutableStop(ctx, in)
}

func (rc Client) ContainerStart(ctx context.Context, in *runtime.Container) (*runtime.ContainerStatus, error) {
	return rc.grpc.ContainerStart(ctx, in)
}

func (rc Client) ContainerStop(ctx context.Context, in *runtime.Container) (*runtime.ContainerStatus, error) {
	return rc.grpc.ContainerStop(ctx, in)
}

func (rc Client) InstanceStart(ctx context.Context, in *runtime.Instance) (*runtime.InstanceStatus, error) {
	return rc.grpc.InstanceStart(ctx, in)
}

func (rc Client) InstanceStop(ctx context.Context, in *runtime.Instance) (*runtime.InstanceStatus, error) {
	return rc.grpc.InstanceStop(ctx, in)
}
