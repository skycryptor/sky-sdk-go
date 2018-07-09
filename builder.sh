#!/usr/bin/env bash
LIBNAME="cryptomagic"
LIBFILE="lib$LIBNAME.a"
SOURCE_DIR="`pwd`"
if [ ! -d "$LIBNAME" ]; then
  git clone git@gitlab.com:skycryptor/cpp-tmp-crypto.git "$LIBNAME"
fi

cd cryptomagic
rm -rf build
mkdir -p build && cd build

cmake ..
make -j4
cp $LIBFILE $SOURCE_DIR
