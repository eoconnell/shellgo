package shell

import(
  "os"
)

func ExampleGenerate_Raw() {
  sh := NewShell()
  sh.Raw("exit 1")

  Generate(sh, os.Stdout)

  // Output:
  // exit 1
}

func ExampleGenerate_Cmd() {
  sh := NewShell()
  sh.Cmd("python --version")

  Generate(sh, os.Stdout)

  // Output:
  // travis_cmd python\ --version
}

func ExampleGenerate_Export() {
  sh := NewShell()
  sh.Export("KEY", "value")

  Generate(sh, os.Stdout)

  // Output:
  // export KEY=value
}

func ExampleGenerate_If() {
  sh := NewShell()
  sh.If("-e requirements.txt", func() {
    sh.Cmd("pip install -r requirements.txt")
  })

  Generate(sh, os.Stdout)

  // Output:
  // if [[ -e requirements.txt ]]; then
  //     travis_cmd pip\ install\ -r\ requirements.txt
  // fi
}

func ExampleGenerate_NoOp() {
  sh := NewShell()
  sh.NoOp()

  Generate(sh, os.Stdout)

  // Output:
  // :
}

func ExampleGenerate_NewLine() {
  sh := NewShell()
  sh.NewLine()

  Generate(sh, os.Stdout)

  // Output:
  //

}
