package shell

import "fmt"

type Export struct {
  key string
  value string
}

func (e Export) handle(g Generator) {
  g.HandleExport(e)
}

func (e Export) String() string {
  return fmt.Sprintf("<export{%s %s}>", e.key, e.value)
}
