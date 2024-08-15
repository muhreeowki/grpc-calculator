package main

import (
	"context"
	"log"
	"net"

	"github.com/muhreeowki/calculator-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, input *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	return &pb.CalculationResponse{
		Result: input.A + input.B,
	}, nil
}

func (s *server) Sub(ctx context.Context, input *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	return &pb.CalculationResponse{
		Result: input.A - input.B,
	}, nil
}

func (s *server) Mul(ctx context.Context, input *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	return &pb.CalculationResponse{
		Result: input.A * input.B,
	}, nil
}

func (s *server) Divide(ctx context.Context, input *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	if input.B == 0 {
		return nil, status.Error(codes.InvalidArgument, "Cannot divide by 0")
	}

	return &pb.CalculationResponse{
		Result: input.A / input.B,
	}, nil
}

func (s *server) Sum(ctx context.Context, input *pb.NumbersRequest) (*pb.CalculationResponse, error) {
	var sum int64

	for i := range len(input.Numbers) {
		sum += input.Numbers[i]
	}

	return &pb.CalculationResponse{
		Result: sum,
	}, nil
}

func main() {
	tcpListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("failed to create listener: ", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterCalculatorServer(s, &server{})
	if err := s.Serve(tcpListener); err != nil {
		log.Fatalln("failed to serve: ", err)
	}
}
