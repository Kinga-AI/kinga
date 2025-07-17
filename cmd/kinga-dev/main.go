package main

import (
  "fmt"
  "os"
  corelang "github.com/kinga-ai/kinga-core/lang/go"
)

func main() {
  fmt.Println("Kinga dev bootstrap")
  fmt.Println("Languages:", corelang.Codes())
  endpoint := os.Getenv("KWELI_ENDPOINT")
  fmt.Println("Kweli endpoint:", endpoint)
}
