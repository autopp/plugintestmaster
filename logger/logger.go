package logger

import (
	"fmt"
)

type Logger struct{}

func (*Logger) Log(args ...interface{}) {
  fmt.Print("LOG: ")
	fmt.Println(args...)
}
