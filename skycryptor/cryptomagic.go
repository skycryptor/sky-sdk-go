package skycryptor

/*
#include "cryptomagic_c.h"
#include "stdlib.h"
#cgo LDFLAGS: ${SRCDIR}/libProxyLib.a -lstdc++ -lssl -lcrypto
#cgo CFLAGS: -O3
*/
import "C"
import (
	"unsafe"
)

// Raw C pointer interface for implementing Garbage collection with it
type CPointer interface {
	Clean()
}

// Generic Key structure for having reference to Public or Private keys
// Because they have mostly the same high level functions
type Key struct {
  CPointer
  pointer unsafe.Pointer
  cm *CryptoMagic
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
func (cm *CryptoMagic) Clean() {
	C.cryptomagic_clear(cm.pointer)
}

// Generating and returning private key object as a random private key
func (cm *CryptoMagic) GeneratePrivateKey() (sk *PrivateKey) {
  sk = &PrivateKey{}
  sk.cm = cm
  sk.pointer = C.cryptomagic_generate_private_key(cm.pointer)
  return sk
}

// Getting Private Key from bytes
func (cm *CryptoMagic) GetPrivateKeyFromBytes(data []byte) (sk *PrivateKey) {
  sk = &PrivateKey{}
  sk.pointer = C.cryptomagic_private_key_from_bytes(cm.pointer, (*C.char)(unsafe.Pointer(&data[0])), C.int(len(data)))
  return sk
}

