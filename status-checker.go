package main

import (
	"fmt"
	"net/http"
)

func main() {
	services := []string{"https://serviceone.com", "https://servicetwo.com"}
	for _, service := range services {
		channel := make(chan int)
		go getHTTPStatusCode(service, channel)
		go logStatus(service, channel)
	}
}

func getHTTPStatusCode(url string, channel chan int) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		channel <- int(response.StatusCode)
	}
}

func logStatus(service string, channel chan int) {
	status := map[string]interface{}{"service": service, "statusCode": isHealthy(<-channel)}
	channel <- status
	fmt.Printf("This will write to the database next %s\n", status)
}

func isHealthy(statusCode int) bool {
	fmt.Printf("isHealthy %d", statusCode)
	if statusCode == 200 {
		return true
	}

	return false
}
