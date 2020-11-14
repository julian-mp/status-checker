package main

import "fmt"

func main() {
	services := []string{"https://google.com", "https://amazon.com", "https://netflix.com"}
	for _, service := range services {
		serviceStatus := checkServiceStatus(service)
		fmt.Println(serviceStatus)
	}
}

func checkServiceStatus(service string) map[string]interface{} {
	statusCode := getHTTPStatusCode(service)
	return logStatus(service, statusCode)
}
