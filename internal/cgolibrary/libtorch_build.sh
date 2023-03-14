#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

pushd $DIR

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
CXX="g++"
LIBRARY_DIR="library"
LIBTORCH_DIR="$LIBRARY_DIR"
GLIBCXX_USE_CXX11_ABI="1"
LOAD="force_load"
LIB_SUFFIX="so"
INSTALL_NAME=""
CUDA_FLAGS=""

function build_linux_cpu() {
    CXX="clang++"
    LIBTORCH_DIR="$LIBTORCH_DIR/linux/libtorch"
    GLIBCXX_USE_CXX11_ABI="0"
}

if [[ "$OS" == "linux" ]]; then
    if $(whereis cuda | cut -f 2 -d ' ')/bin/nvcc --version > /dev/null; then
        CXX="clang++"
        NVCC=$(whereis cuda | cut -f 2 -d ' ')/bin/nvcc
        CUDA_VERSION=$("$NVCC" --version | grep release | grep -Eo "[0-9]+.[0-9]+" | head -1)
        CUDA_FLAGS="$CUDA_FLAGS -DWITH_CUDA -I /usr/local/cuda/include"
        if [[ "$CUDA_VERSION" == "11.7" ]]; then
            echo "Building for Linux with CUDA 11.7";
            LIBTORCH_DIR="$LIBTORCH_DIR/linux-cuda117/libtorch"
        else
            echo "Unknown CUDA version: $CUDA_VERSION. Building for Linux CPU ..."
            build_linux_cpu
        fi
    else
        echo "Building for Linux CPU ...";
        build_linux_cpu
    fi
elif [[ "$OS" == "darwin" ]]; then
    echo "Building for macOS ...";
    CXX="clang++"
    LIBTORCH_DIR="$LIBTORCH_DIR/macos/libtorch"
    LIB_SUFFIX="dylib"
    INSTALL_NAME="-install_name @rpath/\$@"
    LOAD="all_load"
fi

set -o xtrace
make CXX="$CXX" \
     LIB_SUFFIX="$LIB_SUFFIX" \
     INSTALL_NAME="$INSTALL_NAME" \
     LIBTORCH_DIR="$LIBTORCH_DIR" \
     GLIBCXX_USE_CXX11_ABI="$GLIBCXX_USE_CXX11_ABI" \
     LOAD="$LOAD" \
     CUDA_FLAGS="$CUDA_FLAGS" \
     -f Makefile -j

popd