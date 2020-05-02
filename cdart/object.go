package cdart

// #include <dart/abi.h>
import "C"

func (pkt *Packet) Field(key string) (*Packet, error) {
  pkt.maybeBail()
  child := &Packet{}
  err := withTLS(func () C.dart_err_t {
    return C.dart_obj_get_len_err(
      &child.cbuf,
      pkt.rawPtr(),
      C._GoStringPtr(key),
      C.size_t(len(key)),
    )
  })
  return maybeErrReg(child, err)
}

func (pkt *Packet) InsertField(key string, child *Packet) error {
  pkt.maybeBail()
  return withTLS(func () C.dart_err_t {
    return C.dart_obj_insert_dart_len(
      pkt.rawPtr(),
      C._GoStringPtr(key),
      C.size_t(len(key)),
      child.rawPtr(),
    )
  })
}

func (pkt *Packet) InsertStringField(key string, value string) error {
  pkt.maybeBail()
  return withTLS(func () C.dart_err_t {
    return C.dart_obj_insert_str_len(
      pkt.rawPtr(),
      C._GoStringPtr(key),
      C.size_t(len(key)),
      C._GoStringPtr(value),
      C.size_t(len(value)),
    )
  })
}

func (pkt *Packet) InsertIntegerField(key string, value int64) error {
  pkt.maybeBail()
  return withTLS(func () C.dart_err_t {
    return C.dart_obj_insert_int_len(
      pkt.rawPtr(),
      C._GoStringPtr(key),
      C.size_t(len(key)),
      C.int64_t(value),
    )
  })
}

func (pkt *Packet) InsertDecimalField(key string, value float64) error {
  pkt.maybeBail()
  return withTLS(func () C.dart_err_t {
    return C.dart_obj_insert_dcm_len(
      pkt.rawPtr(),
      C._GoStringPtr(key),
      C.size_t(len(key)),
      C.double(value),
    )
  })
}

func (pkt *Packet) InsertBooleanField(key string, value bool) error {
  pkt.maybeBail()
  return withTLS(func () C.dart_err_t {
    return C.dart_obj_insert_bool_len(
      pkt.rawPtr(),
      C._GoStringPtr(key),
      C.size_t(len(key)),
      bool2int(value),
    )
  })
}

func (pkt *Packet) InsertNullField(key string) error {
  pkt.maybeBail()
  return withTLS(func () C.dart_err_t {
    return C.dart_obj_insert_null_len(
      pkt.rawPtr(),
      C._GoStringPtr(key),
      C.size_t(len(key)),
    )
  })
}

func (pkt *Packet) RemoveField(key string) error {
  pkt.maybeBail()
  return withTLS(func () C.dart_err_t {
    return C.dart_obj_erase_len(
      pkt.rawPtr(),
      C._GoStringPtr(key),
      C._GoStringLen(key),
    )
  })
}

func (pkt *Packet) HasField(key string) bool {
  pkt.maybeBail()
  has := C.dart_obj_has_key_len(
    pkt.rawPtr(),
    C._GoStringPtr(key),
    C._GoStringLen(key),
  )
  return int2bool(has)
}
