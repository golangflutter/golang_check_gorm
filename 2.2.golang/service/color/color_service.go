package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
)

type ColorService interface {
    Create(color request.CreateColorRequest)
    Update(color request.UpdateColorRequest)
    Delete(colorID int)
    FindById(colorID int) (response.ColorResponse, error)
    FindAll() ([]response.ColorResponse, error)
    FindByName(colorName string) (response.ColorResponse, error)
    FindAllWithPagination(pageNumber, pageSize int) ([]response.ColorResponse, error)
}
