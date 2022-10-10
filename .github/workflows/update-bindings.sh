#!/bin/sh

set -ex

bazelisk build //tools/graft_package:go_bindings

cd "$(git rev-parse --show-toplevel)" \
  && rm -rf tensorflow \
  && tar xz --strip-components=5 \
     -f bazel-bin/tools/graft_package/go_bindings.tar

bazelisk run //:gazelle
bazelisk run //:gazelle-update-repos
