package main

import (
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/serkanerip/microservices/basket"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	client := redis.NewClient(&redis.Options{
		Addr:     basket.Config.RedisURI,
		Password: basket.Config.RedisPassword,
		DB:       basket.Config.RedisDB,
	})
	repo := basket.NewRedisRepository(client)

	app.Get("/api/v1/basket/:username", func(ctx *fiber.Ctx) error {
		sc, err := repo.GetBasket(ctx.Context(), ctx.Params("username"))
		if err != nil {
			return ctx.JSON(fiber.Map{"error": "Couldn't get basket"})
		}
		return ctx.JSON(sc)
	})

	app.Delete("/api/v1/basket/:username", func(ctx *fiber.Ctx) error {
		err := repo.DeleteBasket(ctx.Context(), ctx.Params("username"))
		if err != nil {
			return ctx.JSON(fiber.Map{"error": "Couldn't delete basket"})
		}
		return nil
	})

	app.Post("/api/v1/basket", func(ctx *fiber.Ctx) error {
		var dto basket.ShoppingCart
		err := ctx.BodyParser(&dto)
		if err != nil {
			ctx.Status(http.StatusPreconditionFailed)
			return ctx.JSON(fiber.Map{"error": "Couldn't parse body"})
		}
		sc, err := repo.UpdateBasket(ctx.Context(), dto)
		if err != nil {
			log.Println(err)
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(fiber.Map{"error": "Couldn't update basket"})
		}
		return ctx.JSON(sc)
	})

	log.Fatal(app.Listen(":8001"))
}
