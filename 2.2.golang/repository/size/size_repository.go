package repository

import "golang-crud-gin/model"

type SizeRepository interface {
    Save(size model.Size) error
    Update(size model.Size) error
    Delete(sizeId int) error
    FindByID(sizeId int) (model.Size, error)
    FindAll() ([]model.Size, error)
    FindBySizeName(sizeName string) (model.Size, error)
    FindSizesByPage(pageNumber, pageSize int) ([]model.Size, error)
}
