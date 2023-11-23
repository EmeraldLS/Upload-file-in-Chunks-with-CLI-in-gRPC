package upload

import (
	"context"
	"fmt"
	"io"
	"os"

	pb "github.com/EmeraldLS/file_upload/proto"
	"github.com/spf13/cobra"
)

func UploadFileInChunks(client pb.FileUploadClient, ctx context.Context) error {
	var allErr error = nil
	stream, err := client.Upload(ctx)
	if err != nil {
		allErr = err
	}

	uploadCmd := cobra.Command{
		Use:   "upload",
		Short: "Upload file to gRPC server",
		Long: `
		Upload file in cuhnks to gRPC server 
		The chunk size is 1024bytes 
		This makes it possible to for poor network latency
		Note: The first argument is the file name
		All files are to be included in the 'file' directory
		`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fileName := args[0]
			var chunkSize = 1024

			if err = RunUpload(stream, fileName, chunkSize); err != nil {
				allErr = err
			}
		},
	}

	if err := uploadCmd.Execute(); err != nil {
		allErr = err
	}

	return allErr
}

func RunUpload(stream pb.FileUpload_UploadClient, fileName string, chunkSize int) error {
	file, err := os.Open("../file/" + fileName)
	if err != nil {
		return fmt.Errorf("error while getting the file: %v", err)
	}
	buf := make([]byte, chunkSize)

	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error while reading file into the buffer: %v", err)
		}

		stream.Send(&pb.FileChunks{Chunks: buf[:n]})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("error closing stream: %v", err)
	}

	fmt.Println("Status: ", res.UploadStatus)

	return nil

}
