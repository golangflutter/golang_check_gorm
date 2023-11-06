package service

import (
	"fmt"
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"

	repository "golang-crud-gin/repository/bill_board"

	"github.com/go-playground/validator/v10"
)

type BillboardServiceImpl struct {
    BillboardRepository repository.BillboardRepository
    Validate            *validator.Validate
}

func NewBillboardServiceImpl(billboardRepository repository.BillboardRepository, validate *validator.Validate) BillboardService {
    return &BillboardServiceImpl{
        BillboardRepository: billboardRepository,
        Validate:            validate,
    }
}

// Create implements BillboardService
func (s *BillboardServiceImpl) Create(billboard request.CreateBillboardRequest) {
    fmt.Printf("createBillboardRequest service billboard:askfdjhaskdhfksadhfkhjds %+v\n", billboard)
    err := s.Validate.Struct(billboard)
    helper.ErrorPanic(err)

    billboardModel := model.Billboard{
        Label:    billboard.Label,
        ImageUrl: billboard.ImageUrl, // Assign the list of strings directly
    }
    fmt.Printf("createBillboardRequest service billboardModel :askfdjhaskdhfksadhfkhjds %+v\n", billboardModel)
    err = s.BillboardRepository.Save(billboardModel)
    helper.ErrorPanic(err)
}

// Update implements BillboardService
func (s *BillboardServiceImpl) Update(billboard request.UpdateBillboardRequest) {
    billboardData, err := s.BillboardRepository.FindByID(billboard.Id)
    helper.ErrorPanic(err)

    // Update the fields directly
    billboardData.Label = billboard.Label
    billboardData.ImageUrl = billboard.ImageUrl

    err = s.BillboardRepository.Update(billboardData)
    helper.ErrorPanic(err)
}

// Delete implements BillboardService
func (s *BillboardServiceImpl) Delete(billboardID int) {
    err := s.BillboardRepository.Delete(billboardID)
    helper.ErrorPanic(err)
}

// FindAll implements BillboardService
func (s *BillboardServiceImpl) FindAll() ([]response.BillboardResponse, error) {
    result, err := s.BillboardRepository.FindAll()
    if err != nil {
        return nil, err
    }

    var billboards []response.BillboardResponse
    for _, value := range result {
        billboard := response.BillboardResponse{
            Id:        value.Id,
            Label:     value.Label,
            ImageUrl:  value.ImageUrl,
            CreatedAt: value.CreatedAt,
        }
        billboards = append(billboards, billboard)
    }

    return billboards, nil
}

// FindByID implements BillboardService
func (s *BillboardServiceImpl) FindByID(billboardID int) (response.BillboardResponse, error) {
    billboardData, err := s.BillboardRepository.FindByID(billboardID)
    if err != nil {
        return response.BillboardResponse{}, err
    }

    billboardResponse := response.BillboardResponse{
        Id:        billboardData.Id,
        Label:     billboardData.Label,
        ImageUrl:  billboardData.ImageUrl,
        CreatedAt: billboardData.CreatedAt,
    }
    return billboardResponse, nil
}

// FindByLabel implements BillboardService
func (s *BillboardServiceImpl) FindByLabel(label string) (response.BillboardResponse, error) {
    billboardData, err := s.BillboardRepository.FindByLabel(label)
    if err != nil {
        return response.BillboardResponse{}, err
    }

    billboardResponse := response.BillboardResponse{
        Id:        billboardData.Id,
        Label:     billboardData.Label,
        ImageUrl:  billboardData.ImageUrl,
        CreatedAt: billboardData.CreatedAt,
    }
    return billboardResponse, nil
}

// FindAllWithPagination implements BillboardService with pagination
func (s *BillboardServiceImpl) FindAllWithPagination(pageNumber, pageSize int) ([]response.BillboardResponse, error) {
    result, err := s.BillboardRepository.FindBillboardsByPage(pageNumber, pageSize)
    if err != nil {
        return nil, err
    }

    var billboards []response.BillboardResponse
    for _, value := range result {
        billboard := response.BillboardResponse{
            Id:        value.Id,
            Label:     value.Label,
            ImageUrl:  value.ImageUrl,
            CreatedAt: value.CreatedAt,
        }
        billboards = append(billboards, billboard)
    }

    return billboards, nil
}
