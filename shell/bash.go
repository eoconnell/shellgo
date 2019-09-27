package shell

import(
  "fmt"
  "io"
)

type BashGenerator struct {
  out io.Writer
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
    for _, n := range node.nodes {
      g.Write("    ") // lazy indentation
      n.handle(g)
    }
  g.Writeln("fi")
}

func (g BashGenerator) HandleExport(node Export) {
  g.Writeln(fmt.Sprintf("export %s=%s", node.key, node.value))
}

func (g BashGenerator) Writeln(str string) {
  g.Write(str + "\n")
}

func (g BashGenerator) Write(str string) {
  g.out.Write([]byte(str))
}

/////////////////////////////

func Generate(sh *Shell, writer io.Writer) {
  g := BashGenerator{writer}
  for _, node := range sh.stack {
    node.handle(g)
  }
}
