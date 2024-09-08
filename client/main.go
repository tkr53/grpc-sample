package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
	"user/client/interceptor"

	pb "user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to listen")
)

func unaryRPC(c pb.UserServiceClient, ctx context.Context) {
	r, err := c.GetUser(ctx, &pb.UserRequest{Id: "test"})
	if err != nil {
		log.Fatalf("could not user: %v", err)
	}
	log.Printf("UserName: %s", r.GetName())
}

func serverStreamingRPC(c pb.UserServiceClient, ctx context.Context) {
	stream, err := c.ServerStreamingGetUser(ctx, &pb.UserRequest{Id: "test"})
	if err != nil {
		log.Fatalf("could not list users: %v", err)
	}
	for {
		user, err := stream.Recv()
		if err != nil {
			break
		}
		log.Printf("UserName: %s", user.GetName())
	}
}

func clientStreamingRPC(c pb.UserServiceClient, ctx context.Context) {
	stream, err := c.ClientStreamingGetUser(ctx)
	if err != nil {
		log.Fatalf("could not list users: %v", err)
	}
	for i := 0; i < 5; i++ {
		err := stream.Send(&pb.UserRequest{Id: fmt.Sprintf("test%d", i)})
		if err != nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	r, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not list users: %v", err)
	}
	log.Printf("UserName: %s", r.GetName())
}

func bidirectionalStreamingRPC(c pb.UserServiceClient, ctx context.Context) {
	sc := bufio.NewScanner(os.Stdin)
	stream, err := c.BidirectionalStreamingGetUser(ctx)
	if err != nil {
		log.Fatalf("could not list users: %v", err)
	}
	end := false
	for !end {
		sc.Scan()
		send := sc.Text()
		log.Printf("Send: %s", send)
		err := stream.Send(&pb.UserRequest{Id: send})
		if err != nil {
			log.Fatalf("could not close stream: %v", err)
		}
		if send == "bye" {
			if err := stream.CloseSend(); err != nil {
				log.Fatalf("could not close stream: %v", err)
			}
			end = true
		}

		user, err := stream.Recv()
		if err != nil {
			log.Fatalf("could not receive: %v", err)
		}
		log.Printf("Receive: %s", user.GetId())
	}
}

func main() {
	flag.Parse()

	sc := bufio.NewScanner(os.Stdin)

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(interceptor.UnaryClientInterceptor), grpc.WithStreamInterceptor(interceptor.StreamingClientInterceptor))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	sc.Scan()

	switch sc.Text() {
	case "unary":
		unaryRPC(c, ctx)
	case "server":
		serverStreamingRPC(c, ctx)
	case "client":
		clientStreamingRPC(c, ctx)
	case "bidirectional":
		bidirectionalStreamingRPC(c, ctx)
	default:
		break
	}
}
