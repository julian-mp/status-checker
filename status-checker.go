package main

import "fmt"

func main() {
	services := []string{"https://google.com", "https://amazon.com", "https://netflix.com"}
	for _, service := range services {
		c := make(chan map[string]interface{})
		go checkServiceStatus(service, c)
		fmt.Println(<-c)
	}
}

func checkServiceStatus(service string, c chan map[string]interface{}) {
	statusCode := getHTTPStatusCode(service)
	c <- logStatus(service, statusCode)
}
