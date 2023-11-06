package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
)

type SizeService interface {
    Create(size request.CreateSizeRequest)
    Update(size request.UpdateSizeRequest)
    Delete(sizeID int)
    FindById(sizeID int) (response.SizeResponse, error)
    FindAll() ([]response.SizeResponse, error)
    FindByName(sizeName string) (response.SizeResponse, error)
    FindAllWithPagination(pageNumber, pageSize int) ([]response.SizeResponse, error)
}
