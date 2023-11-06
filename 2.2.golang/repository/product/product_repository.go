package repository

import "golang-crud-gin/model"

type ProductRepository interface {
    // Create a new product
    Create(product model.Product) error

    // Update an existing product
    Update(product model.Product) error

    // Delete a product by ID
    Delete(id int) error

    // Find a product by its ID
    FindById(id int) (model.Product, error)

    // Find all products
    FindAll() ([]model.Product, error)

    // Find products by category
    FindByCategory(categoryId int) ([]model.Product, error)

    // Find products by color
    FindByColor(colorId int) ([]model.Product, error)

    // Find products by size
    FindBySize(sizeId int) ([]model.Product, error)

    // Find products by characteristic
    FindByCharacteristic(characteristic string) ([]model.Product, error)

        // Find products by name
        FindByProductName(name string) ([]model.Product, error)

    // Pagination and filtering
    FindProductsByPage(pageNumber, pageSize int) ([]model.Product, error)

}
