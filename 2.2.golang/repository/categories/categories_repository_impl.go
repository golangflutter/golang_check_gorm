package repository

import (
	"errors"
	"fmt"
	"golang-crud-gin/data/request"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"log"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	Db *gorm.DB
}

func NewCategoryRepositoryImpl(Db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{Db: Db}
}

// Delete implements CategoryRepository
func (c *CategoryRepositoryImpl) Delete(categoryID int) {
	var category model.Category
	fmt.Printf("Deleting category CategoryRepositoryImpl with ID: %d\n", categoryID)

	result := c.Db.Where("id = ?", categoryID).Delete(&category)
	helper.ErrorPanic(result.Error)
}

// FindAll implements CategoryRepository
func (c *CategoryRepositoryImpl) FindAll() []model.Category {
	var categories []model.Category
	results := c.Db.Find(&categories)
	helper.ErrorPanic(results.Error)
	return categories
}

// FindById implements CategoryRepository
func (c *CategoryRepositoryImpl) FindById(categoryID int) (model.Category, error) {
	log.Println("this is FindById category")
	var category model.Category
	result := c.Db.Find(&category, categoryID)
	if result.Error != nil {
		return category, errors.New("category is not found")
	}
	return category, nil
}

// Save implements CategoryRepository
func (c *CategoryRepositoryImpl) Save(category model.Category) {
	result := c.Db.Create(&category)
	helper.ErrorPanic(result.Error)
}

// Update implements CategoryRepository
func (c *CategoryRepositoryImpl) Update(category model.Category) {
	var updateCategory = request.UpdateCategoryRequest{
		Id:   category.Id,
		Name: category.Name,
		// Add other fields as needed
	}
	result := c.Db.Model(&category).Updates(updateCategory)
	helper.ErrorPanic(result.Error)
}

// FindByCategoryName implements CategoryRepository
func (c *CategoryRepositoryImpl) FindByCategoryName(categoryName string) (model.Category, error) {
	var category model.Category
	result := c.Db.First(&category, "name = ?", categoryName)
	if result.Error != nil {
		return category, errors.New("category not found")
	}
	return category, nil
}

func (c *CategoryRepositoryImpl) FindCategoriesByPage( pageNumber, pageSize int) ([]model.Category, error) {
    // Validate and transform the incoming data


    var categories []model.Category
    offset := (pageNumber - 1) * pageSize
    result := c.Db.Offset(offset).Limit(pageSize).Find(&categories)
    if result.Error != nil {
        return categories, result.Error
    }
    return categories, nil
}

func (c *CategoryRepositoryImpl) FindCategoriesByCharacteristicName(characteristic string, pageNumber, pageSize int) ([]model.Category, error) {
    var categories []model.Category

    // Calculate the offset based on the page number and page size
    offset := (pageNumber - 1) * pageSize

    // Use the Preload method to fetch related data if needed
    result := c.Db.
        Where("name LIKE ?", "%"+characteristic+"%").
        Offset(offset).
        Limit(pageSize).
 
        Find(&categories)

    if result.Error != nil {
        return categories, result.Error
    }

    return categories, nil
}

func (c *CategoryRepositoryImpl) FindCategoriesByCharacteristicFields(characteristic string, pageNumber, pageSize int) ([]model.Category, error) {
    var categories []model.Category

    // Calculate the offset based on the page number and page size
    offset := (pageNumber - 1) * pageSize

    // Use the Preload method to fetch related data if needed
    result := c.Db.
        Where("name LIKE ? OR description LIKE ?", "%"+characteristic+"%", "%"+characteristic+"%").
        Offset(offset).
        Limit(pageSize).
        Order("created_at DESC").
        Find(&categories)

    if result.Error != nil {
        return categories, result.Error
    }

    return categories, nil
}
