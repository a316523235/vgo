package example

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"testing"
)

func TestWindow1(t *testing.T) {
	// show Alert Window
	abool := robotgo.ShowAlert("hello", "robotgo")
	if abool {
		fmt.Println("ok@@@", "ok")
	}
	robotgo.ShowAlert("hello", "robotgo", "Ok", "Cancel")
}
