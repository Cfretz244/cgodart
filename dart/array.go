package dart

// #include <dart/abi.h>
import "C"

func (pkt *Packet) Index(idx int) (*Packet, error) {
  child := &Packet{}
  cidx := C.size_t(idx)
  err := withTLS(func () C.dart_err_t {
    return C.dart_arr_get_err(&child.cbuf, pkt.rawPtr(), cidx)
  })
  return maybeErrReg(child, err)
}

func (pkt *Packet) InsertIndex(idx int, child *Packet) error {
  cidx := C.size_t(idx)
  return withTLS(func () C.dart_err_t {
    return C.dart_arr_insert_dart(
      pkt.rawPtr(),
      cidx,
      child.rawPtr(),
    )
  })
}

func (pkt *Packet) InsertStringIndex(idx int, value string) error {
  cidx := C.size_t(idx)
  return withTLS(func () C.dart_err_t {
    return C.dart_arr_insert_str_len(
      pkt.rawPtr(),
      cidx,
      C._GoStringPtr(value),
      C.size_t(len(value)),
    )
  })
}

func (pkt *Packet) InsertIntegerIndex(idx int, value int64) error {
  cidx := C.size_t(idx)
  return withTLS(func () C.dart_err_t {
    return C.dart_arr_insert_int(
      pkt.rawPtr(),
      cidx,
      C.int64_t(value),
    )
  })
}

func (pkt *Packet) InsertDecimalIndex(idx int, value float64) error {
  cidx := C.size_t(idx)
  return withTLS(func () C.dart_err_t {
    return C.dart_arr_insert_dcm(
      pkt.rawPtr(),
      cidx,
      C.double(value),
    )
  })
}

func (pkt *Packet) InsertBooleanIndex(idx int, value bool) error {
  cidx := C.size_t(idx)
  return withTLS(func () C.dart_err_t {
    return C.dart_arr_insert_bool(
      pkt.rawPtr(),
      cidx,
      bool2int(value),
    )
  })
}

func (pkt *Packet) InsertNullIndex(idx int) error {
  cidx := C.size_t(idx)
  return withTLS(func () C.dart_err_t {
    return C.dart_arr_insert_null(
      pkt.rawPtr(),
      cidx,
    )
  })
}

func (pkt *Packet) SetIndex(idx int, child *Packet) error {
  cidx := C.size_t(idx)
  return withTLS(func () C.dart_err_t {
    return C.dart_arr_set_dart(
      pkt.rawPtr(),
      cidx,
      child.rawPtr(),
    )
  })
}

func (pkt *Packet) SetStringIndex(idx int, value string) error {
  cidx := C.size_t(idx)
  return withTLS(func () C.dart_err_t {
    return C.dart_arr_set_str_len(
      pkt.rawPtr(),
      cidx,
      C._GoStringPtr(value),
      C.size_t(len(value)),
    )
  })
}

func (pkt *Packet) SetIntegerIndex(idx int, value int64) error {
  cidx := C.size_t(idx)
  return withTLS(func () C.dart_err_t {
    return C.dart_arr_set_int(
      pkt.rawPtr(),
      cidx,
      C.int64_t(value),
    )
  })
}

func (pkt *Packet) SetDecimalIndex(idx int, value float64) error {
  cidx := C.size_t(idx)
  return withTLS(func () C.dart_err_t {
    return C.dart_arr_set_dcm(
      pkt.rawPtr(),
      cidx,
      C.double(value),
    )
  })
}

func (pkt *Packet) SetBooleanIndex(idx int, value bool) error {
  cidx := C.size_t(idx)
  return withTLS(func () C.dart_err_t {
    return C.dart_arr_set_bool(
      pkt.rawPtr(),
      cidx,
      bool2int(value),
    )
  })
}

func (pkt *Packet) Resize(length int) error {
  clen := C.size_t(length)
  return withTLS(func () C.dart_err_t {
    return C.dart_arr_resize(pkt.rawPtr(), clen)
  })
}
