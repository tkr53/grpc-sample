package interceptor

import (
	"context"
	"errors"
	"io"
	"log"

	"google.golang.org/grpc"
)

func UnaryServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Println("[pre] unary server interceptor: ", info.FullMethod)
	res, err := handler(ctx, req)
	log.Println("[post] unary server interceptor : ", res)
	return res, err
}

type wrappedStream struct {
	grpc.ServerStream
}

func (s *wrappedStream) SendMsg(m interface{}) error {
	log.Println("[post] send msg: ", m)
	return s.ServerStream.SendMsg(m)
}

func (s *wrappedStream) RecvMsg(m interface{}) error {
	err := s.ServerStream.RecvMsg(m)
	if !errors.Is(err, io.EOF) {
		log.Println("[pre] recv msg: ", m)
	}
	return err
}

func StreamingServerInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	log.Println("[pre] streaming server interceptor: ", info.FullMethod)
	err := handler(srv, &wrappedStream{ss})
	log.Println("[post] streaming server interceptor: ", err)
	return err
}
