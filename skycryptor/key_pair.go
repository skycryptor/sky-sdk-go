package skycryptor

// Making generic key pair for having decoupled key functions available
type KeyPair struct {
  cm *CryptoMagic
}

// Getting public and private keys
func (kp *KeyPair) Generate() (*PrivateKey, *PublicKey) {
  sk := kp.cm.GeneratePrivateKey()
  return sk, sk.GetPublicKey()
}
