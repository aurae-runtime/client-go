package observe

import (
	"context"
	"github.com/aurae-runtime/client-go/pkg/api/v0/observe"
	"google.golang.org/grpc"
)

type Client struct {
	grpc observe.ObserveClient
}

func NewObserveClient(conn *grpc.ClientConn) Client {
	return Client{
		grpc: observe.NewObserveClient(conn),
	}
}

func (oc Client) GetAuraeDaemonLogStream(ctx context.Context, in *observe.GetAuraeDaemonLogStreamRequest) (observe.Observe_GetAuraeDaemonLogStreamClient, error) {
	return oc.grpc.GetAuraeDaemonLogStream(ctx, in)
}

func (oc Client) GetSubProcessStream(ctx context.Context, in *observe.GetSubProcessStreamRequest) (observe.Observe_GetSubProcessStreamClient, error) {
	return oc.grpc.GetSubProcessStream(ctx, in)
}

func (oc Client) Status(ctx context.Context, in *observe.StatusRequest) (*observe.StatusResponse, error) {
	return oc.grpc.Status(ctx, in)
}
