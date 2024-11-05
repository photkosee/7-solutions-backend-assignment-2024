package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	pb "3-pie-fire-dire/services"
)

type testServer struct {
	pb.UnimplementedBeefServiceServer
}

func (s *testServer) GetBeef(ctx context.Context, req *pb.BeefRequest) (*pb.BeefResponse, error) {
	beefs := "Beef bacon beef beef ribs ribs b b b b"
	mapBeef := CountTypeBeef(beefs)

	return &pb.BeefResponse{
		Beef: mapBeef,
	}, nil
}

// Testing the GetBeef function
func TestGetBeef(t *testing.T) {
	// Initialize the test server
	s := &testServer{}

	resp, err := s.GetBeef(context.Background(), &pb.BeefRequest{})
	assert.NoError(t, err)

	expectedCount := map[string]int32{
		"beef":  3,
		"bacon": 1,
		"ribs":  2,
		"b": 4,
	}
	assert.Equal(t, expectedCount, resp.Beef)
}
