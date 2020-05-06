package dart

import (
  "github.com/cfretz244/godart/cdart"
)

type ObjectBuffer struct {
  native *cdart.Packet
  cache map[string] *Buffer
  json string
}

func objFromPacket(pkt *cdart.Packet) *ObjectBuffer {
  if !pkt.IsObject() {
    panic("Native packet of unexpected type passed to ObjectBuffer converter")
  } else if !pkt.IsFinalized() {
    panic("Non-finalized object passed to ObjectBuffer converter")
  }
  return &ObjectBuffer{pkt, make(map[string] *Buffer), ""}
}

func (obj *ObjectBuffer) Field(key string) *Buffer {
  // Short-circuit if we haven't been properly initialized
  if obj.native == nil {
    return nil
  }

  // Lazily load field into cache and return
  if !obj.cache[key].isSet() && obj.native.HasField(key) {
    pkt, err := obj.native.Field(key)
    errCheck(err, "object")
    obj.cache[key] = wrapBuffer(pkt)
  } else if !obj.cache[key].isSet() {
    return nil
  }
  return obj.cache[key]
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
