package skycryptor

//go:generate ../builder.sh

// Main Skycryptor structure for having API functions referenced from it
type SkyCryptor struct {
  CPointer
  cm *CryptoMagic
  Keys KeyPair
}

// Clearing C/C++ allocations
// For example:
//    sc := skycryptor.NewSkycryptor()
//    defer sc.Clean()
func (sc *SkyCryptor) Clean() {
  sc.cm.Clean()
}

// Making new SkyCryptor object with default crypto configurations
func NewSkycryptor() *SkyCryptor {
  cm := NewCryptoMagic()
  return &SkyCryptor{cm: cm, Keys: KeyPair{cm: cm}}
}

// Getting capsule object from given raw data converted from Capsule.ToBytes
func (sc *SkyCryptor) CapsuleFromBytes(capsuleData []byte) *Capsule {
  return CapsuleFromBytes(sc.cm, capsuleData)
}

// Getting private key object from given raw byte data converted with PrivateKey.ToBytes
func (sc *SkyCryptor) PrivateKeyFromBytes(skData []byte) *PrivateKey {
  return PrivateKeyFromBytes(sc.cm, skData)
}

// Getting public key object from given raw byte data converted with PublicKey.ToBytes
func (sc *SkyCryptor) PublicKeyFromBytes(pkData []byte) *PublicKey {
  return PublicKeyFromBytes(sc.cm, pkData)
}

// ReEncrypting capsule
func (sc *SkyCryptor) ReEncrypt(capsule *Capsule, rkk *ReEncryptionKey) *Capsule {
  return rkk.ReEncrypt(capsule)
}
