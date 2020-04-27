package main

import "fmt"
import "github.com/cfretz244/godart/cdart"

func main() {
  pkt, _ := cdart.NewArrayPacket()
  str, _ := cdart.NewStringPacket("hello")
  str2, _ := cdart.NewStringPacket("world")
  pkt.InsertIndex(0, str)
  pkt.InsertIndex(1, str2)
  fmt.Println(pkt.ToJSON())

  _, err := pkt.ToBytes()
  if err == nil {
    panic("Bytes were received for non-finalized array!")
  }
  fmt.Println(err.Error())
  obj, _ := cdart.NewObjectPacket()
  obj.InsertField("arr", pkt)
  obj.Finalize()

  bytes, _ := obj.ToBytes()
  if bytes == nil {
    panic("Bytes were not received for finalzed object!")
  }
  fmt.Println("Byte slice is", len(bytes), "long")

  duppkt, _ := cdart.FromBytes(bytes)
  json, _ := duppkt.ToJSON()
  fmt.Println("Reconstructed packet: ", json)

  it, _ := cdart.NewIterator(pkt)
  for it.Next() {
    val, _ := it.Value()
    fmt.Println(val.ToJSON())
  }
}
