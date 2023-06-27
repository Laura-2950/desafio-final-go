package main

import (
	"database/sql"
	"net/http"

	"github.com/Laura-2950/desafio-final-go/cmd/server/handler"
	"github.com/Laura-2950/desafio-final-go/internal/dentist"
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
	r := gin.Default()

	r.GET("ping", func(ctx *gin.Context) { ctx.String(http.StatusOK, "pong") })
	dentistGroup := r.Group("/dentist")
	{
		dentistGroup.POST("", dentistHandler.NewDentist)
		dentistGroup.GET(":id", dentistHandler.GetById)
		dentistGroup.DELETE(":id", dentistHandler.Delete)
	}

	r.Run(":8080")
}
