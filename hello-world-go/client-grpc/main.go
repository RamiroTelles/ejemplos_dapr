package main

import (
	"context"
	in "hello-world-client-grpc/invoicer"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	address = "localhost:50007"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("fallo en la conexion: %v", err)
	}

	defer conn.Close()
	c := in.NewInvoicerClient(conn)

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

		ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "server")
		r, err := c.Create(ctx, &in.CreateRequest{Message: "Hola Aux elian"})
		if err != nil {
			log.Printf("error al enviar mensaje: %v", err)
			time.Sleep(20 * time.Second)
			continue
		}
		cancel()
		if r.GetSuccess() {
			log.Println("Mensaje Enviado con exito")
		} else {
			log.Println("Error al enviar el mensaje")
		}

		time.Sleep(20 * time.Second)
	}
}
