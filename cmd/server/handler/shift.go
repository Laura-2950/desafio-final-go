package handler

import (
	"errors"
	"net/http"

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

func validateEmptysShift(shift *domain.Shift) (bool, error) {
	if shift.Dentist.ID < 0 || shift.Patient.ID < 0 || shift.DateHour == "" {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}
