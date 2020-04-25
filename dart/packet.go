package dart

/*
#cgo LDFLAGS: -ldart_abi
#include <dart/abi.h>

static inline int dart_type_to_int(dart_type_t type) {
  return (int) type;
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

func destroyPacket(pkt *Packet) {
  C.dart_destroy(pkt.rawPtr())
}

func registerPacket(pkt *Packet) {
  runtime.SetFinalizer(pkt, destroyPacket)
}

func (pkt *Packet) rawPtr() unsafe.Pointer {
  return unsafe.Pointer(&pkt.cbuf)
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

func isOK(err C.dart_err_t) bool {
  return err == C.DART_NO_ERROR;
}

func grabError() error {
  return errors.New(C.GoString(C.dart_get_error()))
}

func newPktRet(pkt *Packet, err C.dart_err_t) (*Packet, error) {
  if isOK(err) {
    registerPacket(pkt)
    return pkt, nil
  } else {
    return nil, grabError()
  }
}

func NewPacket() (*Packet, error) {
  pkt := &Packet {}
  ret := C.dart_init_err(&pkt.cbuf)
  return newPktRet(pkt, ret)
}

func NewObjectPacket() (*Packet, error) {
  pkt := &Packet {}
  ret := C.dart_obj_init_err(&pkt.cbuf)
  return newPktRet(pkt, ret)
}

func NewArrayPacket() (*Packet, error) {
  pkt := &Packet {}
  ret := C.dart_arr_init_err(&pkt.cbuf)
  return newPktRet(pkt, ret)
}

func NewStringPacket(val string) (*Packet, error) {
  pkt := &Packet {}
  ret := C.dart_str_init_len_err(&pkt.cbuf, C._GoStringPtr(val), C._GoStringLen(val))
  return newPktRet(pkt, ret)
}

func NewIntegerPacket(val int64) (*Packet, error) {
  pkt := &Packet {}
  ret := C.dart_int_init_err(&pkt.cbuf, C.int64_t(val))
  return newPktRet(pkt, ret)
}

func NewDecimalPacket(val float64) (*Packet, error) {
  pkt := &Packet {}
  ret := C.dart_dcm_init_err(&pkt.cbuf, C.double(val))
  return newPktRet(pkt, ret)
}

func NewBooleanPacket(val bool) (*Packet, error) {
  pkt := &Packet {}
  conv := 0
  if val {
    conv = 1
  }
  ret := C.dart_bool_init_err(&pkt.cbuf, C.int(conv))
  return newPktRet(pkt, ret)
}

func NewNullPacket() (*Packet, error) {
  return NewPacket()
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

func (pkt *Packet) GetType() int {
  return int(C.dart_type_to_int(C.dart_get_type(pkt.rawPtr())))
}

func FromJSON(val string) *Packet {
  pkt := &Packet {}
  ret := C.dart_from_json_len_err(&pkt.cbuf, C._GoStringPtr(val), C._GoStringLen(val))
  return newPktRet(pkt, ret)
}

func (pkt *Packet) ToJSON() string {
  var length C.size_t
  cstr := C.dart_to_json(pkt.rawPtr(), &length)
  defer C.free(unsafe.Pointer(cstr))
  return C.GoStringN(cstr, C.int(length))
}
