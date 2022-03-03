package service

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

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
		posList = append(posList, []int{x, y})

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
	fmt.Println(posList)
}
