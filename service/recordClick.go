package service

import (
	"fmt"
	"github.com/a316523235/wingo/common"
	"github.com/go-vgo/robotgo"
	"github.com/robotn/gohook"
	"strings"
)

var ClickPosList = [][]int{}

// Deprecated: use the RecordClickPositionV2(),
//
// RecordClickPosition record left mouse click position
func RecordClickPosition()  {
	chMLeft := make(chan bool, 100)
	cnt := 0
	posList := [][]int{}

	go func() {
		for Switch.IsTaskOpen() {
			robotgo.MilliSleep(100)
			chMLeft <- robotgo.AddEvent("mleft")
			robotgo.MilliSleep(200)
		}
	}()

	for {
		select {
		case <-chMLeft:
			x, y := robotgo.GetMousePos()
			cnt++
			fmt.Println(cnt, "mleft pos:", x, y)
			posList = append(posList, []int{x, y})
			if !Switch.IsTaskOpen() || cnt > 10 {
				fmt.Println(posList)
				break
			}
		}
	}
}


// RecordClickPositionV2 record left mouse click position
func RecordClickPositionV2() {
	cnt := 0
	posList := [][]int{}

	same := 1

	for i := 0; i < 100; i++ {
		robotgo.AddEvent("mleft")
		robotgo.MilliSleep(200)
		x, y := robotgo.GetMousePos()
		cnt++
		fmt.Println(cnt, "mleft pos:", x, y)
		posList = append(posList, []int{x, y, 1})	//1 is sleep time

		// if left mouse click same position greater or equal three times, break
		if len(posList) > 1 && posList[len(posList) -1][0] == posList[len(posList) -2][0] && posList[len(posList) -1][1] == posList[len(posList) -2][1] {
			same++
		} else {
			same = 1
		}
		if same >= 3 {
			break
		}
	}
	posStrList := []string{}
	for _, pos := range posList {
		posStrList = append(posStrList, "{" + common.IntJoin(pos) +  "}")
	}
	fmt.Println("{" + strings.Join(posStrList, ",") + "}")
}

func RecordClickPositionV3() {
	cnt := 0
	robotgo.EventHook(hook.MouseDown, []string{}, func(e hook.Event) {
		if e.Button == hook.MouseMap["left"] {
			cnt++
			x, y := robotgo.GetMousePos()
			fmt.Println(cnt, "mleft pos:", x, y)
			ClickPosList = append(ClickPosList, []int{x, y, 2})	//1 is sleep time
		}
	})

	robotgo.EventHook(hook.KeyDown, []string{"enter"}, func(e hook.Event) {
		PrintPos(false)
	})
}

// PrintPos print and copy click position
// isClear is clear position list
func PrintPos(isClear bool)  {
	if len(ClickPosList) > 0 {
		posStrList := []string{}
		for _, pos := range ClickPosList {
			posStrList = append(posStrList, "{" + common.IntJoin(pos) +  "}")
		}
		res := "{" + strings.Join(posStrList, ",") + "}"
		fmt.Println("break record")
		fmt.Println(res)

		_ = robotgo.WriteAll(res)

		if isClear {
			ClickPosList = [][]int{}
		}
	}
}