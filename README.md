# Skycryptor Go SDK
[Introduction](#introduction) | [SDK Features](#sdk-features) | [Installation](#installation) | [Usage Examples](#usage-examples) | [Docs](#docs) | [Support](#support)


## Introduction
[Skycryptor](https://skycryptor.com), SAS, is a Paris, France, based cybersecurity company and a graduate of the Techstars Paris 2017 accelerator program.

We provide "Encryption & Key Management" service in operation with open sourced libraries with support for Javascript, Python, Go, Rust, Java, C++ languages.

[Skycryptor](https://skycryptor.com) GO SDK provides Key Management Service along with fast data encryption capabilities for adding data privacy & 
cryptographic access control layer into your app with just few lines of code. These easy to use SDK and APIs enable to Encrypt, Decrypt and Manage Access 
of all kinds of data by eliminating data exposure risks and also helping to stay compliant with HIPPA, GDPR and other data regulations.

Use these tools for adding privacy by design into your apps starting from storage encryption or KYC platform to password-less authentication and more. 

Skycryptor KMS is based on cutting-edge Proxy Re-Encryption algorithms, which enables to build scalable Public Key Infrastructure with 
powerfull access control capabilities.

## Proxy Re-Encryption Overview

Proxy Re-Encryption is a new type of public key cryptography designed to eliminate some major functional constraints associated with standard 
public key cryptography.It starts from extending the standard public key cryptography setup via the third actor - a proxy service. The Proxy service then 
can be authorized by Alice to take any cyphertext encrypted under Alice's public key, and transform (also named re-encrypt) it under Bob's public key. 
The transformed cyphertext can be decrypted by Bob later. 

For making re-encryption, Proxy Service should be provided a special re-encryption key, which is created by Alice specially for Bob. 
The Re-Encryption key generation requires Alice's private key and Bobs' public key. This means  re-encryption key can be generated only by Alice 
and without any interaction with Bob.
It is very important to note, that the proxy service once given the re-encryptin key from Alice to Bob, can re-encrypt Alice's cyphertexts without being able to decrypt them or
get any extra information about the original plaintext. 

Our Data Encapsulation and Proxy Re-Encryption algorithms are and based on standard ECIES approach and are implemented with [OpenSSL] (https://www.openssl.org/) and [libsodium](https://github.com/jedisct1/libsodium) 
utilizing seckp256k1 elliptic curves and based on standard ECIES approach.


## SDK Features

- Generate and Manage User's Public and Private keys.  
- Enable users to generate Re-Encryption keys for their peers.
- Encapsulate a symmetric encryption key via given Public Key (similar to Diffie-Hellman Key Exchange)
- Perform Re-Encryption for the given ciphertext and the re-encryption key
- Decrypt both the original or transformed ciphertexts in order to reveal the encapsulated symmetric encryption key.

## Installation
This is a standard Go package, but it requires to install OpenSSL package separately.
```bash
~# # Install OpenSSL here, depends on OS you are running
~# go get github.com/skycryptor/skycryptor-sdk-go

# Compile Skycryptor C++ library and combine it with SDK
~# go generate github.com/skycryptor/skycryptor-sdk-go/skyscryptor
```

## Usage Examples
Before using our SDK make sure to successfully complete the [Installation](#installation) step, otherwise the Go compiler will not compile 
without having the `go generate` command successfully completed.

Checking our [Go Tests](https://github.com/skycryptor/skycryptor-sdk-go/tests) is the good start for understanding  to use our SDK.

### Initalization


```go
package main

import (
  "fmt"
  "skycryptor-sdk-go/skycryptor"
)

....

  // Initialize new Skycryptor context from the default encryption context 
  sc := skycryptor.NewSkycryptor()
```

#### Generate User's Public and Private Keys  
  ```
  // generates random private key and creates corresponding public key 
  alice_private_key, alice_public_key := sc.keys.generate()
  ```
#### Generate random symmetric key and encapsulate it with the Alice's Public Key 
```
  // Encapsulate function is a Diffie-Hellman style key exchange with randomly generated temprorary keys and Alice Public Key. 
  // It returns both the exchanged symmetric key and the capsule which can be decapsulated later by the corresponding private key
  // The generated symmetric_key can be used to protect and data object. 
  capsule, symmetricKey := alice_public_key.encapsulate()
  
```

#### The capsule can be unlocked by Alice private key to reveal the locked symmetric key

  //
```

  // Alice can use her private key to unlock the Capsule and recover the symmetric encryptin key.
  capsule, symmetricKey := alice_private_key.decapsulate(Capsule)
  
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
