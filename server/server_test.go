package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pb "github.com/4erneff/alcatraz/pb/proto"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

// Initialize in-memory gRPC server using bufconn for testing.
func init() {
	lis = bufconn.Listen(bufSize)

	s := grpc.NewServer()

	// Register the service
	pb.RegisterFileServiceServer(s, &server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			fmt.Printf("Server exited with error: %v", err)
		}
	}()
}

// Helper function to dial the in-memory server
func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGenerateFile(t *testing.T) {
	// Clean up after the test
	defer os.Remove(filePath)

	err := GenerateFile()
	assert.NoError(t, err, "File should be generated without error")

	// Check file existence and size
	fileInfo, err := os.Stat(filePath)
	assert.NoError(t, err, "File should exist")
	assert.Equal(t, int64(1024*1024*1024), fileInfo.Size(), "File should be 1GB")
}

func TestGetFileMetadata(t *testing.T) {
	// Ensure the file exists for the test
	err := GenerateFile()
	assert.NoError(t, err)
	defer os.Remove(filePath)

	conn, err := grpc.DialContext(context.Background(), "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	assert.NoError(t, err, "Dial should succeed")
	defer conn.Close()

	client := pb.NewFileServiceClient(conn)

	// Call GetFileMetadata
	resp, err := client.GetFileMetadata(context.Background(), &pb.FileMetadataRequest{})
	assert.NoError(t, err, "Metadata retrieval should succeed")
	assert.Equal(t, int64(1024*1024*1024), resp.TotalSize, "Total size should be 1GB")
	assert.Equal(t, int32(1024), resp.TotalChunks, "Total chunks should be 1024")
}

func TestGetFileStream(t *testing.T) {
	// Ensure the file exists for the test
	err := GenerateFile()
	assert.NoError(t, err)
	defer os.Remove(filePath)

	conn, err := grpc.DialContext(context.Background(), "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	assert.NoError(t, err, "Dial should succeed")
	defer conn.Close()

	client := pb.NewFileServiceClient(conn)

	// Request the file stream starting at chunk 0
	stream, err := client.GetFileStream(context.Background(), &pb.FileRequest{StartChunk: 0})
	assert.NoError(t, err, "File stream request should succeed")

	chunk, err := stream.Recv()
	assert.NoError(t, err, "Receiving first chunk should succeed")
	assert.Equal(t, int32(0), chunk.SequenceNumber, "First chunk sequence number should be 0")
	assert.Equal(t, fileChunkSize, len(chunk.ChunkData), "First chunk size should be 1MB")

	// Verify checksum of the chunk
	checksum := sha256.Sum256(chunk.ChunkData)
	assert.Equal(t, fmt.Sprintf("%x", checksum), chunk.Checksum, "Checksum should match")

	totalChunksReceived := 1
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			break
		}
		assert.NoError(t, err, "File stream should not error before EOF")
		totalChunksReceived++
	}

	assert.Equal(t, 1024, totalChunksReceived, "Should receive 1024 chunks in total")
}
