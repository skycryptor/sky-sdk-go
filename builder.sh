#!/usr/bin/env bash
LIBNAME="cryptomagic"
LIBFILE="lib$LIBNAME.a"
SOURCE_DIR="`pwd`"

# Cloning and building C++ library
git clone git@gitlab.com:skycryptor/cpp-tmp-crypto.git "$LIBNAME"
cd cryptomagic
rm -rf build
mkdir -p build && cd build

cmake ..
make -j4
cp "$LIBFILE" "$SOURCE_DIR/cryptomagic"
cd "$SOURCE_DIR" && rm -rf "$LIBNAME"