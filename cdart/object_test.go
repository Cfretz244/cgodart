package cdart_test

import (
  "testing"
  "github.com/cfretz244/godart/cdart"
)

func TestNewObjectPacket(t *testing.T) {
  // Create an object
  obj, err := cdart.NewObjectPacket()
  if err != nil {
    t.Error("Expected no error for default object, got", err)
  }

  // Check its type
  if ctype := obj.GetType(); ctype != cdart.ObjectType {
    t.Error("Expected object type for default object, got", ctype)
  } else if !obj.IsObject() {
    t.Error("Expected default object to claim to be of object type")
  }

  // Check its initial mutability
  if obj.IsFinalized() {
    t.Error("Expected default object to be non-finalized")
  }

  // Check its initial exclusivity
  if rc := obj.Refcount(); rc != 1 {
    t.Error("Expected default object initial refcount to be 1, got", rc)
  }

  // Check its initial size
  if size, err := obj.Size(); size != 0 {
    t.Error("Expected zero size for default object, got", size)
  } else if err != nil {
    t.Error("Expected no error for default object size, got", err)
  }

  // Check its iteration
  it, err := cdart.NewIterator(obj)
  if err != nil {
    t.Error("Expected default object to be iterable")
  }
  for it.Next() {
    t.Error("Expected default object iteration to immediately terminate")
  }

  // Check duplication
  dup, err := cdart.CopyPacket(obj)
  if err != nil {
    t.Error("Expected no error for default copy, got", err)
  }

  // Check equality
  if !obj.Equal(dup) {
    t.Error("Expected copied packets to be equal")
  }

  // Check updated refcount
  if rc := obj.Refcount(); rc != dup.Refcount() || rc != 2 {
    t.Error("Expected copied packets to share a refcount")
  }

  // Check finalized form
  obj.Finalize()
  bytes, err := obj.ToBytes()
  if !obj.IsFinalized() {
    t.Error("Expected finalized object to be finalized")
  } else if err != nil {
    t.Error("Expected finalized object to give a serialized buffer")
  }

  // Reconstruct the packet and check it.
  rebuilt, err := cdart.FromBytes(bytes)
  if err != nil {
    t.Error("Expected finalized object buffer to be reconstructible")
  } else if !rebuilt.Equal(obj) {
    t.Error("Expected rebuilt object to be equal to original")
  }
}

func TestObjectFieldAccess(t *testing.T) {
  // Get an object
  obj, _ := cdart.NewObjectPacket()

  // Check what happens if we access a non-existent field
  fld, err := obj.Field("hello")
  if err != nil {
    t.Error("Expected non-existent field access to return null, not an error")
  } else if !fld.IsNull() {
    t.Error("Expected non-existent field to be null")
  } else if ftype := fld.GetType(); ftype != cdart.NullType {
    t.Error("Expected type of null field to be null")
  }

  // Check what happens if we insert some fields.
  obj.InsertStringField("hello", "world")
  obj.InsertIntegerField("answer", 42)
  obj.InsertDecimalField("c", 2.99792)
  obj.InsertBooleanField("truth", true)
  obj.InsertNullField("lies")
  if size, _ := obj.Size(); size != 5 {
    t.Error("Expected object of size five to report size of 5, got", size)
  }

  // Check the string field
  sfld, _ := obj.Field("hello");
  strval, _ := sfld.String()
  if strval != "world" {
    t.Error("Expected object with key \"hello\" to have value \"world\", got", strval)
  }

  // Check the int field
  ifld, _ := obj.Field("answer")
  intval, _ := ifld.Integer()
  if intval != 42 {
    t.Error("Expected object with key \"answer\" to have value 42, got", intval)
  }

  // Check the decimal field
  dfld, _ := obj.Field("c")
  dcmval, _ := dfld.Decimal()
  if dcmval != 2.99792 {
    t.Error("Expected object with key \"c\" to have value 2.99792, got", dcmval)
  }

  // Check the boolean field
  bfld, _ := obj.Field("truth")
  boolval, _ := bfld.Boolean()
  if !boolval {
    t.Error("Expected object with key \"truth\" to have value true, got", dcmval)
  }

  // Check the null field
  nfld, _ := obj.Field("lies")
  if !nfld.IsNull() {
    t.Error("Expected object with key \"lies\" to have nil value, got", dcmval)
  }

  // Check iteration
  count := 0
  arr := [5]*cdart.Packet{dfld, nfld, sfld, bfld, ifld}
  it, _ := cdart.NewIterator(obj)
  for it.Next() {
    val, _ := it.Value()
    if !val.Equal(arr[count]) {
      t.Error("Expected specific object value")
    }
    count++
  }

  // Check erasure
  obj.RemoveField("lies")
  if shorter, _ := obj.Size(); shorter != 4 {
    t.Error("Expected object with removed field to be of length 4, got", 4)
  }
}
