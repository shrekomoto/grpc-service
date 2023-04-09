package server

import (
	"awesomeProject/service/gen"
	"google.golang.org/grpc"
	"log"
	"net"
)

const server_path = "C://Tages"

type GRPCServer struct {
	gen.FileServiceServer
}

func Run() {
	server := grpc.NewServer()
	srv := &GRPCServer{}
	gen.RegisterFileServiceServer(server, srv)

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
