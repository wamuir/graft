#!/bin/sh

set -ex

BZL_MODULE="$(git rev-parse --show-toplevel)/MODULE.bazel"
TF_VERSION=${1}

if [ -z "${TF_VERSION}" ]; then
    printf 'ERROR: version required as the first positional argument.\n' >&2
    exit 1
fi

sed -i s/\"0\.0\.0\"/\"${TF_VERSION}\"/ ${BZL_MODULE}
