package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/IndominusByte/go-grpc/greet/proto"
	"google.golang.org/grpc/status"
)

func doGreet(c pb.GreetServiceClient) {
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "oman",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Println(res.GetResult())
}

func doGreetManyTimes(c pb.GreetServiceClient) {

	stream, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{
		FirstName: "oman",
	})

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Println("GreetManyTimes:", msg.GetResult())
	}
}

func doLongGreet(c pb.GreetServiceClient) {
	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "nyoman"},
		{FirstName: "pradipta"},
		{FirstName: "dewantara"},
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	fmt.Println("Response: ", res.GetResult())
}

func doGreetEveryone(c pb.GreetServiceClient) {
	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while calling GreetEveryone: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "nyoman"},
		{FirstName: "pradipta"},
		{FirstName: "dewantara"},
	}

	waitc := make(chan bool)

	go func() {
		for _, req := range reqs {
			stream.Send(req)
			fmt.Println("Sending: ", req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			msg, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while reading stream: %v\n", err)
				break
			}

			fmt.Println("Receive: ", msg)
		}
		close(waitc)
	}()

	<-waitc
}

func doLogin(c pb.GreetServiceClient) {
	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Username: "omans",
		Password: "1234",
	})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			fmt.Println(e.Code())
			fmt.Println(e.Message())
			fmt.Println(e.Details())
		} else {
			log.Fatalf("Error send payload: %v\n", err)
		}
	}

	fmt.Println(res.GetResult())
}

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, &pb.GreetRequest{FirstName: "oman"})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			fmt.Println(e)
		} else {
			fmt.Println("Error calling GreetWithDeadline: ", err)
		}
	}

	fmt.Println("Receive: ", res)
}
