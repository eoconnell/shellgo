package script

import (
  "github.com/eoconnell/shell/schema"
  "github.com/eoconnell/shell/shell"
)

func NewPython(config schema.Config) *Python {
  script := Python{}
  script.config = config
  shell := shell.NewShell()
  script.sh = shell
  return &script
}

type Python struct {
  sh *shell.Shell
  config schema.Config
}

func (self Python) Shell() *shell.Shell {
  return self.sh
}

func (self Python) Setup() { self.sh.NoOp() }

func (self Python) Announce() {
  self.sh.Cmd("python3 --version")
  self.sh.Cmd("pip3 --version")
}

func (self Python) BeforeInstall() { self.sh.NoOp() }

func (self Python) Install() {
  self.sh.If("-f requirements.txt", func() {
    self.sh.Cmd("pip3 install -r requirements.txt")
  })
}

func (self Python) BeforeScript() { self.sh.NoOp() }

func (self Python) Script() {
  self.sh.Cmd("echo 'Please override the script: key'")
  self.sh.Cmd("exit 2")
}

