package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/4erneff/alcatraz/client/util"
	pb "github.com/4erneff/alcatraz/pb/proto"
)

const (
	serverAddr     = "localhost:50051"
	outputFile     = "downloaded_file_parallel.mov"
	fileChunkSize  = 1024 * 1024 // 1MB
	numDescriptors = 4           // Number of parallel file descriptors
)

// downloadFile starts or resumes the download from the last known chunk
func downloadFile(client pb.FileServiceClient, startChunk int, totalChunks int32, totalSize int64) (int, error) {
	req := &pb.FileRequest{
		StartChunk: int32(startChunk),
	}

	stream, err := client.GetFileStream(context.Background(), req)
	if err != nil {
		fmt.Println("Failed to start file stream: %w", err)
		return 0, err
	}

	files, mutexes, err := util.CreateFileDescriptors(outputFile, numDescriptors)
	if err != nil {
		return 0, err
	}
	defer func() {
		for _, file := range files {
			file.Close()
		}
	}()

	var wg sync.WaitGroup
	var resultErr error
	lastPercent := 0
	downloadedChunks := startChunk // Tracks the last chunk downloaded

	firstChunk := true

	for {
		chunk, err := stream.Recv()
		if err != nil {
			resultErr = err
			break
		}
		downloadedChunks++

		if firstChunk {
			firstChunk = false
			fmt.Printf("Starting file download/resume, total size: %.2f MB, total chunks: %d\n", float64(totalSize)/(1024*1024), totalChunks)
		}

		// Display progress in percentage
		progress := float64(downloadedChunks) / float64(totalChunks) * 100
		if p := int(progress); p > lastPercent && p%10 == 0 {
			fmt.Printf("\rDownloading... %d%% complete", p)
			lastPercent = p
		}

		wg.Add(1)
		go handleChunk(chunk, &wg, files, mutexes)
	}

	wg.Wait()
	return downloadedChunks, resultErr
}

func handleChunk(chunk *pb.FileChunk, wg *sync.WaitGroup, files []*os.File, mutexes []sync.Mutex) {
	defer wg.Done()

	if !util.VerifyChecksum(chunk.ChunkData, chunk.Checksum) {
		log.Fatalf("Checksum mismatch on chunk %d, ignoring chunk\n", chunk.SequenceNumber)
		return
	}

	// Calculate the offset in the file based on the sequence number
	offset := int64(chunk.SequenceNumber) * int64(fileChunkSize)

	fdIndex := int(chunk.SequenceNumber) % numDescriptors
	file := files[fdIndex]

	mutexes[fdIndex].Lock()

	if _, err := file.WriteAt(chunk.ChunkData, offset); err != nil {
		log.Fatalf("Failed to write chunk at offset %d: %v", offset, err)
	}
	mutexes[fdIndex].Unlock()
}

func main() {
	conn, err := util.GetConn()
	if err != nil {
		log.Fatalf("Failed to start a connection: %v", err)
	}
	defer conn.Close()

	client := pb.NewFileServiceClient(conn)

	metadata, err := client.GetFileMetadata(context.Background(), &pb.FileMetadataRequest{})
	if err != nil {
		log.Fatalf("Failed to fetch file metadata: %v", err)
	}

	var startChunk = 0
	for {
		lastChunk, err := downloadFile(client, startChunk, metadata.TotalChunks, metadata.TotalSize)
		if lastChunk == int(metadata.TotalChunks) {
			fmt.Println("\nFile download complete")
			return
		}

		if err != nil {
			fmt.Println("Error while downloading, retry in 10 seconds: ", err.Error())
			time.Sleep(10 * time.Second)
		}
		startChunk = lastChunk
	}

}
