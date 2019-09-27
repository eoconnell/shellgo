package script

import (
  "github.com/eoconnell/shell/schema"
  "github.com/eoconnell/shell/shell"
)

func NewBashScript(config schema.Config) *Bash {
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

func (bash Bash) Shell() *shell.Shell {
  return bash.sh
}

func (bash Bash) Setup() {}

func (bash Bash) Announce() {
  bash.sh.Cmd("bash --version")
}

func (bash Bash) BeforeInstall() {}

func (bash Bash) Install() {}

func (bash Bash) BeforeScript() {}

func (bash Bash) Script() {
  bash.sh.Cmd("Please override the script: key")
  bash.sh.Cmd("exit 2")
}
