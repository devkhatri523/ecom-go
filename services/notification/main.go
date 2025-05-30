package main

import (
	"github.com/devkhatri523/ecom-go/config/config"
	"github.com/devkhatri523/ecom-go/go-service/event"
	"github.com/devkhatri523/ecom-go/go-service/kafka"
	"github.com/devkhatri523/ecom-go/go-service/logger"
	"github.com/devkhatri523/ecom-go/go-utils/utils"
	service2 "v01/email_service"
	"v01/handler"
)

func main() {
	orderEventConsumer := getEventConsumer(config.Default().GetString("event.order.created.group"))
	orderNotificationService := service2.NewOrderConfirmationEmailSender()
	orderEvent := startOrderEventLooper(orderEventConsumer, orderNotificationService)
	paymentEventConsumer := getEventConsumer(config.Default().GetString("event.payment.created.group"))
	paymentNotificationService := service2.NewPaymentConfirmationEmailSender()
	paymentEvent := startPaymentEventLooper(paymentEventConsumer, paymentNotificationService)
	defer paymentEvent.Stop()
	defer orderEvent.Stop()
}

func getEventConsumer(groupId string) kafka.ConsumerClient {
	consumerClient := kafka.ConsumerClient{
		ConsumerConnectionDetail: kafka.ConsumerConnectionDetail{
			ConnectionDetail: getKafkaConnectionDetail(),
			GroupName:        groupId,
			AutoCommit:       false,
		},
	}
	consumerClient.Start()
	return consumerClient
}

func getKafkaConnectionDetail() kafka.ConnectionDetail {
	return kafka.ConnectionDetail{
		Brokers:  config.Default().GetStringSlice("kafka.brokers"),
		UseSSL:   config.Default().GetBool("kafka.useSSL"),
		ClientId: utils.GenerateUUID(),
	}
}

func startOrderEventLooper(orderEventConsumer kafka.ConsumerClient, service service2.EmailService) event.KafkaLooper {
	orderTopic := config.Default().GetString("event.order.created.topic")
	err := orderEventConsumer.Subscribe(orderTopic)
	if err != nil {
		logger.Default().Errorf("Error while subscribing topic : %s error : %s", orderTopic, err)
		panic(err)
	}
	looper := event.GetNewKafkaLooper(orderEventConsumer, 10000, func() bool {
		return true
	})
	go looper.Loop(handler.NewOrderEventHandler(service))
	return looper
}

func startPaymentEventLooper(orderEventConsumer kafka.ConsumerClient, service service2.EmailService) event.KafkaLooper {
	paymentCreatedTopic := config.Default().GetString("event.payment.created.topic")
	err := orderEventConsumer.Subscribe(paymentCreatedTopic)
	if err != nil {
		logger.Default().Errorf("Error while subscribing topic : %s error : %s", paymentCreatedTopic, err)
		panic(err)
	}
	looper := event.GetNewKafkaLooper(orderEventConsumer, 10000, func() bool {
		return true
	})
	go looper.Loop(handler.NewPaymentEventHandler(service))
	return looper
}
