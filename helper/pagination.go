package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Pagination is PageData interface
type Pagination interface {
	GetOffset() int
	GetPageSize() int
	GetPage() int
}

// PageData is default page return struct of API Request
type PageData struct {
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	Total    int         `json:"total"`
	Records  interface{} `json:"records"`
}

// GetPagination is get pagination infomation by GET Request Query
func GetPagination(c *gin.Context) Pagination {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "15"))
	var pagination Pagination
	pagination = &PageData{Page: page, PageSize: pageSize}
	return pagination
}

// GetOffset is get sql offset in pagination
func (page *PageData) GetOffset() int {
	return (page.Page - 1) * page.PageSize
}

// GetPage is get page in pagination
func (page *PageData) GetPage() int {
	return page.Page
}

// GetPageSize is get pageSize in pagination
func (page *PageData) GetPageSize() int {
	return page.PageSize
}
