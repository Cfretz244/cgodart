package dart

import (
  "github.com/cfretz244/godart/cdart"
)

type NullBuffer struct {}

func nullFromPacket(pkt *cdart.Packet) *NullBuffer {
  if !pkt.IsNull() {
    panic("Native packet of unexpected type passed to NullBuffer converter")
  }
  return &NullBuffer{}
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

func (num *NullBuffer) ToJSON() string {
  return "null"
}
