package dart

/*
#cgo LDFLAGS: -ldart_abi
#include <dart/abi.h>
#include "trampoline.h"
size_t _GoStringLen(_GoString_);
char const* _GoStringPtr(_GoString_);
*/
import "C"
import "unsafe"
import "runtime"

type Packet struct {
  cbuf C.dart_packet_t
}

func destroyPacket(pkt *Packet) {
  C.dart_destroy(punPacket(pkt))
}

func registerPacket(pkt *Packet) {
  runtime.SetFinalizer(pkt, destroyPacket)
}

func punPacket(pkt *Packet) unsafe.Pointer {
  return unsafe.Pointer(&pkt.cbuf)
}

func NewPacket() (pkt *Packet) {
  pkt = &Packet {}
  C.dart_init_err(&pkt.cbuf)
  registerPacket(pkt)
  return pkt
}

func NewObjectPacket() (pkt *Packet) {
  pkt = &Packet {}
  C.dart_obj_init_err(&pkt.cbuf)
  registerPacket(pkt)
  return pkt
}

func NewArrayPacket() (pkt *Packet) {
  pkt = &Packet {}
  C.dart_arr_init_err(&pkt.cbuf)
  registerPacket(pkt)
  return pkt
}

func NewStringPacket(val string) (pkt *Packet) {
  pkt = &Packet {}
  return pkt
}

func (pkt *Packet) ToJSON() string {
  var length C.size_t
  cstr := C.dart_to_json(punPacket(pkt), &length)
  defer C.free(unsafe.Pointer(cstr))
  return C.GoStringN(cstr, C.int(length))
}
