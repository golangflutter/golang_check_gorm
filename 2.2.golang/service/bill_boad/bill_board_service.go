package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
)

type BillboardService interface {
    Create(billboard request.CreateBillboardRequest)
    Update(billboard request.UpdateBillboardRequest)
    Delete(billboardID int)
    FindByID(billboardID int) (response.BillboardResponse, error)
    FindAll() ([]response.BillboardResponse, error)
    FindByLabel(label string) (response.BillboardResponse, error)

    FindAllWithPagination(pageNumber, pageSize int) ([]response.BillboardResponse, error)
}
