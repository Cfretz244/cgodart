package cdart

// #include <dart/abi.h>
import "C"

func (pkt *Packet) Decimal() (float64, error) {
  var val C.double
  err := withTLS(func () C.dart_err_t {
    return C.dart_dcm_get_err(pkt.rawPtr(), &val)
  })
  return float64(val), err
}
