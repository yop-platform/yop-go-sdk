# YOP Go SDK Examples

This directory contains example code demonstrating how to use the YOP Go SDK.

## Examples

### 1. Basic Usage (`basic_usage.go`)

Demonstrates the fundamental operations:
- GET requests with query parameters
- POST requests with JSON payload
- POST requests with form data

**Run the example:**
```bash
cd examples
go run basic_usage.go
```

### 2. Advanced Usage (`advanced_usage.go`)

Shows advanced features and best practices:
- Custom HTTP client configuration
- File upload operations
- Callback decryption
- Signature generation and verification
- Error handling patterns
- Retry mechanisms

**Run the example:**
```bash
cd examples
go run advanced_usage.go
```

## Configuration

Before running the examples, you need to:

1. **Get your credentials** from the YOP platform:
   - App ID
   - Private Key (RSA2048)

2. **Update the example code** with your actual credentials:
   ```go
   appId := "your-actual-app-id"
   privateKey := "your-actual-private-key"
   ```

3. **Set up test environment** (optional):
   ```go
   yopRequest.ServerRoot = "https://ycetest.yeepay.com:30228/yop-center"
   ```

## Common Patterns

### Error Handling

```go
yopResp, err := client.DefaultClient.Request(yopRequest)
if err != nil {
    // Handle network or request building errors
    log.Printf("Request failed: %v", err)
    return
}

// Check business errors
if yopResp.Result != nil {
    result := yopResp.Result.(map[string]interface{})
    if status, ok := result["status"]; ok && status != "SUCCESS" {
        log.Printf("Business error: %v", result["errorMsg"])
        return
    }
}
```

### Custom Timeout

```go
yopRequest := request.NewYopRequest(constants.GET_HTTP_METHOD, "/api/path")
yopRequest.Timeout = 30 * time.Second // Custom timeout
```

### Adding Headers

```go
yopRequest.Headers["Custom-Header"] = "custom-value"
```

### Environment Configuration

```go
// Production (default)
yopRequest.ServerRoot = "https://openapi.yeepay.com/yop-center"

// Test environment
yopRequest.ServerRoot = "https://ycetest.yeepay.com:30228/yop-center"

// YOS file service
yopRequest.ServerRoot = "https://yos.yeepay.com/yop-center"
```

## Testing

To test the examples without making actual API calls, you can:

1. **Mock the client:**
   ```go
   // Create a mock client for testing
   mockClient := &MockYopClient{}
   yopResp, err := mockClient.Request(yopRequest)
   ```

2. **Use test endpoints** (if available in your environment)

3. **Validate request construction** without sending:
   ```go
   // Just build and validate the request
   err := auth.RsaSigner{}.SignRequest(*yopRequest)
   if err != nil {
       log.Printf("Request signing failed: %v", err)
   }
   ```

## Best Practices

1. **Always handle errors** appropriately
2. **Use custom timeouts** for different operations
3. **Implement retry logic** for transient failures
4. **Validate responses** before processing
5. **Keep credentials secure** - never hardcode in production
6. **Use connection pooling** for high-throughput applications
7. **Log requests and responses** for debugging (but mask sensitive data)

## Security Notes

- Never commit real credentials to version control
- Use environment variables or secure configuration management
- Validate all input data before sending requests
- Implement proper error handling to avoid information leakage
- Use HTTPS in production environments

## Support

If you encounter issues with these examples:

1. Check the [main documentation](../README.md)
2. Review the [API reference](https://pkg.go.dev/github.com/yop-platform/yop-go-sdk)
3. Open an [issue](https://github.com/yop-platform/yop-go-sdk/issues) on GitHub
