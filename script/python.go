package script

import (
  "github.com/eoconnell/shell/schema"
  "github.com/eoconnell/shell/shell"
)

func NewPythonScript(config schema.Config) *Python {
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

func (py Python) Shell() *shell.Shell {
  return py.sh
}

func (py Python) Setup() {}

func (py Python) Announce() {
  py.sh.Cmd("python3 --version")
  py.sh.Cmd("pip3 --version")
}

func (py Python) BeforeInstall() {}

func (py Python) Install() {
  py.sh.If("-f requirements.txt", func() {
    py.sh.Cmd("pip3 install -r requirements.txt")
  })
}

func (py Python) BeforeScript() {}

func (py Python) Script() {
  py.sh.Cmd("Please override the script: key")
  py.sh.Cmd("exit 2")
}

