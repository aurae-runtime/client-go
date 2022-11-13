package runtime

import (
	"context"
	"github.com/aurae-runtime/client-go/pkg/api/v0/runtime"
	"google.golang.org/grpc"
)

type Client struct {
	grpc runtime.CoreClient
}

func NewRuntimeClient(conn *grpc.ClientConn) Client {
	return Client{
		grpc: runtime.NewCoreClient(conn),
	}
}

func (rc Client) RunCell(ctx context.Context, in *runtime.Cell) (*runtime.CellStatus, error) {
	return rc.grpc.RunCell(ctx, in)
}

func (rc Client) RunExecutable(ctx context.Context, in *runtime.Executable) (*runtime.ExecutableStatus, error) {
	return rc.grpc.RunExecutable(ctx, in)
}

func (rc Client) RunPod(ctx context.Context, in *runtime.Pod) (*runtime.PodStatus, error) {
	return rc.grpc.RunPod(ctx, in)
}

func (rc Client) RunVirtualMachine(ctx context.Context, in *runtime.VirtualMachine) (*runtime.VirtualMachineStatus, error) {
	return rc.grpc.RunVirtualMachine(ctx, in)
}

func (rc Client) Spawn(ctx context.Context, in *runtime.SpawnRequest) (*runtime.SpawnResponse, error) {
	return rc.grpc.Spawn(ctx, in)
}
