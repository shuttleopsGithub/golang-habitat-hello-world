#!/bin/bash

pkg_version="0.1.0"
pkg_name=golang-habitat-hello-world
pkg_origin=shuttleops-golang-example
pkg_scaffolding="core/scaffolding-go"
pkg_bin_dirs=(bin)

do_check() {
    go test main.go main_test.go
}