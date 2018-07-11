package skycryptor

//go:generate ./builder.sh

// Main Skycryptor structure for having API functions referenced from it
type SkyCryptor struct {
  CPointer
  cm *CryptoMagic
  Keys KeyPair
}

// Clearing C/C++ allocations
func (sc *SkyCryptor) Clean() {
  sc.cm.Clean()
}

func NewSkycryptor() *SkyCryptor {
  cm := NewCryptoMagic()
  return &SkyCryptor{cm: cm, Keys: KeyPair{cm: cm}}
}
