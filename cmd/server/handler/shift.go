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

// NewShift godoc
// @Summary      Create a shift
// @Description  create a shift
// @Tags         Shifts
// @Accept		 json
// @Produce      json
// @Param        token    header   string          true "token"
// @Param        shift  body     domain.Shift true    "Shift to store"
// @Success      201  {object}  domain.Shift
// @Failure      400  {object}  web.ErrorApi
// @Failure      401  {object}  web.ErrorApi
// @Failure      500  {object}  web.ErrorApi
// @Router       /shifts [post]
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

// NewShiftWithCode godoc
// @Summary      Create a shift by dni and registerNumber
// @Description  create a shift by dni and registerNumber
// @Tags         Shifts
// @Accept		 json
// @Produce      json
// @Param        token    header   string          true "token"
// @Param        shift  body     domain.ShiftCode true    "Shift to store"
// @Success      201  {object}  domain.Shift
// @Failure      400  {object}  web.ErrorApi
// @Failure      401  {object}  web.ErrorApi
// @Failure      500  {object}  web.ErrorApi
// @Router       /shifts/code [post]
func (h *ShiftHandler) NewShiftCode(ctx *gin.Context) {
	var shift *domain.ShiftCode

	err := ctx.ShouldBindJSON(&shift)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("invalid body"))
		return
	}
	valid, err := validateEmptysShiftCode(shift)
	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	valid, err = validateDateHour(shift.DateHour)
	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	newShift, err := h.ShiftService.CreateNewShiftCode(shift)
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

// Delete godoc
// @Summary      Delete a shift
// @Description  delete a shift
// @Tags         Shifts
// @Param        token    header   string          true "token"
// @Param        id    path     int    true "Shift id"
// @Success      200  {string}  shift removed successfully
// @Failure      400  {object}  web.ErrorApi
// @Failure      401  {object}  web.ErrorApi
// @Failure      500  {object}  web.ErrorApi
// @Router       /shifts/{id} [delete]
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

// Update godoc
// @Summary      Update a shift
// @Description  update a shift
// @Tags         Shifts
// @Accept       json
// @Produce      json
// @Param        token    header   string          true "token"
// @Param        id    path     int    true "Shift id"
// @Param        Shift  body     domain.Shift true    "Shift to store"
// @Success      200  {object}  domain.Shift
// @Failure      400  {object}  web.ErrorApi
// @Failure      401  {object}  web.ErrorApi
// @Failure      500  {object}  web.ErrorApi
// @Router       /shifts/{id} [put]
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

// UpdatePartial godoc
// @Summary      Update partial a shift
// @Description  update partial a shift
// @Tags         Shifts
// @Accept       json
// @Produce      json
// @Param        token    header   string          true  "token"
// @Param        id       path     int             true "Shift id"
// @Param        Shift  body     domain.RequestShift  true  "Shift to store"
// @Success      200  {object}  domain.Shift
// @Failure      400  {object}  web.ErrorApi
// @Failure      401  {object}  web.ErrorApi
// @Failure      500  {object}  web.ErrorApi
// @Router       /shifts/{id} [patch]
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

// GetById godoc
// @Summary      Show a shift
// @Description  get shift by ID
// @Tags         Shifts
// @Produce      json
// @Param        id   path      int  true  "Shift ID"
// @Success      200  {object}  domain.ResponseShift
// @Failure      400  {object}  web.ErrorApi
// @Failure      500  {object}  web.ErrorApi
// @Router       /shifts/{id} [get]
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

// GetByDni godoc
// @Summary      Show dni shifts
// @Description  get shifts by DNI
// @Tags         Shifts
// @Produce      json
// @Param        dni  query     string  true  "Shift by Dni"
// @Success      200  {array}   domain.ResponseShift
// @Failure      400  {object}  web.ErrorApi
// @Failure      500  {object}  web.ErrorApi
// @Router       /shifts [get]
func (h *ShiftHandler) GetAllByDni(ctx *gin.Context) {
	dniParam := ctx.Query("dni")
	if dniParam == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Invalid DNI"))
		return
	}

	shiftsFounded, err := h.ShiftService.GetAllByDni(dniParam)
	if err != nil {
		if errApi, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &shiftsFounded)
}

func validateEmptysShift(shift *domain.Shift) (bool, error) {
	if shift.Dentist <= 0 || shift.Patient <= 0 || shift.DateHour == "" {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}
func validateEmptysShiftCode(shift *domain.ShiftCode) (bool, error) {
	if shift.Dentist == "" || shift.Patient == "" || shift.DateHour == "" {
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
		return false, errors.New("invalid date, must be in format: dd/mm/yyyy")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("invalid date, must be numbers")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) || (list[1] < 1 || list[1] > 12) || (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("invalid date, date must be between 1 and 31/12/9999")
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
		return false, errors.New("invalid date and hour, date must be between 1 and 31/12/9999, and hour must be between 00:01 and 23:59")
	}
	return true, nil
}
