package main

import (
	"fmt"
	"github.com/autopp/plugintestmaster/logger"
	"os"
	"plugin"
)

type mylogger struct{}

func (*mylogger) Log(args ...interface{}) {
	fmt.Print("LOG: ")
	fmt.Println(args...)
}

func main() {
	l := &mylogger{}
	for _, path := range os.Args[1:] {
		l.Log("run", path)
		p, e := plugin.Open(path)
		if e != nil {
			l.Log("cannot load", path)
			continue
		}
		s, e := p.Lookup("Run")
		if e != nil {
			l.Log("cannot lookup Run from", path)
		}
		f := s.(func(logger.Logger))
		f(l)
	}
}
