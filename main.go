package main

import (
	"fmt"
	"os"

	payment_pb "github.com/Nabwinsaud/microservices-gang/payment"
	"google.golang.org/protobuf/proto"
)

func initiatePayment() *payment_pb.Payment {
	stripe := payment_pb.PaymentMethod_STRIPE.Enum()

	payment := &payment_pb.Payment{
		PaymentId:     "123",
		Status:        "success",
		PaymentMethod: *stripe,
		PaymentDate:   "2025-01-26",
	}
	fmt.Printf("payment: PaymentId=%s, Status=%s, PaymentMethod=%v, PaymentDate=%s\n", payment.PaymentId, payment.Status, payment.PaymentMethod, payment.PaymentDate)
	fmt.Println("stripe", stripe, payment.GetPaymentId())

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
	err := writeFile("payment.txt", payment)
	if err != nil {
		fmt.Println("error occurred while writing file", err)
	}
	newPayment := &payment_pb.Payment{}
	readFile("payment.txt", newPayment)

}
