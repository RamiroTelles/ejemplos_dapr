package main

import (
	"context"
	"hello-world-grpc/invoicer"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(_ context.Context, in *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {

	log.Printf("Mensaje: %s\n", in.GetMessage())

	return &invoicer.CreateResponse{Success: true}, nil
}

func main() {

	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error al levantar servidor: %s", err)
	}

	server := grpc.NewServer()
	service := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(server, service)
	log.Println("Server escuchando en 50051")
	err1 := server.Serve(list)
	if err1 != nil {
		log.Fatalf("error al montar server, %s", err1)
	}

}
