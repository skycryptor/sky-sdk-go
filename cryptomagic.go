package skycryptor_go

/*
#include "CryptoMagic_C.h"
#include "stdlib.h"
#cgo LDFLAGS: ./libcryptomaic.a -lstdc++ -lssl -lcrypto
#cgo CFLAGS: -O3
*/
import "C"
import "unsafe"

type pointerType struct {
	// raw C pointer reference
	pointer unsafe.Pointer
}

type CryptoMagic struct {
	pointerType
}

type PrivateKey struct {
	pointerType
}

type PublicKey struct {
	pointerType
}

func CM_init() {
	C.cryptomagic_init()
	cm := NewCryptoMagic()
	sk := cm.generate_private_key()
	sk.getPublicKey()
	cm.pointer = C.cryptomagic_new()
}

/// Cryptomagic native functions
func NewCryptoMagic() (cm CryptoMagic) {
	cm.pointer = C.cryptomagic_new()
	return cm
}

func (cm *CryptoMagic) free() {
	C.cryptomagic_clear(cm.pointer)
}

func (cm *CryptoMagic) generate_private_key() (sk PrivateKey) {
	sk.pointer = C.cryptomagic_generate_private_key(cm.pointer)
	return sk
}

/// *** PrivateKey native functions
func (sk *PrivateKey) free() {
	C.cryptomagic_private_key_free(sk.pointer)
}

func (sk *PrivateKey) getPublicKey() (pk PublicKey) {
	pk.pointer = C.cryptomagic_get_public_key(sk.pointer)
	return pk
}

/// *** PublicKey native functions
func (sk *PublicKey) free() {
	C.cryptomagic_public_key_free(sk.pointer)
}
