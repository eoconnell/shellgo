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

/////////////////////////////

type Commands struct {
  nodes []Handler
  block func(*Shell)
}

/////////////////////////////

type Cmd struct {
  command string
}

func (e Cmd) handle(g Generator) {
  g.HandleCmd(e)
}

/////////////////////////////

type If struct {
  condition string
  block func(*Shell)
  branches []Commands
}

func (e If) handle(g Generator) {
  g.HandleIf(e)
}

/////////////////////////////

type Export struct {
  key string
  value string
}

func (e Export) handle(g Generator) {
  g.HandleExport(e)
}

/////////////////////////////

type Shell struct {
  body []string
  nodes []Handler
}

func (sh *Shell) cmd(command string) {
  sh.Node(&Cmd{command})
}

func (sh *Shell) if_(condition string, block func(*Shell)) {
  sh.Node(&If{condition: condition, block: block})
}

func (sh *Shell) export(key, value string) {
  sh.Node(&Export{key, value})
}

func (sh *Shell) Node(n Handler) {
  sh.nodes = append(sh.nodes, n)
}

func (sh Shell) String() string {
  return fmt.Sprintf("Shell{%s}", sh.body)
}

/////////////////////////////

type BashGenerator struct {
  out io.Writer
}

func (g BashGenerator) HandleCmd(node Cmd) {
  g.Write("macie_cmd " + shellescape(node.command))
}

func (g BashGenerator) HandleIf(node If) {
  g.Write("if [[ " + node.condition + " ]]; then")
  g.Write("fi")
}

func (g BashGenerator) HandleExport(node Export) {
  g.Write(fmt.Sprintf("export %s=%s", node.key, node.value))
}

func (g BashGenerator) Write(str string) {
  g.out.Write([]byte(str + "\n"))
}

/////////////////////////////

func Generate(sh Shell, writer io.Writer) {
  fmt.Println("#!/bin/bash")
  g := BashGenerator{writer}
  for _, node := range sh.nodes {
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
  sh := Shell{}
  sh.export("FOO", "bar")
  sh.cmd("pip install -r requirements.txt")
  sh.if_("! -f /opt/virtualenv/python3/bin/activate", func(sh *Shell) {
    sh.cmd("terminate")
  })

  fmt.Println(sh)
  fmt.Println()

  Generate(sh, os.Stdout)
}
