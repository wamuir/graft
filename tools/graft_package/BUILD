package(default_visibility = ["//visibility:private"])

genrule(
    name = "libtensorflow_graft",
    srcs = [
        "//tools/lib_package:libtensorflow_go",
        "//tools/graft_package/overlay",
    ],
    outs = ["libtensorflow_graft.tar.gz"],
    cmd = """
set -euo pipefail

# Create a temporary work directory.
TMP=$$(mktemp -d)
trap "rm -rf $${TMP}" EXIT

# Untar Go bindings.
tar xzf $(location //tools/lib_package:libtensorflow_go) --directory=$${TMP}

# Move the bindings.
( \
  cd $${TMP}/src/github.com \
  && mkdir -p wamuir/graft \
  && mv tensorflow/tensorflow/tensorflow/go wamuir/graft/tensorflow \
  && rm -rf tensorflow \
)

# Untar overlay.
tar xzf $(location //tools/graft_package/overlay) --directory=$${TMP}

# Remove extraneous files.
( \
  cd $${TMP}/src/github.com/wamuir/graft/tensorflow \
  && rm -rf README.md test.sh genop op/generate.go  \
)

# Revise import statements.
find $${TMP}/src/github.com/wamuir/graft/tensorflow -type f -name '*.go' -print0 \
  | xargs -0 sed -i'' -e 's|github.com/tensorflow/tensorflow/tensorflow/go|github.com/wamuir/graft/tensorflow|g'


# Tar bindings.
tar czf $@ --mtime='UTC 1970-01-01' -C $${TMP} .
""",
    tags = ["manual"],
    visibility = ["//visibility:public"],
)

sh_test(
    name = "libtensorflow_graft_test",
    size = "small",
    srcs = ["libtensorflow_graft_test.sh"],
    data = [
        "_libtensorflow_graft_hello.go",
        ":libtensorflow_graft",
    ],
)
