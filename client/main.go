package main

import (
	"context"
	"log"

	"github.com/EmeraldLS/file_upload/client/upload"
	pb "github.com/EmeraldLS/file_upload/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":3031", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("An error occured while trying to connect with the gRPC server :: %v", err)
	}
	client := pb.NewFileUploadClient(conn)

	if err := upload.UploadFileInChunks(client, context.TODO()); err != nil {
		log.Printf("%v", err)
	}

}
