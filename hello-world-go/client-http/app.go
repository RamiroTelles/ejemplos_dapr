package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	urlDapr := "http://localhost:3500"

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	for i := 1; i <= 1000; i++ {
		order := `{"mensaje": "Hola aux elian"}`
		req, err := http.NewRequest("POST", urlDapr+"/message", strings.NewReader(order))
		if err != nil {
			log.Fatal(err.Error())
		}
		req.Header.Add("dapr-app-id", "server-http")

		response, err := client.Do(req)
		if err != nil {
			log.Fatal(err.Error())
		}

		result, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		response.Body.Close()

		fmt.Println("Mensaje Enviado:", string(result))
		time.Sleep(5 * time.Second)
	}
}
