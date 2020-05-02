package cdart

// #include <dart/abi.h>
import "C"

func (pkt *Packet) Boolean() (bool, error) {
  var val C.int
  pkt.maybeBail()
  err := withTLS(func () C.dart_err_t {
    return C.dart_bool_get_err(pkt.rawPtr(), &val)
  })
  return int2bool(val), err
}
