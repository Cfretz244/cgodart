package dart

import (
  "fmt"
  "strings"
  "github.com/cfretz244/godart/cdart"
)

type DecimalHeap struct {
  contents float64
}

type DecimalBuffer struct {
  native *cdart.Packet
  json string
  val float64
  set bool
}

func dcmFromPacket(pkt *cdart.Packet) *DecimalBuffer {
  if !pkt.IsDecimal() {
    panic("Native packet of unexpected type passed to DecimalBuffer converter")
  } else if !pkt.IsFinalized() {
    panic("Non-finalized decimal passed to DecimalBuffer converter")
  }
  return &DecimalBuffer{pkt, "", 0.0, false}
}

func (num *DecimalHeap) Decimal() float64 {
  return num.Value()
}

func (num *DecimalBuffer) Decimal() float64 {
  return num.Value()
}

func (num *DecimalHeap) Value() float64 {
  return num.contents
}

func (num *DecimalBuffer) Value() float64 {
  if num.set {
    return num.val
  } else if num.native != nil {
    // Load and verify
    val, err := num.native.Decimal()
    errCheck(err, "decimal")

    // Cache and return
    num.val = val
    num.set = true
    return num.val
  } else {
    return 0.0
  }
}

func (num *DecimalHeap) ctype() *cdart.Packet {
  return nil
}

func (num *DecimalHeap) IsObject() bool {
  return false
}

func (num *DecimalHeap) IsArray() bool {
  return false
}

func (num *DecimalHeap) IsString() bool {
  return false
}

func (num *DecimalHeap) IsInteger() bool {
  return false
}

func (num *DecimalHeap) IsDecimal() bool {
  return true
}

func (num *DecimalHeap) IsBoolean() bool {
  return false
}

func (num *DecimalHeap) IsNull() bool {
  return false
}

func (num *DecimalHeap) IsFinalized() bool {
  return false
}

func (num *DecimalHeap) GetType() int {
  return cdart.DecimalType
}

func (num *DecimalHeap) Refcount() uint64 {
  return 1
}

func (num *DecimalBuffer) ctype() *cdart.Packet {
  return num.native
}

func (num *DecimalBuffer) IsObject() bool {
  return false
}

func (num *DecimalBuffer) IsArray() bool {
  return false
}

func (num *DecimalBuffer) IsString() bool {
  return false
}

func (num *DecimalBuffer) IsInteger() bool {
  return false
}

func (num *DecimalBuffer) IsDecimal() bool {
  return true
}

func (num *DecimalBuffer) IsBoolean() bool {
  return false
}

func (num *DecimalBuffer) IsNull() bool {
  return false
}

func (num *DecimalBuffer) IsFinalized() bool {
  return true
}

func (num *DecimalBuffer) GetType() int {
  return cdart.DecimalType
}

func (num *DecimalBuffer) Refcount() uint64 {
  return num.native.Refcount()
}

func (num *DecimalBuffer) Equal(other *DecimalBuffer) bool {
  // Calling into native extensions will definitely be more expensive
  // than the comparison itself, so use the cache if we can
  return num.Value() == other.Value()
}

func (num *DecimalHeap) toJSON(out *strings.Builder) {
  fmt.Fprintf(out, "%f", num.Value())
}

func (num *DecimalHeap) ToJSON() string {
  var builder strings.Builder
  num.toJSON(&builder)
  return builder.String()
}

func (num *DecimalBuffer) toJSON(out *strings.Builder) {
  out.WriteString(num.ToJSON())
}

func (num *DecimalBuffer) ToJSON() string {
  if len(num.json) > 0 {
    // We've already generated JSON previously
    // Buffers are immutable, so just return it.
    return num.json
  } else if num.native != nil {
    // We haven't generated our JSON before, but we
    // have a native representation, so do it.
    json, err := num.native.ToJSON()
    errCheck(err, "decimal")

    num.json = json
    return num.json
  } else {
    // We're a default initialized numuct
    // Just return a static string
    return "0.0"
  }
}
