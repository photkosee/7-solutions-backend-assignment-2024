package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "3-pie-fire-dire/services"
)

func main() {
	app := fiber.New()

	app.Get("/beef/summary", func(c fiber.Ctx) error {
		// Connect to gRPC server
		conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
 		defer conn.Close()

		// Create gRPC client
		beefServiceClient := pb.NewBeefServiceClient(conn)

		// Call GetBeef on the gRPC server
		data, err := beefServiceClient.GetBeef(context.Background(), &pb.BeefRequest{})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"beef": data.Beef,
		})
	})

	log.Fatal(app.Listen(":3030"))
}
