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
  if !obj.Equals(dup) {
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
  } else if !rebuilt.Equals(obj) {
    t.Error("Expected rebuilt object to be equal to original")
  }
}
