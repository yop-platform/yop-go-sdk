package main

import (
	"fmt"
	"log"

	"github.com/yop-platform/yop-go-sdk/yop/client"
	"github.com/yop-platform/yop-go-sdk/yop/constants"
	"github.com/yop-platform/yop-go-sdk/yop/request"
)

func main() {
	// Example: Basic GET request
	fmt.Println("=== YOP Go SDK Basic Usage Example ===")

	// Configure your credentials
	appId := "your-app-id"
	privateKey := "your-private-key"

	// Example 1: GET Request
	fmt.Println("\n1. GET Request Example:")
	getExample(appId, privateKey)

	// Example 2: POST JSON Request
	fmt.Println("\n2. POST JSON Request Example:")
	postJsonExample(appId, privateKey)

	// Example 3: POST Form Request
	fmt.Println("\n3. POST Form Request Example:")
	postFormExample(appId, privateKey)
}

func getExample(appId, privateKey string) {
	// Create private key object
	priKey := &request.IsvPriKey{
		Value:    privateKey,
		CertType: request.RSA2048,
	}

	// Create GET request
	yopRequest := request.NewYopRequest(constants.GET_HTTP_METHOD, "/rest/v1.0/test/query")
	yopRequest.AppId = appId
	yopRequest.IsvPriKey = *priKey

	// Add query parameters
	yopRequest.AddParam("merchantId", "123456789")
	yopRequest.AddParam("requestId", "req_001")

	// Send request
	yopResp, err := client.DefaultClient.Request(yopRequest)
	if err != nil {
		log.Printf("GET request failed: %v", err)
		return
	}

	fmt.Printf("GET Response: %+v\n", yopResp.Result)
}

func postJsonExample(appId, privateKey string) {
	// Create private key object
	priKey := &request.IsvPriKey{
		Value:    privateKey,
		CertType: request.RSA2048,
	}

	// Create POST request
	yopRequest := request.NewYopRequest(constants.POST_HTTP_METHOD, "/rest/v1.0/test/create")
	yopRequest.AppId = appId
	yopRequest.IsvPriKey = *priKey

	// Set JSON content
	jsonContent := `{
		"merchantId": "123456789",
		"requestId": "req_002",
		"amount": "100.00",
		"currency": "CNY"
	}`
	yopRequest.Content = jsonContent

	// Send request
	yopResp, err := client.DefaultClient.Request(yopRequest)
	if err != nil {
		log.Printf("POST JSON request failed: %v", err)
		return
	}

	fmt.Printf("POST JSON Response: %+v\n", yopResp.Result)
}

func postFormExample(appId, privateKey string) {
	// Create private key object
	priKey := &request.IsvPriKey{
		Value:    privateKey,
		CertType: request.RSA2048,
	}

	// Create POST request
	yopRequest := request.NewYopRequest(constants.POST_HTTP_METHOD, "/rest/v1.0/test/form")
	yopRequest.AppId = appId
	yopRequest.IsvPriKey = *priKey

	// Add form parameters
	yopRequest.AddParam("merchantId", "123456789")
	yopRequest.AddParam("requestId", "req_003")
	yopRequest.AddParam("amount", "200.00")
	yopRequest.AddParam("currency", "CNY")

	// Send request
	yopResp, err := client.DefaultClient.Request(yopRequest)
	if err != nil {
		log.Printf("POST Form request failed: %v", err)
		return
	}

	fmt.Printf("POST Form Response: %+v\n", yopResp.Result)
}
