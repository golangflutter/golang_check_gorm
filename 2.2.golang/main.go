package main

import (
	"golang-crud-gin/config"
	"golang-crud-gin/controller"
	_ "golang-crud-gin/docs"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"
	billboard_repository "golang-crud-gin/repository/bill_board"
	categories_repository "golang-crud-gin/repository/categories"
	color_repository "golang-crud-gin/repository/color"
	product_repository "golang-crud-gin/repository/product"
	size_repository "golang-crud-gin/repository/size" // Import size repository
	"golang-crud-gin/router"
	"golang-crud-gin/service"
	billboard_service "golang-crud-gin/service/bill_boad"
	categories_service "golang-crud-gin/service/categories"
	color_service "golang-crud-gin/service/color"
	product_service "golang-crud-gin/service/product"
	size_service "golang-crud-gin/service/size" // Import size service
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {
    log.Info().Msg("Started Server!")
    // Database
    db := config.DatabaseConnection()
    validate := validator.New()

    db.Table("tags").AutoMigrate(&model.Tags{})
    db.Table("users").AutoMigrate(&model.Users{})
    // AutoMigrate for the "Category" model
    db.Table("categories").AutoMigrate(&model.Category{})
    db.Table("sizes").AutoMigrate(&model.Size{})
    db.Table("colors").AutoMigrate(&model.Color{})
    db.Table("bill_boards").AutoMigrate(&model.Billboard{})
    db.Table("products").AutoMigrate(&model.Product{})
    db.Table("orderItems").AutoMigrate(&model.OrderItem{})
    db.Table("orders").AutoMigrate(&model.Order{})
    // Repository
    productRepository := product_repository.NewProductRepositoryImpl(db)
    tagsRepository := repository.NewTagsREpositoryImpl(db)
    userRepository := repository.NewUsersRepositoryImpl(db)

    categoryRepository := categories_repository.NewCategoryRepositoryImpl(db)
    sizeRepository := size_repository.NewSizeRepositoryImpl(db)
    colorRepository := color_repository.NewColorRepositoryImpl(db)
    billboardRepository := billboard_repository.NewBillboardRepositoryImpl(db)
    // Service
    tagsService := service.NewTagsServiceImpl(tagsRepository, validate)
    authenticationService := service.NewAuthenticationServiceImpl(userRepository, validate)

    categoryService := categories_service.NewCategoryServiceImpl(categoryRepository, validate)
    sizeService := size_service.NewSizeServiceImpl(sizeRepository, validate)
    colorService := color_service.NewColorServiceImpl(colorRepository, validate)
    billboardService := billboard_service.NewBillboardServiceImpl(billboardRepository, validate)
    productService := product_service.NewProductServiceImpl(productRepository, validate)
    // Controller
    productController := controller.NewProductController(productService)
    tagsController := controller.NewTagsController(tagsService)
    authenticationController := controller.NewAuthenticationController(authenticationService)
    usersController := controller.NewUsersController(userRepository)

    categoryController := controller.NewCategoryController(categoryService)
    sizeController := controller.NewSizeController(sizeService)
    colorController := controller.NewColorController(colorService)
    billboardController := controller.NewBillboardController(billboardService)
    // Create the RefreshTokenController
    refreshTokenController := controller.RefreshTokenController

    // Image Upload Controller
    imageController := controller.NewImageUploadController()

    // Router
    routes := router.NewRouter(tagsController, userRepository, authenticationController, usersController, refreshTokenController, imageController, categoryController, sizeController, colorController, billboardController, productController)

    // Server configuration
    server := &http.Server{
        Addr:    ":8888",
        Handler: routes,
    }

    err := server.ListenAndServe()
    helper.ErrorPanic(err)
}
