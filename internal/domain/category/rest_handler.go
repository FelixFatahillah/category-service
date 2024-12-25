package category

import (
	"category-service/internal/domain/category/services"
	"category-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type handlerRESTCategory struct {
	service services.CategoryService
}

func NewHandlerRESTCategory(service services.CategoryService, router *fiber.App) {
	handler := handlerRESTCategory{
		service,
	}

	routerV1 := router.Group("/api/v1")
	routerProtected := routerV1.Group("/private/categories", middleware.RoleMiddleware())

	routerProtected.Delete("/:id", handler.handlerDelete)
	routerProtected.Get("/:id", handler.handlerFindById)
	routerProtected.Put("/:id", handler.handlerUpdate)
	routerProtected.Post("/", handler.handlerCreate)
	routerProtected.Get("/", handler.handlerGetAll)
}
