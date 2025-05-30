package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/devkhatri523/ecom-go/config/config"
	"github.com/devkhatri523/ecom-go/go-service/http"
	"github.com/devkhatri523/order-service/data/request"
	response2 "github.com/devkhatri523/order-service/data/response"
	"log"
)

var purchaseServiceUrl = config.Default().GetString("product.service.url")

func GetPurchaseProducts(orderRequest request.OrderRequest) ([]response2.PurchaseResponse, error) {
	jsonRequest, err := json.Marshal(orderRequest.Products)
	if err != nil {
		log.Fatal(err)
	}
	body := bytes.NewReader(jsonRequest)
	purchaseServiceUrl = fmt.Sprintf("%s/purchase", purchaseServiceUrl)
	req, err := http.GetHttpPostRequest(purchaseServiceUrl, body)
	if err != nil {
		log.Fatal(err)
	}
	response, err := http.ExecuteHttpRequest(req)
	if err != nil {
		log.Fatal(err)
	}
	var purchaseProduct []response2.PurchaseResponse
	err = http.ReadHttpBodyAsJson(&purchaseProduct, response)

	if err != nil {
		log.Fatal(err)
	}
	return purchaseProduct, nil
}
