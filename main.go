package main

import "fmt"
import "github.com/cfretz244/godart/dart"

func main() {
  var obj dart.ObjectHeap
  fmt.Println(obj.ToJSON())
}
