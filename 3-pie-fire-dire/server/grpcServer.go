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
	"time"

	"github.com/patrickmn/go-cache"
	"google.golang.org/grpc"

	pb "3-pie-fire-dire/services"
)

type server struct {
	pb.UnimplementedBeefServiceServer
}

// Cache for 10 minutes, with a cleanup interval of 20 minutes
var beefCache = cache.New(10*time.Minute, 20*time.Minute)

func CountTypeBeef(text string) map[string]int32 {
	mapBeef := make(map[string]int32)
	// Regular expression to split words
	re := regexp.MustCompile(`\s+|\.|,`)

	// Process and count each type of beef (case insensitive)
	for _, word := range re.Split(strings.ToLower(text), -1) {
		if word != "" {
			mapBeef[word]++
		}
	}

	return mapBeef
}

func (s *server) GetBeef(ctx context.Context, req *pb.BeefRequest) (*pb.BeefResponse, error) {
	// Check if the beef data is already cached
	if cachedData, found := beefCache.Get("beef_data"); found {
		return cachedData.(*pb.BeefResponse), nil
	}

	// If not found in cache, fetch from the external API
	resp, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch beef data: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Count each type of beef
	mapBeef := CountTypeBeef(string(body))

	res := &pb.BeefResponse{
		Beef: mapBeef,
	}

	// Cache the response for future requests
	beefCache.Set("beef_data", res, cache.DefaultExpiration)

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
