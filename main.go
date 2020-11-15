package main

import (
	"fmt"
	// "sync"
)

func main() {
	config := getConfig("./config.yaml")

	for service := range config.Services {
		fmt.Println(service)
	}

	// wg := sync.WaitGroup{}
	// services := []string{"https://amazon.com", "https://google.com", "https://netflix.com"}
	// for i := range services {
	// 	wg.Add(1)
	// 	go func(service string) {
	// 		defer wg.Done()
	// 		checkServiceStatus(service)
	// 	}(services[i])
	// }
	// wg.Wait()
}

func checkServiceStatus(service string) {
	statusCode := getHTTPStatusCode(service)
	fmt.Println(logStatus(service, statusCode))
}
