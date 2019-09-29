package script

import (
  "github.com/eoconnell/shell/schema"
  "github.com/eoconnell/shell/shell"
)

func NewBash(config schema.Config) *Bash {
  script := Bash{}
  script.config = config
  shell := shell.NewShell()
  script.sh = shell
  return &script
}

type Bash struct {
  sh *shell.Shell
  config schema.Config
}

func (self Bash) Shell() *shell.Shell {
  return self.sh
}

func (self Bash) Setup() { self.sh.NoOp() }

func (self Bash) Announce() {
  self.sh.Cmd("bash --version")
}

func (self Bash) BeforeInstall() { self.sh.NoOp() }

func (self Bash) Install() { self.sh.NoOp() }

func (self Bash) BeforeScript() { self.sh.NoOp() }

func (self Bash) Script() {
  self.sh.Cmd("Please override the script: key")
  self.sh.Cmd("exit 2")
}
