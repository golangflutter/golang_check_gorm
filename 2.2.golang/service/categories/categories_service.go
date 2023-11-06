package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
)

type CategoryService interface {
    Create(category request.CreateCategoryRequest)
    Update(category request.UpdateCategoryRequest)
    Delete(categoryID int)
    FindById(categoryID int) response.CategoryResponse
    FindAll() []response.CategoryResponse
    FindByName(categoryName string) (response.CategoryResponse, error)
    FindAllWithPagination(pageNumber, pageSize int) ([]response.CategoryResponse, error)
    FindCategoriesByCharacteristicName(characteristic string, pageNumber, pageSize int) ([]response.CategoryResponse, error) // Add FindCategoriesByCharacteristicName method
}
