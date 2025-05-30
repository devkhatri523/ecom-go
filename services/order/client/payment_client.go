package client

import (
	"bytes"
	"encoding/json"
	"github.com/devkhatri523/ecom-go/config/config"
	"github.com/devkhatri523/ecom-go/go-service/http"
	"github.com/devkhatri523/order-service/data/request"
	"log"
)

var paymentServiceUrl = config.Default().GetString("payment.service.url")

func RequestPayment(paymentRequest request.PaymentRequest) error {
	jsonRequest, err := json.Marshal(paymentRequest)
	if err != nil {
		log.Fatal(err)
	}
	body := bytes.NewReader(jsonRequest)
	req, err := http.GetHttpPostRequest(paymentServiceUrl, body)
	if err != nil {
		log.Fatal(err)
	}
	_, err = http.ExecuteHttpRequest(req)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
