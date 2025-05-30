package main

import (
	"fmt"
	"github.com/devkhatri523/ecom-go/config/config"
	"github.com/devkhatri523/ecom-go/config/database"
	"github.com/devkhatri523/ecom-go/go-service/event"
	"github.com/devkhatri523/ecom-go/go-service/kafka"
	utils2 "github.com/devkhatri523/ecom-go/go-utils/utils"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"time"
	"v01/controller"
	pymentKafka "v01/kafka"
	"v01/repository"
	"v01/router"
	"v01/service"
	"v01/utils"
)

func main() {
	//Database
	db, err := OpenDb()
	if err != nil {
		fmt.Sprintf("Error while connecting database %s", err)
	}
	fmt.Println(db)
	validate := validator.New()
	paymentRepository := repository.NewPaymentRepositoryImpl(db.OrmInstance)
	kafkaProducer := openKafkaProducer()
	publisher := event.NewKafkaPublisher(kafkaProducer)
	orderConfirmationPublisher := pymentKafka.NewPaymentConfirmationPublisher(publisher)
	paymentService, err := service.NewPaymentServiceImpl(paymentRepository, validate)
	if err != nil {
		// Handle error appropriately, such as logging and exiting
		log.Fatalf("Error initializing payment service: %v", err)
	}
	paymentController := controller.NewProductController(paymentService)
	routes := router.PaymentRouter(paymentController)

	server := &http.Server{
		Addr:           ":8030",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = server.ListenAndServe()
	utils.ErrorPanic(err)

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
		ClientId: utils2.GenerateUUID(),
	}
}
