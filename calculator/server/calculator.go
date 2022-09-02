package main

import (
	"context"
	"io"
	"log"

	pb "github.com/IndominusByte/go-grpc/calculator/proto"
)

func (s *Server) Calculator(ctx context.Context, req *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	return &pb.CalculatorResponse{
		Result: req.GetNum_1() + req.GetNum_2(),
	}, nil
}

func (s *Server) Primes(req *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	var (
		k int64 = 2
		n int64 = req.GetNumber()
	)

	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: k,
			})
			n /= k
		} else {
			k++
		}
	}
	return nil
}

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {

	ans, count := 0, 0

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(ans) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while receive msg: %v\n", err)
		}

		ans += int(msg.GetNumber())
		count++
	}
}

func GetMax(ins []int) int {
	if len(ins) == 0 {
		return 0
	}

	max := ins[0]
	for _, in := range ins {
		if in > max {
			max = in
		}
	}

	return max
}

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	input := []int{}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while receive msg: %v\n", err)
		}

		input = append(input, int(msg.GetNumber()))

		err = stream.Send(&pb.MaxResponse{
			Result: int64(GetMax(input)),
		})

		if err != nil {
			log.Fatalln("Error while sending msg: ", err)
		}
	}
}
