package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/IndominusByte/go-grpc/calculator/proto"
)

func doCalculator(c pb.CalculatorServiceClient) {
	res, err := c.Calculator(context.Background(), &pb.CalculatorRequest{
		Num_1: 10,
		Num_2: 3,
	})

	if err != nil {
		log.Fatalf("Could not calculator: %v\n", err)
	}

	log.Println(res.GetResult())
}

func doPrimes(c pb.CalculatorServiceClient) {
	stream, err := c.Primes(context.Background(), &pb.PrimeRequest{
		Number: 120,
	})

	if err != nil {
		log.Fatalf("Error while calling Primes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Println("Primes:", msg.GetResult())
	}

}

func doAvg(c pb.CalculatorServiceClient) {
	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Avg: %v\n", err)
	}

	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receive response: %v\n", err)
	}
	fmt.Println(response.GetResult())
}

func doMax(c pb.CalculatorServiceClient) {
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Max: %v\n", err)
	}

	reqs := []int{
		1, 5, 3, 6, 2, 20,
	}

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending: ", req)
			stream.Send(&pb.MaxRequest{Number: int64(req)})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	waitc := make(chan bool)

	go func() {
		for {
			msg, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error receive response: %v\n", err)
			}

			fmt.Println("Receive: ", msg.GetResult())
		}
		close(waitc)
	}()

	<-waitc
}
