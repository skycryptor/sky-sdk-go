package skycryptor_go

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

type PointerType interface {
	free()
}

type CryptoMagic struct {
	PointerType

	// raw C pointer reference
	pointer unsafe.Pointer
}

type PrivateKey struct {
	PointerType

	// raw C pointer reference
	pointer unsafe.Pointer
}

type PublicKey struct {
	PointerType

	// raw C pointer reference
	pointer unsafe.Pointer
}

func CM_init() {
	C.cryptomagic_init()
	cm := NewCryptoMagic()
	sk := cm.generate_private_key()
	sk.getPublicKey()
	sk.free()
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

///
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
