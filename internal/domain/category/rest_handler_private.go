package category

import (
	"category-service/internal/domain/category/dtos"
	"category-service/pkg/helper"
	"category-service/pkg/response"
	"category-service/pkg/validation"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (handler handlerRESTCategory) handlerCreate(ctx *fiber.Ctx) error {
	var body dtos.CreateCategoryDto
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}
	if err := validation.Validate(body); err != nil {
		return err
	}

	data, err := handler.service.Create(ctx.Context(), body)
	if err != nil {
		return err
	}

	return response.Respond(ctx, fiber.StatusOK, "success", data, nil, nil)
}

func (handler handlerRESTCategory) handlerGetAll(ctx *fiber.Ctx) error {
	paginate := helper.Pagination{
		Page:  1,
		Limit: 10,
	}
	err := ctx.QueryParser(&paginate)
	if paginate.Limit >= 100 {
		paginate.Limit = 100
	}

	filter := dtos.CategoryFilter{
		Pagination: paginate,
	}
	_ = ctx.QueryParser(&filter)

	data, meta, err := handler.service.GetAll(ctx.Context(), filter)
	if err != nil {
		return err
	}

	return response.Respond(ctx, fiber.StatusOK, "success", data, nil, meta)
}

func (handler handlerRESTCategory) handlerFindById(ctx *fiber.Ctx) error {
	data, err := handler.service.FindById(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return response.Respond(ctx, fiber.StatusOK, "success", data, nil, nil)
}

func (handler handlerRESTCategory) handlerUpdate(ctx *fiber.Ctx) error {
	var body dtos.UpdateCategoryDto
	body.ID = ctx.Params("id")
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}
	if err := validation.Validate(body); err != nil {
		return err
	}

	err := handler.service.Update(ctx.Context(), body)
	if err != nil {
		return err
	}

	return response.Respond(ctx, fiber.StatusOK, fmt.Sprintf("success update %s", ctx.Params("id")), nil, nil, nil)
}

func (handler handlerRESTCategory) handlerDelete(ctx *fiber.Ctx) error {
	err := handler.service.Delete(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return response.Respond(ctx, fiber.StatusOK, fmt.Sprintf("success deleted %s", ctx.Params("id")), nil, nil, nil)
}
