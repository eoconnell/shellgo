package shell

import "fmt"

type Cmds struct {
  nodes []Handler
  block func(*Commands)
}

func (n *Cmds) Nodes() []Handler {
  return n.nodes
}

func (n *Cmds) Insert(h Handler) {
  n.nodes = append(n.nodes, h)
}

func (n Cmds) handle(g Generator) {
  for _, node := range n.nodes {
    node.handle(g)
  }
}

func (n Cmds) String() string {
  return fmt.Sprintf("Cmds{%s}", n.nodes)
}
