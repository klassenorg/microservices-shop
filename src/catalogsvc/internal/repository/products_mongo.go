package repository

import (
	"catalogsvc/internal/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductsRepo struct {
	db *mongo.Collection
}

func (p *ProductsRepo) GetAll(ctx context.Context) ([]domain.Product, error) {
	products := make([]domain.Product, 0, 12)

	cur, err := p.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var res domain.Product
		if err := cur.Decode(&res); err != nil {
			return nil, err
		}
		products = append(products, res)
	}

	return products, nil
}

func (p *ProductsRepo) GetByID(ctx context.Context, id int) (domain.Product, error) {
	var product domain.Product
	err := p.db.FindOne(ctx, bson.M{"id": id}).Decode(&product)

	return product, err
}

func NewProductsRepo(db *mongo.Database) *ProductsRepo {
	return &ProductsRepo{db: db.Collection(productsCollection)}
}
