package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "os"

  "github.com/eoconnell/shell/shell"
)

func main() {
  functions := []string { "travis_cmd.bash" }

  var example string
  if len(os.Args) > 1 {
    example = os.Args[1]
  } else {
    log.Fatal("missing example argument")
  }
  sh := shell.NewShell()

  if (example == "bash") {
    sh.Export("FOO", "\"this is a variable\"")
    sh.Raw("echo \"$FOO\"")
    sh.If("-f myscript", func() {
      sh.Cmd("./myscript")
    })
  }
  if (example == "python3") {
    sh.Cmd("python3 -V")
    sh.Cmd("pip3 -V")
    sh.If("-f requirements.txt", func() {
      sh.Cmd("echo 'installing dependencies'")
      sh.Cmd("pip3 install -r requirements.txt")
    })
  }

  fmt.Println(sh)
  fmt.Println()

  file, err := os.Create("./examples/" + example + "/build.sh")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  defer file.Close()

  for _, filename := range functions {
    data, err := ioutil.ReadFile("./functions/"+filename)
    if err != nil {
      log.Fatal("error reading file", err)
    }
    file.Write(data)
  }

  //shell.Generate(sh, os.Stdout)
  shell.Generate(sh, file)
}
