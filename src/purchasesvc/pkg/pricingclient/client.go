package pricingclient

import (
	"context"
	"google.golang.org/grpc"
	pb "purchasesvc/gen/proto"
)

type Client struct {
	client pb.PricingServiceClient
	conn   *grpc.ClientConn
}

func NewPricingClient(addr string) (*Client, error) {

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewPricingServiceClient(conn)

	return &Client{client: client, conn: conn}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Calculate(ctx context.Context, userID string) (string, error) {
	res, err := c.client.Calculate(ctx, &pb.CalculateRequest{Id: userID})
	if err != nil {
		return "", err
	}
	return res.Price, nil
}
