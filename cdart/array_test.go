package cdart_test

import (
  "testing"
  "github.com/cfretz244/godart/cdart"
)

func TestNewArrayPacket(t *testing.T) {
  // Create an array
  arr, err := cdart.NewArrayPacket()
  if err != nil {
    t.Error("Expected no error for default array, got", err)
  }

  // Check its type
  if ctype := arr.GetType(); ctype != cdart.ArrayType {
    t.Error("Expected array type for default array, got", ctype)
  } else if !arr.IsArray() {
    t.Error("Expected default array to claim to be of array type")
  }

  // Check its initial mutability
  if arr.IsFinalized() {
    t.Error("Expected default array to be non-finalized")
  }

  // Check its initial exclusivity
  if rc := arr.Refcount(); rc != 1 {
    t.Error("Expected default array initial refcount to be 1, got", rc)
  }

  // Check its initial size
  if size, err := arr.Size(); size != 0 {
    t.Error("Expected zero size for default array , got", size)
  } else if err != nil {
    t.Error("Expected no error for default array size, got", err)
  }

  // Check its iteration
  it, err := cdart.NewIterator(arr)
  if err != nil {
    t.Error("Expected default array to be iterable")
  }
  for it.Next() {
    t.Error("Expected default array iteration to immediately terminate")
  }

  // Check duplication
  dup, err := cdart.CopyPacket(arr)
  if err != nil {
    t.Error("Expected no error for default copy, got", err)
  }

  // Check equality
  if !arr.Equal(dup) {
    t.Error("Expected copied packets to be equal")
  }

  // Check updated refcount
  if rc := arr.Refcount(); rc != dup.Refcount() || rc != 2 {
    t.Error("Expected copied packets to share a refcount")
  }
}

func TestArrayFieldAccess(t *testing.T) {
  // Get an array
  arr, _ := cdart.NewArrayPacket()

  // Check what happens if we access a non-existent field
  fld, err := arr.Index(0)
  if err != nil {
    t.Error("Expected non-existent index access to return null, not an error")
  } else if !fld.IsNull() {
    t.Error("Expected non-existent index to be null")
  } else if ftype := fld.GetType(); ftype != cdart.NullType {
    t.Error("Expected type of null index to be null")
  }

  // Check what happens if we insert some fields.
  arr.InsertStringIndex(0, "world")
  arr.InsertIntegerIndex(1, 42)
  arr.InsertDecimalIndex(2, 2.99792)
  arr.InsertBooleanIndex(3, true)
  arr.InsertNullIndex(4)
  if size, _ := arr.Size(); size != 5 {
    t.Error("Expected array of size five to report size of 5, got", size)
  }

  // Check the string field
  sfld, _ := arr.Index(0);
  strval, _ := sfld.String()
  if strval != "world" {
    t.Error("Expected array with key \"hello\" to have value \"world\", got", strval)
  }

  // Check the int field
  ifld, _ := arr.Index(1)
  intval, _ := ifld.Integer()
  if intval != 42 {
    t.Error("Expected array with key \"answer\" to have value 42, got", intval)
  }

  // Check the decimal field
  dfld, _ := arr.Index(2)
  dcmval, _ := dfld.Decimal()
  if dcmval != 2.99792 {
    t.Error("Expected array with key \"c\" to have value 2.99792, got", dcmval)
  }

  // Check the boolean field
  bfld, _ := arr.Index(3)
  boolval, _ := bfld.Boolean()
  if !boolval {
    t.Error("Expected array with key \"truth\" to have value true, got", dcmval)
  }

  // Check the null field
  nfld, _ := arr.Index(4)
  if !nfld.IsNull() {
    t.Error("Expected array with key \"lies\" to have nil value, got", dcmval)
  }

  // Check iteration
  count := 0
  goarr := [5]*cdart.Packet{sfld, ifld, dfld, bfld, nfld}
  it, _ := cdart.NewIterator(arr)
  for it.Next() {
    val, _ := it.Value()
    if !val.Equal(goarr[count]) {
      t.Error("Expected specific array value")
    }
    count++
  }

  // Check erasure
  arr.RemoveIndex(2)
  if shorter, _ := arr.Size(); shorter != 4 {
    t.Error("Expected array with removed field to be of length 4, got", 4)
  }
}

func TestArrayInitialization(t *testing.T) {
  // cdart.Packet contains a raw C type,
  // which MUST be initialized by one of
  // the NewPacket functions prior to use
  // An unitialized packet instance.
  // DON'T DO ANY OF THIS
  pkt := &cdart.Packet{}

  assertPanic(t, func () {
    pkt.Index(0)
  })
  assertPanic(t, func () {
    pkt.InsertIndex(0, &cdart.Packet{})
  })
  assertPanic(t, func () {
    pkt.InsertStringIndex(0, "nope")
  })
  assertPanic(t, func () {
    pkt.InsertIntegerIndex(0, -1)
  })
  assertPanic(t, func () {
    pkt.InsertDecimalIndex(0, -1.0)
  })
  assertPanic(t, func () {
    pkt.InsertBooleanIndex(0, false)
  })
  assertPanic(t, func () {
    pkt.InsertNullIndex(0)
  })
  assertPanic(t, func () {
    pkt.SetIndex(0, &cdart.Packet{})
  })
  assertPanic(t, func () {
    pkt.SetStringIndex(0, "nope")
  })
  assertPanic(t, func () {
    pkt.SetIntegerIndex(0, -1)
  })
  assertPanic(t, func () {
    pkt.SetDecimalIndex(0, -1.0)
  })
  assertPanic(t, func () {
    pkt.SetBooleanIndex(0, false)
  })
  assertPanic(t, func () {
    pkt.SetNullIndex(0)
  })
  assertPanic(t, func () {
    pkt.RemoveIndex(0)
  })
  assertPanic(t, func () {
    pkt.Resize(10)
  })
  assertPanic(t, func () {
    pkt.Reserve(10)
  })
}
