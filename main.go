package main

import (
	"os"
	"plugin"
  "github.com/autopp/plugintestmaster/logger"
)

func main() {
  l := &logger.Logger{}
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
    f := s.(func (*logger.Logger))
    f(l)
	}
}
