package main

import(
    "fmt"
    "os"
    "os/exec"
)

// docker run <container> command args
// go run main.go run command args

func main() {
   switch os.Args[1] {
    case "run":
        run()
    default:
       panic("Oops!")
    }
}

func run() {
   fmt.Printf("running %v\n", os.Args[2:])
   cmd := exec.Command(os.Args[2], os.Args[3:]...)
   cmd.Stdin = os.Stdin
   cmd.Stderr = os.Stderr
   cmd.Stdout = os.Stdout
   must(cmd.Run())
}

func must(err error) {
  if err != nil {
     panic(err)
  }
}
