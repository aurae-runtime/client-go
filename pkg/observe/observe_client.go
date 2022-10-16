package observe

import (
	"context"
	"github.com/aurae-runtime/client-go/pkg/stdlib/v0/observe"
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

func (oc Client) Status(ctx context.Context, in *observe.StatusRequest) (*observe.StatusResponse, error) {
	return oc.grpc.Status(ctx, in)
}
