load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "cpp_shape_inference_go_proto",
    srcs = ["cpp_shape_inference.pb.go"],
    importpath = "github.com/wamuir/graft/tensorflow/python/framework/cpp_shape_inference_go_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//tensorflow/core/framework/full_type_go_proto",
        "//tensorflow/core/framework/tensor_shape_go_proto",
        "//tensorflow/core/framework/types_go_proto",
        "@org_golang_google_protobuf//reflect/protoreflect:go_default_library",
        "@org_golang_google_protobuf//runtime/protoimpl:go_default_library",
    ],
)

alias(
    name = "go_default_library",
    actual = ":cpp_shape_inference_go_proto",
    visibility = ["//visibility:public"],
)
