package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
	pb "user"
	"user/server/interceptor"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnsafeUserServiceServer
}

func (s *server) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("Get request received: %s", in.GetId())
	return &pb.UserResponse{
		Id:    in.Id,
		Name:  "testUser",
		Age:   24,
		Email: "test@example.com",
	}, nil
}

func (s *server) ServerStreamingGetUser(in *pb.UserRequest, stream pb.UserService_ServerStreamingGetUserServer) error {
	log.Printf("Server streaming request received: %s", in.GetId())
	for i := 0; i < 5; i++ {
		err := stream.Send(&pb.UserResponse{
			Id:    in.Id,
			Name:  fmt.Sprintf("testUser%d", i),
			Age:   24,
			Email: "test@example.com",
		})
		if err != nil {
			return err
		}
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

func (s *server) ClientStreamingGetUser(stream pb.UserService_ClientStreamingGetUserServer) error {
	log.Printf("Client streaming request received")
	var ids []string
	for {
		in, err := stream.Recv()
		if err != nil {
			break
		}
		log.Printf("receive: %s", in.GetId())
		ids = append(ids, in.GetId())
	}
	return stream.SendAndClose(&pb.UserResponse{
		Id:    fmt.Sprintf("%v", ids),
		Name:  "testUserSendCompleted!!",
		Age:   24,
		Email: "test@example.com",
	})
}

func (s *server) BidirectionalStreamingGetUser(stream pb.UserService_BidirectionalStreamingGetUserServer) error {
	log.Printf("Bidirectional streaming request received")
	for {
		in, err := stream.Recv()
		if err != nil {
			return nil
		}
		log.Printf("receive: %s", in.GetId())
		err = stream.Send(&pb.UserResponse{
			Id:    in.GetId(),
			Name:  "testUser",
			Age:   24,
			Email: "test@example.com",
		})
		if err != nil {
			return err
		}
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(interceptor.UnaryServerInterceptor), grpc.StreamInterceptor((interceptor.StreamingServerInterceptor)))
	pb.RegisterUserServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
