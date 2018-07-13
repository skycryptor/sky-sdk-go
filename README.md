# Skycryptor Go SDK
[Introduction](#introduction) | [SDK Features](#sdk-features) | [Installation](#installation) | [Usage Examples](#usage-examples) | [Docs](#docs) | [Support](#support)


## Introduction
[Skycryptor](https://skycryptor.com), SAS, is a Paris, France, based cybersecurity company and a graduate of the Techstars Paris 2017 accelerator program.

We provide "Encryption & Key Management" service in operation with open sourced libraries with support for Javascript, Python, Go, Rust, Java, C++ languages.

[Skycryptor](https://skycryptor.com) GO SDK provides Key Management Service along with fast data encryption capabilities for adding data privacy & cryptographic access control 
layer into your application with few lines of code. 
These easy to use SDK and APIs enable to Encrypt, Decrypt and Manage Access of all kinds of data by eliminating any data exposure risk and helping to stay 
compliant with HIPPA, GDPR and other data regulations.

It can empower applications from storage encryption or KYC apps to password-less authentication and more. 


Skycryptor KMS is based on cutting-edge Proxy Re-Encryption algorithms, which enables to build 
scalable Public Key Infrastructure with powerfull access control capabilities.

## Proxy Re-Encryption Overview

Proxy Re-Encryption is a new type of public key cryptography designed to eliminate some major functional constraints associated with standard 
public key cryptography.It starts from extending the standard public key cryptography setup via the third actor - a proxy service. The Proxy service then 
can be authorized by Alice to take any cyphertext encrypted under Alice's public key, and transform (also named re-encrypt) it under Bob's public key. 
The transformed cyphertext can be decrypted by Bob later. 

For making re-encryption, Proxy Service should be provided a special re-encryption key, which is created by Alice specially for Bob. 
The Re-Encryption key generation requires Alice's private key as well Bobs' public key, which means that re-encryption key can be created only by Alice 
but all without her interaction with Bob.
After being provide the re-encryption key, it is very important that the proxy service can re-encrypt Alice's cyphertexts without being able to decrypt them or
get any extra information about the original plaintext. 

Our Data Encapsulation and Proxy Re-Encryption algorithms are and based on standard ECIES approach and are implemented with [OpenSSL] (https://www.openssl.org/) and [libsodium](https://github.com/jedisct1/libsodium) 
utilizing seckp256k1 elliptic curves and based on standard ECIES approach.


## SDK Features
- Generate Public and Private key pairs for Alice, Bob,... 
- Enable Alice to generate the Re-Encryption key for Bob
- Perform Symmetric Key Encapsulation for the given Public Key (Similar to standard Diffie-Hellman key exchange)
- Given the Capsule and Re-Encryption Key, transform the capsule (Re-Encrypt) for the targeted recipient
- Perform Decapsulation and recovering of the symmetric key
- 

## Installation
This is a standard Go package, but it requires to install OpenSSL package separately.
```bash
~# # Install OpenSSL here, depends on OS you are running
~# go get github.com/skycryptor/skycryptor-sdk-go

# Compile Skycryptor C++ library and combine it with SDK
~# go generate github.com/skycryptor/skycryptor-sdk-go/skyscryptor
```

## Usage Examples
Before using our SDK make sure that you have successfully completed [Installation](#installation) step, 
as Go compiler will not compile without having the `go generate` command successful completed.

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

## Use Cases
- KYC Applications
- End-to-Enc encrypted cloud collaboration
- Decentralized Supply-Chain management
- Keeping data private from the subset of peers in Hyperledge Fabric Channels
## Support
Our developer support team is here to help you. Find out more information on our [Help Center](https://help.skycryptor.com/).
