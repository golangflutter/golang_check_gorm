package controller

import (
	"fmt"
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	service "golang-crud-gin/service/bill_boad"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type BillboardController struct {
	billboardService service.BillboardService // Updated service interface
}

func NewBillboardController(service service.BillboardService) *BillboardController {
	return &BillboardController{
		billboardService: service,
	}
}

func (controller *BillboardController) CreateBillboard(ctx *gin.Context) {
    log.Info().Msg("create billboard")
    createBillboardRequest := request.CreateBillboardRequest{}
    if err := ctx.ShouldBindJSON(&createBillboardRequest); err != nil {
        helper.ErrorPanic(err)
    }

    // Log the createBillboardRequest for debugging
    fmt.Printf("createBillboardRequest controller :askfdjhaskdhfksadhfkhjds %+v\n", createBillboardRequest)

    // Now you have the ImageUrl as a slice of strings
    // You can access them using createBillboardRequest.ImageUrl
    // and use it as needed in your service.
    controller.billboardService.Create(createBillboardRequest)

    webResponse := response.Response{
        Code:   http.StatusOK,
        Status: "Ok",
        Data:   nil,
    }
    ctx.JSON(http.StatusOK, webResponse)
}


func (controller *BillboardController) UpdateBillboard(ctx *gin.Context) {
	log.Info().Msg("update billboard")
	billboardID := ctx.Param("billboardId")
	id, err := strconv.Atoi(billboardID)
	helper.ErrorPanic(err)

	updateBillboardRequest := request.UpdateBillboardRequest{}
	if err := ctx.ShouldBindJSON(&updateBillboardRequest); err != nil {
		helper.ErrorPanic(err)
	}
	updateBillboardRequest.Id = id

	controller.billboardService.Update(updateBillboardRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BillboardController) DeleteBillboard(ctx *gin.Context) {
	log.Info().Msg("delete billboard")
	billboardID := ctx.Param("billboardId")
	id, err := strconv.Atoi(billboardID)
	helper.ErrorPanic(err)

	controller.billboardService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BillboardController) FindBillboardByID(ctx *gin.Context) {
	log.Info().Msg("find billboard by ID")
	billboardID := ctx.Param("billboardId")
	id, err := strconv.Atoi(billboardID)
	helper.ErrorPanic(err)

	billboardResponse, err := controller.billboardService.FindByID(id)
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   "Billboard not found",
		}
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   billboardResponse,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BillboardController) FindAllBillboards(ctx *gin.Context) {
	log.Info().Msg("find all billboards")
	billboardResponses, err := controller.billboardService.FindAll()
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   "No billboards found",
		}
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   billboardResponses,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BillboardController) FindBillboardByLabel(ctx *gin.Context) {
	log.Info().Msg("find billboard by label controller")
	label := ctx.Query("label")

	if label == "" {
		webResponse := response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "Billboard label is required",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	billboardResponse, err := controller.billboardService.FindByLabel(label)
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   "Billboard not found",
		}
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   billboardResponse,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BillboardController) FindBillboardsByPage(ctx *gin.Context) {
	log.Info().Msg("find billboards by page controller")

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	billboardResponses, err := controller.billboardService.FindAllWithPagination(page, pageSize)
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   "No billboards found for the specified page",
		}
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   billboardResponses,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
