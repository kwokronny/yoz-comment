package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Pagination 分页结构对应接口
type Pagination interface {
	GetOffset() int
	GetPageSize() int
	GetPage() int
}

// PageData 通用分页结构
type PageData struct {
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	Total    int64       `json:"total"`
	Records  interface{} `json:"records"`
}

// GetPagination 获取请求中的分页传参
func GetPagination(c *gin.Context) Pagination {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "15"))
	var pagination Pagination
	pagination = &PageData{Page: page, PageSize: pageSize}
	return pagination
}

// GetOffset 获取页码对应偏移值
func (page *PageData) GetOffset() int {
	return (page.Page - 1) * page.PageSize
}

// GetPage 获取页码
func (page *PageData) GetPage() int {
	return page.Page
}

// GetPageSize 获取单页数量
func (page *PageData) GetPageSize() int {
	return page.PageSize
}
