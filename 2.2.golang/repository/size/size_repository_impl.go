package repository

import (
	"errors"
	"golang-crud-gin/model"

	"gorm.io/gorm"
)

type SizeRepositoryImpl struct {
	Db *gorm.DB
}

func NewSizeRepositoryImpl(Db *gorm.DB) SizeRepository {
	return &SizeRepositoryImpl{Db: Db}
}

func (s *SizeRepositoryImpl) Save(size model.Size) error {
	result := s.Db.Create(&size)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *SizeRepositoryImpl) Update(size model.Size) error {
	result := s.Db.Save(&size) // Use Save to update all fields
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *SizeRepositoryImpl) Delete(sizeId int) error {
	result := s.Db.Where("id = ?", sizeId).Delete(&model.Size{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *SizeRepositoryImpl) FindByID(sizeId int) (model.Size, error) {
	var size model.Size
	result := s.Db.Find(&size, sizeId)
	if result.Error != nil {
		return size, errors.New("size is not found")
	}
	return size, nil
}

func (s *SizeRepositoryImpl) FindAll() ([]model.Size, error) {
	var sizes []model.Size
	result := s.Db.Find(&sizes)
	if result.Error != nil {
		return nil, result.Error
	}
	return sizes, nil
}

func (s *SizeRepositoryImpl) FindBySizeName(sizeName string) (model.Size, error) {
	var size model.Size
	result := s.Db.First(&size, "name = ?", sizeName)
	if result.Error != nil {
		return size, errors.New("size not found")
	}
	return size, nil
}

func (s *SizeRepositoryImpl) FindSizesByPage(pageNumber, pageSize int) ([]model.Size, error) {
	var sizes []model.Size
	offset := (pageNumber - 1) * pageSize
	result := s.Db.Offset(offset).Limit(pageSize).Find(&sizes)
	if result.Error != nil {
		return nil, result.Error
	}
	return sizes, nil
}
