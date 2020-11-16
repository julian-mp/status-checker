package main

import (
	"fmt"
	"sync"
)

func main() {
	config := getConfig("./config.yaml")

	wg := sync.WaitGroup{}
	for i := range config.Services {
		wg.Add(1)
		go func(service ServiceConfig) {
			defer wg.Done()
			checkServiceStatus(service)
		}(config.Services[i])
	}
	wg.Wait()
}

func checkServiceStatus(service ServiceConfig) {
	statusCode := getHTTPStatusCode(service)
	fmt.Println(logStatus(service, statusCode))
}
