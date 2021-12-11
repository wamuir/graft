load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "log_memory_go_proto",
    srcs = ["log_memory.pb.go"],
    importpath = "github.com/wamuir/graft/tensorflow/core/framework/log_memory_go_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//tensorflow/core/framework/tensor_description_go_proto",
        "@org_golang_google_protobuf//reflect/protoreflect:go_default_library",
        "@org_golang_google_protobuf//runtime/protoimpl:go_default_library",
    ],
)

alias(
    name = "go_default_library",
    actual = ":log_memory_go_proto",
    visibility = ["//visibility:public"],
)
