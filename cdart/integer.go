package cdart

// #include <dart/abi.h>
import "C"

func (pkt *Packet) Integer() (int64, error) {
  pkt.maybeBail()
  var val C.int64_t
  err := withTLS(func () C.dart_err_t {
    return C.dart_int_get_err(pkt.rawPtr(), &val)
  })
  return int64(val), err
}
