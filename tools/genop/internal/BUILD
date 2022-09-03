load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:private"])

go_library(
    name = "internal",
    srcs = [
        "api_def_map.go",
        "genop.go",
        "lib.go",
    ],
    cdeps = ["//third_party/org_tensorflow:libtensorflow"],  # keep
    cgo = True,
    clinkopts = [],  # keep
    copts = [],  # keep
    importpath = "github.com/wamuir/graft/tools/genop/internal",
    visibility = ["//tools/genop:__pkg__"],
    deps = [
        "//tensorflow/core/framework/api_def_go_proto",
        "//tensorflow/core/framework/op_def_go_proto",
        "@org_golang_google_protobuf//encoding/prototext:go_default_library",
        "@org_golang_google_protobuf//proto:go_default_library",
    ],
)

go_test(
    name = "internal_test",
    size = "small",
    srcs = ["genop_test.go"],
    embed = [":internal"],
    deps = [
        "//tensorflow/core/framework/api_def_go_proto",
        "//tensorflow/core/framework/op_def_go_proto",
        "@org_golang_google_protobuf//encoding/prototext:go_default_library",
    ],
)

alias(
    name = "go_default_library",
    actual = ":internal",
    visibility = ["//tools/genop:__pkg__"],
)
