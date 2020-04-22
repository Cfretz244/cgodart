package main

import "fmt"
import "github.com/cfretz244/godart/dart"

func main() {
  pkt := dart.NewObjectPacket()
  fmt.Println(pkt.ToJSON())
}
