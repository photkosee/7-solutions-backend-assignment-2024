package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"

	"google.golang.org/grpc"

	pb "3-pie-fire-dire/services"
)

type server struct {
	pb.UnimplementedBeefServiceServer
}

func (s *server) GetBeef(ctx context.Context, req *pb.BeefRequest) (*pb.BeefResponse, error) {
	resp, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch beef data: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Initialize map to count word occurrences
	mapBeef := make(map[string]int32)
	// Regular expression to split words
	re := regexp.MustCompile(`\s+|\.|,`)

	// Process and count words (case insensitive)
	for _, word := range re.Split(strings.ToLower(string(body)), -1) {
		if word != "" {
			mapBeef[word]++
		}
	}

	res := &pb.BeefResponse{
		Beef: mapBeef,
	}

	return res, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	
	s := grpc.NewServer()
	pb.RegisterBeefServiceServer(s, &server{})

	fmt.Println("gRPC server listening on port 50051")

	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
