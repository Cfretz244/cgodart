package dart

import (
  "github.com/cfretz244/godart/cdart"
)

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

func (str *StringBuffer) String() string {
  return str.Value()
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

func (str *StringBuffer) Equal(other *StringBuffer) bool {
  // Calling into native extensions will likely be more expensive
  // than the string comparison itself, so use the cache if we can
  return str.Value() == other.Value()
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
