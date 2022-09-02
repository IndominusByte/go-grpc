package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/IndominusByte/go-grpc/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func doCreateBlog(c pb.BlogServiceClient) string {
	res, err := c.CreateBlog(context.Background(), &pb.Blog{
		Title:   "kasus sambo",
		Author:  "oman",
		Content: "indoo oh indoo",
	})

	if err != nil {
		log.Fatalf("Cannot calling create blog: %v\n", err)
	}

	return res.Id
}

func doReadBlog(c pb.BlogServiceClient, id string) {
	res, err := c.ReadBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("Cannot calling read blog: %v\n", err)
	}

	fmt.Println(res)
}

func doUpdateBlog(c pb.BlogServiceClient, id string) {
	data := pb.Blog{
		Id:      id,
		Author:  "edit bro",
		Title:   "edit titleee",
		Content: "edit conteent",
	}

	_, err := c.UpdateBlog(context.Background(), &data)

	if err != nil {
		log.Fatalf("Cannot calling update blog: %v\n", err)
	}

	fmt.Println("success update blog")
}

func doListBlogs(c pb.BlogServiceClient) {
	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Cannot calling list blogs: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while read data: %v\n", err)
		}
		fmt.Println(msg)
	}
}

func doDeleteBlog(c pb.BlogServiceClient, id string) {
	_, err := c.DeleteBlogs(context.Background(), &pb.BlogId{
		Id: id,
	})
	if err != nil {
		log.Fatalf("Cannot calling delete blog: %v\n", err)
	}

	fmt.Println("Success delete id: ", id)
}
