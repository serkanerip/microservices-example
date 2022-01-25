package catalog

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ProductsController struct {
	r ProductRepository
}

func NewProductsController(r ProductRepository) *ProductsController {
	return &ProductsController{r: r}
}

func (p *ProductsController) GetAllProducts(c *fiber.Ctx) error {
	products, err := p.r.GetProducts(c.Context())
	if err != nil {
		log.Println(err)
		return c.JSON(map[string]string{"msg": "Error"})
	}
	return c.JSON(products)
}

func (p *ProductsController) GetCategoryProducts(c *fiber.Ctx) error {
	products, err := p.r.GetProductsByCategory(c.Context(), c.Params("category"))
	if err != nil {
		return c.JSON(map[string]string{"msg": "Error"})
	}
	return c.JSON(products)
}

func (p *ProductsController) GetProduct(c *fiber.Ctx) error {
	products, err := p.r.GetProduct(c.Context(), c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.JSON(map[string]string{"msg": "Error"})
	}
	return c.JSON(products)
}

func (p *ProductsController) CreteProduct(c *fiber.Ctx) error {
	var product CreateProductDTO
	if c.BodyParser(&product) != nil {
		c.Status(http.StatusPreconditionFailed)
		return c.JSON(map[string]string{"msg": "Invalid request body"})
	}
	err := p.r.CreateProduct(c.Context(), product)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(map[string]string{"msg": "Error"})
	}
	c.Status(http.StatusCreated)
	return c.JSON(map[string]string{"msg": "Product created"})
}

func (p *ProductsController) UpdateProduct(c *fiber.Ctx) error {
	var product ProductDTO
	if c.BodyParser(&product) != nil {
		c.Status(http.StatusPreconditionFailed)
		return c.JSON(map[string]string{"msg": "Invalid request body"})
	}
	err := p.r.UpdateProduct(c.Context(), product)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err.Error())
		return c.JSON(map[string]string{"msg": "Error"})
	}
	c.Status(http.StatusOK)
	return c.JSON(map[string]string{"msg": "Product updated"})
}

func (p *ProductsController) DeleteProduct(c *fiber.Ctx) error {
	err := p.r.DeleteProduct(c.Context(), c.Params("id"))
	if err != nil {
		return c.JSON(map[string]string{"msg": "Error"})
	}
	c.Status(http.StatusOK)
	return c.Send([]byte{})
}
