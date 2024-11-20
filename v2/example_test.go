package v2

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println(0 / 1)

	robotgo.Scroll(0, -10)
	robotgo.Scroll(100, 0)

	robotgo.MilliSleep(100)

	robotgo.Move(10, 20)
	robotgo.MoveRelative(0, -10)
	robotgo.DragSmooth(10, 10)

	robotgo.Click("wheelRight")
	robotgo.Click("left", true)
	robotgo.MoveSmooth(100, 200, 1.0, 10.0)
}
