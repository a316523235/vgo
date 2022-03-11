package example

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"testing"
	"time"
)

func TestGetScaleSize(t *testing.T)  {
	time.Sleep(2 * time.Second)
	width, height := robotgo.GetScaleSize()
	fmt.Println("get scale screen size: ", width, height)

	sx, sy := robotgo.GetScreenSize()
	fmt.Println("get screen size: ", sx, sy)

	// system 1920 1080
	// left 4800 1350    125%
	// right 3840 1080 100%
}