package service

import (
	"context"
	"easyshop/utils"
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

func SendPushNotification(isAdmin bool, notification *messaging.Notification) {
	db := utils.GetDB()

	var tipe string
	if isAdmin {
		tipe = "ADMIN"
	} else {
		tipe = "CUST"
	}

	deviceTokens := make([]string, 0)
	db.Select("token").Table("firebase_token").Where("is_delete = FALSE AND type = ?", tipe).Scan(&deviceTokens)

	_, err := fcmClient.SendMulticast(context.Background(), &messaging.MulticastMessage{
		// Notification: &messaging.Notification{
		// 	Title: "Congratulations!!",
		// 	Body:  "You have just implemented push notification",
		// },
		Notification: notification,
		Tokens:       deviceTokens, // it's an array of device tokens
	})

	if err != nil {
		fmt.Println("error send push notification ", err.Error())
	}
}
