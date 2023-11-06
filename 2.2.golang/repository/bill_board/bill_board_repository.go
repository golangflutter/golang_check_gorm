package repository

import "golang-crud-gin/model"



type BillboardRepository interface {
    Save(billboard model.Billboard) error
    Update(billboard model.Billboard) error
    Delete(billboardId int) error
    FindByID(billboardId int) (model.Billboard, error)
    FindAll() ([]model.Billboard, error)
    // Add any additional methods you need for billboard repository here

    // For example, you can add a method to find billboards by label or image URL:
    FindByLabel(label string) (model.Billboard, error)

    FindBillboardsByPage(pageNumber, pageSize int) ([]model.Billboard, error)
}
