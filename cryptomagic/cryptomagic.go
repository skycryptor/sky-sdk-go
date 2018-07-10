package cryptomagic

/*
#include "cryptomagic_c.h"
#include "stdlib.h"
#cgo LDFLAGS: ${SRCDIR}/libcryptomagic.a -lstdc++ -lssl -lcrypto
#cgo CFLAGS: -O3
*/
import "C"
import (
	"unsafe"
)

// Raw C pointer interface for implementing Garbage collection with it
type CPointer interface {
	free()
}

// PrivateKey object for which is Go implementation of existing C/C++ library interface
type PrivateKey struct {
  CPointer
  pointer unsafe.Pointer
  cm *CryptoMagic
}

// Cleaning up C/C++ allocations for private key object
func (sk *PrivateKey) free() {
  C.cryptomagic_private_key_free(sk.pointer);
}

// Main Crypto operations structure, which is a Go implementation of existing C/C++ library interface
type CryptoMagic struct {
	CPointer
	pointer unsafe.Pointer
}

// Making new CryptoMagic root object to perform cryptographic operations
func NewCryptoMagic() (cm *CryptoMagic) {
  cm = &CryptoMagic{}
	cm.pointer = C.cryptomagic_new()
	return cm
}

// Cleaning up C/C++ allocated memory for CryptoMagic object
func (cm *CryptoMagic) free() {
	C.cryptomagic_clear(cm.pointer)
}

// Generating and returning private key object as a random private key
func (cm *CryptoMagic) GeneratePrivateKey() (sk *PrivateKey) {
  sk = &PrivateKey{}
  sk.cm = cm
  sk.pointer = C.cryptomagic_generate_private_key(cm.pointer)
  return sk
}

func (cm *CryptoMagic) GetPrivateKeyFromBytes(data []byte) (sk *PrivateKey) {
  sk = &PrivateKey{}
  sk.pointer = C.cryptomagic_private_key_from_bytes(cm.pointer, (*C.char)(unsafe.Pointer(&data[0])), C.int(len(data)))
  return sk
}

func (sk *PrivateKey) ToBytes() []byte {
  var c_buffer *C.char
  c_buffer_len := C.int(0)
  C.cryptomagic_private_key_to_bytes(sk.pointer, &c_buffer,&c_buffer_len)
  retBuf := C.GoBytes(unsafe.Pointer(c_buffer), c_buffer_len)
  C.free(unsafe.Pointer(c_buffer))
  return retBuf
}

