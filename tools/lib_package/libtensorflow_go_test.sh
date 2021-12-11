#!/usr/bin/env bash
# Copyright 2021 William Muir. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# =============================================================================

set -exo pipefail

# Sanity test for the package Go library archive.
# - Unarchive
# - Compile a trivial Go file that uses the archive
# - Run it

# Tools needed: go and tar
[ -z "${GO}" ] && GO=$(command -v go)
[ -z "${TAR}" ] && TAR=$(command -v tar)

# Bazel tests run with ${PWD} set to the root of the bazel workspace
GOTESTFILE="${PWD}/tools/lib_package/libtensorflow_go_hello.go"
LIBARCHIVE="${PWD}/tools/lib_package/libtensorflow_go.tar.gz"

# Go needs GOPATH and GOCACHE set; assign a temp directory for each
export GOPATH="$(mktemp --directory --tmpdir="${TEST_TMPDIR}")"
export GOCACHE="$(mktemp --directory --tmpdir="${TEST_TMPDIR}")"

# Extract the libtensorflow_go archive into GOPATH
${TAR} -xzf "${LIBARCHIVE}" -C"${GOPATH}"
(
  cd "${GOPATH}/src/github.com/tensorflow/tensorflow" \
  && ${GO} mod init github.com/tensorflow/tensorflow \
  && ${GO} mod tidy
)

# Build and execute the trivial Go application, with the TensorFlow Go bindings
# and C API (shared library) available. Note that DYLD_LIBRARY_PATH is used on
# OS X, LD_LIBRARY_PATH on Linux.
#
# The tests for GPU require CUDA libraries to be accessible, which are in
# DYLD_LIBRARY_PATH in the test harness for OS X.
export DYLD_LIBRARY_PATH="${GOPATH}/lib:${DYLD_LIBRARY_PATH}"
export LD_LIBRARY_PATH="${GOPATH}/lib:${LD_LIBRARY_PATH}"

BINDINGS="${GOPATH}/src/github.com/tensorflow/tensorflow"
GOTESTPATH="${GOPATH}/src/github.com/tensorflow/lib_package/libtensorflow_go_test"
install -D -m400  "${GOTESTFILE}" "${GOTESTPATH}/main.go"
(
  cd "${GOTESTPATH}" \
  && ${GO} mod init libtensorflow_go \
  && ${GO} mod edit -require github.com/tensorflow/tensorflow@v0.0.0 \
  && ${GO} mod edit -replace github.com/tensorflow/tensorflow="${BINDINGS}" \
  && ${GO} mod tidy \
  && CGO_CFLAGS="-I${GOPATH}/include" \
     CGO_LDFLAGS="-L${GOPATH}/lib" \
     ${GO} build -o a.out \
  && ./a.out
)
