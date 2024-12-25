package models

import (
	"category-service/internal/shared"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID string `json:"id" gorm:"type:uuid;primaryKey;not null"`
	shared.BaseModel
	Category    string `json:"category,omitempty" gorm:"type:varchar(255);not null"`
	Description string `json:"description,omitempty" gorm:"type:varchar(255);not null"`
}

func (u *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String() // Generate UUID
	}
	return
}
