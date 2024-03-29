package service

import (
	"time"

	"github.com/robfig/cron/v3"
)

func InitScheduler(listAction []func()) {
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))

	for _, action := range listAction {
		scheduler.AddFunc("0 0 * * 0", action)
	}

	// start scheduler
	scheduler.Start()

	// // trap SIGINT untuk trigger shutdown.
	// sig := make(chan os.Signal, 1)
	// signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	// <-sig
}
