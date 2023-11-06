package controller

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	service "golang-crud-gin/service/color"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ColorController struct {
	colorService service.ColorService
}

func NewColorController(service service.ColorService) *ColorController {
	return &ColorController{
		colorService: service,
	}
}

func (controller *ColorController) CreateColor(ctx *gin.Context) {
	log.Info().Msg("create color")
	createColorRequest := request.CreateColorRequest{}
	if err := ctx.ShouldBindJSON(&createColorRequest); err != nil {
		helper.ErrorPanic(err)
	}

	controller.colorService.Create(createColorRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ColorController) UpdateColor(ctx *gin.Context) {
	log.Info().Msg("update color")
	colorID := ctx.Param("colorId")
	id, err := strconv.Atoi(colorID)
	helper.ErrorPanic(err)

	updateColorRequest := request.UpdateColorRequest{}
	if err := ctx.ShouldBindJSON(&updateColorRequest); err != nil {
		helper.ErrorPanic(err)
	}
	updateColorRequest.Id = id

	controller.colorService.Update(updateColorRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ColorController) DeleteColor(ctx *gin.Context) {
	log.Info().Msg("delete color")
	colorID := ctx.Param("colorId")
	id, err := strconv.Atoi(colorID)
	helper.ErrorPanic(err)

	controller.colorService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ColorController) FindColorByID(ctx *gin.Context) {
	log.Info().Msg("find color by ID")
	colorID := ctx.Param("colorId")
	id, err := strconv.Atoi(colorID)
	helper.ErrorPanic(err)

	colorResponse, err := controller.colorService.FindById(id)
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   "Color not found",
		}
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   colorResponse,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ColorController) FindAllColors(ctx *gin.Context) {
	log.Info().Msg("find all colors")
	colorResponses, err := controller.colorService.FindAll()
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   "No colors found",
		}
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   colorResponses,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ColorController) FindColorByName(ctx *gin.Context) {
	log.Info().Msg("find color by name controller")
	colorName := ctx.Query("name")

	if colorName == "" {
		webResponse := response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "Color name is required",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	colorResponse, err := controller.colorService.FindByName(colorName)
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   "Color not found",
		}
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   colorResponse,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ColorController) FindColorsByPage(ctx *gin.Context) {
	log.Info().Msg("find colors by page controller")

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	colorResponses, err := controller.colorService.FindAllWithPagination(page, pageSize)
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   "No colors found for the specified page",
		}
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   colorResponses,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
