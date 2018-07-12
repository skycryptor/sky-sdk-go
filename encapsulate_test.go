package skycryptor_go

import (
  "testing"
  "skycryptor-go/skycryptor"
)

func TestEncapsulateDecapsulate(t *testing.T) {
  sc := skycryptor.NewSkycryptor()
  defer sc.Clean()

  skA, pkA := sc.Keys.Generate()
  defer skA.Clean()
  defer pkA.Clean()

  capsule1, symmetricKey1 := pkA.Encapsulate()
  defer capsule1.Clean()

  cData := capsule1.ToBytes()
  capsule1 = sc.CapsuleFromBytes(cData)

  symmetricKey2 := skA.Decapsulate(capsule1)
  if string(symmetricKey2) != string(symmetricKey1) {
    t.Errorf("Symmetric Key1: %s", string(symmetricKey1))
    t.Errorf("Symmetric Key2: %s", string(symmetricKey2))
    t.Fail()
    return
  }
}
