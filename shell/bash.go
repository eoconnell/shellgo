package shell

import(
  "fmt"
  "io"
)

type BashGenerator struct {
  out io.Writer
  level int
}

func (g BashGenerator) HandleCmd(node Cmd) {
  if node.raw {
    g.Writeln(node.command)
  } else {
    g.Writeln("travis_cmd " + shellescape(node.command))
  }
}

func (g BashGenerator) HandleIf(node If) {
  g.Writeln("if [[ " + node.condition + " ]]; then")
  g.indent(func() {
    for _, n := range node.nodes {
      n.handle(g)
    }
  })
  g.Writeln("fi")
}

func (g BashGenerator) HandleExport(node Export) {
  g.Writeln(fmt.Sprintf("export %s=%s", node.key, node.value))
}

func (g BashGenerator) Writeln(str string) {
  for i := 0; i < g.level; i++ {
    g.Write("    ")
  }
  g.Write(str + "\n")
}

func (g BashGenerator) Write(str string) {
  g.out.Write([]byte(str))
}

func (g *BashGenerator) indent(block func()) {
  g.level++
  block()
  g.level--
}

/////////////////////////////

func Generate(sh *Shell, writer io.Writer) {
  g := BashGenerator{writer, 0}
  for _, node := range sh.stack {
    node.handle(g)
  }
}
