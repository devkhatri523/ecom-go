package main

import (
	"fmt"
	"github.com/devkhatri523/ecom-go/config/config"
	"github.com/devkhatri523/ecom-go/config/database"
	"github.com/devkhatri523/ecom-go/go-service/event"
	"github.com/devkhatri523/ecom-go/go-service/kafka"
	"github.com/devkhatri523/ecom-go/go-utils/utils"
	"github.com/devkhatri523/order-service/controller"
	kafka2 "github.com/devkhatri523/order-service/kafka"
	"github.com/devkhatri523/order-service/repository"
	"github.com/devkhatri523/order-service/router"
	"github.com/devkhatri523/order-service/service"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"time"
)

func main() {
	db, err := OpenDb()
	if err != nil {
		fmt.Sprintf("Error while connecting database %s", err)
	}
	validate := validator.New()
	orderRepository := repository.NewOrderRepositoryImpl(db.OrmInstance)
	orderLineRepository := repository.NewOrderLineRepositoryImpl(db.OrmInstance)
	kafkaProducer := openKafkaProducer()
	publisher := event.NewKafkaPublisher(kafkaProducer)
	orderConfirmationPublisher := kafka2.NewOrderConfirmationPublisher(publisher)
	orderService, err := service.NewOrderServiceImpl(orderRepository, validate, orderLineRepository, orderConfirmationPublisher)
	if err != nil {
		// Handle error appropriately, such as logging and exiting
		log.Fatalf("Error initializing order service: %v", err)
	}

	orderLineService, err := service.NewOrderLineServiceImpl(orderLineRepository, validate)
	if err != nil {
		// Handle error appropriately, such as logging and exiting
		log.Fatalf("Error initializing order line service: %v", err)
	}

	orderController := controller.NewOrderController(orderService)
	orderLineController := controller.NewOrderLineController(orderLineService)
	routes := router.OrderRouter(orderController, orderLineController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = server.ListenAndServe()
}

func OpenDb() (*database.OrmDB, error) {
	host := config.Default().GetString("db.postgres.host")
	user := config.Default().GetString("db.postgres.username")
	password := config.Default().GetString("db.postgres.password")
	dbName := config.Default().GetString("db.postgres.database")
	port := config.Default().GetInt("db.postgres.port")
	orm, err := database.OpenORM(host, port, user, password, dbName)
	if err != nil {
		return nil, err
	}
	return orm, nil
}

func openKafkaProducer() kafka.ProducerClient {
	k := kafka.ProducerClient{
		ProducerConnectionDetail: kafka.ProducerConnectionDetail{
			ConnectionDetail: getKafkaConnectionDetail(),
		},
	}
	k.Start()
	return k
}
func getKafkaConnectionDetail() kafka.ConnectionDetail {
	return kafka.ConnectionDetail{
		Brokers:  config.Default().GetStringSlice("kafka.brokers"),
		UseSSL:   config.Default().GetBool("kafka.useSSL"),
		ClientId: utils.GenerateUUID(),
	}
}
