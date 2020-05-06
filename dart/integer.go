package dart

import (
  "github.com/cfretz244/godart/cdart"
)

type IntegerBuffer struct {
  native *cdart.Packet
  json string
  val int64
  set bool
}

func intFromPacket(pkt *cdart.Packet) *IntegerBuffer {
  if !pkt.IsInteger() {
    panic("Native packet of unexpected type passed to IntegerBuffer converter")
  } else if !pkt.IsFinalized() {
    panic("Non-finalized integer passed to IntegerBuffer converter")
  }
  return &IntegerBuffer{pkt, "", 0, false}
}

func (num *IntegerBuffer) ctype() *cdart.Packet {
  return num.native
}

func (num *IntegerBuffer) IsObject() bool {
  return false
}

func (num *IntegerBuffer) IsArray() bool {
  return false
}

func (num *IntegerBuffer) IsString() bool {
  return false
}

func (num *IntegerBuffer) IsInteger() bool {
  return true
}

func (num *IntegerBuffer) IsDecimal() bool {
  return false
}

func (num *IntegerBuffer) IsBoolean() bool {
  return false
}

func (num *IntegerBuffer) IsNull() bool {
  return false
}

func (num *IntegerBuffer) IsFinalized() bool {
  return true
}

func (num *IntegerBuffer) GetType() int {
  return cdart.IntegerType
}

func (num *IntegerBuffer) Refcount() uint64 {
  return num.native.Refcount()
}

func (num *IntegerBuffer) equal(other wrapper) bool {
  return false
}

func (num *IntegerBuffer) ToJSON() string {
  if len(num.json) > 0 {
    // We've already generated JSON previously
    // Buffers are immutable, so just return it.
    return num.json
  } else if num.native != nil {
    // We haven't generated our JSON before, but we
    // have a native representation, so do it.
    json, err := num.native.ToJSON()
    errCheck(err, "integer")

    num.json = json
    return num.json
  } else {
    // We're a default initialized numuct
    // Just return a static string
    return "0"
  }
}
