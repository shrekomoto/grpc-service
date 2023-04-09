package server

import (
	"awesomeProject/service/gen"
	"context"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"log"
	"net"
)

const server_path = "C://Tages"

type GRPCServer struct {
	gen.FileServiceServer
}

func Run() {
	unaryLimiter := rate.NewLimiter(100, 100)
	streamLimiter := rate.NewLimiter(10, 10)
	server := grpc.NewServer(grpc.StreamInterceptor(func(srv interface{}, stream grpc.ServerStream,
		info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

		if streamLimiter.Allow() {
			return handler(srv, stream)
		}
		log.Print("Too many connections, wait a bit")
		return nil
	}),

		grpc.UnaryInterceptor(func(ctx context.Context, req interface{},
			info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

			if unaryLimiter.Allow() {
				return handler(ctx, req)
			}
			log.Print("Too many connections, wait a bit")
			return nil, nil
		}),
	)

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
