package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Laura-2950/desafio-final-go/internal/domain"
	"github.com/Laura-2950/desafio-final-go/internal/shift"
	"github.com/Laura-2950/desafio-final-go/pkg/web"
	"github.com/gin-gonic/gin"
)

type ShiftHandler struct {
	ShiftService shift.IService
}

func (h *ShiftHandler) NewShift(ctx *gin.Context) {
	var shift *domain.Shift

	err := ctx.ShouldBindJSON(&shift)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("invalid body"))
		return
	}
	valid, err := validateEmptysShift(shift)
	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	newShift, err := h.ShiftService.CreateNewShift(shift)
	if err != nil {
		if errApi, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, newShift)
}

func (h *ShiftHandler) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Invalid ID"))
		return
	}

	erro := h.ShiftService.Delete(id)
	if erro != nil {
		if errApi, ok := erro.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, erro)
		return
	}

	ctx.JSON(http.StatusOK, "shift removed successfully")
}

func (h *ShiftHandler) UpdateShift(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Invalid ID"))
		return
	}

	var sh *domain.Shift
	err = ctx.ShouldBindJSON(&sh)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("invalid body"))
		return
	}

	valid, err := validateEmptysShift(sh)
	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	shift, err := h.ShiftService.UpdateShift(id, sh)
	if err != nil {
		if errApi, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, shift)
}

func (h *ShiftHandler) UpdatePartialShift(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Invalid ID"))
		return
	}

	var request *domain.RequestShift
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("invalid body"))
		return
	}
	shiftUpdate := domain.Shift{
		Patient:  request.Patient,
		Dentist:  request.Dentist,
		DateHour: request.DateHour,
	}

	shift, err := h.ShiftService.UpdateShift(id, &shiftUpdate)
	if err != nil {
		if errApi, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, shift)
}

func (h *ShiftHandler) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Invalid ID"))
		return
	}

	shiftFounded, err := h.ShiftService.GetShiftByID(id)
	if err != nil {
		if errApi, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, &shiftFounded)
}

func validateEmptysShift(shift *domain.Shift) (bool, error) {
	if shift.Dentist <= 0 || shift.Patient <= 0 || shift.DateHour == "" {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}
