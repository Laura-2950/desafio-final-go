package main

import (
	"database/sql"
	"net/http"

	"github.com/Laura-2950/desafio-final-go/cmd/server/handler"
	"github.com/Laura-2950/desafio-final-go/internal/dentist"
	"github.com/Laura-2950/desafio-final-go/internal/patient"
	"github.com/Laura-2950/desafio-final-go/pkg/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/dental_clinic")
	if err != nil {
		panic(err.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		panic(errPing.Error())
	}

	storage := store.SqlStore{DB: db}
	repo := dentist.Repository{Storage: &storage}
	serv := dentist.Service{Repository: &repo}
	dentistHandler := handler.DentistHandler{DentistService: &serv}

	repoPatient := patient.Repository{Storage: &storage}
	servPatient := patient.Service{Repository: &repoPatient}
	patientHandler := handler.PatienttHandler{PatientService: &servPatient}

	r := gin.Default()

	r.GET("ping", func(ctx *gin.Context) { ctx.String(http.StatusOK, "pong") })
	dentistGroup := r.Group("/dentists")
	{
		dentistGroup.POST("", dentistHandler.NewDentist)
		dentistGroup.GET(":id", dentistHandler.GetById)
		dentistGroup.DELETE(":id", dentistHandler.Delete)
		dentistGroup.PUT(":id", dentistHandler.Update)
		dentistGroup.PATCH(":id", dentistHandler.UpdatePartial)
	}
	patientGroup := r.Group("/patients")
	{
		patientGroup.GET(":id", patientHandler.GetById)
		patientGroup.PUT(":id", patientHandler.Update)
		patientGroup.PATCH(":id", patientHandler.UpdatePartial)
		patientGroup.POST("", patientHandler.NewPatient)
		patientGroup.DELETE(":id", patientHandler.DeletePatient)
	}

	r.Run(":8080")
}
