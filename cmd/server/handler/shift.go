package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	valid, err = validateDateHour(shift.DateHour)
	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	newShift, err := h.ShiftService.CreateNewShift(shift)
	if err != nil {
		if errApi, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
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

	ctx.JSON(http.StatusOK, "patient removed successfully")
}


func validateEmptysShift(shift *domain.Shift) (bool, error) {
	if shift.Dentist <= 0 || shift.Patient <=  0 || shift.DateHour == "" {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

func validateDateHour(exp string) (bool, error) {
	dateHour := strings.Split(exp, " ")
	fmt.Println(dateHour)
	dates := strings.Split(dateHour[0], "/")
	fmt.Println(dates)
	hour := strings.Split(dateHour[1], ":")
	fmt.Println(hour)
//-------------Date-------------------------//
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("invalid registration date, must be in format: dd/mm/yyyy")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("invalid registration date, must be numbers")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) || (list[1] < 1 || list[1] > 12) || (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("invalid registration date, date must be between 1 and 31/12/9999")
	}

//-------------Hour-------------------------//
	list2 := []int{}
	if len(hour) != 2 {
		return false, errors.New("invalid hour, must be in format: hh:mm")
	}
	for value := range hour {
		number, err := strconv.Atoi(hour[value])
		if err != nil {
			return false, errors.New("invalid hour, must be numbers")
		}
		list2 = append(list2, number)
	}
	condition2 := (list2[0] < 0 || list2[0] > 23) || (list2[1] < 0 || list2[1] > 59)
	if condition2 {
		return false, errors.New("invalid registration date and hour, date must be between 1 and 31/12/9999, and hour must be between 00:01 and 23:59")
	}
	return true, nil
}