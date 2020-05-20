package dart

import (
  "fmt"
  "strings"
  "github.com/cfretz244/godart/cdart"
)

type ObjectHeap struct {
  contents map[string] *Heap
}

type ObjectBuffer struct {
  native *cdart.Packet
  cache map[string] *Buffer
  json string
  size uint
}

func objFromPacket(pkt *cdart.Packet) *ObjectBuffer {
  // Check invariants
  if !pkt.IsObject() {
    panic("Native packet of unexpected type passed to ObjectBuffer converter")
  } else if !pkt.IsFinalized() {
    panic("Non-finalized object passed to ObjectBuffer converter")
  }

  // Load basic properties and return
  size, err := pkt.Size()
  if err != nil {
    panic("Failed to check native object size in ObjectBuffer converter")
  }
  return &ObjectBuffer{pkt, make(map[string] *Buffer), "", size}
}

func (obj *ObjectBuffer) Field(key string) *Buffer {
  // Short-circuit if we haven't been properly initialized
  if obj.native == nil {
    return nullBuffer
  }

  // Lazily load field into cache and return
  if !obj.cache[key].isSet() && obj.native.HasField(key) {
    pkt, err := obj.native.Field(key)
    errCheck(err, "object")
    obj.cache[key] = wrapBuffer(pkt)
  } else if !obj.cache[key].isSet() {
    return nullBuffer
  }
  return obj.cache[key]
}

func (obj *ObjectBuffer) Iterator() *BufferIterator {
  it := &BufferIterator{}

  // Load the iterator if we've been initialized
  if obj.native != nil {
    tmp, err := cdart.NewIterator(obj.native)
    errCheck(err, "iterator")
    it.native = tmp
  }
  return it
}

func (obj *ObjectBuffer) KeyIterator() *BufferIterator {
  it := &BufferIterator{}

  // Load the iterator if we've been initialized
  if obj.native != nil {
    tmp, err := cdart.NewKeyIterator(obj.native)
    errCheck(err, "iterator")
    it.native = tmp
  }
  return it
}

func (obj *ObjectHeap) ctype() *cdart.Packet {
  return nil
}

func (obj *ObjectHeap) Size() uint {
  return uint(len(obj.contents))
}

func (obj *ObjectHeap) IsObject() bool {
  return true
}

func (obj *ObjectHeap) IsArray() bool {
  return false
}

func (obj *ObjectHeap) IsString() bool {
  return false
}

func (obj *ObjectHeap) IsInteger() bool {
  return false
}

func (obj *ObjectHeap) IsDecimal() bool {
  return false
}

func (obj *ObjectHeap) IsBoolean() bool {
  return false
}

func (obj *ObjectHeap) IsNull() bool {
  return false
}

func (obj *ObjectHeap) IsFinalized() bool {
  return false
}

func (obj *ObjectHeap) GetType() int {
  return cdart.ObjectType
}

func (obj *ObjectHeap) Refcount() uint64 {
  return 1
}

func (obj *ObjectBuffer) Size() uint {
  return obj.size
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

func (obj *ObjectBuffer) Equal(other *ObjectBuffer) bool {
  // Recursively checking equality in Go would be slow,
  // but in C this operation is literally a memcmp,
  // so hand off to extensions unconditionally
  us, them := obj.ctype(), other.ctype()
  if us == them {
    return true
  } else if us == nil || them == nil {
    return false
  } else {
    return us.Equal(them)
  }
}

func (obj *ObjectHeap) toJSON(out *strings.Builder) {
  if obj.contents != nil {
    // Get a string builder
    out.WriteRune('{')

    // Add in all of our key-value pairs
    first := true
    for k, v := range obj.contents {
      if !first {
        out.WriteRune(',')
      }
      fmt.Fprintf(out, "\"%s\":", k)
      v.toJSON(out)
      first = false
    }
    out.WriteRune('}')
  } else {
    out.WriteString("{}")
  }
}

func (obj *ObjectHeap) ToJSON() string {
  var builder strings.Builder
  obj.toJSON(&builder)
  return builder.String()
}

func (obj *ObjectBuffer) toJSON(out *strings.Builder) {
  out.WriteString(obj.ToJSON())
}

func (obj *ObjectBuffer) ToJSON() string {
  if len(obj.json) > 0 {
    // We've already generated JSON previously
    // Buffers are immutable, so just return it.
    return obj.json
  } else if obj.native != nil {
    // We haven't generated our JSON before, but we
    // have a native representation, so do it.
    json, err := obj.native.ToJSON()
    errCheck(err, "object")

    obj.json = json
    return obj.json
  } else {
    // We're a default initialized struct
    // Just return a static string
    return "{}"
  }
}
