package script

import "testing"

import "github.com/eoconnell/shell/schema"

func TestGoName(t *testing.T) {
  result := goName("some_name")
  if result != "SomeName" {
    t.Errorf("Expected `SomeName` but got `%s`", result)
  }
}

func TestCustomCommands(t *testing.T) {
  config := schema.Config{}
  config.Build.Install = []string{"some_command"}

  result := customCommands(config, "install")

  if len(result) != 1 {
    t.Errorf("Expected size `[1]` but got size `[%d]`", len(result))
  }
  if result[0] != "some_command" {
    t.Errorf("Expected `some_command` but got `%s`", result[0])
  }
}

func TestCustomCommands_UnknownStage(t *testing.T) {
  config := schema.Config{}

  result := customCommands(config, "unknown_stage")

  if len(result) != 0 {
    t.Errorf("Expected empty `[]` but got size `[%d]`", len(result))
  }
}
