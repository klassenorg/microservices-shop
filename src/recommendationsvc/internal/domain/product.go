package domain

type Product struct {
	ProductID   int    `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Price       int    `json:"price" bson:"price"`
	ImagePath   string `json:"image_path,omitempty" bson:"imagePath,omitempty"`
}
