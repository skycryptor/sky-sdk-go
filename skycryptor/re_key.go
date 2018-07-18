package skycryptor
/*
#include "cryptomagic_c.h"
#include "stdlib.h"
 */
import "C"
import "unsafe"

// ReEncryptionKey is a definiton and functionality for having re-encryption keys from C/C++ library
type ReEncryptionKey struct {
  Key
}

// Cleaning up memory allocated from C/C++ library for this object pointer
func (rk *ReEncryptionKey) Clean() {
  C.cryptomagic_re_encryption_key_free(rk.pointer)
}

// Running re-encryption for given capsule and returning re-encrypted capsule
func (rk *ReEncryptionKey) ReEncrypt(capsule *Capsule) *Capsule {
  reCapsule := &Capsule{}
  reCapsule.cm = rk.cm
  reCapsule.pointer = C.cryptomagic_get_re_encryption_capsule(rk.cm.pointer, capsule.pointer, rk.pointer)
  return reCapsule
}

// Encoding re-encryption key into bytes
func (rk *ReEncryptionKey) ToBytes() []byte {
  var c_buffer *C.char
  c_buffer_len := C.int(0)
  C.cryptomagic_re_encryption_to_bytes(rk.pointer, &c_buffer,&c_buffer_len)
  retBuf := C.GoBytes(unsafe.Pointer(c_buffer), c_buffer_len)
  C.free(unsafe.Pointer(c_buffer))
  return retBuf
}

// Decoding and returning re-encryption key from encoded byte array
func ReEncryptionKeyFromBytes(cm *CryptoMagic, data []byte) *ReEncryptionKey {
  rk := &ReEncryptionKey{}
  rk.cm = cm
  rk.pointer = C.cryptomagic_get_re_encryption_from_bytes(cm.pointer, (*C.char)(unsafe.Pointer(&data[0])), C.int(len(data)))
  return rk
}
