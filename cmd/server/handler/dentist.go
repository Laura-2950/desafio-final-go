package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Laura-2950/desafio-final-go/internal/dentist"
	"github.com/Laura-2950/desafio-final-go/internal/domain"
	"github.com/Laura-2950/desafio-final-go/pkg/web"
	"github.com/gin-gonic/gin"
)

type DentistHandler struct {
	DentistService dentist.IService
}

func (h *DentistHandler) NewDentist(ctx *gin.Context) {
	var dentist *domain.Dentist

	err := ctx.ShouldBindJSON(&dentist)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid dentist"})
		return
	}
	valid, err := validateEmptys(dentist)
	if !valid {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newDentist, err := h.DentistService.CreateNewDentist(dentist)
	if err != nil {
		if errApi, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, newDentist)
}

func (h *DentistHandler) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Invalid ID"))
		return
	}

	dentistFounded, err := h.DentistService.GetDentistByID(id)
	if err != nil {
		if errApi, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, &dentistFounded)
}

func (h *DentistHandler) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Invalid ID"))
		return
	}

	erro := h.DentistService.DeleteDentist(id)
	if erro != nil {
		if errApi, ok := erro.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, erro)
		return
	}

	ctx.JSON(http.StatusOK, "dentist removed successfully")
}

func (h *DentistHandler) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Invalid ID"))
		return
	}

	_, err = h.DentistService.GetDentistByID(id)
	if err != nil {
		if errApi, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	var dent *domain.Dentist
	err = ctx.ShouldBindJSON(&dent)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid body"})
		return
	}

	valid, err := validateEmptysUpdate(dent)
	if !valid {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	dentist, err := h.DentistService.UpdateDentist(id, dent)
	if err != nil {
		if errApi, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, dentist)
}

func (h *DentistHandler) UpdatePartial(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Invalid ID"))
		return
	}

	_, err = h.DentistService.GetDentistByID(id)
	if err != nil {
		if errApi, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	var request *domain.RequestDentist
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid json"})
		return
	}
	dentistUpdate := domain.Dentist{
		Name:               request.Name,
		LastName:           request.LastName,
		RegistrationNumber: request.RegistrationNumber,
	}

	dentist, err := h.DentistService.UpdateDentist(id, &dentistUpdate)
	if err != nil {
		if errApi, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, dentist)
}

// validateEmptys valida que los campos no esten vacios
func validateEmptys(dentist *domain.Dentist) (bool, error) {
	if dentist.Name == "" || dentist.LastName == "" || dentist.RegistrationNumber == "" {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

func validateEmptysUpdate(dentist *domain.Dentist) (bool, error) {
	if dentist.Name == "" || dentist.LastName == "" {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}
