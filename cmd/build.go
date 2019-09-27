package main

import (
  "encoding/json"
  "flag"
  "io/ioutil"
  "log"
  "os"

  "github.com/eoconnell/shell/schema"
  "github.com/eoconnell/shell/script"
  "github.com/eoconnell/shell/shell"
)

func main() {

  var file = flag.String("file", "config.json", "config file")
  flag.Parse()
  config := configFromFile(*file)

  var s script.Script
  if config.Build.Language == "python" {
    s = script.NewPythonScript(config)
  } else if config.Build.Language == "bash" {
    s = script.NewBashScript(config)
  } else {
    s = nil
  }

  script.Run(config, s)
  shell.Generate(s.Shell(), os.Stdout)
}

func configFromFile(file string) schema.Config {
    data, err := ioutil.ReadFile(file)
    if err != nil {
      log.Fatal("error reading file", err)
    }
    config := schema.Config{}
    json.Unmarshal([]byte(data), &config)
    return config
}
