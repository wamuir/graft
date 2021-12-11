load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "api_def_go_proto",
    srcs = ["api_def.pb.go"],
    importpath = "github.com/wamuir/graft/tensorflow/core/framework/api_def_go_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//tensorflow/core/framework/attr_value_go_proto",
        "@org_golang_google_protobuf//reflect/protoreflect:go_default_library",
        "@org_golang_google_protobuf//runtime/protoimpl:go_default_library",
    ],
)

alias(
    name = "go_default_library",
    actual = ":api_def_go_proto",
    visibility = ["//visibility:public"],
)
