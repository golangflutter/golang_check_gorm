package repository

import (
	"fmt"
	"golang-crud-gin/model"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
    Db *gorm.DB
}

func NewProductRepositoryImpl(Db *gorm.DB) ProductRepository {
    return &ProductRepositoryImpl{Db: Db}
}

func (p *ProductRepositoryImpl) Create(product model.Product) error {
    fmt.Printf("createProductRequest repository:1111111111111" )
    result := p.Db.Create(&product)
    if result.Error != nil {
        return result.Error
    }
	fmt.Printf("createProductRequest controller: %+v\n", product)
    return nil
}

func (p *ProductRepositoryImpl) Update(product model.Product) error {
    result := p.Db.Save(&product)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (p *ProductRepositoryImpl) Delete(id int) error {
    result := p.Db.Where("id = ?", id).Delete(&model.Product{})
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (p *ProductRepositoryImpl) FindById(id int) (model.Product, error) {
    var product model.Product
    fmt.Printf("createProductRequest repository product: %+v\n", product)

    result := p.Db.Preload("Category").Preload("Color").Preload("Size").First(&product, id)

    // result := p.Db.First(&product, id)
    if result.Error != nil {
        return product, result.Error
    }

    fmt.Printf("Product retrieved from the database: %+v\n", product)
    return product, nil
}


func (p *ProductRepositoryImpl) FindAll() ([]model.Product, error) {

 
    var products []model.Product
    result := p.Db.Preload("Category").Preload("Color").Preload("Size").Find(&products)
    if result.Error != nil {
        return nil, result.Error
    }
    fmt.Printf("Product retrieved from the database repository: %+v\n", products)

    return products, nil
}

func (p *ProductRepositoryImpl) FindByCategory(categoryId int) ([]model.Product, error) {
    var products []model.Product
    result := p.Db.Where("category_id = ?", categoryId).Find(&products)
    if result.Error != nil {
        return nil, result.Error
    }
    return products, nil
}

func (p *ProductRepositoryImpl) FindByColor(colorID int) ([]model.Product, error) {
    var products []model.Product
    result := p.Db.Where("color_id = ?", colorID).Find(&products)
    if result.Error != nil {
        return nil, result.Error
    }
    return products, nil
}

func (p *ProductRepositoryImpl) FindBySize(sizeID int) ([]model.Product, error) {
    var products []model.Product
    result := p.Db.Where("size_id = ?", sizeID).Find(&products)
    if result.Error != nil {
        return nil, result.Error
    }
    return products, nil
}

func (p *ProductRepositoryImpl) FindByCharacteristic(characteristic string) ([]model.Product, error) {
    var products []model.Product
    result := p.Db.Where("characteristic = ?", characteristic).Find(&products)
    if result.Error != nil {
        return nil, result.Error
    }
    return products, nil
}

func (p *ProductRepositoryImpl) FindProductsByPage(pageNumber, pageSize int) ([]model.Product, error) {
    var products []model.Product
    offset := (pageNumber - 1) * pageSize
    result := p.Db.Offset(offset).Limit(pageSize).Find(&products)
    if result.Error != nil {
        return nil, result.Error
    }
    return products, nil
}


func (p *ProductRepositoryImpl) FindByProductName(name string) ([]model.Product, error) {

    fmt.Printf("repository  FindProductsByProductName:111111 %+v\n", name)
    var products []model.Product
    result := p.Db.Where("name = ?", name).Find(&products)
    if result.Error != nil {
        return nil, result.Error
    }

    fmt.Printf("repository  product result%+v\n", &result)
    return products, nil
}
