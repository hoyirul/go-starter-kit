package utils

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pagination struct {
	Page       int 		`json:"page"`
	Limit      int 		`json:"limit"`
	TotalRows  int64 	`json:"total_rows"`
	TotalPages int 		`json:"total_pages"`
}

func Paginate(c *gin.Context, db *gorm.DB, model any, result any) (*Pagination, error) {
	
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	var totalRows int64
	if err := db.Model(model).Count(&totalRows).Error; err != nil {
		return nil, err
	}

	if err := db.Limit(limit).Offset(offset).Find(result).Error; err != nil {
		return nil, err
	}

	pagination := &Pagination{
		Page:       page,
		Limit:      limit,
		TotalRows:  totalRows,
		TotalPages: int(math.Ceil(float64(totalRows) / float64(limit))),
	}

	return pagination, nil
}
