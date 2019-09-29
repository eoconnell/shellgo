package shell

import(
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

type Shell struct {
  body []string
  nodes []Handler
  stack []Commands
}

func NewShell() *Shell {
  sh := Shell{}
  sh.stack = append(sh.stack, &Cmds{})
  return &sh
}

func (sh *Shell) Cmd(command string) {
  n := &Cmd{command, false}
  sh.last().Insert(n)
  sh.node(n)
}

func (sh *Shell) Raw(command string) {
  n := &Cmd{command, true}
  sh.last().Insert(n)
  sh.node(n)
}

func (sh *Shell) NewLine() {
  sh.Raw("")
}

func (sh *Shell) NoOp() {
  sh.Raw(":")
}

func (sh *Shell) If(condition string, block func()) {
  b := sh.withNode(block)
  n := NewIf(condition, b)
  sh.last().Insert(n)
  sh.node(n)
}

func (sh *Shell) Export(key, value string) {
  n := &Export{key, value}
  sh.last().Insert(n)
  sh.node(n)
}

func (sh *Shell) node(n Handler) {
  sh.nodes = append(sh.nodes, n)
}

func (sh *Shell) withNode(block func()) func(Commands) {
  return func(n Commands) {
    sh.stack = append(sh.stack, n)
    block()
    sh.stack = sh.stack[:len(sh.stack)-1]
  }
}

func (sh *Shell) last() Commands {
  return sh.stack[len(sh.stack)-1]
}

func shellescape(str string) string {
  var pattern = regexp.MustCompile(`([^A-Za-z0-9_\-.,:\/@\n])`)
  return pattern.ReplaceAllString(str, "\\$1")
}
