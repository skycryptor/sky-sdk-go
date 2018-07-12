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
  cm *CryptoMagic
}

func (c *Capsule) Clean() {
  C.cryptomagic_capsule_free(c.pointer)
}

func (c *Capsule) ToBytes() []byte {
  var c_buffer *C.char
  c_buffer_len := C.int(0)
  C.cryptomagic_capsule_to_bytes(c.pointer, &c_buffer, &c_buffer_len)
  retBuf := C.GoBytes(unsafe.Pointer(c_buffer), c_buffer_len)
  C.free(unsafe.Pointer(c_buffer))
  return retBuf
}

func CapsuleFromBytes(cm *CryptoMagic, capsuleData []byte) *Capsule {
  capsule := &Capsule{cm: cm}
  capsule.pointer = C.cryptomagic_capsule_from_bytes(cm.pointer, (*C.char)(unsafe.Pointer(&capsuleData[0])), C.int(len(capsuleData)))
  return capsule
}
