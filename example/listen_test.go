package example

import (
	"fmt"
	"github.com/a316523235/wingo/common"
	"github.com/a316523235/wingo/service"
	"github.com/go-vgo/robotgo"
	"github.com/robotn/gohook"
	"testing"
	"time"
)

func TestListenRMouse(t *testing.T)  {
	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {

		fmt.Println(e.String())
		if e.Button == hook.MouseMap["left"] {
			fmt.Println("mouse left")
		}
	})

	s := hook.Start()
	<- hook.Process(s)
}

func TestWriteByWord(t *testing.T) {
	//time.Sleep(3 * time.Second)
	robotgo.AddEvents("4", "alt")
	service.ReadWord()
}

func TestCopy(t *testing.T)  {
	time.Sleep(3 * time.Second)
	robotgo.KeyTap("c","ctrl")
	robotgo.MilliSleep(100)
	str, err := robotgo.ReadAll()
	fmt.Println(str)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestWriteEnter(t *testing.T)  {
	str := robotgo.PasteStr("123")
	fmt.Println(str)
}

func TestToField(t *testing.T)  {
	str := common.ToField("user_id")
	fmt.Println(str)
}

func TestAltEnter(t *testing.T)  {
	robotgo.SetDelay(100, 100)
	robotgo.AddEvent("enter")
	robotgo.TypeStr("here")
}

func TestAlt4(t *testing.T)  {
	hook.Register(hook.KeyDown, []string{"4", "ctrl"}, func(e hook.Event) {
		fmt.Println("alt 4")
	})

	s := hook.Start()
	<-hook.Process(s)
}

func TestSleepOne(t *testing.T)  {
	robotgo.EventHook(hook.KeyDown, []string{"h"}, func(e hook.Event) {
		fmt.Println("h")
		time.Sleep(3 * time.Second)
	})

	robotgo.EventHook(hook.KeyDown, []string{"w"}, func(e hook.Event) {
		fmt.Println("w")
	})

	//kk := []string{
	//	"num_enter",
	//	"capslock",
	//	"space",
	//	"print",
	//	"printscreen",
	//}

	robotgo.EventHook(hook.KeyDown, []string{"num_enter"}, func(e hook.Event) {
		fmt.Println("num_enter")
	})
	//robotgo.EventHook(hook.KeyDown, []string{"capslock"}, func(e hook.Event) {
	//	fmt.Println("capslock")
	//})
	//robotgo.EventHook(hook.KeyDown, []string{"space"}, func(e hook.Event) {
	//	fmt.Println("space")
	//})
	//robotgo.EventHook(hook.KeyDown, []string{"print"}, func(e hook.Event) {
	//	fmt.Println("print")
	//})
	//robotgo.EventHook(hook.KeyDown, []string{"printscreen"}, func(e hook.Event) {
	//	fmt.Println("printscreen")
	//})


	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}