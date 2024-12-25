package delivery

import (
	"category-service/internal/domain/category"
	servicesCategory "category-service/internal/domain/category/services"
	"category-service/pkg/exception"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type handler struct {
	serviceCategory servicesCategory.CategoryService
}

func NewHandler(serviceCategory servicesCategory.CategoryService) *handler {
	return &handler{
		serviceCategory: serviceCategory,
	}
}

const idleTimeout = 60 * time.Second

func (handler *handler) Init() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.FiberErrorHandler,
		IdleTimeout:  idleTimeout,
		JSONEncoder:  sonic.Marshal,
		JSONDecoder:  sonic.Unmarshal,
	})

	// Middleware
	app.Use(logger.New())
	app.Use(etag.New())
	app.Use(requestid.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	category.NewHandlerRESTCategory(handler.serviceCategory, app)

	return app
}
