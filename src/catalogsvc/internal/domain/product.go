package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	ProductID   int                `json:"id" bson:"id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Price       int                `json:"price" bson:"price"`
	ImagePath   string             `json:"image_path,omitempty" bson:"imagePath,omitempty"`
}

// db.products.insertOne({"id":100000, "name":"Lemon", "Description":"The lemon is a round, slightly elongated fruit, it has a strong and resistant skin, with an intense bright yellow colour when it is totaly ripe, giving off a special aroma when it is cut.", "price": 30, "imagePath":"/static/img/products/lemon.jpg"})
