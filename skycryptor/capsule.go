package skycryptor

/*
#include "cryptomagic_c.h"
#include "stdlib.h"
 */
import "C"
import "unsafe"

// Cryptographic capsule referenced from C/C++ library implementation
type Capsule struct {
  CPointer
  pointer unsafe.Pointer
}

func (c *Capsule) Clean() {
  C.cryptomagic_capsule_free(c.pointer)
}
