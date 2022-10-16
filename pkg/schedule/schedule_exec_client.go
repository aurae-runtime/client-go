package schedule

import (
	"context"
	"github.com/aurae-runtime/client-go/pkg/stdlib/v0/runtime"
	"github.com/aurae-runtime/client-go/pkg/stdlib/v0/schedule"
	"google.golang.org/grpc"
)

type ExecClient struct {
	grpc schedule.ScheduleExecutableClient
}

func NewScheduleExecClient(conn *grpc.ClientConn) ExecClient {
	return ExecClient{
		grpc: schedule.NewScheduleExecutableClient(conn),
	}
}

func (ec *ExecClient) Enable(ctx context.Context, in *runtime.Executable) (*schedule.ExecutableEnableResponse, error) {
	return ec.grpc.Enable(ctx, in)
}

func (ec *ExecClient) Disable(ctx context.Context, in *runtime.Executable) (*schedule.ExecutableDisableResponse, error) {
	return ec.grpc.Disable(ctx, in)
}

func (ec *ExecClient) Destroy(ctx context.Context, in *runtime.Executable) (*schedule.ExecutableDestroyResponse, error) {
	return ec.grpc.Destroy(ctx, in)
}
