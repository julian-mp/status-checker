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
		go func(serviceURL string) {
			defer wg.Done()
			checkServiceStatus(serviceURL)
		}(config.Services[i].URL)
	}
	wg.Wait()
}

func checkServiceStatus(service string) {
	statusCode := getHTTPStatusCode(service)
	fmt.Println(logStatus(service, statusCode))
}
