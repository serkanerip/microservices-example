package catalog

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Category    string             `bson:"category" json:"category"`
	Summary     string             `bson:"summary" json:"summary"`
	Description string             `bson:"description" json:"description"`
	ImageFile   string             `bson:"image_file" json:"image_file"`
	Price       float64            `bson:"price" json:"price"`
}

type ProductDTO struct {
	ID          string  `bson:"-" json:"id"`
	Name        string  `bson:"name" json:"name"`
	Category    string  `bson:"category" json:"category"`
	Summary     string  `bson:"summary" json:"summary"`
	Description string  `bson:"description" json:"description"`
	ImageFile   string  `bson:"image_file" json:"image_file"`
	Price       float64 `bson:"price" json:"price"`
}

type CreateProductDTO struct {
	Name        string  `bson:"name" json:"name"`
	Category    string  `bson:"category" json:"category"`
	Summary     string  `bson:"summary" json:"summary"`
	Description string  `bson:"description" json:"description"`
	ImageFile   string  `bson:"image_file" json:"image_file"`
	Price       float64 `bson:"price" json:"price"`
}
