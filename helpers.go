package main

import (
	"fmt"
	"net/http"
)

func getHTTPStatusCode(url string) int {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	return int(response.StatusCode)
}

func logStatus(service string, statusCode int) map[string]interface{} {
	return map[string]interface{}{"service": service, "statusCode": isHealthy(statusCode)}
}

func isHealthy(statusCode int) bool {
	if statusCode == 200 {
		return true
	}

	return false
}
