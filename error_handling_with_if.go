package main

import (
  "fmt"
  "os"
)

func main() {
  // error is a standard type in Go
  // If the function passes, the error code is nil.
  _, err := os.Open("/Users/z0019z3/go/src/testing/error_handling_with_if.go")
  // "_" means ignore

  if err != nil {
    fmt.Println("Error returned was: ", err)
  // If the file doesn't exist, it'll return "no such file or directory"
  }
}
