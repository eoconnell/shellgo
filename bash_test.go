package main

import (
  "os"
)

func ExampleGenerate_Raw() {
  sh := NewShell()
  sh.raw("exit 1")

  Generate(sh, os.Stdout)

  // Output:
  // exit 1
}

func ExampleGenerate_Cmd() {
  sh := NewShell()
  sh.cmd("python --version")

  Generate(sh, os.Stdout)

  // Output:
  // travis_cmd python\ --version
}

func ExampleGenerate_Export() {
  sh := NewShell()
  sh.export("KEY", "value")

  Generate(sh, os.Stdout)

  // Output:
  // export KEY=value
}

func ExampleGenerate_If() {
  sh := NewShell()
  sh.if_("-e requirements.txt", func() {
    sh.cmd("pip install -r requirements.txt")
  })

  Generate(sh, os.Stdout)

  // Output:
  // if [[ -e requirements.txt ]]; then
  //     travis_cmd pip\ install\ -r\ requirements.txt
  // fi
}
