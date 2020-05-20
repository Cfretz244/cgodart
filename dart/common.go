package dart

import (
  "strings"
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

  toJSON(out *strings.Builder)
  ToJSON() string
}
