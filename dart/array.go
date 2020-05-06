package dart

import (
  "github.com/cfretz244/godart/cdart"
)

type ArrayBuffer struct {
  native *cdart.Packet
  cache []Buffer
  json string
}

func (arr *ArrayBuffer) validate() {
  if arr.native == nil || arr.cache == nil {
    panic("dart.ArrayBuffer instances must be initialized by a factory function")
  }
}

func (arr *ArrayBuffer) Index(idx int64) *Buffer {
  arr.validate()
  if arr.cache[idx].isSet() {
    return &arr.cache[idx]
  } else {
    return nil
  }
}

func (arr *ArrayBuffer) ctype() *cdart.Packet {
  return arr.native
}

func (arr *ArrayBuffer) IsObject() bool {
  return false
}

func (arr *ArrayBuffer) IsArray() bool {
  return true
}

func (arr *ArrayBuffer) IsString() bool {
  return false
}

func (arr *ArrayBuffer) IsInteger() bool {
  return false
}

func (arr *ArrayBuffer) IsDecimal() bool {
  return false
}

func (arr *ArrayBuffer) IsBoolean() bool {
  return false
}

func (arr *ArrayBuffer) IsNull() bool {
  return false
}

func (arr *ArrayBuffer) IsFinalized() bool {
  return true
}

func (arr *ArrayBuffer) GetType() int {
  return cdart.ArrayType
}

func (arr *ArrayBuffer) Refcount() uint64 {
  if arr.native == nil {
    return 0
  } else {
    return arr.native.Refcount()
  }
}

func (arr *ArrayBuffer) equal(other wrapper) bool {
  return false
}

func (arr *ArrayBuffer) ToJSON() string {
  if len(arr.json) > 0 {
    // We've already generated JSON previously
    // Buffers are immutable, so just return it.
    return arr.json
  } else if arr.native != nil {
    // We haven't generated our JSON before, but we
    // have a native representation, so do it.
    arr.json, _ = arr.native.ToJSON()
    return arr.json
  } else {
    // We're a default initialized struct
    // Just return a static string
    return "[]"
  }
}
