load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "tensor_go_proto",
    srcs = ["tensor.pb.go"],
    importpath = "github.com/wamuir/graft/tensorflow/core/framework/tensor_go_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//tensorflow/core/framework/resource_handle_go_proto",
        "//tensorflow/core/framework/tensor_shape_go_proto",
        "//tensorflow/core/framework/types_go_proto",
        "@org_golang_google_protobuf//reflect/protoreflect:go_default_library",
        "@org_golang_google_protobuf//runtime/protoimpl:go_default_library",
    ],
)

alias(
    name = "go_default_library",
    actual = ":tensor_go_proto",
    visibility = ["//visibility:public"],
)
