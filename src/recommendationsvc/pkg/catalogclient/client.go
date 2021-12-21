package catalogclient

import (
	"context"
	"google.golang.org/grpc"
	pb "recommendationsvc/gen/proto"
)

type Client struct {
	client pb.CatalogServiceClient
	conn   *grpc.ClientConn
}

func NewCartClient(addr string) (*Client, error) {

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewCatalogServiceClient(conn)

	return &Client{client: client, conn: conn}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

type Product struct {
	ProductID   int    `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Price       int    `json:"price" bson:"price"`
	ImagePath   string `json:"image_path,omitempty" bson:"imagePath,omitempty"`
}

func (c *Client) GetAllProducts(ctx context.Context) ([]Product, error) {
	res, err := c.client.GetAllProducts(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}

	out := make([]Product, 0, len(res.Products))

	for _, product := range res.Products {
		out = append(out, Product{
			ProductID:   int(product.Id),
			Name:        product.Name,
			Description: product.Description,
			Price:       int(product.Price),
			ImagePath:   product.ImagePath,
		})
	}

	return out, nil
}

func (c *Client) GetProductByID(ctx context.Context, id int) (*Product, error) {
	res, err := c.client.GetProductByID(ctx, &pb.GetProductByIDRequest{Id: uint32(id)})
	if err != nil {
		return nil, err
	}

	return &Product{
		ProductID:   int(res.Id),
		Name:        res.Name,
		Description: res.Description,
		Price:       int(res.Price),
		ImagePath:   res.ImagePath,
	}, nil
}
