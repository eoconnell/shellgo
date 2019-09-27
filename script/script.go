package script

import (
  "github.com/eoconnell/shell/schema"
  "github.com/eoconnell/shell/shell"
)

type Script interface {
  Setup()
  Announce()
  BeforeInstall()
  Install()
  BeforeScript()
  Script()
  Shell() *shell.Shell
}

func BuiltinPhase(sh *shell.Shell, user []string, script func()) {
  if len(user) != 0 {
    for _, command := range user {
      sh.Cmd(command)
    }
  } else {
    script()
  }
}

func Run(config schema.Config, script Script) {
  sh := script.Shell()
  script.Setup()
  script.Announce()
  script.BeforeInstall()
  BuiltinPhase(sh, config.Build.Install, script.Install)
  script.BeforeScript()
  BuiltinPhase(sh, config.Build.Script, script.Script)
}

