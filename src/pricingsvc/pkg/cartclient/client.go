package cartclient

import (
	"context"
	"google.golang.org/grpc"
	pb "pricingsvc/gen/proto"
)

type CartClient struct {
	client pb.CartServiceClient
	conn   *grpc.ClientConn
}

func NewCartClient(addr string) (*CartClient, error) {

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewCartServiceClient(conn)

	return &CartClient{client: client, conn: conn}, nil
}

func (c *CartClient) Close() error {
	return c.conn.Close()
}

func (c *CartClient) GetCart(ctx context.Context, id string) (map[string]string, error) {
	res, err := c.client.GetCart(ctx, &pb.CartRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return res.Cart, nil
}
