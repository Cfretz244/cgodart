package dart

import (
  "github.com/cfretz244/godart/cdart"
)

type DecimalBuffer struct {
  native *cdart.Packet
  json string
  val float64
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

func (num *DecimalBuffer) equal(other wrapper) bool {
  return false
}

func (num *DecimalBuffer) ToJSON() string {
  if len(num.json) > 0 {
    // We've already generated JSON previously
    // Buffers are immutable, so just return it.
    return num.json
  } else if num.native != nil {
    // We haven't generated our JSON before, but we
    // have a native representation, so do it.
    num.json, _ = num.native.ToJSON()
    return num.json
  } else {
    // We're a default initialized numuct
    // Just return a static string
    return "0.0"
  }
}
