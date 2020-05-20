package dart

import (
  "strings"
  "github.com/cfretz244/godart/cdart"
)

type StringHeap struct {
  contents string
}

type StringBuffer struct {
  native *cdart.Packet
  json, val string
  set bool
}

func strFromPacket(pkt *cdart.Packet) *StringBuffer {
  if !pkt.IsString() {
    panic("Native packet of unexpected type passed to StringBuffer converter")
  } else if !pkt.IsFinalized() {
    panic("Non-finalized string passed to StringBuffer converter")
  }
  return &StringBuffer{pkt, "", "", false}
}

func (str *StringHeap) String() string {
  return str.Value()
}

func (str *StringBuffer) String() string {
  return str.Value()
}

func (str *StringHeap) Value() string {
  return str.contents
}

func (str *StringBuffer) Value() string {
  if str.set {
    return str.val
  } else if str.native != nil {
    // Load and verify
    val, err := str.native.String()
    errCheck(err, "string")

    // Cache and return
    str.val = val
    str.set = true
    return str.val
  } else {
    return ""
  }
}

func (str *StringHeap) ctype() *cdart.Packet {
  return nil
}

func (str *StringHeap) Size() uint {
  return uint(len(str.contents))
}

func (str *StringHeap) IsObject() bool {
  return false
}

func (str *StringHeap) IsArray() bool {
  return false
}

func (str *StringHeap) IsString() bool {
  return true
}

func (str *StringHeap) IsInteger() bool {
  return false
}

func (str *StringHeap) IsDecimal() bool {
  return false
}

func (str *StringHeap) IsBoolean() bool {
  return false
}

func (str *StringHeap) IsNull() bool {
  return false
}

func (str *StringHeap) IsFinalized() bool {
  return false
}

func (str *StringHeap) GetType() int {
  return cdart.StringType
}

func (str *StringHeap) Refcount() uint64 {
  return 1
}

func (str *StringBuffer) Size() uint {
  return uint(len(str.Value()))
}

func (str *StringBuffer) ctype() *cdart.Packet {
  return str.native
}

func (str *StringBuffer) IsObject() bool {
  return false
}

func (str *StringBuffer) IsArray() bool {
  return false
}

func (str *StringBuffer) IsString() bool {
  return true
}

func (str *StringBuffer) IsInteger() bool {
  return false
}

func (str *StringBuffer) IsDecimal() bool {
  return false
}

func (str *StringBuffer) IsBoolean() bool {
  return false
}

func (str *StringBuffer) IsNull() bool {
  return false
}

func (str *StringBuffer) IsFinalized() bool {
  return true
}

func (str *StringBuffer) GetType() int {
  return cdart.StringType
}

func (str *StringBuffer) Refcount() uint64 {
  return str.native.Refcount()
}

func (str *StringHeap) Equal(other *StringHeap) bool {
  return str.contents == other.contents
}

func (str *StringBuffer) Equal(other *StringBuffer) bool {
  // Calling into native extensions will likely be more expensive
  // than the string comparison itself, so use the cache if we can
  return str.Value() == other.Value()
}

func (str *StringHeap) toJSON(out *strings.Builder) {
  out.WriteRune('"')
  out.WriteString(str.contents)
  out.WriteRune('"')
}

func (str *StringHeap) ToJSON() string {
  var builder strings.Builder
  str.toJSON(&builder)
  return builder.String()
}

func (str *StringBuffer) toJSON(out *strings.Builder) {
  out.WriteString(str.ToJSON())
}

func (str *StringBuffer) ToJSON() string {
  if len(str.json) > 0 {
    // We've already generated JSON previously
    // Buffers are immutable, so just return it.
    return str.json
  } else if str.native != nil {
    // We haven't generated our JSON before, but we
    // have a native representation, so do it.
    json, err := str.native.ToJSON()
    errCheck(err, "string")

    str.json = json
    return str.json
  } else {
    // We're a default initialized struct
    // Just return a static string
    return "\"\""
  }
}
