package main

import (
  "skycryptor-go/cryptomagic"
  "fmt"
  "encoding/hex"
)

func main() {
  cm := cryptomagic.NewCryptoMagic()
  defer cm.Free()

  sk := cm.GeneratePrivateKey()
  defer sk.Free()

  data := sk.ToBytes()
  fmt.Println(hex.EncodeToString(data))
}
