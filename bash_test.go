package main

import "testing"

import (
  "bytes"
  "bufio"
)

func TestGenerateCmd(t *testing.T) {
  var buf bytes.Buffer
  out := bufio.NewWriter(&buf)
  sh := NewShell()
  sh.cmd("python --version")

  Generate(sh, out)

  out.Flush()
  expected := "travis_cmd python\\\\ --version\n"
  actual := buf.String()
  if actual != expected {
    t.Errorf("Expected `%s` but got `%s`", expected, actual)
  }
}

func TestGenerateExport(t *testing.T) {
  var buf bytes.Buffer
  out := bufio.NewWriter(&buf)
  sh := NewShell()
  sh.export("KEY", "value")

  Generate(sh, out)

  out.Flush()
  expected := "export KEY=value\n"
  actual := buf.String()
  if actual != expected {
    t.Errorf("Expected `%s` but got `%s`", expected, actual)
  }
}

func TestGenerateIf(t *testing.T) {
  var buf bytes.Buffer
  out := bufio.NewWriter(&buf)
  sh := NewShell()
  sh.if_("-e requirements.txt", func() {
    sh.cmd("pip install -r requirements.txt")
  })

  Generate(sh, out)

  out.Flush()
  expected := `if [[ -e requirements.txt ]]; then
travis_cmd pip\\ install\\ -r\\ requirements.txt
fi
`
  actual := buf.String()
  if actual != expected {
    t.Errorf("Expected `%s` but got `%s`", expected, actual)
  }
}
