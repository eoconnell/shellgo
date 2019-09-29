package main

import (
  "encoding/json"
  "flag"
  "io/ioutil"
  "log"
  "os"

  "github.com/eoconnell/shell/schema"
  "github.com/eoconnell/shell/script"
)

func main() {

  var file = flag.String("file", "config.json", "config file")
  flag.Parse()
  config := configFromFile(*file)


  script.Compile(config, os.Stdout)
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
