package tests

import (
  "testing"
  "skycryptor-go/skycryptor"
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
}


