package server

import (
	"awesomeProject/service/gen"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"path/filepath"
)

//	}
//	boo, err := gen.SaveFileRequest{}
//	file_name := req.File
//	return
//}

func (s *GRPCServer) SaveFile(ctx context.Context, req *gen.SaveFileRequest) (*gen.SaveFileResponse, error) {
	var newFilePath string
	if files, _ := os.ReadDir(server_path); files != nil {
		err := os.Mkdir(server_path, 0750)
		if err != nil && !os.IsExist(err) {
			log.Fatal(err)
		}
	}
	// получаем информацию о файле
	fileInfo := req.File.Info
	// получаем содержимое файла
	content := req.File.Content
	// создаемся временный файл и помещаем в него содержимое
	tmpfile, err := os.CreateTemp("", "*"+filepath.Ext(fileInfo.Name))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to create tmp file: %v", err)
	}

	// закрываем файл и удаляем его после использования
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			log.Fatal(err)
		}
	}(tmpfile.Name())

	_, err = tmpfile.Write(content)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to write to tmp file: %v", err)
	}
	// перемещаем созданный временный файл в искомую директорию с указанным именем
	if err := os.Rename(tmpfile.Name(), fileInfo.Name); err != nil {
		return nil, status.Errorf(codes.Internal, "unable to move file to directory: %v", err)
	}
	// возвращаем ответ метода
	return &gen.SaveFileResponse{}, nil
}
