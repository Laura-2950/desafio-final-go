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

type PatientHandler struct {
	PatientService patient.IService
}

// GetById godoc
// @Summary      Show a patient
// @Description  get patient by ID
// @Tags         Patients
// @Produce      json
// @Param        id   path      int  true  "Patient ID"
// @Success      200  {object}  domain.Patient
// @Failure      400  {object}  web.ErrorApi
// @Failure      500  {object}  web.ErrorApi
// @Router       /patients/{id} [get]
func (h *PatientHandler) GetById(ctx *gin.Context) {
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

// Update godoc
// @Summary      Update a patient
// @Description  update a patient
// @Tags         Patients
// @Accept       json
// @Produce      json
// @Param        token    header   string          true "token"
// @Param        id    path     int    true "Patient id"
// @Param        patient  body     domain.Patient true    "Patient to store"
// @Success      200  {object}  domain.Patient
// @Failure      400  {object}  web.ErrorApi
// @Failure      401  {object}  web.ErrorApi
// @Failure      500  {object}  web.ErrorApi
// @Router       /patients/{id} [put]
func (h *PatientHandler) Update(ctx *gin.Context) {
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

// UpdatePartial godoc
// @Summary      Update partial a patient
// @Description  update partial a patient
// @Tags         Patients
// @Accept       json
// @Produce      json
// @Param        token    header   string          true  "token"
// @Param        id       path     int             true "Patient id"
// @Param        patient  body     domain.RequestPatient  true  "Patient to store"
// @Success      200  {object}  domain.Patient
// @Failure      400  {object}  web.ErrorApi
// @Failure      401  {object}  web.ErrorApi
// @Failure      500  {object}  web.ErrorApi
// @Router       /patients/{id} [patch]
func (h *PatientHandler) UpdatePartial(ctx *gin.Context) {
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

// NewPatient godoc
// @Summary      Create a patient
// @Description  create a patient
// @Tags         Patients
// @Accept		 json
// @Produce      json
// @Param        token    header   string          true "token"
// @Param        patient  body     domain.Patient true    "Patient to store"
// @Success      201  {object}  domain.Patient
// @Failure      400  {object}  web.ErrorApi
// @Failure      401  {object}  web.ErrorApi
// @Failure      500  {object}  web.ErrorApi
// @Router       /patients [post]
func (h *PatientHandler) NewPatient(ctx *gin.Context) {
	var patient *domain.Patient

	err := ctx.ShouldBindJSON(&patient)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("invalid body"))
		return
	}
	valid, err := validEmptysPatient(patient)
	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	valid, err = validateRegistrationDate(patient.RegistrationDate)
	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
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

// Delete godoc
// @Summary      Delete a Patient
// @Description  delete a patient
// @Tags         Patients
// @Param        token    header   string          true "token"
// @Param        id    path     int    true "Patient id"
// @Success      200  string    patient removed successfully
// @Failure      400  {object}  web.ErrorApi
// @Failure      401  {object}  web.ErrorApi
// @Failure      500  {object}  web.ErrorApi
// @Router       /patients/{id} [delete]
func (h *PatientHandler) DeletePatient(ctx *gin.Context) {
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
