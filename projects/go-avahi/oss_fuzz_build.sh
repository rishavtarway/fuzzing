#!/bin/bash -eu
# Build script for go-avahi integration
sudo apt-get update && sudo apt-get install -y libavahi-common-dev libavahi-client-dev
go get github.com/OpenPrinting/go-avahi
compile_native_go_fuzzer $SRC/openprinting-fuzzing/projects/go-avahi/fuzzer/fuzz_domain_test.go FuzzDomainNormalize fuzz_domain_normalize
compile_native_go_fuzzer $SRC/openprinting-fuzzing/projects/go-avahi/fuzzer/fuzz_domain_test.go FuzzDomainSplit fuzz_domain_split
