package cdart_test

import (
  "testing"
  "github.com/cfretz244/godart/cdart"
)

// Some machinery to safely detect if a call panicked.
func try(impl func ()) bool {
  panicked := false
  func () {
    defer func () {
      panicked = recover() != nil
    }()
    impl()
  }()
  return panicked
}

// Check that something crashes
func assertPanic(t *testing.T, impl func ()) {
  if !try(impl) {
    t.Error("Expected panic for invoking method of uninitialized packet")
  }
}

func TestPacketInitialization(t *testing.T) {
  // cdart.Packet contains a raw C type,
  // which MUST be initialized by one of
  // the NewPacket functions prior to use
  // An unitialized packet instance.
  // DON'T DO ANY OF THIS
  pkt := &cdart.Packet{}

  assertPanic(t, func () {
    cdart.CopyPacket(pkt)
  })
  assertPanic(t, func () {
    cdart.NewIterator(pkt)
  })
  assertPanic(t, func () {
    cdart.NewKeyIterator(pkt)
  })
  assertPanic(t, func () {
    pkt.IsObject()
  })
  assertPanic(t, func () {
    pkt.IsArray()
  })
  assertPanic(t, func () {
    pkt.IsAggregate()
  })
  assertPanic(t, func () {
    pkt.IsString()
  })
  assertPanic(t, func () {
    pkt.IsInteger()
  })
  assertPanic(t, func () {
    pkt.IsDecimal()
  })
  assertPanic(t, func () {
    pkt.IsBoolean()
  })
  assertPanic(t, func () {
    pkt.IsNull()
  })
  assertPanic(t, func () {
    pkt.IsFinalized()
  })
  assertPanic(t, func () {
    pkt.GetType()
  })
  assertPanic(t, func () {
    pkt.Refcount()
  })
  assertPanic(t, func () {
    pkt.Size()
  })
  assertPanic(t, func () {
    pkt.Clear()
  })
  assertPanic(t, func () {
    pkt.Equal(&cdart.Packet{})
  })
  assertPanic(t, func () {
    pkt.Finalize()
  })
  assertPanic(t, func () {
    pkt.Lower()
  })
  assertPanic(t, func () {
    pkt.Definalize()
  })
  assertPanic(t, func () {
    pkt.Lift()
  })
  assertPanic(t, func () {
    pkt.ToBytes()
  })
  assertPanic(t, func () {
    pkt.ToJSON()
  })
}
