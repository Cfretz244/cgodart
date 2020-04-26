package main

import "fmt"
import "github.com/cfretz244/godart/dart"

func main() {
  pkt, _ := dart.NewArrayPacket()
  str, _ := dart.NewStringPacket("hello")
  str2, _ := dart.NewStringPacket("world")
  pkt.InsertIndex(0, str)
  pkt.InsertIndex(1, str2)
  fmt.Println(pkt.ToJSON())

  it, _ := dart.NewIterator(pkt)
  for it.Next() {
    val, _ := it.Value()
    fmt.Println(val.ToJSON())
  }
}
