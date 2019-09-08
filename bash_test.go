package main

import "testing"

import (
  "bytes"
  "bufio"
)

func TestGenerateCmd(t *testing.T) {
  var buf bytes.Buffer
  out := bufio.NewWriter(&buf)
  sh := Shell{}
  sh.cmd("python --version")

  Generate(sh, out)

  out.Flush()
  expected := "travis_cmd python\\\\ --version\n"
  actual := buf.String()
  if actual != expected {
    t.Errorf("Expected `%s` but got `%s`", expected, actual)
  }
}
