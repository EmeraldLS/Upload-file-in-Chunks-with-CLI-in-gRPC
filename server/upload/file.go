package upload

import (
	"io"
	"log"

	pb "github.com/EmeraldLS/file_upload/proto"
)

type FileUpload struct {
	pb.FileUploadServer
}

func (f *FileUpload) Upload(stream pb.FileUpload_UploadServer) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Status{UploadStatus: "File stream recieved"})
		}
		if err != nil {
			log.Printf("An error occured recieving file chunks :: %v", err)
			return err
		}

		log.Printf("Recieved chunks %vkb", len(req.GetChunks()))
	}

}
