package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	repository "golang-crud-gin/repository/size"

	"github.com/go-playground/validator/v10"
)

type SizeServiceImpl struct {
	SizeRepository repository.SizeRepository
	Validate      *validator.Validate
}

func NewSizeServiceImpl(sizeRepository repository.SizeRepository, validate *validator.Validate) SizeService {
	return &SizeServiceImpl{
		SizeRepository: sizeRepository,
		Validate:      validate,
	}
}

// Create implements SizeService
func (s *SizeServiceImpl) Create(size request.CreateSizeRequest) {
	err := s.Validate.Struct(size)
	helper.ErrorPanic(err)
	sizeModel := model.Size{
		Name:  size.Name,
		Value: size.Value,
	}
	err = s.SizeRepository.Save(sizeModel)
	helper.ErrorPanic(err)
}

// Delete implements SizeService
func (s *SizeServiceImpl) Delete(sizeID int) {
	err := s.SizeRepository.Delete(sizeID)
	helper.ErrorPanic(err)
}

// FindAll implements SizeService
func (s *SizeServiceImpl) FindAll() ([]response.SizeResponse, error) {
	result, err := s.SizeRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var sizes []response.SizeResponse
	for _, value := range result {
		size := response.SizeResponse{
			Id:        value.Id,
			Name:      value.Name,
			Value:     value.Value,
			CreatedAt: value.CreatedAt,
		}
		sizes = append(sizes, size)
	}

	return sizes, nil
}

// FindById implements SizeService
func (s *SizeServiceImpl) FindById(sizeID int) (response.SizeResponse, error) {
	sizeData, err := s.SizeRepository.FindByID(sizeID)
	if err != nil {
		return response.SizeResponse{}, err
	}

	sizeResponse := response.SizeResponse{
		Id:        sizeData.Id,
		Name:      sizeData.Name,
		Value:     sizeData.Value,
		CreatedAt: sizeData.CreatedAt,
	}
	return sizeResponse, nil
}

// Update implements SizeService
func (s *SizeServiceImpl) Update(size request.UpdateSizeRequest) {
	sizeData, err := s.SizeRepository.FindByID(size.Id)
	helper.ErrorPanic(err)
	sizeData.Name = size.Name
	sizeData.Value = size.Value
	err = s.SizeRepository.Update(sizeData)
	helper.ErrorPanic(err)
}

// FindByName implements SizeService
func (s *SizeServiceImpl) FindByName(sizeName string) (response.SizeResponse, error) {
	sizeData, err := s.SizeRepository.FindBySizeName(sizeName)
	if err != nil {
		return response.SizeResponse{}, err
	}

	sizeResponse := response.SizeResponse{
		Id:        sizeData.Id,
		Name:      sizeData.Name,
		Value:     sizeData.Value,
		CreatedAt: sizeData.CreatedAt,
	}
	return sizeResponse, nil
}

// FindAllWithPagination implements SizeService with pagination
func (s *SizeServiceImpl) FindAllWithPagination(pageNumber, pageSize int) ([]response.SizeResponse, error) {
	result, err := s.SizeRepository.FindSizesByPage(pageNumber, pageSize)
	if err != nil {
		return nil, err
	}

	var sizes []response.SizeResponse
	for _, value := range result {
		size := response.SizeResponse{
			Id:        value.Id,
			Name:      value.Name,
			Value:     value.Value,
			CreatedAt: value.CreatedAt,
		}
		sizes = append(sizes, size)
	}

	return sizes, nil
}
