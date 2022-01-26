package catalog

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/serkanerip/microservices/catalog/test/containers"
)

var ctx = context.Background()

func getRepo(connParams GetMongoDBConnectionParams) ProductRepository {
	var err error
	client, err := GetMongoDBConnection(context.Background(), connParams)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("test", nil).Collection("products", nil)

	return NewProductMongoRepository(collection)
}

func TestProductMongoRepository_CreateProduct(t *testing.T) {
	mongoC, err := containers.SetupMongoContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer mongoC.Terminate(ctx)

	repo := getRepo(GetMongoDBConnectionParams{URI: fmt.Sprintf("mongodb://localhost:%s/", mongoC.Port)})

	p := CreateProductDTO{
		Name:        "Macbook PRO 16'",
		Category:    "Laptops",
		Summary:     "Best laptop ever",
		Description: "Best for your work",
		ImageFile:   "",
		Price:       2499.99,
	}
	createdProduct, err := repo.CreateProduct(ctx, p)
	if err != nil {
		t.Fatal(err)
	}

	actualProduct, err := repo.GetProduct(ctx, createdProduct.ID.Hex())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, createdProduct, actualProduct)
}

func TestProductMongoRepository_DeleteProduct(t *testing.T) {
	mongoC, err := containers.SetupMongoContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer mongoC.Terminate(ctx)

	repo := getRepo(GetMongoDBConnectionParams{URI: fmt.Sprintf("mongodb://localhost:%s/", mongoC.Port)})

	p := CreateProductDTO{
		Name:        "Macbook PRO 16'",
		Category:    "Laptops",
		Summary:     "Best laptop ever",
		Description: "Best for your work",
		ImageFile:   "",
		Price:       2499.99,
	}
	createdProduct, err := repo.CreateProduct(ctx, p)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.DeleteProduct(ctx, createdProduct.ID.Hex())
	if err != nil {
		t.Fatal(err)
	}

	products, err := repo.GetProducts(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(products) != 0 {
		t.Fatal("product not deleted!")
	}
}

func TestProductMongoRepository_GetProductsByName(t *testing.T) {
	mongoC, err := containers.SetupMongoContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer mongoC.Terminate(ctx)

	repo := getRepo(GetMongoDBConnectionParams{URI: fmt.Sprintf("mongodb://localhost:%s/", mongoC.Port)})

	dto := CreateProductDTO{
		Name:        "Macbook PRO 16'",
		Category:    "Laptops",
		Summary:     "Best laptop ever",
		Description: "Best for your work",
		ImageFile:   "",
		Price:       2499.99,
	}
	_, err = repo.CreateProduct(ctx, dto)
	if err != nil {
		t.Fatal(err)
	}

	products, err := repo.GetProductsByName(ctx, dto.Name)
	if err != nil {
		t.Fatal(err)
	}

	if len(products) != 1 {
		t.Fatal("products length should be one!")
	}

	products, err = repo.GetProductsByName(ctx, dto.Name+"x")
	if err != nil {
		t.Fatal(err)
	}

	if len(products) != 0 {
		t.Fatal("products length should be zero!")
	}
}
