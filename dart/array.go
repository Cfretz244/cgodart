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

func (arr *ArrayHeap) grow(idx uint) {
  // Check if we need to do anything
  max := uint(cap(arr.contents))
  if int(idx) < len(arr.contents) {
    return
  }

  // Check if we need to reallocate
  if idx + 1 > max {
    var size uint = 0
    if idx + 1 > max * 2 {
      size = idx + 1
    } else {
      size = max * 2
    }
    tmp := make([]*Heap, size)
    copy(tmp, arr.contents)
    arr.contents = tmp
  } else {
    arr.contents = arr.contents[:idx + 1]
  }
}

func (arr *ArrayHeap) pushUp(idx uint) {
  end := len(arr.contents)
  for end > int(idx) {
    if end == len(arr.contents) {
      arr.contents = append(arr.contents, arr.contents[end - 1])
    } else {
      arr.contents[end] = arr.contents[end - 1]
    }
    end--
  }
}

func (arr *ArrayHeap) Index(idx uint) *Heap {
  if arr.contents == nil || int(idx) >= len(arr.contents) {
    return nullHeap
  }
  return arr.contents[idx]
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

func (arr *ArrayHeap) InsertIndex(idx uint, val *Heap) {
  // Shift everything above
  arr.pushUp(idx)

  // Length check
  arr.grow(idx)

  // Update it
  arr.contents[idx] = val
}

func (arr *ArrayHeap) InsertStringIndex(idx uint, val string) {
  // Shift everything above
  arr.pushUp(idx)

  // Length check
  arr.grow(idx)

  // Update it
  arr.contents[idx] = &Heap{&StringHeap{val}}
}

func (arr *ArrayHeap) InsertIntegerIndex(idx uint, val int64) {
  // Shift everything above
  arr.pushUp(idx)

  // Length check
  arr.grow(idx)

  // Update it
  arr.contents[idx] = &Heap{&IntegerHeap{val}}
}

func (arr *ArrayHeap) InsertDecimalIndex(idx uint, val float64) {
  // Shift everything above
  arr.pushUp(idx)

  // Length check
  arr.grow(idx)

  // Update it
  arr.contents[idx] = &Heap{&DecimalHeap{val}}
}

func (arr *ArrayHeap) InsertBooleanIndex(idx uint, val bool) {
  // Shift everything above
  arr.pushUp(idx)

  // Length check
  arr.grow(idx)

  // Update it
  arr.contents[idx] = &Heap{&BooleanHeap{val}}
}

func (arr *ArrayHeap) InsertNullIndex(idx uint) {
  // Shift everything above
  arr.pushUp(idx)

  // Length check
  arr.grow(idx)

  // Update it
  arr.contents[idx] = &Heap{&NullHeap{}}
}

func (arr *ArrayHeap) SetIndex(idx uint, val *Heap) {
  // Length check
  arr.grow(idx)

  // Update it
  arr.contents[idx] = val
}

func (arr *ArrayHeap) SetStringIndex(idx uint, val string) {
  // Length check
  arr.grow(idx)

  // Update it
  arr.contents[idx] = &Heap{&StringHeap{val}}
}

func (arr *ArrayHeap) SetIntegerIndex(idx uint, val int64) {
  // Length check
  arr.grow(idx)

  // Update it
  arr.contents[idx] = &Heap{&IntegerHeap{val}}
}

func (arr *ArrayHeap) SetDecimalIndex(idx uint, val float64) {
  // Length check
  arr.grow(idx)

  // Update it
  arr.contents[idx] = &Heap{&DecimalHeap{val}}
}

func (arr *ArrayHeap) SetBooleanIndex(idx uint, val bool) {
  // Length check
  arr.grow(idx)

  // Update it
  arr.contents[idx] = &Heap{&BooleanHeap{val}}
}

func (arr *ArrayHeap) SetNullIndex(idx uint) {
  // Length check
  arr.grow(idx)

  // Update it
  arr.contents[idx] = &Heap{&NullHeap{}}
}

func (arr *ArrayHeap) Iterator() *HeapIterator {
  if arr.contents == nil {
    return &HeapIterator{}
  }

  // Create our implementation closure
  var i uint = 0
  impl := func () *Heap {
    if i < arr.Size() {
      tmp := arr.contents[i]
      i++
      if tmp != nil {
        return tmp
      } else {
        return nullHeap
      }
    } else {
      return nil
    }
  }
  return &HeapIterator{nil, impl}
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

      if val == nil {
        out.WriteString("null")
      } else {
        val.toJSON(out)
      }
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
