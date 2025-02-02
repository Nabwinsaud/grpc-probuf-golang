package main

import (
	"context"
	"fmt"
	"log"
	"net"

	payment_pb "github.com/Nabwinsaud/microservices-gang/proto/gen/proto"
	"google.golang.org/grpc"
)

type server struct {
	payment_pb.UnimplementedPaymentServiceServer
}

func (s *server) CreatePayment(ctx context.Context, req *payment_pb.PaymentRequest) (*payment_pb.PaymentResponse, error) {
	// clonedPayment := proto.Clone(req).(*payment_pb.Payment)
	// fmt.Println("payment cloned:", clonedPayment)
	fmt.Println("Payment body", req.PaymentId, req.CreatedAt, req.PaymentMethod)
	return &payment_pb.PaymentResponse{
		Message: "Payment created successfully",
		Payment: &payment_pb.Payment{
			PaymentId:     req.PaymentId,
			Status:        req.Status,
			CreatedAt:     req.CreatedAt,
			PaymentMethod: req.PaymentMethod,
			PaymentDate:   req.PaymentDate,
			PaymentDetails: &payment_pb.PaymentDetails{
				StripePaymentId: req.PaymentDetails.StripePaymentId,
			},
			PaymentLinkExpiry: req.PaymentLinkExpiry,
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5051")
	if err != nil {
		log.Fatal("error while listening to port", err)
	}
	grpcServer := grpc.NewServer()
	payment_pb.RegisterPaymentServiceServer(grpcServer, &server{})
	fmt.Println("server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("error while  serving the server", err)
	}
}
