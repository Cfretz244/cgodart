package main

import "fmt"
import "github.com/cfretz244/godart/dart"

func main() {
  pkt, _ := dart.NewPacket()
  fmt.Println(pkt.ToJSON())
}
