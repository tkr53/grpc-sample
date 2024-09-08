package interceptor

import (
	"context"
	"errors"
	"io"
	"log"

	"google.golang.org/grpc"
)

func UnaryClientInterceptor(ctx context.Context, method string, req, res interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Printf("[pre] unary client interceptor: %s", method)
	err := invoker(ctx, method, req, res, cc, opts...)
	log.Printf("[post] unary client interceptor: %s", res)
	return err
}

type wrappedStream struct {
	grpc.ClientStream
}

func (s *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("[pre] send msg: %v", m)
	return s.ClientStream.SendMsg(m)
}

func (s *wrappedStream) RecvMsg(m interface{}) error {
	err := s.ClientStream.RecvMsg(m)

	if !errors.Is(err, io.EOF) {
		log.Printf("[post] recv msg: %v", m)
	}
	return err
}

func (s *wrappedStream) CloseSend() error {
	err := s.ClientStream.CloseSend()

	log.Println("[post] streaming client interceptor")
	return err
}

func StreamingClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	log.Printf("[pre] streaming client interceptor: %s", method)
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return &wrappedStream{s}, nil
}
