package main

import (
	"log"

	pb "github.com/IndominusByte/go-grpc/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr string = ":8008"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)
	valid := doCreateBlog(c)
	// doReadBlog(c, valid)
	// // doReadBlog(c, "63106ed614b498d3a1d51240")
	// doUpdateBlog(c, valid)
	// doUpdateBlog(c, "63106ed614b498d3a1d51240")
	doListBlogs(c)
	doDeleteBlog(c, valid)
	doListBlogs(c)
	doDeleteBlog(c, "63106ed614b498d3a1d51240")
}
