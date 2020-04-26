package dart

/*
#cgo LDFLAGS: -ldart_abi
#include <dart/abi.h>

static inline int dart_type_as_int(dart_packet_t const* pkt) {
  return (int) dart_get_type(pkt);
}
*/
import "C"
import (
  "unsafe"
  "runtime"
  "errors"
)

const (
  ObjectType = iota + 1
  ArrayType
  StringType
  IntegerType
  DecimalType
  BooleanType
  NullType
)

type Packet struct {
  cbuf C.dart_packet_t
}

type Iterator struct {
  initialCheck bool
  cbuf C.dart_iterator_t
}

func destroyPacket(pkt *Packet) {
  C.dart_destroy(pkt.rawPtr())
}

func destroyIterator(it *Iterator) {
  C.dart_iterator_destroy(&it.cbuf)
}

func registerCObj(cobj interface{}) {
  switch obj := cobj.(type) {
  case *Packet:
    runtime.SetFinalizer(obj, destroyPacket)
  case *Iterator:
    runtime.SetFinalizer(obj, destroyIterator)
  default:
    panic("Invalid type passed to dart.registerCObj")
  }
}

func (pkt *Packet) rawPtr() unsafe.Pointer {
  return unsafe.Pointer(&pkt.cbuf)
}

func (it *Iterator) rawPtr() unsafe.Pointer {
  return unsafe.Pointer(&it.cbuf)
}

func int2bool(val C.int) bool {
  // Go is a bit overzealous about bools not being ints if you ask me
  // Like, I appreciate the thinking, but this is silly
  if val != 0 {
    return true
  } else {
    return false
  }
}

func bool2int(val bool) C.int {
  if val {
    return 1
  } else {
    return 0
  }
}

func isOK(err C.dart_err_t) bool {
  return err == C.DART_NO_ERROR
}

func grabError() error {
  return errors.New(C.GoString(C.dart_get_error()))
}

func maybeErr(pkt *Packet, err error) (*Packet, error) {
  if err == nil {
    return pkt, nil
  } else {
    return nil, err
  }
}

func maybeErrReg(pkt *Packet, err error) (*Packet, error) {
  pkt, err = maybeErr(pkt, err)
  if pkt != nil {
    registerCObj(pkt)
  }
  return pkt, err
}

func withTLS(impl func () C.dart_err_t) error {
  var err error
  runtime.LockOSThread()
  ret := impl()
  if !isOK(ret) {
    err = grabError()
  }
  runtime.UnlockOSThread()
  return err
}

func NewPacket() (*Packet, error) {
  pkt := &Packet {}
  err := withTLS(func () C.dart_err_t {
    return C.dart_init_err(&pkt.cbuf)
  })
  return maybeErrReg(pkt, err)
}

func NewObjectPacket() (*Packet, error) {
  pkt := &Packet {}
  err := withTLS(func () C.dart_err_t {
    return C.dart_obj_init_err(&pkt.cbuf)
  })
  return maybeErrReg(pkt, err)
}

func NewArrayPacket() (*Packet, error) {
  pkt := &Packet {}
  err := withTLS(func () C.dart_err_t {
    return C.dart_arr_init_err(&pkt.cbuf)
  })
  return maybeErrReg(pkt, err)
}

func NewStringPacket(val string) (*Packet, error) {
  pkt := &Packet {}
  err := withTLS(func () C.dart_err_t {
    return C.dart_str_init_len_err(&pkt.cbuf, C._GoStringPtr(val), C._GoStringLen(val))
  })
  return maybeErrReg(pkt, err)
}

func NewIntegerPacket(val int64) (*Packet, error) {
  pkt := &Packet {}
  err := withTLS(func () C.dart_err_t {
    return C.dart_int_init_err(&pkt.cbuf, C.int64_t(val))
  })
  return maybeErrReg(pkt, err)
}

func NewDecimalPacket(val float64) (*Packet, error) {
  pkt := &Packet {}
  err := withTLS(func () C.dart_err_t {
    return C.dart_dcm_init_err(&pkt.cbuf, C.double(val))
  })
  return maybeErrReg(pkt, err)
}

func NewBooleanPacket(val bool) (*Packet, error) {
  pkt := &Packet {}
  conv := 0
  if val {
    conv = 1
  }
  err := withTLS(func () C.dart_err_t {
    return C.dart_bool_init_err(&pkt.cbuf, C.int(conv))
  })
  return maybeErrReg(pkt, err)
}

func NewNullPacket() (*Packet, error) {
  return NewPacket()
}

func CopyPacket(pkt *Packet) (*Packet, error) {
  dup := &Packet{}
  err := withTLS(func () C.dart_err_t {
    return C.dart_copy_err(dup.rawPtr(), pkt.rawPtr())
  })
  return maybeErrReg(dup, err)
}

func NewIterator(pkt *Packet) (*Iterator, error) {
  it := &Iterator{true, C.dart_iterator_t{}}
  err := withTLS(func () C.dart_err_t {
    return C.dart_iterator_init_from_err(&it.cbuf, pkt.rawPtr())
  })
  if err == nil {
    registerCObj(it)
  } else {
    it = nil
  }
  return it, err
}

func NewKeyIterator(pkt *Packet) (*Iterator, error) {
  it := &Iterator{true, C.dart_iterator_t{}}
  err := withTLS(func () C.dart_err_t {
    return C.dart_iterator_init_key_from_err(&it.cbuf, pkt.rawPtr())
  })
  if err == nil {
    registerCObj(it)
  } else {
    it = nil
  }
  return it, err
}

func CopyIterator(it *Iterator) (*Iterator, error) {
  dup := &Iterator{}
  err := withTLS(func () C.dart_err_t {
    return C.dart_iterator_copy_err(&dup.cbuf, &it.cbuf)
  })
  if err == nil {
    registerCObj(dup)
  } else {
    dup = nil
  }
  return dup, err
}

func (pkt *Packet) IsObject() bool {
  retval := C.dart_is_obj(pkt.rawPtr())
  return int2bool(retval)
}

func (pkt *Packet) IsArray() bool {
  retval := C.dart_is_arr(pkt.rawPtr())
  return int2bool(retval)
}

func (pkt *Packet) IsAggregate() bool {
  return pkt.IsObject() || pkt.IsArray()
}

func (pkt *Packet) IsString() bool {
  retval := C.dart_is_str(pkt.rawPtr())
  return int2bool(retval)
}

func (pkt *Packet) IsInteger() bool {
  retval := C.dart_is_int(pkt.rawPtr())
  return int2bool(retval)
}

func (pkt *Packet) IsDecimal() bool {
  retval := C.dart_is_dcm(pkt.rawPtr())
  return int2bool(retval)
}

func (pkt *Packet) IsBoolean() bool {
  retval := C.dart_is_bool(pkt.rawPtr())
  return int2bool(retval)
}

func (pkt *Packet) IsNull() bool {
  retval := C.dart_is_null(pkt.rawPtr())
  return int2bool(retval)
}

func (pkt *Packet) IsFinalized() bool {
  retval := C.dart_is_finalized(pkt.rawPtr())
  return int2bool(retval)
}

func (pkt *Packet) GetType() int {
  return int(C.dart_type_as_int(&pkt.cbuf))
}

func (pkt *Packet) Refcount() uint64 {
  return uint64(C.dart_refcount(pkt.rawPtr()))
}

func (pkt *Packet) Size() (int, error) {
  var size C.size_t
  errVal := ^C.size_t(0)
  err := withTLS(func () C.dart_err_t {
    size = C.dart_size(pkt.rawPtr())
    if size != errVal {
      return C.DART_NO_ERROR
    } else {
      return C.DART_CLIENT_ERROR
    }
  })

  if err != nil {
    size = 0
  }
  return int(size), err
}

func (pkt *Packet) Clear() error {
  if pkt.IsObject() {
    C.dart_obj_clear(pkt.rawPtr())
  } else if pkt.IsArray() {
    C.dart_arr_clear(pkt.rawPtr())
  } else {
    return errors.New("dart::packet is not an aggregate, and has no values to clear")
  }
  return nil
}

func (pkt *Packet) Equals(other *Packet) bool {
  if pkt == other {
    return true
  } else {
    return int2bool(C.dart_equal(pkt.rawPtr(), other.rawPtr()))
  }
}

func (pkt *Packet) Finalize() error {
  return withTLS(func () C.dart_err_t {
    // Create a temporary packet to swap with.
    // This step could conceivably fail
    tmp := &Packet{}
    err := C.dart_finalize_err(&tmp.cbuf, pkt.rawPtr())
    if err != C.DART_NO_ERROR {
      return err
    }

    // Swap packet instances
    // This step can't fail even though the signatures
    // would suggest otherwise
    C.dart_destroy(pkt.rawPtr())
    C.dart_move_err(pkt.rawPtr(), tmp.rawPtr())
    C.dart_destroy(tmp.rawPtr())
    return C.DART_NO_ERROR
  })
}

func (pkt *Packet) Lower() error {
  return pkt.Finalize()
}

func (pkt *Packet) Definalize() error {
  return withTLS(func () C.dart_err_t {
    // Create a temporary packet to swap with.
    // This step could conceivably fail
    tmp := &Packet{}
    err := C.dart_definalize_err(&tmp.cbuf, pkt.rawPtr())
    if err != C.DART_NO_ERROR {
      return err
    }

    // Swap packet instances
    // This step can't fail even though the signatures
    // would suggest otherwise
    C.dart_destroy(pkt.rawPtr())
    C.dart_move_err(pkt.rawPtr(), tmp.rawPtr())
    C.dart_destroy(tmp.rawPtr())
    return C.DART_NO_ERROR
  })
}

func (pkt *Packet) Lift() error {
  return pkt.Definalize()
}

func (it *Iterator) Next() bool {
  if it.initialCheck {
    it.initialCheck = false
    if !int2bool(C.dart_iterator_done(&it.cbuf)) {
      return true
    } else {
      return false
    }
  } else {
    C.dart_iterator_next(&it.cbuf)
    return !int2bool(C.dart_iterator_done(&it.cbuf))
  }
}

func (it *Iterator) Value() (*Packet, error) {
  pkt := &Packet{}
  err := withTLS(func () (C.dart_err_t) {
    return C.dart_iterator_get_err(&pkt.cbuf, &it.cbuf)
  })
  return maybeErrReg(pkt, err)
}

func FromJSON(val string) (*Packet, error) {
  pkt := &Packet {}
  err := withTLS(func () C.dart_err_t {
    return C.dart_from_json_len_err(&pkt.cbuf, C._GoStringPtr(val), C._GoStringLen(val))
  })
  return maybeErrReg(pkt, err)
}

func (pkt *Packet) ToJSON() (string, error) {
  var cstr *C.char
  var length C.size_t
  err := withTLS(func () C.dart_err_t {
    cstr = C.dart_to_json(pkt.rawPtr(), &length)
    if cstr != nil {
      return C.DART_NO_ERROR
    } else {
      return C.DART_CLIENT_ERROR
    }
  })
  if err == nil {
    defer C.free(unsafe.Pointer(cstr))
    return C.GoStringN(cstr, C.int(length)), nil
  } else {
    return "", err
  }
}
