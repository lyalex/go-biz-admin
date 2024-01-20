package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Entity interface {
	Count(db *gorm.DB) int64
	Take(db *gorm.DB, limit int, offset int) interface{}
}

func Paginate(db *gorm.DB, entity Entity, page int) gin.H {
	limit := 10
	offset := (page - 1) * limit

	data := entity.Take(db, limit, offset)
	total := entity.Count(db)
	return gin.H{
		"data": data,
		"metadata": gin.H{
			"total": total,
			"page":  page,
		},
	}
}
