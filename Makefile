.DEFAULT_GOAL := default

.PHONY: default
default: libtorch_install build

.PHONY: build
build:
	go build -o service ./cmd

.PHONY: build_cuda
build_cuda:
	CGO_CPPFLAG="-DWITH_CUDA -I/usr/local/cuda/include" \
	go build -o service ./cmd

.PHONY: libtorch_clean
libtorch_clean:
	make --directory=./internal/cgolibrary clean

.PHONY: libtorch_install
libtorch_install: libtorch_clean
	./internal/cgolibrary/libtorch_install.sh

.PHONY: libtorch_build
libtorch_build: libtorch_install
	./internal/cgolibrary/libtorch_build.sh