#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

pushd $DIR

OS=$(uname | tr '[:upper:]' '[:lower:]')
LIBRARY_DIR="library"
LIBTORCH_DIR="$LIBRARY_DIR"

rm -rf "$DIR/$LIBRARY_DIR"

if [[ ! -d "$DIR/$LIBRARY_DIR" ]]; then
    mkdir -p "$DIR/$LIBRARY_DIR"
fi

function install_linux_cpu() {
    LIBTORCH_DIR="$LIBTORCH_DIR/linux/libtorch"
    if [[ ! -d "$DIR/$LIBTORCH_DIR" ]]; then
        curl -LsO 'https://download.pytorch.org/libtorch/cpu/libtorch-shared-with-deps-1.13.1%2Bcpu.zip'
        unzip -qq -o libtorch-shared-with-deps-1.13.1%2Bcpu.zip -d "$DIR/$LIBRARY_DIR"/linux
        rm libtorch-shared-with-deps-1.13.1%2Bcpu.zip
    fi
}

if [[ "$OS" == "linux" ]]; then
    if $(whereis cuda | cut -f 2 -d ' ')/bin/nvcc --version > /dev/null; then
        NVCC=$(whereis cuda | cut -f 2 -d ' ')/bin/nvcc
        CUDA_VERSION=$("$NVCC" --version | grep release | grep -Eo "[0-9]+.[0-9]+" | head -1)
        if [[ "$CUDA_VERSION" == "11.7" ]]; then
            echo "Installing for Linux with CUDA 11.7";
            LIBTORCH_DIR="$LIBTORCH_DIR/linux-cuda117/libtorch"
            if [[ ! -d "$DIR/$LIBTORCH_DIR" ]]; then
                curl -Lso libtorch-cxx11-abi-shared-with-deps-1.13.1-cu117.zip 'https://download.pytorch.org/libtorch/cu117/libtorch-cxx11-abi-shared-with-deps-1.13.1%2Bcu117.zip'
                unzip -qq -o libtorch-cxx11-abi-shared-with-deps-1.13.1-cu117.zip -d "$DIR/$LIBRARY_DIR"/linux-cuda117
                rm libtorch-cxx11-abi-shared-with-deps-1.13.1-cu117.zip
            fi
        else
            echo "Unknown CUDA version: $CUDA_VERSION. Installing for Linux CPU ..."
            install_linux_cpu
        fi
    else
        echo "Installing for Linux CPU ...";
        install_linux_cpu
    fi
elif [[ "$OS" == "darwin" ]]; then
    echo "Installing for macOS ...";
    if [[ ! -d "$DIR/$LIBTORCH_DIR" ]]; then
        curl -LsO https://download.pytorch.org/libtorch/cpu/libtorch-macos-1.13.1.zip
        unzip -qq -o libtorch-macos-1.13.1.zip -d "$DIR/$LIBRARY_DIR"/macos
        rm libtorch-macos-1.13.1.zip
    fi
fi

rm -f libtorch
ln -s ${LIBTORCH_DIR} libtorch

popd