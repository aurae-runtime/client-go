package pkg

import (
	"github.com/aurae-runtime/client-go/pkg/api/v0/schedule"
	"github.com/aurae-runtime/client-go/pkg/observe"
	"github.com/aurae-runtime/client-go/pkg/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	conn *grpc.ClientConn
}

func NewAuraeClient(addr string, credentials credentials.TransportCredentials) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(credentials))
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil
}

func (c *Client) Observe() observe.Client {
	return observe.NewObserveClient(c.conn)
}

func (c *Client) Runtime() runtime.Client {
	return runtime.NewRuntimeClient(c.conn)
}

func (c *Client) Schedule() schedule.ScheduleClient {
	return schedule.NewScheduleClient(c.conn)
}

func (c *Client) ScheduleExecutable() schedule.ScheduleExecutableClient {
	return schedule.NewScheduleExecutableClient(c.conn)
}
