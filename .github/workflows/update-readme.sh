#!/bin/sh

set -ex

BRANCH=$(git symbolic-ref --short HEAD)
TF_VERSION=${1}

if [ -z "${TF_VERSION}" ]; then
    printf 'ERROR: version required as the first positional argument.\n' >&2
    exit 1
fi

cd "$(git rev-parse --show-toplevel)"
README=$(awk 'NR>4' README.md)
cat << EOF > README.md
# Graft
[![tensorflow version](https://img.shields.io/badge/tf-v${TF_VERSION}-FF6F00?logo=tensorflow&logoColor=FF6F00)](https://github.com/tensorflow/tensorflow/tree/v${TF_VERSION})
[![build](https://img.shields.io/github/actions/workflow/status/wamuir/graft/build-and-test-bindings.yml?branch=${BRANCH}&label=build&logo=github)](https://github.com/wamuir/graft/actions/workflows/test.yml?query=branch%3A${BRANCH})
[![go.dev reference](https://pkg.go.dev/badge/wamuir/graft)](https://pkg.go.dev/github.com/wamuir/graft/tensorflow)
EOF
printf "%s\n" "${README}" >> README.md
