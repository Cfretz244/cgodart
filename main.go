package main

import "fmt"
import "github.com/cfretz244/godart/dart"

func main() {
  str, _ := dart.NewIntegerPacket(5)
  pkt, _ := dart.FromJSON("{\"hello\":\"world\",\"yes\":\"no\"}")
  pkt.InsertField("hello", str)
  pkt.Finalize()
  if !pkt.IsFinalized() {
    panic("Packet isn't finalized!")
  }

  _, err := str.Size()
  if (err == nil) {
    panic("Integer allowed string call!")
  }
  fmt.Println(err.Error())

  it, _ := dart.NewIterator(pkt)
  keyIt, _ := dart.NewKeyIterator(pkt)
  for it.Next() {
    keyIt.Next()
    val, _ := it.Value()
    key, _ := keyIt.Value()
    fmt.Println(key.ToJSON())
    fmt.Println(val.ToJSON())
  }
}
