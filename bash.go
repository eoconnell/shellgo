package main

import (
  "fmt"
  "io"
  "os"
  "regexp"
)

type Generator interface {
  HandleCmd(e Cmd)
  HandleIf(e If)
  HandleExport(e Export)
}

type Handler interface {
  handle(g Generator)
}

type Commands interface {
  Nodes() []Handler
  Insert(h Handler)
  handle(g Generator)
}

/////////////////////////////

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

/////////////////////////////

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

/////////////////////////////

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

/////////////////////////////

type Export struct {
  key string
  value string
}

func (e Export) handle(g Generator) {
  g.HandleExport(e)
}

func (e Export) String() string {
  return fmt.Sprintf("<export{%s %s}>", e.key, e.value)
}

/////////////////////////////

type Shell struct {
  body []string
  nodes []Handler
  stack []Commands
}

func NewShell() Shell {
  sh := Shell{}
  sh.stack = append(sh.stack, &Cmds{})
  return sh
}

func (sh *Shell) cmd(command string) {
  n := &Cmd{command, false}
  sh.last().Insert(n)
  sh.Node(n)
}

func (sh *Shell) raw(command string) {
  n := &Cmd{command, true}
  sh.last().Insert(n)
  sh.Node(n)
}

func (sh *Shell) if_(condition string, block func()) {
  b := sh.WithNode(block)
  n := NewIf(condition, b)
  sh.last().Insert(n)
  sh.Node(n)
}

func (sh *Shell) export(key, value string) {
  n := &Export{key, value}
  sh.last().Insert(n)
  sh.Node(n)
}

func (sh *Shell) Node(n Handler) {
  sh.nodes = append(sh.nodes, n)
}

func (sh *Shell) WithNode(block func()) func(Commands) {
  return func(node Commands) {
    sh.stack = append(sh.stack, node)
    block()
    sh.stack = sh.stack[:len(sh.stack)-1]
  }
}

func (sh *Shell) last() Commands {
  return sh.stack[len(sh.stack)-1]
}

func (sh Shell) String() string {
  return fmt.Sprintf("Shell{%s}", sh.stack)
}

/////////////////////////////

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

func Generate(sh Shell, writer io.Writer) {
  g := BashGenerator{writer}
  for _, node := range sh.stack {
    node.handle(g)
  }
}

/////////////////////////////

func shellescape(str string) string {
  var pattern = regexp.MustCompile(`([^A-Za-z0-9_\-.,:\/@\n])`)
  return pattern.ReplaceAllString(str, "\\\\$1")
}

/////////////////////////////

func main() {
  sh := NewShell()
  sh.export("FOO", "bar")
  sh.cmd("pip install -r requirements.txt")
  sh.if_("! -f /opt/virtualenv/python3/bin/activate", func() {
    sh.cmd("terminate")
  })

  fmt.Println(sh)
  fmt.Println()

  Generate(sh, os.Stdout)
}
