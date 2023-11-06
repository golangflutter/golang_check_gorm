package repository

import (
	"errors"
	"golang-crud-gin/model"

	"gorm.io/gorm"
)

type ColorRepositoryImpl struct {
	Db *gorm.DB
}

func NewColorRepositoryImpl(Db *gorm.DB) ColorRepository {
	return &ColorRepositoryImpl{Db: Db}
}

func (c *ColorRepositoryImpl) Save(color model.Color) error {
	result := c.Db.Create(&color)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *ColorRepositoryImpl) Update(color model.Color) error {
	result := c.Db.Save(&color)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *ColorRepositoryImpl) Delete(colorId int) error {
	result := c.Db.Where("id = ?", colorId).Delete(&model.Color{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *ColorRepositoryImpl) FindByID(colorId int) (model.Color, error) {
	var color model.Color
	result := c.Db.Find(&color, colorId)
	if result.Error != nil {
		return color, errors.New("color is not found")
	}
	return color, nil
}

func (c *ColorRepositoryImpl) FindAll() ([]model.Color, error) {
	var colors []model.Color
	result := c.Db.Find(&colors)
	if result.Error != nil {
		return nil, result.Error
	}
	return colors, nil
}

func (c *ColorRepositoryImpl) FindByColorName(colorName string) (model.Color, error) {
	var color model.Color
	result := c.Db.First(&color, "name = ?", colorName)
	if result.Error != nil {
		return color, errors.New("color not found")
	}
	return color, nil
}

func (c *ColorRepositoryImpl) FindColorsByPage(pageNumber, pageSize int) ([]model.Color, error) {
	var colors []model.Color
	offset := (pageNumber - 1) * pageSize
	result := c.Db.Offset(offset).Limit(pageSize).Find(&colors)
	if result.Error != nil {
		return nil, result.Error
	}
	return colors, nil
}
