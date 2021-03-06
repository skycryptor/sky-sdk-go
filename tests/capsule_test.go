package tests

import (
  "testing"
  "skycryptor-sdk-go/skycryptor"
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
  }
}

func TestCapsuleFromAndToBytes(t *testing.T) {
  sc := skycryptor.NewSkycryptor()
  defer sc.Clean()

  skA, pkA := sc.Keys.Generate()
  defer skA.Clean()
  defer pkA.Clean()

  capsule1, _ := pkA.Encapsulate()
  defer capsule1.Clean()

  cData1 := capsule1.ToBytes()
  capsule2 := sc.CapsuleFromBytes(cData1)
  defer capsule2.Clean()

  cData2 := capsule2.ToBytes()
  if string(cData2) != string(cData1) {
    t.Errorf("Capsule Bytes1: %s", string(cData1))
    t.Errorf("Capsule Bytes2: %s", string(cData2))
    t.Fail()
  }
}
