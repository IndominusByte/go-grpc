package main

import (
	"log"
	"net"

	pb "github.com/IndominusByte/go-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
	pb.GreetServiceServer
}

const addr string = "bhaktirahayugroup.co.id:8008"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("Listening on: %v\n", addr)

	opts := []grpc.ServerOption{}
	tls := true

	if tls {
		certFile := "ssl/bhaktirahayugroup.co.id/fullchain.pem"
		keyFile := "ssl/bhaktirahayugroup.co.id/privkey.pem"

		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificate: %v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
