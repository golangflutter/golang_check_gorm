package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
)

type ProductService interface {
    CreateProduct(request request.CreateProductRequest) error
    UpdateProduct(request request.UpdateProductRequest) error
    DeleteProduct(id int) error
    FindProductByID(id int) (response.ProductResponse, error)
    FindAllProducts() ([]response.ProductResponse, error)
    FindProductsByCategory(categoryID int) ([]response.ProductResponse, error)
    FindProductsByColor(colorID int) ([]response.ProductResponse, error)
    FindProductsBySize(sizeID int) ([]response.ProductResponse, error)
    FindProductsByCharacteristic(characteristic string) ([]response.ProductResponse, error)
    FindProductsByPage(pageNumber, pageSize int) ([]response.ProductResponse, error)
    FindProductsByProductName(name string) ([]response.ProductResponse, error)
}
