package catalog

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	GetProducts(ctx context.Context) ([]Product, error)
	GetProduct(ctx context.Context, id string) (*Product, error)
	GetProductsByName(ctx context.Context, name string) ([]Product, error)
	GetProductsByCategory(ctx context.Context, category string) ([]Product, error)
	CreateProduct(ctx context.Context, p CreateProductDTO) error
	UpdateProduct(ctx context.Context, p ProductDTO) error
	DeleteProduct(ctx context.Context, id string) error
}

type ProductMongoRepository struct {
	c *mongo.Collection
}

func NewProductMongoRepository(c *mongo.Collection) ProductRepository {
	return &ProductMongoRepository{c: c}
}

func (p *ProductMongoRepository) GetProducts(ctx context.Context) ([]Product, error) {
	return p.findProducts(ctx, bson.M{})
}

func (p *ProductMongoRepository) GetProduct(ctx context.Context, id string) (*Product, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	one := p.c.FindOne(ctx, bson.M{"_id": objectID})
	if err := one.Err(); err != nil {
		return nil, err
	}
	var product *Product
	if err := one.Decode(&product); err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductMongoRepository) findProducts(ctx context.Context, filter interface{}) ([]Product, error) {
	cursor, err := p.c.Find(ctx, filter)
	products := make([]Product, 0)
	if err != nil {
		return products, err
	}

	if err := cursor.All(ctx, &products); err != nil {
		return products, errors.Wrap(err, "cannot unmarshall products")
	}

	return products, nil
}

func (p *ProductMongoRepository) GetProductsByName(ctx context.Context, name string) ([]Product, error) {
	return p.findProducts(ctx, bson.D{{"name", bson.D{{"$eq", name}}}})
}

func (p *ProductMongoRepository) GetProductsByCategory(ctx context.Context, category string) ([]Product, error) {
	return p.findProducts(ctx, bson.M{"category": category})
}

func (p *ProductMongoRepository) CreateProduct(ctx context.Context, product CreateProductDTO) error {
	_, err := p.c.InsertOne(ctx, product, nil)
	return err
}

func (p *ProductMongoRepository) UpdateProduct(ctx context.Context, product ProductDTO) error {
	objectID, _ := primitive.ObjectIDFromHex(product.ID)
	_, err := p.c.UpdateOne(ctx, bson.D{{"_id", objectID}}, bson.D{{"$set", product}})
	return err
}

func (p *ProductMongoRepository) DeleteProduct(ctx context.Context, id string) error {
	objectID, _ := primitive.ObjectIDFromHex(id)
	_, err := p.c.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
