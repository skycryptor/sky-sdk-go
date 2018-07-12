package skycryptor
/*
#include "cryptomagic_c.h"
#include "stdlib.h"
 */
import "C"


// ReEncryptionKey is a definiton and functionality for having re-encryption keys from C/C++ library
type ReEncryptionKey struct {
  Key
}

// Cleaning up memory allocated from C/C++ library for this object pointer
func (rk *ReEncryptionKey) Clean() {
  C.cryptomagic_re_encryption_key_free(rk.pointer)
}

func (rk *ReEncryptionKey) ReEncrypt(capsule *Capsule) *Capsule {
  reCapsule := &Capsule{}
  reCapsule.cm = rk.cm
  reCapsule.pointer = C.cryptomagic_get_re_encryption_capsule(rk.cm.pointer, capsule.pointer, rk.pointer)
  return reCapsule
}
