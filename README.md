godart
============
[![Build Status](https://travis-ci.com/Cfretz244/godart.svg?branch=master)](https://travis-ci.com/github/Cfretz244/godart)

**GoDart** is a set of experimental, actively being developed,
**Go** bindings for [libdart](https://github.com/target/libdart).

## State of the Repository

At the moment, **GoDart** contains a single package, `cdart`,
which implements a very thin wrapper over **Dart's C API**.

`cdart` will likely receive some extension in the future, but this
package can be considered reasonably stable for the time being.
For actively developing against, use the `v1.0.0-alpha` snapshot.

`cdart` is rather raw at the moment, mapping directly from **C**
types/functions into **Go** types/functions, and dispatching into
**C** on every individual call.

A basic example:
```go
package main

import "fmt"
import "github.com/cfretz244/godart/cdart"

func main() {
  pkt, _ := cdart.NewObjectPacket()
  pkt.InsertStringField("hello", "world")
  pkt.InsertIntegerField("answer", 42)
  pkt.InsertDecimalField("pi", 3.14159)
  pkt.Finalize()

  bytes, _ := pkt.ToBytes()
  rebuilt, _ := cdart.FromBytes(bytes)

  str, _ := pkt.ToJSON()
  rstr, _ := rebuilt.ToJSON()
  fmt.Println(str)
  fmt.Println(rstr)
}

// => {"pi":3.14159,"hello":"world","answer":42}
// => {"pi":3.14159,"hello":"world","answer":42}
```
More detailed examples of `cdart` usage can be found in the tests.

## Design Moving Forward

Invoking **C** code presents significant call overhead in **Go**,
and so the intention is for a second package to be added, `dart`,
which will serve as the client-facing **API** layer.

The purpose of the `dart` package is both to present an idiomatic
**Go** interface, but also to decrease the cost of calling into
**C** extensions by lazily constructing and caching objects received
from the `cdart` package.
