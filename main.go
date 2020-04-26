package main

import "fmt"
import "github.com/cfretz244/godart/dart"

func main() {
  pkt, _ := dart.FromJSON("{\"hello\":\"world\"}")
  fmt.Println(pkt.ToJSON())
}
