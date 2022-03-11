package service

import (
	"encoding/json"
	"fmt"
	"github.com/a316523235/wingo/models"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"os"
	"time"
)

var Switch = &models.Switch{ TaskSwitch: true }


func Start()  {
	fmt.Println("--- Please press alt + q to stop hook ---")
	robotgo.EventHook(hook.KeyDown, []string{"q", "alt"}, func(e hook.Event) {
		fmt.Println("alt-q")
		robotgo.EventEnd()	//exit listen
	})

	//fmt.Println("--- Please press alt + o to start hook ---")
	//hook.Register(hook.KeyDown, []string{"o", "alt"}, func(e hook.Event) {
	//	fmt.Println("alt-o")
	//	door = true
	//	//hook.End()	//exit listen
	//})
	//
	//fmt.Println("--- Please press alt + c to close hook ---")
	//hook.Register(hook.KeyDown, []string{"c", "alt"}, func(e hook.Event) {
	//	fmt.Println("alt-c")
	//	door = false
	//	//hook.End()	//exit listen
	//})

	fmt.Println("--- Please press esc to break task---")
	robotgo.EventHook(hook.KeyDown, []string{"esc"}, func(e hook.Event) {
		Switch.CloseTask()
		fmt.Println("esc")
		robotgo.EventEnd()
		PrintPos(true) // copy pos by record click

		time.Sleep(3 * time.Second)
		Start()
	})

	fmt.Println("--- Please press alt 1 to start GotoMergerPage task---")
	robotgo.EventHook(hook.KeyDown, []string{"1", "alt"}, func(e hook.Event) {
		Switch.OpenTask()
		fmt.Println("alt 1")
		go GotoMergerPage()
	})

	fmt.Println("--- Please press alt 2 to start RecordClickPosition task---")
	robotgo.EventHook(hook.KeyDown, []string{"2", "alt"}, func(e hook.Event) {
		Switch.OpenTask()
		fmt.Println("alt 2")
		go RecordClickPositionV3()
	})

	fmt.Println("--- Please press alt 3 to start GotoMergerLastSubmitToRelease task---")
	robotgo.EventHook(hook.KeyDown, []string{"3", "alt"}, func(e hook.Event) {
		Switch.OpenTask()
		fmt.Println("alt 3")
		go GotoMergerLastSubmitToRelease()
	})

	fmt.Println("--- Please press alt 4 to start ReadWord task---")
	robotgo.EventHook(hook.KeyDown, []string{"4", "alt"}, func(e hook.Event) {
		Switch.OpenTask()
		fmt.Println("alt 4")
		go ReadWord()
	})

	fmt.Println("--- Please press alt 5 to start Print Key ---")
	robotgo.EventHook(hook.KeyDown, []string{"5", "alt"}, func(e hook.Event) {
		Switch.OpenTask()
		fmt.Println("alt-5")
		go PrintAllKeyCode()
	})

	fmt.Println("--- Please press alt 6 to start Booking Key ---")
	robotgo.EventHook(hook.KeyDown, []string{"6", "alt"}, func(e hook.Event) {
		Switch.OpenTask()
		fmt.Println("alt-6")
		go Booking()
	})

	fmt.Println("--- Please press alt 7 to start PrintPosition Key ---")
	robotgo.EventHook(hook.KeyDown, []string{"7", "alt"}, func(e hook.Event) {
		Switch.OpenTask()
		fmt.Println("alt-7")
		go PrintPosition()
	})

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

// Esc exit script
func Esc() {
	chEsc := make(chan bool)
	out := time.After(30 * time.Second)

	go func() {
		chEsc <- robotgo.AddEvent("esc")
	}()

	select {
	case <-chEsc:
		fmt.Println("esc over")
		os.Exit(1)
	case <-out:
		fmt.Println("timeout over")
		os.Exit(1)
	}
}

func PrintAllKeyCode()  {
	robotgo.EventHook(hook.KeyHold, []string{}, func(e hook.Event) {
		str, _ := json.Marshal(e)
		fmt.Println(string(str))
		//fmt.Printf("%#v", e)
	})
	time.Sleep(1 * time.Millisecond)
}

func PrintPosition()  {
	robotgo.EventHook(hook.KeyHold, []string{"alt"}, func(e hook.Event) {
		x, y := robotgo.GetMousePos()
		fmt.Printf("{%d, %d},\n\r", x, y)
	})
	time.Sleep(1 * time.Millisecond)
}