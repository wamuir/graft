load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "model_go_proto",
    srcs = ["model.pb.go"],
    importpath = "github.com/wamuir/graft/tensorflow/core/framework/model_go_proto",
    visibility = ["//visibility:public"],
    deps = [
        "@org_golang_google_protobuf//reflect/protoreflect:go_default_library",
        "@org_golang_google_protobuf//runtime/protoimpl:go_default_library",
    ],
)

alias(
    name = "go_default_library",
    actual = ":model_go_proto",
    visibility = ["//visibility:public"],
)
