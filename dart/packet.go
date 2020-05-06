package dart

import (
  "github.com/cfretz244/godart/cdart"
)

type wrapper interface {
  ctype() *cdart.Packet

  IsObject() bool
  IsArray() bool
  IsString() bool
  IsInteger() bool
  IsDecimal() bool
  IsBoolean() bool
  IsNull() bool
  IsFinalized() bool

  GetType() int
  Refcount() uint64
  equal(other wrapper) bool

  ToJSON() string
}

type Buffer struct {
  impl wrapper
}

func (buf *Buffer) asObject() *ObjectBuffer {
  return buf.impl.(*ObjectBuffer)
}

func (buf *Buffer) asArray() *ArrayBuffer {
  return buf.impl.(*ArrayBuffer)
}

func (buf *Buffer) isSet() bool {
  switch buf.impl.(type) {
  case *ObjectBuffer:
    return true
  case *ArrayBuffer:
    return true
  default:
    return false
  }
}
