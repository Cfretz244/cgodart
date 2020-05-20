package dart

import (
  "strings"
  "github.com/cfretz244/godart/cdart"
)

type ArrayHeap struct {
  contents []*Heap
}

type ArrayBuffer struct {
  native *cdart.Packet
  cache []*Buffer
  json string
  size uint
}

func arrFromPacket(pkt *cdart.Packet) *ArrayBuffer {
  // Make sure type is as expected
  if !pkt.IsArray() {
    panic("Native packet of unexpected type passed to ArrayBuffer converter")
  } else if !pkt.IsFinalized() {
    panic("Non-finalized array passed to ArrayBuffer converter")
  }
  size, err := pkt.Size()
  errCheck(err, "array")
  return &ArrayBuffer{pkt, make([]*Buffer, size), "", size}
}

func (arr *ArrayBuffer) Index(idx uint) *Buffer {
  // Short-circuit if we haven't been properly initialized
  if arr.native == nil || int(idx) >= len(arr.cache) {
    return nullBuffer
  }

  // Lazily load index into cache and return
  if !arr.cache[idx].isSet() {
    pkt, err := arr.native.Index(idx)
    errCheck(err, "array")
    arr.cache[idx] = wrapBuffer(pkt)
  }
  return arr.cache[idx]
}

func (arr *ArrayBuffer) Iterator() *BufferIterator {
  it := &BufferIterator{}
  
  // Load the iterator if we've been initialized
  if arr.native != nil {
    tmp, err := cdart.NewIterator(arr.native)
    errCheck(err, "iterator")
    it.native = tmp
  }
  return it
}

func (arr *ArrayHeap) ctype() *cdart.Packet {
  return nil
}

func (arr *ArrayHeap) Size() uint {
  return uint(len(arr.contents))
}

func (arr *ArrayHeap) IsObject() bool {
  return false
}

func (arr *ArrayHeap) IsArray() bool {
  return true
}

func (arr *ArrayHeap) IsString() bool {
  return false
}

func (arr *ArrayHeap) IsInteger() bool {
  return false
}

func (arr *ArrayHeap) IsDecimal() bool {
  return false
}

func (arr *ArrayHeap) IsBoolean() bool {
  return false
}

func (arr *ArrayHeap) IsNull() bool {
  return false
}

func (arr *ArrayHeap) IsFinalized() bool {
  return false
}

func (arr *ArrayHeap) GetType() int {
  return cdart.ArrayType
}

func (arr *ArrayHeap) Refcount() uint64 {
  return 1
}

func (arr *ArrayHeap) Equal(other *ArrayHeap) bool {
  // Recursively checking equality in Go would be slow,
  // but in C this operation is literally a memcmp,
  // so hand off to extensions unconditionally
  us, them := arr.ctype(), other.ctype()
  if us == them {
    return true
  } else if us == nil || them == nil {
    return false
  } else {
    return us.Equal(them)
  }
}

func (arr *ArrayBuffer) ctype() *cdart.Packet {
  return arr.native
}

func (arr *ArrayBuffer) Size() uint {
  return arr.size
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

func (arr *ArrayBuffer) Equal(other *ArrayBuffer) bool {
  // Recursively checking equality in Go would be slow,
  // but in C this operation is literally a memcmp,
  // so hand off to extensions unconditionally
  us, them := arr.ctype(), other.ctype()
  if us == them {
    return true
  } else if us == nil || them == nil {
    return false
  } else {
    return us.Equal(them)
  }
}

func (arr *ArrayHeap) toJSON(out *strings.Builder) {
  if arr.contents != nil {
    // Get a string builder
    out.WriteRune('[')

    // Add in all our elements
    first := true
    for _, val := range arr.contents {
      if !first {
        out.WriteRune(',')
      }
      val.toJSON(out)
      first = false
    }
    out.WriteRune(']')
  } else {
    out.WriteString("[]")
  }
}

func (arr *ArrayHeap) ToJSON() string {
  var builder strings.Builder
  arr.toJSON(&builder)
  return builder.String()
}

func (arr *ArrayBuffer) toJSON(out *strings.Builder) {
  out.WriteString(arr.ToJSON())
}

func (arr *ArrayBuffer) ToJSON() string {
  if len(arr.json) > 0 {
    // We've already generated JSON previously
    // Buffers are immutable, so just return it.
    return arr.json
  } else if arr.native != nil {
    // We haven't generated our JSON before, but we
    // have a native representation, so do it.
    json, err := arr.native.ToJSON()
    errCheck(err, "array")

    arr.json = json
    return arr.json
  } else {
    // We're a default initialized struct
    // Just return a static string
    return "[]"
  }
}
