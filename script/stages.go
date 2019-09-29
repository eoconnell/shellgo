package script

import (
  "github.com/eoconnell/shell/schema"
  "reflect"
  "strings"
)

var stages = []string {
  "announce",
  "before_install",
  "install",
  "before_script",
  "script",
  "after_success",
  "after_failure" }

func Stages(config schema.Config, script Script) {
  sh := script.Shell()

  sh.Raw("# START_FUNCS")
  for _, stage := range(stages) {
    sh.Raw("function travis_run_"+stage+"() {")
    // addons before_{stage}

    cmds := customCommands(config, stage)
    if len(cmds) > 0 {
      // custom stage
      for _, cmd := range(cmds) {
        sh.Cmd(cmd)
        if stage == "script" {
          sh.Raw("travis_result $?")
        }
      }
    } else {
      // builtin stage
      method := reflect.ValueOf(script).MethodByName(goName(stage))
      if method.IsValid() {
        method.Call([]reflect.Value{})
      } else {
        sh.NoOp()
      }
    }

    // addons after_{stage}
    sh.Raw("}")
    sh.NewLine()
  }
  sh.Raw("# END_FUNCS")

  sh.NewLine()

	sh.Raw("travis_setup_env")
  for _, stage := range(stages) {
    sh.Raw("travis_run_"+stage)
  }

  sh.NewLine()
}

func customCommands(config schema.Config, stage string) []string {
    field := reflect.ValueOf(config.Build).FieldByName(goName(stage))
    if field.IsValid() {
      return field.Interface().([]string)
    } else {
      return []string{}
    }
}

func goName(stage string) string {
  s := strings.ReplaceAll(stage, "_", " ")
  s = strings.Title(s)
  s = strings.ReplaceAll(s, " ", "")
  return s
}

