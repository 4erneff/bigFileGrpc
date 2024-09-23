package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	pb "github.com/4erneff/alcatraz/pb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port          = ":50051"
	fileChunkSize = 1024 * 1024 // 1MB per chunk
	filePath      = "large_file.bin"
)

// Server is the gRPC server
type server struct {
	pb.UnimplementedFileServiceServer
}

// GenerateFile creates a large file (1GB) on the server
func GenerateFile() error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data := make([]byte, 1024*1024) // 1MB buffer
	for i := 0; i < 1024; i++ {     // Write 1024MB to make 1GB
		if _, err := file.Write(data); err != nil {
			return err
		}
	}

	return nil
}

// GetFileMetadata returns the total size and total number of chunks
func (s *server) GetFileMetadata(
	ctx context.Context,
	req *pb.FileMetadataRequest,
) (*pb.FileMetadataResponse, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	totalSize := fileInfo.Size()
	totalChunks := int32((totalSize + int64(fileChunkSize) - 1) / int64(fileChunkSize)) // Calculate total number of chunks

	return &pb.FileMetadataResponse{
		TotalSize:   totalSize,
		TotalChunks: totalChunks,
	}, nil
}

// GetFileStream sends the file in chunks to the client
func (s *server) GetFileStream(req *pb.FileRequest, stream pb.FileService_GetFileStreamServer) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get the total size of the file
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	totalSize := fileInfo.Size()

	// Calculate total number of chunks
	totalChunks := int32((totalSize + int64(fileChunkSize) - 1) / int64(fileChunkSize))

	buffer := make([]byte, fileChunkSize)
	sequenceNumber := req.StartChunk

	_, err = file.Seek(int64(sequenceNumber)*int64(fileChunkSize), io.SeekStart)
	if err != nil {
		return err
	}

	for {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		checksum := sha256.Sum256(buffer[:bytesRead])
		chunk := &pb.FileChunk{
			SequenceNumber: sequenceNumber,
			ChunkData:      buffer[:bytesRead],
			Checksum:       fmt.Sprintf("%x", checksum),
			TotalSize:      totalSize,
			TotalChunks:    totalChunks,
		}

		if err := stream.Send(chunk); err != nil {
			return err
		}

		sequenceNumber++
	}

	return nil
}

func main() {
	// Generate the large file on the server
	if err := GenerateFile(); err != nil {
		log.Fatalf("Failed to generate file: %v", err)
	}
	fmt.Println("File generated successfully")

	creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		log.Fatalf("Failed to load TLS keys: %v", err)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterFileServiceServer(s, &server{})

	log.Printf("Server listening on %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
