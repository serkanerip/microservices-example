package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/serkanerip/microservices/discount"
)

var ctx = context.Background()

func main() {
	app := fiber.New()
	conn := discount.GetPGDBConnection()
	defer conn.Close()

	if err := conn.PingContext(ctx); err != nil {
		log.Fatal("couldn't connect to pg db err is ", err)
	}

	repo := discount.NewPGRepository(conn)

	app.Get("/api/v1/discount/:productName", func(c *fiber.Ctx) error {
		coupon, err := repo.GetDiscount(ctx, c.Params("productName"))
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{"error": "Cannot get discount"})
		}
		return c.JSON(coupon)
	})

	app.Post("/api/v1/discount", func(c *fiber.Ctx) error {
		var d discount.CreateCouponDTO
		if err := c.BodyParser(&d); err != nil {
			log.Println(err)
			c.Status(http.StatusPreconditionFailed)
			return c.JSON(fiber.Map{"error": "Cannot parse request body"})
		}
		coupon, err := repo.CreateDiscount(c.Context(), d)
		if err != nil {
			log.Println(err)
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{"error": "Cannot create discount"})
		}
		c.Status(http.StatusCreated)
		return c.JSON(coupon)
	})

	app.Put("/api/v1/discount", func(c *fiber.Ctx) error {
		var d discount.Coupon
		if err := c.BodyParser(&d); err != nil {
			log.Println(err)
			c.Status(http.StatusPreconditionFailed)
			return c.JSON(fiber.Map{"error": "Cannot parse request body"})
		}
		err := repo.UpdateDiscount(c.Context(), d)
		if err != nil {
			log.Println(err)
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{"error": "Cannot update discount"})
		}
		return nil
	})

	app.Delete("/api/v1/discount/:productName", func(c *fiber.Ctx) error {
		err := repo.DeleteDiscount(c.Context(), c.Params("productName"))
		if err != nil {
			log.Println(err)
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{"error": "Cannot delete discount"})
		}
		return nil
	})

	log.Fatal(app.Listen(":8003"))
}
