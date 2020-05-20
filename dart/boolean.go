package dart

import (
  "strings"
  "github.com/cfretz244/godart/cdart"
)

type BooleanHeap struct {
  contents bool
}

type BooleanBuffer struct {
  native *cdart.Packet
  json string
  val, set bool
}

func boolFromPacket(pkt *cdart.Packet) *BooleanBuffer {
  if !pkt.IsBoolean() {
    panic("Native packet of unexpected type passed to BooleanBuffer converter")
  } else if !pkt.IsFinalized() {
    panic("Non-finalized boolean passed to BooleanBuffer converter")
  }
  return &BooleanBuffer{pkt, "", false, false}
}

func (num *BooleanHeap) Boolean() bool {
  return num.Value()
}

func (num *BooleanBuffer) Boolean() bool {
  return num.Value()
}

func (num *BooleanHeap) Value() bool {
  return num.contents
}

func (num *BooleanBuffer) Value() bool {
  if num.set {
    return num.val
  } else if num.native != nil {
    // Load and verify
    val, err := num.native.Boolean()
    errCheck(err, "boolean")

    // Cache and return
    num.val = val
    num.set = true
    return num.val
  } else {
    return false
  }
}

func (num *BooleanHeap) ctype() *cdart.Packet {
  return nil
}

func (num *BooleanHeap) IsObject() bool {
  return false
}

func (num *BooleanHeap) IsArray() bool {
  return false
}

func (num *BooleanHeap) IsString() bool {
  return false
}

func (num *BooleanHeap) IsInteger() bool {
  return false
}

func (num *BooleanHeap) IsDecimal() bool {
  return false
}

func (num *BooleanHeap) IsBoolean() bool {
  return true
}

func (num *BooleanHeap) IsNull() bool {
  return false
}

func (num *BooleanHeap) IsFinalized() bool {
  return false
}

func (num *BooleanHeap) GetType() int {
  return cdart.BooleanType
}

func (num *BooleanHeap) Refcount() uint64 {
  return 1
}

func (num *BooleanBuffer) ctype() *cdart.Packet {
  return num.native
}

func (num *BooleanBuffer) IsObject() bool {
  return false
}

func (num *BooleanBuffer) IsArray() bool {
  return false
}

func (num *BooleanBuffer) IsString() bool {
  return false
}

func (num *BooleanBuffer) IsInteger() bool {
  return false
}

func (num *BooleanBuffer) IsDecimal() bool {
  return false
}

func (num *BooleanBuffer) IsBoolean() bool {
  return true
}

func (num *BooleanBuffer) IsNull() bool {
  return false
}

func (num *BooleanBuffer) IsFinalized() bool {
  return true
}

func (num *BooleanBuffer) GetType() int {
  return cdart.BooleanType
}

func (num *BooleanBuffer) Refcount() uint64 {
  return num.native.Refcount()
}

func (num *BooleanBuffer) Equal(other *BooleanBuffer) bool {
  // Calling into native extensions will definitely be more expensive
  // than the comparison itself, so use the cache if we can
  return num.Value() == other.Value()
}

func (num *BooleanHeap) toJSON(out *strings.Builder) {
  if num.Value() {
    out.WriteString("true")
  } else {
    out.WriteString("false")
  }
}

func (num *BooleanHeap) ToJSON() string {
  var builder strings.Builder
  num.toJSON(&builder)
  return builder.String()
}

func (num *BooleanBuffer) toJSON(out *strings.Builder) {
  out.WriteString(num.ToJSON())
}

func (num *BooleanBuffer) ToJSON() string {
  if len(num.json) > 0 {
    // We've already generated JSON previously
    // Buffers are immutable, so just return it.
    return num.json
  } else if num.native != nil {
    // We haven't generated our JSON before, but we
    // have a native representation, so do it.
    json, err := num.native.ToJSON()
    errCheck(err, "boolean")

    num.json = json
    return num.json
  } else {
    // We're a default initialized numuct
    // Just return a static string
    return "false"
  }
}
