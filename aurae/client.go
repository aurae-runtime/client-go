package aurae

import (
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
	return c.Observe()
}

func (c *Client) Runtime() runtime.Client {
	return c.Runtime()
}
