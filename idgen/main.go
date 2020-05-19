package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func main() {
	opt := option.WithCredentialsFile("../keys_DO_NOT_MODIFY/midnight-31686-firebase-adminsdk-obui8-fe0e06993f.json")
	midnightAuth, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Errorf("error initializing app: %v", err)
	}

	defaultClient, err := midnightAuth.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

}
