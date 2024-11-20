package service

import (
	"fmt"
	"github.com/a316523235/wingo/models"
	"github.com/go-toast/toast"
	"log"
	"time"
)

// origin example
func notifyExample() {
	notification := toast.Notification{
		AppID:   "Example App",
		Title:   "My notification",
		Message: "Some message about how important something is...",
		Icon:    "", // This file must exist (remove this line if no icon is desired)
		Actions: []toast.Action{
			{"protocol", "Yes", ""},
			{"protocol", "No !", ""},
		},
	}

	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}

// origin example
func Notify(msg string) {
	notification := toast.Notification{
		AppID:   "Example App",
		Title:   "My notification",
		Message: msg,
		Icon:    "", // This file must exist (remove this line if no icon is desired)
		Actions: []toast.Action{
			{"protocol", "Yes", ""},
			{"protocol", "No !", ""},
			//{"protocol", "No 3 !", ""},
			//{"protocol", "No 4 !", ""},
			//{"protocol", "No 5 !", ""},
			// 最多5个， 超过不提醒，不报错
		},
	}

	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}

func StartNotify()  {
	go func() {
		d := 1 * time.Second
		tick := time.NewTicker(d)
		for {
			select {
			case <-tick.C:
				checkCronNotify()
			}
		}
	}()
}

func checkCronNotify()  {
	//location, _ := time.LoadLocation("Asia/Shanghai")
	arr := []models.CronTip{
		{
			Ts: "2023-11-17 15:30:00",
			Tip: "now is ready 1",
		},
		{
			Ts: "2023-11-17 15:31:00",
			Tip: "now is ready 2",
		},
		{
			Ts: "2023-11-17 18:25:00",
			Tip: "体检报告和抽纸",
		},
	}

	now := time.Now().Format("2006-01-02 15:04:05")

	for _, v := range arr  {
		if now == v.Ts {
			fmt.Println("命中：" + v.Ts)
			Notify(v.Tip)
		}
	}
}