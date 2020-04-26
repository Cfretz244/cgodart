package main

import "fmt"
import "github.com/cfretz244/godart/dart"

func main() {
  str, _ := dart.NewStringPacket("life!")
  pkt, _ := dart.FromJSON("{\"hello\":\"world\",\"yes\":\"no\"}")
  err := pkt.Insert("hello", str)
  if err != nil {
    panic("insert failed")
  }
  pkt.Finalize()
  if !pkt.IsFinalized() {
    panic("Packet isn't finalized!")
  }

  it, iterr := dart.NewIterator(pkt)
  if iterr != nil {
    panic("Failed to initialize iterator")
  }
  for it.Next() {
    val, valerr := it.Value()
    if valerr != nil {
      panic("Failed to grab value")
    }
    fmt.Println(val.ToJSON())
  }
}
