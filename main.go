package main

import "fmt"
import "github.com/cfretz244/godart/dart"

func main() {
  buf, _ := dart.BufferFromJSON("{\"hello\":\"world\"}")
  obj := buf.AsObject()
  fmt.Println(obj.ToJSON())
}
