package dart

// #include <dart/abi.h>
import "C"

func (pkt *Packet) Field(key string) (*Packet, error) {
  child := &Packet{}
  err := withTLS(func () C.dart_err_t {
    return C.dart_obj_get_len_err(
      &child.cbuf,
      pkt.rawPtr(),
      C._GoStringPtr(key),
      C._GoStringLen(key),
    )
  })
  return maybeErrReg(child, err)
}

func (pkt *Packet) InsertField(key string, child *Packet) error {
  return withTLS(func () C.dart_err_t {
    return C.dart_obj_insert_dart_len(
      pkt.rawPtr(),
      C._GoStringPtr(key),
      C._GoStringLen(key),
      child.rawPtr(),
    )
  })
}

func (pkt *Packet) RemoveField(key string) error {
  return withTLS(func () C.dart_err_t {
    return C.dart_obj_erase_len(pkt.rawPtr(), C._GoStringPtr(key), C._GoStringLen(key))
  })
}

func (pkt *Packet) HasField(key string) bool {
  has := C.dart_obj_has_key_len(pkt.rawPtr(), C._GoStringPtr(key), C._GoStringLen(key))
  return int2bool(has)
}
