package schedule

import (
	"context"
	"github.com/aurae-runtime/client-go/pkg/api/v0/schedule"
	"google.golang.org/grpc"
)

type Client struct {
	grpc schedule.ScheduleClient
}

func NewScheduleClient(conn *grpc.ClientConn) Client {
	return Client{
		grpc: schedule.NewScheduleClient(conn),
	}
}

func (c Client) ShowEnabled(ctx context.Context, in *schedule.ShowEnabledRequest) (*schedule.ShowEnabledResponse, error) {
	return c.grpc.ShowEnabled(ctx, in)
}

func (c Client) ShowDisabled(ctx context.Context, in *schedule.ShowDisabledRequest) (*schedule.ShowDisabledResponse, error) {
	return c.grpc.ShowDisabled(ctx, in)
}
