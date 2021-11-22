package catalogclient

import (
	"context"
	"google.golang.org/grpc"
	pb "pricingsvc/gen/proto"
	"strconv"
)

type CatalogClient struct {
	client pb.CatalogServiceClient
	conn   *grpc.ClientConn
}

func NewCartClient(addr string) (*CatalogClient, error) {

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewCatalogServiceClient(conn)

	return &CatalogClient{client: client, conn: conn}, nil
}

func (c *CatalogClient) Close() error {
	return c.conn.Close()
}

func (c *CatalogClient) GetProductPriceByID(ctx context.Context, id string) (int, error) {
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return -1, err
	}
	product, err := c.client.GetProductByID(ctx, &pb.GetProductByIDRequest{Id: uint32(idInt)})
	if err != nil {
		return -1, err
	}
	return int(product.Price), nil
}
