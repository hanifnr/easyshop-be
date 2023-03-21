package service

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var fcmClient *messaging.Client

func InitFirebase() {
	SERVICE_ACCOUNT_PATH := os.Getenv("SERVICE_ACCOUNT_PATH")

	// Initialize firebase app
	opt := option.WithCredentialsFile(SERVICE_ACCOUNT_PATH)
	config := &firebase.Config{ProjectID: "easy-shop-364408"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		fmt.Println("error init firebase ", err.Error())
		return
	}

	fcmClient, err = app.Messaging(context.Background())
	if err != nil {
		fmt.Println("error init firebase ", err.Error())
		return
	}
}

func SendPushNotification(tokens []string, notification *messaging.Notification) {

	_, err := fcmClient.SendMulticast(context.Background(), &messaging.MulticastMessage{
		Notification: notification,
		Tokens:       tokens, // it's an array of device tokens
	})

	if err != nil {
		fmt.Println("error send push notification ", err.Error())
	}
}
