package main

import (
	"Lecture7/api"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	port = ":8000"
)

func main() {
	ctx := context.Background()

	connStartTime := time.Now()
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("could not connect to %s: %v", port, err)
	}
	log.Printf("connected in %d microsec", time.Now().Sub(connStartTime).Microseconds())

	mailSenderClient := api.NewMailServiceClient(conn)
	personalAccountClient := api.NewPersonalAccountServiceClient(conn)

	mailSendStartTime := time.Now()
	_, err = mailSenderClient.MailSend(ctx, &api.MailSendRequest{
		To:      "adilzanj@mail.ru",
		Message: "please, send password",
	})
	if err != nil {
		log.Fatalf("could not send mail: %v", err)
	}

	log.Printf("got response in %d microsec", time.Now().Sub(mailSendStartTime).Microseconds())

	validAccountID := int64(1)
	invalidAccountID := int64(123)

	validAccount, err := personalAccountClient.PersonalAccount(ctx, &api.PersonalAccountRequest{Id: validAccountID})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("got account %+v", validAccount)

	_, err = personalAccountClient.PersonalAccount(ctx, &api.PersonalAccountRequest{Id: invalidAccountID})
	if err != nil {
		log.Println("got err: %v", err)
	}
}
