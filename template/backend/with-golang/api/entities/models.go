package entities

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id" binding:"required"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type PaginationModel[T any] struct {
	Pagination Pagination `json:"pagination" binding:"required"`
	Data       T          `json:"data" binding:"required"`
}

type Pagination struct {
	Total  int64 `json:"total" binding:"required"`
	Offset int   `json:"offSet" binding:"required"`
	Limit  int   `json:"limit" binding:"required"`
}
