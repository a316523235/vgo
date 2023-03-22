package example

import (
	"fmt"
	"github.com/a316523235/wingo/conf"
	"strings"
	"testing"
)

func TestReadEn(t *testing.T)  {
	lineArr := strings.Split(conf.EN1, "\n")
	fmt.Println(len(lineArr))
	for i := 0; i < 3; i++ {
		fmt.Println(lineArr[i])
	}
}
