package dart

import (
  "errors"
  "github.com/cfretz244/godart/cdart"
)

const (
  ObjectType = cdart.ObjectType
  ArrayType = cdart.ArrayType
  StringType = cdart.StringType
  IntegerType = cdart.IntegerType
  DecimalType = cdart.DecimalType
  BooleanType = cdart.BooleanType
  NullType = cdart.NullType
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

  ToJSON() string
}

type Buffer struct {
  impl wrapper
}

type BufferIterator struct {
  native *cdart.Iterator
  cache *Buffer
}

var nullBuffer *Buffer

func init() {
  nullBuffer = &Buffer{&NullBuffer{}}
}

func errCheck(err error, kind string) {
  if err != nil {
    msg := "Encountered unexpected error: \"" + err.Error() + "\" while interacting with " + kind
    panic(msg)
  }
}

func wrapBuffer(pkt *cdart.Packet) *Buffer {
  switch pkt.GetType() {
  case cdart.ObjectType:
    return &Buffer{objFromPacket(pkt)}
  case cdart.ArrayType:
    return &Buffer{arrFromPacket(pkt)}
  case cdart.StringType:
    return &Buffer{strFromPacket(pkt)}
  case cdart.IntegerType:
    return &Buffer{intFromPacket(pkt)}
  case cdart.DecimalType:
    return &Buffer{dcmFromPacket(pkt)}
  case cdart.BooleanType:
    return &Buffer{boolFromPacket(pkt)}
  case cdart.NullType:
    return &Buffer{nullFromPacket(pkt)}
  default:
    panic("Unexpected packet type passed into dart.wrapBuffer")
  }
}

func BufferFromJSON(val string) (*Buffer, error) {
  pkt, err := cdart.FastFromJSON(val)
  if err != nil {
    return nil, err
  }
  return wrapBuffer(pkt), nil
}

func (buf *Buffer) IsObject() bool {
  return buf.impl.IsObject()
}

func (buf *Buffer) IsArray() bool {
  return buf.impl.IsArray()
}

func (buf *Buffer) IsString() bool {
  return buf.impl.IsString()
}

func (buf *Buffer) IsInteger() bool {
  return buf.impl.IsInteger()
}

func (buf *Buffer) IsDecimal() bool {
  return buf.impl.IsDecimal()
}

func (buf *Buffer) IsBoolean() bool {
  return buf.impl.IsBoolean()
}

func (buf *Buffer) IsNull() bool {
  return buf.impl.IsNull()
}

func (buf *Buffer) IsFinalized() bool {
  return true
}

func (buf *Buffer) GetType() int {
  return buf.impl.GetType()
}

func (buf *Buffer) Refcount() uint64 {
  return buf.impl.Refcount()
}

func (buf *Buffer) Equal(other *Buffer) bool {
  us, them := buf.impl.ctype(), other.impl.ctype()
  if us == them {
    return true
  } else if us == nil || them == nil {
    return false
  } else {
    return us.Equal(them)
  }
}

func (buf *Buffer) ToJSON() string {
  return buf.impl.ToJSON()
}

func (buf *Buffer) ToObject() (*ObjectBuffer, error) {
  if buf.IsObject() {
    return buf.impl.(*ObjectBuffer), nil
  } else {
    return nil, errors.New("dart.Buffer is not an object and cannot be coerced as such")
  }
}

func (buf *Buffer) ToArray() (*ArrayBuffer, error) {
  if buf.IsArray() {
    return buf.impl.(*ArrayBuffer), nil
  } else {
    return nil, errors.New("dart.Buffer is not an array and cannot be coerced as such")
  }
}

func (buf *Buffer) ToString() (*StringBuffer, error) {
  if buf.IsString() {
    return buf.impl.(*StringBuffer), nil
  } else {
    return nil, errors.New("dart.Buffer is not a string and cannot be coerced as such")
  }
}

func (buf *Buffer) ToInteger() (*IntegerBuffer, error) {
  if buf.IsInteger() {
    return buf.impl.(*IntegerBuffer), nil
  } else {
    return nil, errors.New("dart.Buffer is not an integer and cannot be coerced as such")
  }
}

func (buf *Buffer) ToDecimal() (*DecimalBuffer, error) {
  if buf.IsDecimal() {
    return buf.impl.(*DecimalBuffer), nil
  } else {
    return nil, errors.New("dart.Buffer is not a decimal and cannot be coerced as such")
  }
}

func (buf *Buffer) ToBoolean() (*BooleanBuffer, error) {
  if buf.IsBoolean() {
    return buf.impl.(*BooleanBuffer), nil
  } else {
    return nil, errors.New("dart.Buffer is not a boolean and cannot be coerced as such")
  }
}

func (buf *Buffer) ToNull() (*NullBuffer, error) {
  if buf.IsNull() {
    return buf.impl.(*NullBuffer), nil
  } else {
    return nil, errors.New("dart.Buffer is not null and cannot be coerced as such")
  }
}

func (buf *Buffer) AsObject() *ObjectBuffer {
  obj, err := buf.ToObject()
  if err != nil {
    panic(err.Error())
  }
  return obj
}

func (buf *Buffer) AsArray() *ArrayBuffer {
  arr, err := buf.ToArray()
  if err != nil {
    panic(err.Error())
  }
  return arr
}

func (buf *Buffer) AsString() *StringBuffer {
  str, err := buf.ToString()
  if err != nil {
    panic(err.Error())
  }
  return str
}

func (buf *Buffer) AsInteger() *IntegerBuffer {
  num, err := buf.ToInteger()
  if err != nil {
    panic(err.Error())
  }
  return num
}

func (buf *Buffer) AsDecimal() *DecimalBuffer {
  dcm, err := buf.ToDecimal()
  if err != nil {
    panic(err.Error())
  }
  return dcm
}

func (buf *Buffer) AsBoolean() *BooleanBuffer {
  boolean, err := buf.ToBoolean()
  if err != nil {
    panic(err.Error())
  }
  return boolean
}

func (buf *Buffer) AsNull() *NullBuffer {
  null, err := buf.ToNull()
  if err != nil {
    panic(err.Error())
  }
  return null
}

func (buf *Buffer) isSet() bool {
  // First make sure the receiver is set
  if buf == nil {
    return false
  }

  // Then check if its implementation is set
  switch buf.impl.(type) {
  case *ObjectBuffer:
    return true
  case *ArrayBuffer:
    return true
  case *StringBuffer:
    return true
  case *IntegerBuffer:
    return true
  case *DecimalBuffer:
    return true
  case *BooleanBuffer:
    return true
  case *NullBuffer:
    return true
  default:
    return false
  }
}

func (it *BufferIterator) Next() bool {
  if it.native != nil {
    // Moving iterator forward invalidates cache
    it.cache = nil
    return it.native.Next()
  } else {
    return false
  }
}

func (it *BufferIterator) Value() *Buffer {
  if it.cache != nil {
    return it.cache
  } else if it.native != nil {
    // Load and verify
    val, err := it.native.Value()
    errCheck(err, "iterator")

    // Cache and return
    it.cache = wrapBuffer(val)
    return it.cache
  } else {
    return nullBuffer
  }
}
