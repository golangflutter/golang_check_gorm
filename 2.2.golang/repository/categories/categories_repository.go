package repository

import "golang-crud-gin/model"

type CategoryRepository interface {
    Save(category model.Category)
    Update(category model.Category)
    Delete(categoryID int)
    FindById(categoryID int) (model.Category, error)
    FindAll() []model.Category
    FindByCategoryName(categoryName string) (model.Category, error)
    FindCategoriesByPage(pageNumber, pageSize int) ([]model.Category, error)
    FindCategoriesByCharacteristicName(characteristic string, pageNumber, pageSize int) ([]model.Category, error)
}
