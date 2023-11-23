package main

import (
	"log"
	"net"

	pb "github.com/EmeraldLS/file_upload/proto"
	"github.com/EmeraldLS/file_upload/server/upload"
	"google.golang.org/grpc"
)

func main() {
	list, err := net.Listen("tcp", ":3031")
	if err != nil {
		log.Fatalf("An error occured while opening the listener :: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterFileUploadServer(server, &upload.FileUpload{})

	log.Printf("Server started @ :: %v", list.Addr())

	if err := server.Serve(list); err != nil {
		log.Printf("An error occured starting the gRPC server :: %v", err)
	}
}
