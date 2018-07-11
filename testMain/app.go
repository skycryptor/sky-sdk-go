package main

import (
  "skycryptor-go/skycryptor"
  "fmt"
)

func main() {
  sc := skycryptor.NewSkycryptor()
  skA, pkA := sc.Keys.Generate()
  capsule1, _ := pkA.Encapsulate()

  skA.Decapsulate(capsule1)
  fmt.Println(string(symmetricKey1))
  fmt.Println(string(symmetricKey2))
  fmt.Println(string(symmetricKey2) == string(symmetricKey1))

  capsule1.Clean()
  skA.Clean()
  pkA.Clean()
  sc.Clean()
}
