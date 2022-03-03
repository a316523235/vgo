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

var keyHoldChan chan hook.Event = make(chan hook.Event, 100)

func TestListenAll(t *testing.T) {
	hook.Register(hook.KeyHold, []string{}, func(e hook.Event) {
		//fmt.Println(e.String())
		keyHoldChan <- e
	})

	go func() {
		for  {
			select {
			case s := <-keyHoldChan:
				fmt.Println(s.String())
			}
		}
	}()

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