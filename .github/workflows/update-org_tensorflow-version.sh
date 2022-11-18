#!/bin/sh

set -ex

VERSION_BZL="$(git rev-parse --show-toplevel)/third_party/org_tensorflow/version.bzl"
TF_VERSION=${1}

if [ -z "${TF_VERSION}" ]; then
    printf 'ERROR: version required as the first positional argument.\n' >&2
    exit 1
fi

cat <<EOF > "${VERSION_BZL}"
VERSION_MAJOR=$(echo "${TF_VERSION}" | cut -d'.' -f1)
VERSION_MINOR=$(echo "${TF_VERSION}" | cut -d'.' -f2)
VERSION_PATCH=$(echo "${TF_VERSION}" | cut -d'.' -f3)
EOF
