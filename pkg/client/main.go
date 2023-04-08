package client

import (
	"awesomeProject/service/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	x :=
	client := gen.NewFileServiceClient(conn)
	file, err := client.SaveFile()
	if err != nil {
		log.Fatal(err)
	}
}
