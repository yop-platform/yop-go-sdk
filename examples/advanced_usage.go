package main

import (
	"crypto"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/yop-platform/yop-go-sdk/yop/client"
	"github.com/yop-platform/yop-go-sdk/yop/constants"
	"github.com/yop-platform/yop-go-sdk/yop/request"
	"github.com/yop-platform/yop-go-sdk/yop/utils"
)

func main() {
	fmt.Println("=== YOP Go SDK Advanced Usage Examples ===")

	// Example 1: Custom HTTP Client
	fmt.Println("\n1. Custom HTTP Client Example:")
	customClientExample()

	// Example 2: File Upload
	fmt.Println("\n2. File Upload Example:")
	fileUploadExample()

	// Example 3: Callback Decryption
	fmt.Println("\n3. Callback Decryption Example:")
	callbackDecryptionExample()

	// Example 4: Signature Operations
	fmt.Println("\n4. Signature Operations Example:")
	signatureExample()

	// Example 5: Error Handling
	fmt.Println("\n5. Error Handling Example:")
	errorHandlingExample()
}

func customClientExample() {
	// Create custom HTTP client with timeout and connection pooling
	customClient := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}

	// Create YOP client with custom HTTP client
	yopClient := client.YopClient{Client: customClient}

	// Create request
	priKey := &request.IsvPriKey{
		Value:    "your-private-key",
		CertType: request.RSA2048,
	}

	yopRequest := request.NewYopRequest(constants.GET_HTTP_METHOD, "/rest/v1.0/test/query")
	yopRequest.AppId = "your-app-id"
	yopRequest.IsvPriKey = *priKey
	yopRequest.Timeout = 15 * time.Second // Custom timeout for this request

	// Send request with custom client
	yopResp, err := yopClient.Request(yopRequest)
	if err != nil {
		log.Printf("Custom client request failed: %v", err)
		return
	}

	fmt.Printf("Custom client response: %+v\n", yopResp.Result)
}

func fileUploadExample() {
	// Open file for upload
	file, err := os.Open("test_file.txt")
	if err != nil {
		log.Printf("Failed to open file: %v", err)
		return
	}
	defer file.Close()

	// Create request
	priKey := &request.IsvPriKey{
		Value:    "your-private-key",
		CertType: request.RSA2048,
	}

	yopRequest := request.NewYopRequest(constants.POST_HTTP_METHOD, "/rest/v1.0/file/upload")
	yopRequest.AppId = "your-app-id"
	yopRequest.IsvPriKey = *priKey

	// Add file
	yopRequest.AddFile("file", file)

	// Add additional parameters
	yopRequest.AddParam("description", "Test file upload")

	// Send request
	yopResp, err := client.DefaultClient.Request(yopRequest)
	if err != nil {
		log.Printf("File upload failed: %v", err)
		return
	}

	fmt.Printf("File upload response: %+v\n", yopResp.Result)
}

func callbackDecryptionExample() {
	// Example callback data (this would come from YOP platform)
	callbackData := "encrypted_callback_data_here"

	// Platform public key (provided by YOP)
	platformPubKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA..."

	// Your private key
	isvPriKey := "your-private-key"

	// Decrypt callback
	decryptedContent, err := utils.DecryptCallback(platformPubKey, isvPriKey, callbackData)
	if err != nil {
		log.Printf("Callback decryption failed: %v", err)
		return
	}

	fmt.Printf("Decrypted callback content: %s\n", decryptedContent)
}

func signatureExample() {
	// Example data to sign
	data := "merchantId=123456789&amount=100.00&currency=CNY"
	privateKey := "your-private-key"

	// Generate signature
	signature, err := utils.RsaSignBase64(data, privateKey, crypto.SHA256)
	if err != nil {
		log.Printf("Signature generation failed: %v", err)
		return
	}

	fmt.Printf("Generated signature: %s\n", signature)

	// Verify signature
	publicKey := "platform-public-key"
	isValid := utils.VerifySign(data, signature, publicKey, crypto.SHA256)
	fmt.Printf("Signature verification result: %t\n", isValid)
}

func errorHandlingExample() {
	// Create request with invalid configuration to demonstrate error handling
	priKey := &request.IsvPriKey{
		Value:    "invalid-private-key",
		CertType: request.RSA2048,
	}

	yopRequest := request.NewYopRequest(constants.POST_HTTP_METHOD, "/rest/v1.0/test/error")
	yopRequest.AppId = "invalid-app-id"
	yopRequest.IsvPriKey = *priKey

	// Send request
	yopResp, err := client.DefaultClient.Request(yopRequest)
	if err != nil {
		// Handle network or request building errors
		log.Printf("Request error: %v", err)
		return
	}

	// Check for business errors in response
	if yopResp.Result != nil {
		if result, ok := yopResp.Result.(map[string]interface{}); ok {
			if status, exists := result["status"]; exists && status != "SUCCESS" {
				// Handle business errors
				errorMsg := result["errorMsg"]
				log.Printf("Business error: %v", errorMsg)
				return
			}
		}
	}

	fmt.Printf("Request successful: %+v\n", yopResp.Result)
}

// Helper function to demonstrate retry mechanism
func requestWithRetry(yopRequest *request.YopRequest, maxRetries int) error {
	var lastErr error

	for i := 0; i <= maxRetries; i++ {
		yopResp, err := client.DefaultClient.Request(yopRequest)
		if err == nil {
			fmt.Printf("Request successful on attempt %d: %+v\n", i+1, yopResp.Result)
			return nil
		}

		lastErr = err
		if i < maxRetries {
			waitTime := time.Duration(i+1) * time.Second
			fmt.Printf("Request failed on attempt %d, retrying in %v...\n", i+1, waitTime)
			time.Sleep(waitTime)
		}
	}

	return fmt.Errorf("request failed after %d retries: %v", maxRetries, lastErr)
}
