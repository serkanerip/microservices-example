package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/serkanerip/microservices/catalog"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := catalog.GetMongoDBConnection(ctx, catalog.GetMongoDBConnectionParams{URI: catalog.Config.MongoURI})
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	productsCollection := client.Database("test", nil).Collection("products", nil)

	controller := catalog.NewProductsController(catalog.NewProductMongoRepository(productsCollection))
	app := fiber.New()

	app.Get("/api/v1/catalog/get-product-by-category/:category", controller.GetCategoryProducts)
	app.Get("/api/v1/catalog", controller.GetAllProducts)
	app.Post("/api/v1/catalog", controller.CreteProduct)
	app.Get("/api/v1/catalog/:id", controller.GetProduct)
	app.Put("/api/v1/catalog/", controller.UpdateProduct)
	app.Delete("/api/v1/catalog/:id", controller.DeleteProduct)

	log.Fatal(app.Listen(":3000"))
}
