package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Laura-2950/desafio-final-go/cmd/server/handler"
	"github.com/Laura-2950/desafio-final-go/internal/dentist"
	"github.com/Laura-2950/desafio-final-go/internal/patient"
	"github.com/Laura-2950/desafio-final-go/internal/shift"
	"github.com/Laura-2950/desafio-final-go/pkg/middleware"
	"github.com/Laura-2950/desafio-final-go/pkg/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


// @title           API Dental Clinic
// @version         1.0
// @description     This is a API to register a dental shift.

// @contact.name   API Support
// @contact.url    http://www.dentalclinic.com
// @contact.email  support@dentalclinic.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080

// @externalDocs.url          http://localhost:8080/swagger/index.html
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error al intentar cargar archivo .env")
	}
	username := os.Getenv("USER_MYSQL")
	password := os.Getenv("PASS_MYSQL")
	dbName := os.Getenv("DB_MYSQL")

	connectionString := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", username, password, dbName)

	
	db, err := sql.Open("mysql", connectionString)
	//db, err := sql.Open("mysql", "user:root@tcp(localhost:3306)/dental_clinic")
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
	patientHandler := handler.PatientHandler{PatientService: &servPatient}

	repoShift := shift.Repository{Storage: &storage}
	servShift := shift.Service{Repository: &repoShift}
	shiftHandler := handler.ShiftHandler{ShiftService: &servShift}

	r := gin.Default()

	r.GET("ping", func(ctx *gin.Context) { ctx.String(http.StatusOK, "pong") })
	dentistGroup := r.Group("/dentists")
	{
		dentistGroup.POST("", middleware.Authentication(), dentistHandler.NewDentist)
		dentistGroup.GET(":id", dentistHandler.GetById)
		dentistGroup.DELETE(":id", middleware.Authentication(), dentistHandler.Delete)
		dentistGroup.PUT(":id", middleware.Authentication(), dentistHandler.Update)
		dentistGroup.PATCH(":id", middleware.Authentication(), dentistHandler.UpdatePartial)
	}
	patientGroup := r.Group("/patients")
	{
		patientGroup.GET(":id", patientHandler.GetById)
		patientGroup.PUT(":id", middleware.Authentication(), patientHandler.Update)
		patientGroup.PATCH(":id", middleware.Authentication(), patientHandler.UpdatePartial)
		patientGroup.POST("", middleware.Authentication(), patientHandler.NewPatient)
		patientGroup.DELETE(":id", middleware.Authentication(), patientHandler.DeletePatient)
	}
	shiftGroup := r.Group("/shifts")
	{
		shiftGroup.POST("", middleware.Authentication(), shiftHandler.NewShift)
		shiftGroup.POST("/code", middleware.Authentication(), shiftHandler.NewShiftCode)
		shiftGroup.GET(":id", shiftHandler.GetById)
		shiftGroup.DELETE(":id", middleware.Authentication(), shiftHandler.Delete)
		shiftGroup.PUT(":id", middleware.Authentication(), shiftHandler.UpdateShift)
		shiftGroup.PATCH(":id", middleware.Authentication(), shiftHandler.UpdatePartialShift)
		shiftGroup.GET("", shiftHandler.GetAllByDni)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
