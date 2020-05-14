package dart_test

import (
  "testing"
  "github.com/cfretz244/godart/dart"
)

func TestParseObject(t *testing.T) {
  // Get a string to parse
  str := "{\"hello\":\"world\",\"pi\":3.14159,\"truth\":true,\"answer\":42}"
  buff, err := dart.BufferFromJSON(str)
  if err != nil {
    t.Error("Expected no error for well formed JSON string, got", err)
  }

  // Check its type
  if buftype := buff.GetType(); buftype != dart.ObjectType {
    t.Error("Expected object type for JSON object, got", buftype)
  } else if !buff.IsObject() {
    t.Error("Expected JSON object to claim to be of object type")
  }

  // Check its mutability
  if !buff.IsFinalized() {
    t.Error("Expected Buffer type to be finalized by definition")
  }

  // Check its initial exclusivity
  if rc := buff.Refcount(); rc != 1 {
    t.Error("Expected initial object refcount to be 1, got", rc)
  }

  // Attempt bad casts
  _, err1 := buff.ToArray()
  _, err2 := buff.ToString()
  _, err3 := buff.ToInteger()
  _, err4 := buff.ToDecimal()
  _, err5 := buff.ToBoolean()
  _, err6 := buff.ToNull()
  if err1 == nil || err2 == nil || err3 == nil || err4 == nil || err5 == nil || err6 == nil {
    t.Error("Expected object buffer to only be castable to object type")
  }

  // Cast it
  obj, err := buff.ToObject()
  if err != nil {
    t.Error("Expected object buffer to be castable to object type, got", err)
  }

  // Check its initial size
  if size := obj.Size(); size != 4 {
    t.Error("Expected size 4 for JSON object of size 4, got", size)
  }

  // Check individual fields
  world, pi, truth, ans := obj.Field("hello"), obj.Field("pi"), obj.Field("truth"), obj.Field("answer")
  if worldtype := world.GetType(); worldtype != dart.StringType {
    t.Error("Expected string type for string field, got", worldtype)
  } else if !world.IsString() {
    t.Error("Expected string field to claim to be of string type")
  } else if pitype := pi.GetType(); pitype != dart.DecimalType {
    t.Error("Expected decimaml type for decimal field, got", pitype)
  } else if !pi.IsDecimal() {
    t.Error("Expected decimal field to claim to be of decimal type")
  } else if truthtype := truth.GetType(); truthtype != dart.BooleanType {
    t.Error("Expected boolean type for boolean field, got", truthtype)
  } else if !truth.IsBoolean() {
    t.Error("Expected boolean field to claim to be of boolean type")
  } else if anstype := ans.GetType(); anstype != dart.IntegerType {
    t.Error("Expected integer type for integer field, got", anstype)
  } else if !ans.IsInteger() {
    t.Error("Expected integer field to claim to be of integer type")
  }

  // Cast the string
  pstr, err := world.ToString()
  if err != nil {
    t.Error("Expected string field to be castable to string type, got", err)
  }

  // Check its value
  if val := pstr.Value(); val != "world" {
    t.Error("Expected string field to have value \"world\", got", val)
  }

  // Cast the decimal
  pdcm, err := pi.ToDecimal()
  if err != nil {
    t.Error("Expected decimal field to be castable to decimal type, got", err)
  }

  // Check its value
  if val := pdcm.Value(); val != 3.14159 {
    t.Error("Expected decimal field to have value 3.14159, got", val)
  }

  // Cast the boolean
  pbool, err := truth.ToBoolean()
  if err != nil {
    t.Error("Expected boolean field to be castable to boolean type, got", err)
  }

  // Check its value
  if val := pbool.Value(); val != true {
    t.Error("Expected boolean field to have value true, got", val)
  }

  // Cast the integer
  pint, err := ans.ToInteger()
  if err != nil {
    t.Error("Expected integer field to be castable to integer type, got", err)
  }

  // Check its value
  if val := pint.Value(); val != 42 {
    t.Error("Expected integer field to have value 42, got", val)
  }

  // Check comparison
  buff2, err := dart.BufferFromJSON(str)
  if err != nil {
    t.Error("Expected no error for well formed JSON string, got", err)
  } else if !buff.Equal(buff2) {
    t.Error("Expected equivalent objects to be equal")
  }
}

func TestIterateObject(t *testing.T) {
  // Get an object
  str := "{\"hello\":\"world\",\"pi\":3.14159,\"truth\":true,\"answer\":42}"
  buff, err := dart.BufferFromJSON(str)
  if err != nil {
    t.Error("Expected no error for well formed JSON string, got", err)
  }
  obj, err := buff.ToObject()
  if err != nil {
    t.Error("Expected object to be castable to object type, got", err)
  }

  // Get an iterator
  it := obj.Iterator()

  // Check iteration order
  count := 0
  arr := [4]*dart.Buffer{obj.Field("pi"), obj.Field("hello"), obj.Field("truth"), obj.Field("answer")}
  for it.Next() {
    if !arr[count].Equal(it.Value()) {
      t.Error("Expected specific object value")
    }
    count++
  }
}

func TestObjectInitialization(t *testing.T) {
  // A default initialized object.
  obj := &dart.ObjectBuffer{}

  // Make sure we can pull a field out of it safely
  null := obj.Field("not here")
  if ntype := null.GetType(); ntype != dart.NullType {
    t.Error("Expected null type for non-existent field, got", ntype)
  } else if !null.IsNull() {
    t.Error("Expected non-existent field to claim to be null type")
  }

  // Make sure we can iterate safely
  it := obj.Iterator()
  for it.Next() {
    t.Error("Expected no iteration for empty object")
  }

  // Do something stupid and make sure it works
  null = it.Value()
  if ntype := null.GetType(); ntype != dart.NullType {
    t.Error("Expected null type for non-existent field, got", ntype)
  } else if !null.IsNull() {
    t.Error("Expected non-existent field to claim to be null type")
  }
}
