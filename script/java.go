package script

import (
  "github.com/eoconnell/shell/schema"
  "github.com/eoconnell/shell/shell"
)

func NewJava(config schema.Config) *Java {
  script := Java{}
  script.config = config
  shell := shell.NewShell()
  script.sh = shell
  return &script
}

type Java struct {
  sh *shell.Shell
  config schema.Config
}

func (self Java) Shell() *shell.Shell {
  return self.sh
}

func (self Java) Setup() { self.sh.NoOp() }

func (self Java) Announce() {
  self.sh.Cmd("java --version")
  self.sh.Cmd("javac --version")
  self.sh.Cmd("mvn --version")
}

func (self Java) BeforeInstall() { self.sh.NoOp() }

func (self Java) Install() {
  self.sh.If("-f pom.xml", func() {
    self.sh.Cmd("mvn dependency:resolve-plugins -B")
    self.sh.Cmd("mvn dependency:resolve -B")
  })
}

func (self Java) BeforeScript() { self.sh.NoOp() }

func (self Java) Script() {
  self.sh.If("-f pom.xml", func() {
    self.sh.Cmd("mvn test -B")
  })
}
