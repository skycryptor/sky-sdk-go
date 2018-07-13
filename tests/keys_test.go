package tests

import (
  "testing"
  "skycryptor-sdk-go/skycryptor"
)

func TestPublicKeyFromToBytes(t *testing.T) {
  sc := skycryptor.NewSkycryptor()
  defer sc.Clean()

  privateKey, publicKey := sc.Keys.Generate()
  defer privateKey.Clean()
  defer publicKey.Clean()

  pkData1 := publicKey.ToBytes()
  publicKey2 := sc.PublicKeyFromBytes(pkData1)
  defer publicKey2.Clean()

  pkData2 := publicKey2.ToBytes()
  if string(pkData1) != string(pkData2) {
    t.Errorf("Public Key Bytes1: %s", string(pkData1))
    t.Errorf("Public Key Bytes2: %s", string(pkData2))
    t.Fail()
  }
}

func TestPrivateKeyFromToBytes(t *testing.T) {
  sc := skycryptor.NewSkycryptor()
  defer sc.Clean()

  privateKey, publicKey := sc.Keys.Generate()
  defer privateKey.Clean()
  defer publicKey.Clean()

  skData1 := privateKey.ToBytes()
  privateKey2 := sc.PrivateKeyFromBytes(skData1)
  defer privateKey2.Clean()

  skData2 := privateKey2.ToBytes()
  if string(skData1) != string(skData2) {
    t.Errorf("Private Key Bytes1: %s", string(skData1))
    t.Errorf("Private Key Bytes2: %s", string(skData2))
    t.Fail()
  }
}
