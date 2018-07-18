package tests

import (
  "testing"
  "skycryptor-sdk-go/skycryptor"
)

func TestGenerateReEncryptionKey(t *testing.T) {
  sc := skycryptor.NewSkycryptor()
  defer sc.Clean()

  skA, pkA := sc.Keys.Generate()
  defer skA.Clean()
  defer pkA.Clean()

  skB, pkB := sc.Keys.Generate()
  defer skB.Clean()
  defer pkB.Clean()

  reEncryptionKey := skA.GenerateReKey(pkB)
  defer reEncryptionKey.Clean()

  buffer := reEncryptionKey.ToBytes()
  rkk2 := sc.ReEncryptionKeyFromBytes(buffer)
  buffer2 := rkk2.ToBytes()
  if string(buffer2) != string(buffer2) {
    t.Errorf("Reencryption encode/decoding error")
    t.Fail()
  }
}


