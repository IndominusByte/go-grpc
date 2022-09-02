package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/IndominusByte/go-grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Println(req)
	return &pb.GreetResponse{
		Result: "Hello " + req.GetFirstName(),
	}, nil
}

func (s *Server) GreetManyTimes(req *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", req.GetFirstName(), i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {

	res := ""
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: "Hello " + res,
			})
		}

		if err != nil {
			log.Fatalln("Error receiving msg: ", err)
		}

		fmt.Println("receiving: ", msg)
		res += msg.GetFirstName() + ", "
	}
}

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalln("Error receiving msg: ", err)
		}

		res := "Hello " + msg.GetFirstName() + "!"
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalln("Error while send msg: ", err)
		}
	}
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.GetUsername() != "oman" && req.GetPassword() != "123" {
		st := status.Newf(
			codes.InvalidArgument,
			"Username or password invalid.",
		)
		st, _ = st.WithDetails(req)
		return nil, st.Err()
	}

	return &pb.LoginResponse{
		Result: "Login success.",
	}, nil
}

func (s *Server) GreetWithDeadline(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, status.Error(codes.Canceled, "request cancel from client!")
		}
		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello " + req.GetFirstName(),
	}, nil
}
