package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/Laura-2950/desafio-final-go/internal/domain"
	"github.com/Laura-2950/desafio-final-go/internal/patient"
	"github.com/Laura-2950/desafio-final-go/pkg/web"
	"github.com/gin-gonic/gin"
)

type PatienttHandler struct {
	PatientService patient.IService
}

func (h *PatienttHandler) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Invalid ID"))
		return
	}

	patientFounded, err := h.PatientService.GetPatientByID(id)
	if err != nil {
		if errApi, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, &patientFounded)
}

func (h *PatienttHandler) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Invalid ID"))
		return
	}

	var pat *domain.Patient
	err = ctx.ShouldBindJSON(&pat)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("invalid body"))
		return
	}

	valid, err := validEmptysPatient(pat)
	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	patient, err := h.PatientService.UpdatePatient(id, pat)
	if err != nil {
		if errApi, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, patient)
}

func (h *PatienttHandler) UpdatePartial(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Invalid ID"))
		return
	}

	var request *domain.RequestPatient
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("invalid body"))
		return
	}
	patientUpdate := domain.Patient{
		Name:             request.Name,
		LastName:         request.LastName,
		Address:          request.Address,
		Dni:              request.Dni,
		RegistrationDate: request.RegistrationDate,
	}

	patient, err := h.PatientService.UpdatePatient(id, &patientUpdate)
	if err != nil {
		if errApi, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, patient)
}

func (h *PatienttHandler) NewPatient(ctx *gin.Context) {
	var patient *domain.Patient

	err := ctx.ShouldBindJSON(&patient)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("invalid body"))
		return
	}
	valid, err := validEmptysPatient(patient)
	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	valid, err = validateRegistrationDate(patient.RegistrationDate)
	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	newPatient, err := h.PatientService.CreateNewPatient(patient)
	if err != nil {
		if errApi, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, newPatient)
}

func (h *PatienttHandler) DeletePatient(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Invalid ID"))
		return
	}

	erro := h.PatientService.DeletePatient(id)
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

func validEmptysPatient(patient *domain.Patient) (bool, error) {
	if patient.Name == "" || patient.LastName == "" || patient.Address == "" || patient.Dni == "" || patient.RegistrationDate == "" {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

func validateRegistrationDate(exp string) (bool, error) {
	dates := strings.Split(exp, "/")
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
	return true, nil
}
