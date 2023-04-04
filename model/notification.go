package model

import (
	"easyshop/service"
	"easyshop/utils"
	"encoding/json"
	"fmt"
	"time"

	"firebase.google.com/go/messaging"
	"github.com/go-redis/redis"
)

type Notification struct {
	Id        int64     `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Body      string    `json:"body,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Read      bool
}

type UserNotification struct {
	Uid           string          `json:"uid,omitempty"`
	Notifications []*Notification `json:"notification,omitempty"`
}

func CreateNotification(title, body string, listUserId []string) {
	client := service.GetRedisClient()

	for _, userId := range listUserId {
		exist, userNotification := GetUserNotification(userId)
		if !exist {
			userNotification.Uid = userId
		}
		userNotification.Notifications = append(userNotification.Notifications, &Notification{
			Id:        client.Incr("notif").Val(),
			Title:     title,
			Body:      body,
			CreatedAt: time.Now(),
			Read:      false,
		})

		notificationJSON, err := json.Marshal(userNotification)
		if err != nil {
			fmt.Println(err)
		}

		SaveNotification(client, userId, string(notificationJSON))
	}
	service.CloseRedisClient(client)
}

func GetUserNotification(userId string) (bool, *UserNotification) {
	client := service.GetRedisClient()
	var notificationFromCache UserNotification
	// Get the notification from cache
	notificationString, err := client.Get(userId).Result()
	if err == redis.Nil {
		return false, &UserNotification{}
	} else if err != nil {
		fmt.Println(err.Error())
	} else {
		// Unmarshal the notification from JSON
		err = json.Unmarshal([]byte(notificationString), &notificationFromCache)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	service.CloseRedisClient(client)

	return true, &notificationFromCache
}

func GetNotification(userId string) map[string]interface{} {
	exist, retval := GetUserNotification(userId)
	if !exist {
		return utils.MessageErr(false, utils.ErrExist, "Notification not found")
	}
	return utils.MessageData(true, retval)
}

func ReadNotification(userId string, listNotifId []int64) map[string]interface{} {
	client := service.GetRedisClient()

	exist, userNotification := GetUserNotification(userId)
	if !exist {
		return utils.MessageErr(false, utils.ErrExist, "User not found")
	}
	for _, notif := range userNotification.Notifications {
		for _, notifId := range listNotifId {
			if notif.Id == notifId {
				notif.Read = true
			}
		}
	}

	notificationJSON, err := json.Marshal(userNotification)
	if err != nil {
		fmt.Println(err)
	}

	SaveNotification(client, userId, string(notificationJSON))
	service.CloseRedisClient(client)

	return utils.MessageData(true, utils.StatusReturnOK())
}

func SaveNotification(client *redis.Client, key string, value interface{}) {
	err := client.Set(key, value, 24*7*time.Hour).Err()
	if err != nil {
		fmt.Println("Redis error set nofication ", err)
	}
}

func SendPushNotification(title, body string, isAdmin bool) {
	db := utils.GetDB()

	var tipe string
	if isAdmin {
		tipe = "ADMIN"
	} else {
		tipe = "CUST"
	}

	firebaseToken := make([]*FirebaseToken, 0)
	db.Where("is_delete = FALSE AND type = ?", tipe).Find(&firebaseToken)

	tokens := make([]string, 0)
	for _, t := range firebaseToken {
		tokens = append(tokens, t.Token)
	}

	uids := make([]string, 0)
	for _, t := range firebaseToken {
		uids = append(uids, t.Uid)
	}

	service.SendPushNotification(tokens, &messaging.Notification{
		Title: "New order received",
		Body:  body,
	})
	// CreateNotification(title, body, uids)
}
