package dtos

type CreateCategoryDto struct {
	Category    string `json:"category" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreateCategoryResponseDto struct {
	Category    string `json:"category"`
	Description string `json:"description"`
}

type UpdateCategoryDto struct {
	ID          string `json:"id"`
	Category    string `json:"category"`
	Description string `json:"description"`
}
