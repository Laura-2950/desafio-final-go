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

// NewDentist godoc
// @Summary      Create a dentist
// @Description  create a dentist
// @Tags         Dentists
// @Accept		 json
// @Produce      json
// @Param        dentist  body     domain.Dentist true    "Dentist to store"
// @Param        token    header   string          true "token"
// @Success      201  {object}  domain.Dentist
// @Failure      400  {object}  web.ErrorApi
// @Failure      401  {object}  web.ErrorApi
// @Failure      500  {object}  web.ErrorApi
// @Router       /dentists [post]
func (h *DentistHandler) NewDentist(ctx *gin.Context) {
	var dentist *domain.Dentist

	err := ctx.ShouldBindJSON(&dentist)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("invalid body"))
		return
	}
	valid, err := validateEmptys(dentist)
	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
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

	ctx.JSON(http.StatusCreated, newDentist)
}

// GetById godoc
// @Summary      Show a dentist
// @Description  get dentist by ID
// @Tags         Dentists
// @Produce      json
// @Param        id   path      int  true  "Dentist ID"
// @Success      200  {object}  domain.Dentist
// @Failure      400  {object}  web.ErrorApi
// @Failure      401  {object}  web.ErrorApi
// @Failure      500  {object}  web.ErrorApi
// @Router       /dentists/{id} [get]
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

// Delete godoc
// @Summary      Delete a dentist
// @Description  delete a dentist
// @Tags         Dentists
// @Param        id    path     int    true "Dentist id"
// @Param        token    header   string          true "token"
// @Success      200  {string}  dentist removed successfully
// @Failure      400  {object}  web.ErrorApi
// @Failure      401  {object}  web.ErrorApi
// @Failure      500  {object}  web.ErrorApi
// @Router       /dentists/{id} [delete]
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

// Update godoc
// @Summary      Update a dentist
// @Description  update a dentist
// @Tags         Dentists
// @Accept       json
// @Produce      json
// @Param        id    path     int    true "Dentist id"
// @Param        token    header   string          true "token"
// @Param        dentist  body     domain.Dentist true    "Dentist to store"
// @Success      200  {object}  domain.Dentist
// @Failure      400  {object}  web.ErrorApi
// @Failure      401  {object}  web.ErrorApi
// @Failure      500  {object}  web.ErrorApi
// @Router       /dentists/{id} [put]
func (h *DentistHandler) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Invalid ID"))
		return
	}

	var dent *domain.Dentist
	err = ctx.ShouldBindJSON(&dent)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("invalid body"))
		return
	}

	valid, err := validateEmptys(dent)
	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
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

// UpdatePartial godoc
// @Summary      Update partial a dentist
// @Description  update partial a dentist
// @Tags         Dentists
// @Accept       json
// @Produce      json
// @Param        id       path     int             true "Dentist id"
// @Param        token    header   string          true  "token"
// @Param        dentist  body     domain.RequestDentist  true  "Dentist to store"
// @Success      200  {object}  domain.Dentist
// @Failure      400  {object}  web.ErrorApi
// @Failure      401  {object}  web.ErrorApi
// @Failure      500  {object}  web.ErrorApi
// @Router       /dentists/{id} [patch]
func (h *DentistHandler) UpdatePartial(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Invalid ID"))
		return
	}

	var request *domain.RequestDentist
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("invalid body"))
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
