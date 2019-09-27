package shell

import "fmt"

type Cmd struct {
  command string
  raw bool
}

func (e Cmd) handle(g Generator) {
  g.HandleCmd(e)
}

func (e Cmd) String() string {
  return fmt.Sprintf("<cmd{%s raw=%v}>", e.command, e.raw)
}

