package dart

import (
  "github.com/cfretz244/godart/cdart"
)

type StringBuffer struct {
  native *cdart.Packet
  json string
}

func (str *StringBuffer) ctype() *cdart.Packet {
  return str.native
}

func (str *StringBuffer) IsStringBuffer() bool {
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

func (str *StringBuffer) equal(other wrapper) bool {
  return false
}

func (str *StringBuffer) ToJSON() string {
  if len(str.json) > 0 {
    // We've already generated JSON previously
    // Buffers are immutable, so just return it.
    return str.json
  } else if str.native != nil {
    // We haven't generated our JSON before, but we
    // have a native representation, so do it.
    str.json, _ = str.native.ToJSON()
    return str.json
  } else {
    // We're a default initialized struct
    // Just return a static string
    return "\"\""
  }
}
