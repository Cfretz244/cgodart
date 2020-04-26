package dart

// #include <dart/abi.h>
import "C"
import "unsafe"

func (pkt *Packet) String() (string, error) {
  var cstr *C.char
  var length C.size_t
  err := withTLS(func () C.dart_err_t {
    cstr = C.dart_str_get_len(pkt.rawPtr(), &length)
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
