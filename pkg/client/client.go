package client

import (
	"awesomeProject/service/gen"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"os"
	"time"
)

func main() {
	// устанавливаем соединение с сервером
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("unable to connect to server: %v", err)
	}
	defer conn.Close()
	// создаем клиент для сервиса FileService
	client := gen.NewFileServiceClient(conn)

	// читаем содержимое файла в переменную content
	content, err := os.ReadFile("file.jpg")
	if err != nil {
		log.Fatalf("unable to read file content: %v", err)
	}
	// создаем переменную fileInfo с информацией о файле
	fileInfo := &gen.FileInfo{
		Name:      "file.jpg",
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
	}
	// создаем переменную req с информацией о файле и его содержимым
	req := &gen.SaveFileRequest{
		File: &gen.File{
			Info:    fileInfo,
			Content: content,
		},
	}
	// отправляем запрос на сервер и получаем ответ
	res, err := client.SaveFile(context.Background(), req)
	if err != nil {
		log.Fatalf("unable to save file: %v", err)
	}
	log.Printf("File saved successfully: %v", res)
}
