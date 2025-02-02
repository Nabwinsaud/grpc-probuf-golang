package main

import (
	"fmt"
	"os"
	"time"

	payment_pb "github.com/Nabwinsaud/microservices-gang/proto/gen/proto"
	"google.golang.org/protobuf/proto"
	duration_pb "google.golang.org/protobuf/types/known/durationpb"
	timestamp_pb "google.golang.org/protobuf/types/known/timestamppb"
)

func initiatePayment() *payment_pb.Payment {
	stripe := payment_pb.PaymentMethod_STRIPE.Enum()

	payment := &payment_pb.Payment{
		PaymentId:     "123",
		Status:        "success",
		PaymentMethod: *stripe,
		PaymentDate:   "2025-01-26",
		PaymentDetails: &payment_pb.PaymentDetails{
			StripePaymentId: "stripe123",
		},
		CreatedAt: timestamp_pb.New(time.Now()),
		PaymentLinkExpiry: &duration_pb.Duration{
			Seconds: int64(time.Second * 5),
		},
	}
	fmt.Printf("payment: PaymentId=%s, Status=%s, PaymentMethod=%v, PaymentDate=%s\n", payment.PaymentId, payment.Status, payment.PaymentMethod, payment.PaymentDate)
	fmt.Println("stripe", stripe, payment.GetPaymentId())
	fmt.Println("payment created_at", payment.GetCreatedAt()) // it return the timestamp in this format {seconds: 1611630000, nanos: 204114511}
	fmt.Println("payment details", payment.PaymentDetails.GetStripePaymentId())

	//* payment link expiry check - it does not really need to be done in the code, it is just for demonstration
	time.AfterFunc(time.Second*5, func() {
		fmt.Println("payment link expired after 5 seconds", payment.GetPaymentLinkExpiry())
	})
	time.Sleep(time.Second * 6)
	return payment

}

func readFile(filename string, pb proto.Message) error {
	message, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error while reading the file:", err)
		return err
	}
	err = proto.Unmarshal(message, pb)
	if err != nil {
		fmt.Println("could not unmarshal the message:", err)
		return err
	}

	fmt.Println("payment read file :", pb)

	return nil
}

func writeFile(filename string, pb proto.Message) error {
	message, err := proto.Marshal(pb)
	if err != nil {
		fmt.Println("could not serialize the proto message:", err)
		return err
	}
	err = os.WriteFile(filename, message, 0644)
	if err != nil {
		fmt.Println("error while writing the file to disk:", err)
		return err
	}

	return nil
}

func main() {
	payment := initiatePayment()
	fmt.Println("payment", payment)
	err := writeFile("payment.bin", payment)
	if err != nil {
		fmt.Println("error occurred while writing file", err)
	}
	newPayment := &payment_pb.Payment{}
	readFile("payment.bin", newPayment)

}
