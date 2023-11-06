package controller

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	service "golang-crud-gin/service/size"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type SizeController struct {
	sizeService service.SizeService
}

func NewSizeController(service service.SizeService) *SizeController {
	return &SizeController{
		sizeService: service,
	}
}

func (controller *SizeController) CreateSize(ctx *gin.Context) {
	log.Info().Msg("create size")
	createSizeRequest := request.CreateSizeRequest{}
	if err := ctx.ShouldBindJSON(&createSizeRequest); err != nil {
		helper.ErrorPanic(err)
	}

	controller.sizeService.Create(createSizeRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *SizeController) UpdateSize(ctx *gin.Context) {
	log.Info().Msg("update size")
	sizeID := ctx.Param("sizeId")
	id, err := strconv.Atoi(sizeID)
	helper.ErrorPanic(err)

	updateSizeRequest := request.UpdateSizeRequest{}
	if err := ctx.ShouldBindJSON(&updateSizeRequest); err != nil {
		helper.ErrorPanic(err)
	}
	updateSizeRequest.Id = id

	controller.sizeService.Update(updateSizeRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *SizeController) DeleteSize(ctx *gin.Context) {
	log.Info().Msg("delete size")
	sizeID := ctx.Param("sizeId")
	id, err := strconv.Atoi(sizeID)
	helper.ErrorPanic(err)

	controller.sizeService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *SizeController) FindSizeByID(ctx *gin.Context) {
	log.Info().Msg("find size by ID")
	sizeID := ctx.Param("sizeId")
	id, err := strconv.Atoi(sizeID)
	helper.ErrorPanic(err)

	sizeResponse, err := controller.sizeService.FindById(id)
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   "Size not found",
		}
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   sizeResponse,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *SizeController) FindAllSizes(ctx *gin.Context) {
	log.Info().Msg("find all sizes")
	sizeResponses, err := controller.sizeService.FindAll()
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   "No sizes found",
		}
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   sizeResponses,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *SizeController) FindSizeByName(ctx *gin.Context) {
	log.Info().Msg("find size by name controller")
	sizeName := ctx.Query("name")

	if sizeName == "" {
		webResponse := response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "Size name is required",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	sizeResponse, err := controller.sizeService.FindByName(sizeName)
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   "Size not found",
		}
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   sizeResponse,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *SizeController) FindSizesByPage(ctx *gin.Context) {
	log.Info().Msg("find sizes by page controller")

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	sizeResponses, err := controller.sizeService.FindAllWithPagination(page, pageSize)
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   "No sizes found for the specified page",
		}
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   sizeResponses,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
