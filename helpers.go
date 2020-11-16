package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func getHTTPStatusCode(service ServiceConfig) int {
	switch method := service.Method; method {
	case "GET":
		response, err := http.Get(service.URL)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}

		return int(response.StatusCode)

	case "POST":
		payloadByte, _ := json.Marshal(service.Payload)
		response, err := http.Post(service.URL, "application/json", bytes.NewReader(payloadByte))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}

		return int(response.StatusCode)

	default:
		panic("Unknown method")
	}
}

func logStatus(service ServiceConfig, statusCode int) map[string]interface{} {
	return map[string]interface{}{"service": service.URL, "statusCode": isHealthy(statusCode)}
}

func isHealthy(statusCode int) bool {
	if statusCode == 200 {
		return true
	}

	return false
}
