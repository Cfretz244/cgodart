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

func (pkt *Packet) Insert(key string, child *Packet) error {
  return withTLS(func () C.dart_err_t {
    return C.dart_obj_insert_dart_len(
      pkt.rawPtr(),
      C._GoStringPtr(key),
      C._GoStringLen(key),
      child.rawPtr(),
    )
  })
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
