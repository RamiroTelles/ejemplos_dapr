package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getMessage(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error recibiendo:", err.Error())
	}
	fmt.Println("Mensaje:", string(data))
	_, err = w.Write(data)
	if err != nil {
		log.Println("Error al escribir respuesta:", err.Error())
	}
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/message", getMessage).Methods("POST")

	err := http.ListenAndServe(":6006", r)
	if !errors.Is(err, http.ErrServerClosed) {
		log.Println("Error al iniciar server")
	}
}
