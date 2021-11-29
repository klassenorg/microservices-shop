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

type Product struct {
	ProductID   int    `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Price       int    `json:"price" bson:"price"`
	ImagePath   string `json:"image_path,omitempty" bson:"imagePath,omitempty"`
}

func (c *CatalogClient) GetAllProducts(ctx context.Context) ([]Product, error) {
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

func (c *CatalogClient) GetProductByID(ctx context.Context, id int) (*Product, error) {
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
