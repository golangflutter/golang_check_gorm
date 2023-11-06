package controller

import (
	"fmt"
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	service "golang-crud-gin/service/categories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type CategoryController struct {
    categoryService service.CategoryService
}

func NewCategoryController(service service.CategoryService) *CategoryController {
    return &CategoryController{
        categoryService: service,
    }
}

// CreateCategory		godoc
// @Summary				Create category
// @Description			Save category data in Db.
// @Param				category body request.CreateCategoryRequest true "Create category"
// @Produce				application/json
// @Tags				categories
// @Success				200 {object} response.Response{}
// @Router				/categories [post]
func (controller *CategoryController) Create(ctx *gin.Context) {
    log.Info().Msg("create category")
    createCategoryRequest := request.CreateCategoryRequest{}
    fmt.Println("create category request", createCategoryRequest) 

    err := ctx.ShouldBindJSON(&createCategoryRequest)
    if err != nil {
        fmt.Println("Error binding JSON:", err) // Add this line for error debugging
        helper.ErrorPanic(err)
    }

    controller.categoryService.Create(createCategoryRequest)

    webResponse := response.Response{
        Code:   http.StatusOK,
        Status: "Ok",
        Data:   nil,
    }
    ctx.Header("Content-Type", "application/json")
    ctx.JSON(http.StatusOK, webResponse)
}


// UpdateCategory		godoc
// @Summary				Update category
// @Description			Update category data.
// @Param				categoryId path string true "update category by id"
// @Param				category body request.UpdateCategoryRequest true  "Update category"
// @Tags				categories
// @Produce				application/json
// @Success				200 {object} response.Response{}
// @Router				/categories/{categoryId} [patch]
func (controller *CategoryController) Update(ctx *gin.Context) {
    log.Info().Msg("update category")
    updateCategoryRequest := request.UpdateCategoryRequest{}
    err := ctx.ShouldBindJSON(&updateCategoryRequest)
    helper.ErrorPanic(err)

    categoryId := ctx.Param("categoryId")
    id, err := strconv.Atoi(categoryId)
    helper.ErrorPanic(err)
    updateCategoryRequest.Id = id

    controller.categoryService.Update(updateCategoryRequest)

    webResponse := response.Response{
        Code:   http.StatusOK,
        Status: "Ok",
        Data:   nil,
    }
    ctx.Header("Content-Type", "application/json")
    ctx.JSON(http.StatusOK, webResponse)
}

// DeleteCategory		godoc
// @Summary				Delete category
// @Description			Remove category data by id.
// @Produce				application/json
// @Tags				categories
// @Success				200 {object} response.Response{}
// @Router				/categories/{categoryId} [delete]
func (controller *CategoryController) Delete(ctx *gin.Context) {
    log.Info().Msg("delete category")
    categoryId := ctx.Param("categoryId")
    id, err := strconv.Atoi(categoryId)
    helper.ErrorPanic(err)

    controller.categoryService.Delete(id)

    webResponse := response.Response{
        Code:   http.StatusOK,
        Status: "Ok",
        Data:   nil,
    }
    ctx.Header("Content-Type", "application/json")
    ctx.JSON(http.StatusOK, webResponse)
}

// FindByIdCategory 		godoc
// @Summary					Get Single category by id.
// @Param					categoryId path string true "update category by id"
// @Description				Return the category whose categoryId value matches id.
// @Produce					application/json
// @Tags					categories
// @Success					200 {object} response.Response{}
// @Router					/categories/{categoryId} [get]
func (controller *CategoryController) FindById(ctx *gin.Context) {
    log.Info().Msg("findbyid category aodsjflaksdf")
    categoryId := ctx.Param("categoryId")
    id, err := strconv.Atoi(categoryId)
    helper.ErrorPanic(err)

    categoryResponse := controller.categoryService.FindById(id)

    webResponse := response.Response{
        Code:   http.StatusOK,
        Status: "Ok",
        Data:   categoryResponse,
    }
    ctx.Header("Content-Type", "application/json")
    ctx.JSON(http.StatusOK, webResponse)
}

// FindAllCategories 		godoc
// @Summary				Get All categories.
// @Description			Return a list of categories.
// @Tags				categories
// @Success				200 {object} response.Response{}
// @Router				/categories [get]
func (controller *CategoryController) FindAll(ctx *gin.Context) {
    log.Info().Msg("findAll categories")
    categoryResponse := controller.categoryService.FindAll()

    webResponse := response.Response{
        Code:   http.StatusOK,
        Status: "Ok",
        Data:   categoryResponse,
    }
    ctx.Header("Content-Type", "application/json")
    ctx.JSON(http.StatusOK, webResponse)
}


// GetCategoryByName    godoc
// @Summary               Get category by name
// @Description           Return a category by its name.
// @Param                  name query string true "Category name"
// @Produce               application/json
// @Tags                  categories
// @Success               200 {object} response.Response{}
// @Router                /categories [get]
func (controller *CategoryController) FindByCategoryName(ctx *gin.Context) {
    log.Info().Msg("get category by name controller 1")
    categoryName := ctx.Query("Name") // Retrieve the category name from the query parameters
    log.Info().Msg("get category by name controller 2")
 
    if categoryName == "" {
        webResponse := response.Response{
            Code:   http.StatusBadRequest,
            Status: "Bad Request",
            Data:   "Category name is required",
        }
        ctx.Header("Content-Type", "application/json")
        ctx.JSON(http.StatusBadRequest, webResponse)
        return
    }
    log.Info().Msg("get category by name controller 3")
    // Call the service function to get the category by name
    categoryResponse, err := controller.categoryService.FindByName(categoryName)
    log.Info().Msg("get category by name controller 4")
    if err != nil {
        webResponse := response.Response{
            Code:   http.StatusNotFound,
            Status: "Not Found",
            Data:   "Category not found",
        }
        ctx.Header("Content-Type", "application/json")
        ctx.JSON(http.StatusNotFound, webResponse)
        return
    }
    log.Info().Msg("get category by name controller 5")
    webResponse := response.Response{
        Code:   http.StatusOK,
        Status: "Ok",
        Data:   categoryResponse,
    }
    log.Info().Msg("get category by name controller 7")
    ctx.Header("Content-Type", "application/json")
    ctx.JSON(http.StatusOK, webResponse)
}
func (controller *CategoryController) FindCategoriesByPage(ctx *gin.Context) {
    log.Info().Msg("get categories bFindCategoriesByPage 1")

    // Retrieve pagination parameters from the query string
    page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
    if err != nil || page < 1 {
        page = 1
    }

    pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
    if err != nil || pageSize < 1 {
        pageSize = 10
    }


    log.Info().Msg("get categories bFindCategoriesByPage 2")
    // Call the service function to get categories by name with pagination
    categoryResponses, err := controller.categoryService.FindAllWithPagination( page, pageSize)

    if err != nil {
        webResponse := response.Response{
            Code:   http.StatusNotFound,
            Status: "Not Found",
            Data:   "No categories found",
        }
        ctx.Header("Content-Type", "application/json")
        ctx.JSON(http.StatusNotFound, webResponse)
        return
    }
    log.Info().Msg("get categories bFindCategoriesByPage 3")

    webResponse := response.Response{
        Code:   http.StatusOK,
        Status: "Ok",
        Data:   categoryResponses,
    }
    log.Info().Msg("get categories bFindCategoriesByPage 4")
    log.Info().Msg("get categories by page controller 5")
    ctx.Header("Content-Type", "application/json")
    ctx.JSON(http.StatusOK, webResponse)
}
func (controller *CategoryController) FindCategoriesByCharacteristicName(ctx *gin.Context) {
    log.Info().Msg("find categories by characteristic name")
    
    // Retrieve characteristic, page, and pageSize from the query string
    characteristic := ctx.Query("characteristic")
    page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
    if err != nil || page < 1 {
        page = 1
    }

    pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
    if err != nil || pageSize < 1 {
        pageSize = 10
    }

    // Call the service function to get categories by characteristic with pagination
    categoryResponses, err := controller.categoryService.FindCategoriesByCharacteristicName(characteristic, page, pageSize)

    if err != nil {
        webResponse := response.Response{
            Code:   http.StatusNotFound,
            Status: "Not Found",
            Data:   "No categories found for the specified characteristic",
        }
        ctx.Header("Content-Type", "application/json")
        ctx.JSON(http.StatusNotFound, webResponse)
        return
    }

    webResponse := response.Response{
        Code:   http.StatusOK,
        Status: "Ok",
        Data:   categoryResponses,
    }
    ctx.Header("Content-Type", "application/json")
    ctx.JSON(http.StatusOK, webResponse)
}
