package script

import (
  "io"
  "io/ioutil"
  "log"

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

var functions = []string {
  "travis_setup_env.bash",
  "travis_cmd.bash",
  "travis_result.bash" }

func Compile(config schema.Config, out io.Writer) {
  s := byLang(config)

  sh := s.Shell()
  sh.Raw("#!/bin/bash")

  for _, filename := range functions {
    data, err := ioutil.ReadFile("functions/"+filename)
    if err != nil {
      log.Fatal("error reading file", err)
    }
    sh.Raw(string(data))
  }

  Stages(config, s)
  shell.Generate(sh, out)
}

func byLang(config schema.Config) Script {
  var s Script
  switch lang := config.Build.Language; lang {
  case "python":
    s = NewPython(config)
  case "bash":
    s = NewBash(config)
  case "java":
    s = NewJava(config)
  default:
    s = nil
  }
  return s
}
