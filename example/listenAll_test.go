package example

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"testing"
)

var keyHoldChan = make(chan hook.Event, 100)

func TestListenAll(t *testing.T) {
	hook.Register(hook.KeyHold, []string{}, func(e hook.Event) {
		//fmt.Println(e.String())
		keyHoldChan <- e
	})

	hook.Register(hook.KeyHold, []string{"enter", "alt"}, func(e hook.Event) {
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

func TestExample(t *testing.T)  {
	fmt.Println("test begin...")
	addEvent()

	addMouse()

	lowLevel()
}

func addEvent() {
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	robotgo.EventHook(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		robotgo.EventEnd()
	})

	fmt.Println("--- Please press w---")
	robotgo.EventHook(hook.KeyDown, []string{"w"}, func(e hook.Event) {
		fmt.Println("w")
	})

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

func addMouse() {
	fmt.Println("--- Please press left mouse button to see it's position and the right mouse button to exit ---")
	robotgo.EventHook(hook.MouseDown, []string{}, func(e hook.Event) {
		if e.Button == hook.MouseMap["left"] {
			fmt.Printf("mouse left @ %v - %v\n", e.X, e.Y)
		} else if e.Button == hook.MouseMap["right"] {
			robotgo.EventEnd()
		}
	})

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

func lowLevel() {
	////////////////////////////////////////////////////////////////////////////////
	// Global event listener
	////////////////////////////////////////////////////////////////////////////////
	fmt.Println("Press q to stop event gathering")
	evChan := robotgo.EventStart()
	for e := range evChan {
		fmt.Println(e)
		if e.Keychar == 'q' {
			robotgo.EventEnd()
			// break
		}
	}
}
