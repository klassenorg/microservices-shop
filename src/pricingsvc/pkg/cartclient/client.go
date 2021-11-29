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

func (c *CartClient) GetCart(ctx context.Context, userID string) (map[string]string, error) {
	res, err := c.client.GetCart(ctx, &pb.CartRequest{Id: userID})
	if err != nil {
		return nil, err
	}

	return res.Cart, nil
}

func (c *CartClient) AddToCart(ctx context.Context, userID, productID string, count int) error {
	_, err := c.client.AddToCart(ctx, &pb.CartUpdateRequest{UserId: userID, ProductId: productID, Count: uint32(count)})

	return err
}

func (c *CartClient) RemoveFromCart(ctx context.Context, userID, productID string, count int) error {
	_, err := c.client.RemoveFromCart(ctx, &pb.CartUpdateRequest{UserId: userID, ProductId: productID, Count: uint32(count)})

	return err
}

func (c *CartClient) RemoveAllFromCart(ctx context.Context, userID string) error {
	_, err := c.client.RemoveAllFromCart(ctx, &pb.CartRequest{Id: userID})

	return err
}
