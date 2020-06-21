#!/bin/bash

pkg_name=golang-habitat-hello-world
pkg_version=0.1.0
pkg_build_deps=(
    core/go
    core/git
)
pkg_deps=(
    core/curl
)
pkg_bin_dirs=(
    bin
)

do_build() {
    export GOBIN="${GOPATH}/bin"
    go get
    go build -o "${pkg_name}"
}

do_install() {
    mkdir -p "${pkg_prefix}/bin"
    mv "${SRC_PATH}/${pkg_name}" "${pkg_prefix}/bin/"
}