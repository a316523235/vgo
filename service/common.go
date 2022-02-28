package service

import (
	"fmt"
	"github.com/a316523235/wingo/common"
	"github.com/a316523235/wingo/models"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"os"
	"time"
)

var Switch = &models.Switch{ TaskSwitch: true }

func Start()  {
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
		go RecordClickPosition()
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

// TestGotoMergerPage open wait merge page
func GotoMergerPage() {
	posList := [][]int{{3822, 58, 1}, {3566, 190, 1}, {2072, 95, 1}, {2135, 394, 2}, {3613, 121, 2}, {3459, 213, 2}, {2291, 345, 2}, {2583, 298, 2}}
	for i, pos := range posList {
		if !Switch.IsTaskOpen() {
			break
		}
		robotgo.MoveClick(common.To100(pos[0], pos[1]))
		x, y := robotgo.GetMousePos()
		fmt.Println(i ,"mleft pos:", x, y)
		robotgo.Sleep(pos[2])
	}
}

func RecordClickPosition()  {
	//chMLeft := make(chan bool)
	cnt := 0
	posList := [][]int{}

	for Switch.IsTaskOpen() {
		time.Sleep(100 * time.Millisecond)
		robotgo.AddEvent("mleft")
		x, y := robotgo.GetMousePos()
		fmt.Println(cnt, "mleft pos:", x, y)
		cnt++
		posList = append(posList, []int{x, y})
	}

	fmt.Println(posList)
}