package server

import (
	"awesomeProject/service/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io/fs"
	"log"
	"os"
)

// метод GetFileList для получения списка файлов на сервере
func (s *GRPCServer) GetFileList(empty *emptypb.Empty, stream gen.FileService_GetFileListServer) error {
	// Получаем список файлов в директории
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	infos := make([]fs.FileInfo, 0, len(files))
	for _, entry := range files {
		info, err := entry.Info()
		if err != nil {
			log.Fatal(err)
		}
		infos = append(infos, info)
	}
	if err != nil {
		return status.Errorf(codes.Internal, "unable to read directory: %v", err)
	}
	// Отправляем список файлов через стрим
	arr := make([]string, len(infos))
	for _, file := range infos {
		// изменяем формат даты
		createdAt := timestamppb.New(file.ModTime())
		updatedAt := timestamppb.New(file.ModTime())
		info := &gen.FileInfo{
			Name:      file.Name(),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		arr = append(arr, info.String())
	}
	if err := stream.Send(&gen.GetFileListResponse{List: arr}); err != nil {
		return status.Errorf(codes.Internal, "unable to send file info: %v", err)
	}
	return nil
}
