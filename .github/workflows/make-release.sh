#!/bin/sh

set -e

TF_VERSION=

while :; do
    case $1 in
        -v|--version)       # Takes an option argument, ensuring it has been specified.
            if [ -n "$2" ]; then
                TF_VERSION=$2
                shift
            else
                printf 'ERROR: "--version" requires a non-empty option argument.\n' >&2
                exit 1
            fi
            ;;
        --version=?*)
            TF_VERSION=${1#*=} # Delete everything up to "=" and assign the remainder.
            ;;
        --version=)         # Handle the case of an empty --version=
            printf 'ERROR: "--version" requires a non-empty option argument.\n' >&2
            exit 1
            ;;
        --)              # End of all options.
            shift
            break
            ;;
        -?*)
            printf 'WARN: Unknown option (ignored): %s\n' "$1" >&2
            ;;
        *)               # Default case: If no more options then break out of the loop.
            break
    esac

    shift
done

if [ -z "${TF_VERSION}" ]; then
    printf 'ERROR: "--version" requires a non-empty option argument.\n' >&2
    exit 1
fi

cd "$(git rev-parse --show-toplevel)" \
   && .github/workflows/update-org_tensorflow-workspace.sh "${TF_VERSION}" \
   && .github/workflows/update-readme.sh "${TF_VERSION}" \
   && .github/workflows/update-bindings.sh
