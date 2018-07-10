package main

import (
  "skycryptor-go/cryptomagic"
  "fmt"
  "encoding/hex"
)

func main() {
  cm := cryptomagic.NewCryptoMagic()
  sk := cm.GeneratePrivateKey()
  data := sk.ToBytes()
  fmt.Println(hex.EncodeToString(data))
}
