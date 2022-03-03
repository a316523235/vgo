package service

import (
	"fmt"
	"github.com/a316523235/wingo/models"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"os"
	"time"
)

var Switch = &models.Switch{ TaskSwitch: true }

func Start()  {
	robotgo.SetDelay(200, 200)

	fmt.Println("--- Please press alt + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "alt"}, func(e hook.Event) {
		fmt.Println("alt-q")
		hook.End()	//exit listen
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
	hook.Register(hook.KeyDown, []string{"esc"}, func(e hook.Event) {
		Switch.CloseTask()
		fmt.Println("esc")
	})

	fmt.Println("--- Please press alt 1 to start GotoMergerPage task---")
	hook.Register(hook.KeyDown, []string{"1", "alt"}, func(e hook.Event) {
		Switch.OpenTask()
		fmt.Println("alt 1")
		go GotoMergerPage()
	})

	fmt.Println("--- Please press alt 2 to start RecordClickPosition task---")
	hook.Register(hook.KeyDown, []string{"2", "alt"}, func(e hook.Event) {
		Switch.OpenTask()
		fmt.Println("alt 2")
		//go RecordClickPosition()
		hook.StopEvent()
		time.Sleep(100 * time.Millisecond)
		RecordClickPositionV2()
		time.Sleep(100 * time.Millisecond)
		hook.End()
		Start()
	})

	fmt.Println("--- Please press alt 3 to start GotoMergerLastSubmitToRelease task---")
	hook.Register(hook.KeyDown, []string{"3", "alt"}, func(e hook.Event) {
		Switch.OpenTask()
		fmt.Println("alt 3")
		go GotoMergerLastSubmitToRelease()
	})

	fmt.Println("--- Please press alt 4 to start ReadWord task---")
	hook.Register(hook.KeyUp, []string{"4", "ctrl"}, func(e hook.Event) {
		fmt.Println("alt 4")
		go ReadWord()
	})

	s := hook.Start()
	<-hook.Process(s)
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