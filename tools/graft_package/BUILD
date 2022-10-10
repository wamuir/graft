load("@rules_pkg//:pkg.bzl", "pkg_tar")

package(default_visibility = ["//visibility:private"])

pkg_tar(
    name = "libtensorflow_graft",
    strip_prefix = "/",
    visibility = ["//visibility:public"],
    deps = [
        ":go_bindings",
        "//tools/lib_package:libtensorflow",
    ],
)

genrule(
    name = "go_bindings",
    srcs = [
        "//third_party/org_tensorflow:go_bindings",
        "//tools/genproto:go_protos",
        "//tools/graft_package/overlay",
    ],
    outs = ["go_bindings.tar"],
    cmd = """
set -euo pipefail

# Create a temporary work directory.
TMP=$$(mktemp -d)
trap "rm -rf $${TMP}" EXIT

# Untar Go bindings and protos.
tar xf $(location //third_party/org_tensorflow:go_bindings) --directory=$${TMP}
tar xf $(location //tools/genproto:go_protos) --directory=$${TMP}

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
tar cf $@ --mtime='UTC 1970-01-01' -C $${TMP} .
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
