package main

import (
	"context"
	"fmt"
	"log"
	"time"

	payment_pb "github.com/Nabwinsaud/microservices-gang/proto/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	//* grpc.Dial and grpc.WithInsecure is deprecated
	conn, err := grpc.NewClient("localhost:5051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error occurred while connecting to grpc server", err)
	}
	defer conn.Close()

	client := payment_pb.NewPaymentServiceClient(conn)
	payment := &payment_pb.PaymentRequest{
		PaymentId:     "123",
		Status:        "success",
		PaymentMethod: payment_pb.PaymentMethod_STRIPE,
		PaymentDate:   "2025-01-26",
		PaymentDetails: &payment_pb.PaymentDetails{
			StripePaymentId: "stripe123",
		},
		CreatedAt: timestamppb.New(time.Now()),
		PaymentLinkExpiry: &durationpb.Duration{
			Seconds: int64(time.Second * 60),
		},
	}

	res, err := client.CreatePayment(context.Background(), payment)
	if err != nil {
		log.Fatal("error while creating payment", err)
	}
	if res == nil {
		log.Fatal("payment is nil")
		return
	}
	fmt.Println("payment created successfully:", res.Payment)

}
