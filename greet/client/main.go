package main

import (
	"log"
	"time"

	pb "github.com/IndominusByte/go-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const addr string = "bhaktirahayugroup.co.id:8008"

func main() {

	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/bhaktirahayugroup.co.id/cert.pem"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while read certfile: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	// doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetEveryone(c)
	// doLogin(c)
	doGreetWithDeadline(c, 5*time.Second)
}
