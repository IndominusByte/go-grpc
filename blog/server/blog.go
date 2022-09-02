package main

import (
	"context"
	"fmt"

	pb "github.com/IndominusByte/go-grpc/blog/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BlogItem struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Author  string             `bson:"author"`
	Title   string             `bson:"title"`
	Content string             `bson:"content"`
}

func (s *Server) CreateBlog(ctx context.Context, req *pb.Blog) (*pb.BlogId, error) {
	data := BlogItem{
		Author:  req.Author,
		Title:   req.Title,
		Content: req.Content,
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Failed to insert data: %v\n", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to OID",
		)
	}

	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil
}

func (s *Server) ReadBlog(ctx context.Context, req *pb.BlogId) (*pb.Blog, error) {
	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot parse id")
	}
	result := &pb.Blog{}

	if err := collection.FindOne(ctx, bson.M{"_id": oid}).Decode(result); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Blog not found!")
	}

	return result, nil
}

func (s *Server) UpdateBlog(ctx context.Context, req *pb.Blog) (*empty.Empty, error) {
	oid, err := primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot parse id")
	}

	data := BlogItem{
		Author:  req.Author,
		Title:   req.Title,
		Content: req.Content,
	}
	result, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": data})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cannot update document: %v\n", err)
	}

	if result.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Blog not found!")
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) ListBlogs(req *empty.Empty, stream pb.BlogService_ListBlogsServer) error {
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return status.Errorf(codes.Internal, "Unknown Error: %v", err)
	}

	for cursor.Next(context.TODO()) {
		var result BlogItem
		if err := cursor.Decode(&result); err != nil {
			return status.Errorf(codes.Internal, "Unknown Error: %v", err)
		}
		stream.Send(&pb.Blog{
			Id:      result.ID.Hex(),
			Author:  result.Author,
			Title:   result.Title,
			Content: result.Content,
		})
	}

	return nil
}

func (s *Server) DeleteBlogs(ctx context.Context, req *pb.BlogId) (*empty.Empty, error) {
	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot parse id")
	}

	result, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while delete data: %v", err)
	}

	if result.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Blog not found!")
	}
	return &emptypb.Empty{}, nil
}
