package server

import (
	"awesomeProject/service/gen"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"os"
	"time"
)

// метод DownloadFile для скачивания выбранного файла
func (s *GRPCServer) DownloadFile(ctx context.Context, req *gen.DownloadFileRequest) (*gen.DownloadFileResponse, error) {
	// Получаем информацию о файле
	fileInfo := &gen.FileInfo{}

	file, err := os.Open("savedFiles\\" + req.Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to open file: %v", err)
	}
	defer file.Close()
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	// Получаем содержимое файла
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to read file content: %v", err)
	}
	// Создаем ответ с файлом
	fileInfo.Name = req.Name
	fileInfo.CreatedAt = timestamppb.New(time.Now())
	fileInfo.UpdatedAt = timestamppb.New(time.Now())
	fileProto := &gen.File{
		Info:    fileInfo,
		Content: content,
	}
	// Возвращаем файл
	return &gen.DownloadFileResponse{File: fileProto}, nil
}
