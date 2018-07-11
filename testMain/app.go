package main

import (
  "skycryptor-go/skycryptor"
  "fmt"
  "encoding/hex"
)

func main() {
  cm := skycryptor.NewCryptoMagic()
  defer cm.Clean()

  sk := cm.GeneratePrivateKey()
  defer sk.Clean()

  data := sk.ToBytes()
  fmt.Println(hex.EncodeToString(data))
}
