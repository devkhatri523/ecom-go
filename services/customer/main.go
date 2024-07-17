package main

import (
	"github.com/devkhatri523/ecom-go/config/config"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
	database "v01/config"
	"v01/controller"
	"v01/repository"
	"v01/router"
	"v01/service"
	"v01/utils"
)

func main() {
	client, err := database.Connect(config.Default().GetString("db.mongo.url"))
	if err != nil {
		panic(err)
	}
	db := client.Database(config.Default().GetString("db.mongo.database"))
	validate := validator.New()
	repo := repository.NewCustomerRepositoryImpl(db)
	customerService := service.NewCustomerServiceImpl(repo, validate)
	customerController := controller.NewCustomerController(customerService)
	routes := router.CustomerRouter(customerController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = server.ListenAndServe()
	utils.ErrorPanic(err)

}
