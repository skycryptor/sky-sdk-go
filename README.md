# Skycryptor Go SDK
[Introduction](#introduction) | [SDK Features](#sdk-features) | [Installation](#installation) | [Usage Examples](#usage-examples) | [Docs](#docs) | [Support](#support)


## Introduction
[Skycryptor](https://skycryptor.com) provides scalable Key Management Service with powerful encryption integrations for adding security and data privacy layer to any application.
Using our tools you can encrypt, decrypt and share any kind of data easily, or write a middleware for your storage encryption, or provide password-less authentication for your application, etc...
Our flexible SDK and API will give you powerful tools to become HIPAA and GDPR compliant and match more.

## SDK Features
- Easily generate Public and Private keys 
- Generate and store Re encryption key from given private key
- Make an encryption capsule from Public key
- Create re-encryption capsule for making dasha sharing match easier

## Installation
This is a standard Go package, but because we are using our C++ library as an integration, to provide more flexible API interface and performance
you need also to have OpenSSL installed.
```bash
~# # Install OpenSSL here, depends on OS you are running
~# go get github.com/skycryptor/skycryptor-sdk-go

# Compile Skycryptor C++ library and combine it with SDK
~# go generate github.com/skycryptor/skycryptor-sdk-go/skyscryptor
```

## Usage Examples
Before using our SDK make sure that you have successfully completed [Installation](#installation) step, because Go compiler would not compile without having `go generate` command successful completion.

Checking our [Go Tests](https://github.com/skycryptor/skycryptor-sdk-go/tests) is a good start to understand how to use our SDK.

```go
package main

import (
  "fmt"
  "skycryptor-sdk-go/skycryptor"
)

func main() {
  // Making new Skycryptor object from default encryption context 
  sc := skycryptor.NewSkycryptor()
  defer sc.Clean()

  // Generating Private, Public keys randomly 
  skA, pkA := sc.Keys.Generate()
  defer skA.Clean()
  defer pkA.Clean()
  fmt.Println("Private Key bytes: ", string(skA.ToBytes()))
  fmt.Println("Public Key bytes: ", string(pkA.ToBytes()))
  
  // Encapsulating and getting encryption capsule and symmetric key bytes 
  capsule, symmetricKey := pkA.Encapsulate()
  fmt.Println("Capsule Buffer: ", string(capsule.ToBytes()))
  fmt.Println("Symmetric Key Buffer: ", string(symmetricKey.ToBytes()))
}
```
This basic example demonstrates how to generate random keys, get them as a basic byte arrays, get encryption capsule and symmetric key as a byte buffer out of the generated public key.

## Docs
This Skycryptor Go SDK documentation available in GoDocs https://godoc.org/github.com/skycryptor/skycryptor-sdk-go

## Support
Our developer support team is here to help you. Find out more information on our [Help Center](https://help.skycryptor.com/).
