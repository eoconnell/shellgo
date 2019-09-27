package shell

import "fmt"

type If struct {
  condition string
  nodes []Handler
  branches []Commands
}

func NewIf(condition string, block func(Commands)) *If {
  n := &If{condition: condition}
  block(n)
  return n
}

func (e *If) Nodes() []Handler {
  return e.nodes
}

func (e *If) Insert(h Handler) {
  e.nodes = append(e.nodes, h)
}

func (e If) handle(g Generator) {
  g.HandleIf(e)
}

func (e If) String() string {
  return fmt.Sprintf("<if{%s then=%s}>", e.condition, e.nodes)
}

