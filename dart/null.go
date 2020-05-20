package dart

import (
  "strings"
  "github.com/cfretz244/godart/cdart"
)

type NullHeap struct {}

type NullBuffer struct {}

func nullFromPacket(pkt *cdart.Packet) *NullBuffer {
  if !pkt.IsNull() {
    panic("Native packet of unexpected type passed to NullBuffer converter")
  }
  return &NullBuffer{}
}

func (num *NullHeap) ctype() *cdart.Packet {
  return nil
}

func (num *NullHeap) IsObject() bool {
  return false
}

func (num *NullHeap) IsArray() bool {
  return false
}

func (num *NullHeap) IsString() bool {
  return false
}

func (num *NullHeap) IsInteger() bool {
  return false
}

func (num *NullHeap) IsDecimal() bool {
  return false
}

func (num *NullHeap) IsBoolean() bool {
  return false
}

func (num *NullHeap) IsNull() bool {
  return true
}

func (num *NullHeap) IsFinalized() bool {
  return false
}

func (num *NullHeap) GetType() int {
  return cdart.NullType
}

func (num *NullHeap) Refcount() uint64 {
  return 0
}

func (num *NullHeap) Equal(other *NullHeap) bool {
  return true
}

func (num *NullHeap) toJSON(out *strings.Builder) {
  out.WriteString("null")
}

func (num *NullHeap) ToJSON() string {
  var builder strings.Builder
  num.toJSON(&builder)
  return builder.String()
}

func (num *NullBuffer) ctype() *cdart.Packet {
  return nil
}

func (num *NullBuffer) IsObject() bool {
  return false
}

func (num *NullBuffer) IsArray() bool {
  return false
}

func (num *NullBuffer) IsString() bool {
  return false
}

func (num *NullBuffer) IsInteger() bool {
  return false
}

func (num *NullBuffer) IsDecimal() bool {
  return false
}

func (num *NullBuffer) IsBoolean() bool {
  return false
}

func (num *NullBuffer) IsNull() bool {
  return true
}

func (num *NullBuffer) IsFinalized() bool {
  return true
}

func (num *NullBuffer) GetType() int {
  return cdart.NullType
}

func (num *NullBuffer) Refcount() uint64 {
  return 0
}

func (num *NullBuffer) Equal(other *NullBuffer) bool {
  return true
}

func (num *NullBuffer) toJSON(out *strings.Builder) {
  out.WriteString(num.ToJSON())
}

func (num *NullBuffer) ToJSON() string {
  return "null"
}
