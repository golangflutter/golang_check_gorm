package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	repository "golang-crud-gin/repository/color"

	"github.com/go-playground/validator/v10"
)

type ColorServiceImpl struct {
	ColorRepository repository.ColorRepository
	Validate        *validator.Validate
}

func NewColorServiceImpl(colorRepository repository.ColorRepository, validate *validator.Validate) ColorService {
	return &ColorServiceImpl{
		ColorRepository: colorRepository,
		Validate:        validate,
	}
}

// Create implements ColorService
func (c *ColorServiceImpl) Create(color request.CreateColorRequest) {
	err := c.Validate.Struct(color)
	helper.ErrorPanic(err)
	colorModel := model.Color{
		Name:  color.Name,
		Value: color.Value,
	}
	err = c.ColorRepository.Save(colorModel)
	helper.ErrorPanic(err)
}

// Delete implements ColorService
func (c *ColorServiceImpl) Delete(colorID int) {
	err := c.ColorRepository.Delete(colorID)
	helper.ErrorPanic(err)
}

// FindAll implements ColorService
func (c *ColorServiceImpl) FindAll() ([]response.ColorResponse, error) {
	result, err := c.ColorRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var colors []response.ColorResponse
	for _, value := range result {
		color := response.ColorResponse{
			Id:        value.Id,
			Name:      value.Name,
			Value:     value.Value,
			CreatedAt: value.CreatedAt,
		}
		colors = append(colors, color)
	}

	return colors, nil
}

// FindById implements ColorService
func (c *ColorServiceImpl) FindById(colorID int) (response.ColorResponse, error) {
	colorData, err := c.ColorRepository.FindByID(colorID)
	if err != nil {
		return response.ColorResponse{}, err
	}

	colorResponse := response.ColorResponse{
		Id:        colorData.Id,
		Name:      colorData.Name,
		Value:     colorData.Value,
		CreatedAt: colorData.CreatedAt,
	}
	return colorResponse, nil
}

// Update implements ColorService
func (c *ColorServiceImpl) Update(color request.UpdateColorRequest) {
	colorData, err := c.ColorRepository.FindByID(color.Id)
	helper.ErrorPanic(err)
	colorData.Name = color.Name
	colorData.Value = color.Value
	err = c.ColorRepository.Update(colorData)
	helper.ErrorPanic(err)
}

// FindByName implements ColorService
func (c *ColorServiceImpl) FindByName(colorName string) (response.ColorResponse, error) {
	colorData, err := c.ColorRepository.FindByColorName(colorName)
	if err != nil {
		return response.ColorResponse{}, err
	}

	colorResponse := response.ColorResponse{
		Id:        colorData.Id,
		Name:      colorData.Name,
		Value:     colorData.Value,
		CreatedAt: colorData.CreatedAt,
	}
	return colorResponse, nil
}

// FindAllWithPagination implements ColorService with pagination
func (c *ColorServiceImpl) FindAllWithPagination(pageNumber, pageSize int) ([]response.ColorResponse, error) {
	result, err := c.ColorRepository.FindColorsByPage(pageNumber, pageSize)
	if err != nil {
		return nil, err
	}

	var colors []response.ColorResponse
	for _, value := range result {
		color := response.ColorResponse{
			Id:        value.Id,
			Name:      value.Name,
			Value:     value.Value,
			CreatedAt: value.CreatedAt,
		}
		colors = append(colors, color)
	}

	return colors, nil
}


