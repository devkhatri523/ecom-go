package client

import (
	"fmt"
	"github.com/devkhatri523/ecom-go/config/config"
	"github.com/devkhatri523/ecom-go/go-service/http"
	"github.com/devkhatri523/order-service/data/response"
	"log"
)

var customerServiceUrl = config.Default().GetString("customer.service.url")

func GetCustomerDetails(customerId string) (response.CustomerResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", customerServiceUrl, customerId)
	req, err := http.GetHttpGetRequest(endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := http.ExecuteHttpRequest(req)
	if err != nil {
		log.Fatal(err)
	}
	var customer response.CustomerResponse
	_ = http.ReadHttpBodyAsJson(&customer, res)
	fmt.Println(customer)
	if err != nil {
		log.Fatal(err)
	}
	return customer, nil
}
