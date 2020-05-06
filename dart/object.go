package dart

import (
  "github.com/cfretz244/godart/cdart"
)

type ObjectBuffer struct {
  native *cdart.Packet
  cache map[string] Buffer
  json string
}

func (obj *ObjectBuffer) ctype() *cdart.Packet {
  return obj.native
}

func (obj *ObjectBuffer) IsObject() bool {
  return true
}

func (obj *ObjectBuffer) IsArray() bool {
  return false
}

func (obj *ObjectBuffer) IsString() bool {
  return false
}

func (obj *ObjectBuffer) IsInteger() bool {
  return false
}

func (obj *ObjectBuffer) IsDecimal() bool {
  return false
}

func (obj *ObjectBuffer) IsBoolean() bool {
  return false
}

func (obj *ObjectBuffer) IsNull() bool {
  return false
}

func (obj *ObjectBuffer) IsFinalized() bool {
  return true
}

func (obj *ObjectBuffer) GetType() int {
  return cdart.ObjectType
}

func (obj *ObjectBuffer) Refcount() uint64 {
  if obj.native == nil {
    return 0
  } else {
    return obj.native.Refcount()
  }
}

func (obj *ObjectBuffer) equal(other wrapper) bool {
  return false
}

func (obj *ObjectBuffer) ToJSON() string {
  if len(obj.json) > 0 {
    // We've already generated JSON previously
    // Buffers are immutable, so just return it.
    return obj.json
  } else if obj.native != nil {
    // We haven't generated our JSON before, but we
    // have a native representation, so do it.
    obj.json, _ = obj.native.ToJSON()
    return obj.json
  } else {
    // We're a default initialized struct
    // Just return a static string
    return "{}"
  }
}
