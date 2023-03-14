package service

import (
	"context"
	"easyshop/utils"
	"encoding/base64"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var fcmClient *messaging.Client

func getDecodedFireBaseKey() ([]byte, error) {

	fireBaseAuthKey := os.Getenv("FIREBASE_AUTH_KEY")

	decodedKey, err := base64.StdEncoding.DecodeString(fireBaseAuthKey)
	if err != nil {
		return nil, err
	}

	return decodedKey, nil
}

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

func SendPushNotification() {
	db := utils.GetDB()

	deviceTokens := make([]string, 0)
	db.Select("token").Table("token").Where("is_deleted = FALSE").Scan(&deviceTokens)

	_, err := fcmClient.SendMulticast(context.Background(), &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: "Congratulations!!",
			Body:  "You have just implemented push notification",
		},
		Tokens: deviceTokens, // it's an array of device tokens
	})

	if err != nil {
		fmt.Println("error send push notification ", err.Error())
	}
}
