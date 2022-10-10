load("@rules_pkg//:pkg.bzl", "pkg_tar")

package(default_visibility = ["//visibility:private"])

alias(
    name = "libtensorflow_proto",
    actual = "//third_party/org_tensorflow:libtensorflow_proto",
    visibility = ["//visibility:public"],
)

alias(
    name = "libtensorflow",
    actual = "//third_party/org_tensorflow:libtensorflow_pkg",
    visibility = ["//visibility:public"],
)

pkg_tar(
    name = "libtensorflow_go",
    strip_prefix = "/",
    visibility = ["//visibility:public"],
    deps = [
        ":libtensorflow",
        "//third_party/org_tensorflow:go_bindings",
        "//tools/genproto:go_protos",
    ],
)

sh_test(
    name = "libtensorflow_go_test",
    size = "small",
    srcs = ["libtensorflow_go_test.sh"],
    data = [
        "_libtensorflow_go_hello.go",
        ":libtensorflow_go",
    ],
)
