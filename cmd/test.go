package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}

	res, err := client.Get("http://localhost:9105/fetchAppsSelect")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("StatusCode:", res.StatusCode)
	defer res.Body.Close()
}
