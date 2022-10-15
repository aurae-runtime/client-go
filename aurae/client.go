package aurae

import (
	"github.com/aurae-runtime/client-go/pkg/stdlib/v0/observe"
	"github.com/aurae-runtime/client-go/pkg/stdlib/v0/runtime"
	schedule "github.com/aurae-runtime/client-go/pkg/stdlib/v0/schedule"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	conn *grpc.ClientConn
}

func NewClient(addr string, credentials credentials.TransportCredentials) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(credentials))
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil
}

func (c *Client) Observe() observe.ObserveClient {
	return observe.NewObserveClient(c.conn)
}

func (c *Client) Runtime() runtime.RuntimeClient {
	return runtime.NewRuntimeClient(c.conn)
}

func (c *Client) Schedule() schedule.ScheduleClient {
	return schedule.NewScheduleClient(c.conn)
}
