package repository

import "golang-crud-gin/model"

type ColorRepository interface {
    Save(color model.Color) error
    Update(color model.Color) error
    Delete(colorId int) error
    FindByID(colorId int) (model.Color, error)
    FindAll() ([]model.Color, error)
    FindByColorName(colorName string) (model.Color, error)
    FindColorsByPage(pageNumber, pageSize int) ([]model.Color, error)
}
