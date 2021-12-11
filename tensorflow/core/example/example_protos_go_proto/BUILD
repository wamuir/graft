load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "example_protos_go_proto",
    srcs = [
        "example.pb.go",
        "feature.pb.go",
    ],
    importpath = "github.com/wamuir/graft/tensorflow/core/example/example_protos_go_proto",
    visibility = ["//visibility:public"],
    deps = [
        "@org_golang_google_protobuf//reflect/protoreflect:go_default_library",
        "@org_golang_google_protobuf//runtime/protoimpl:go_default_library",
    ],
)

alias(
    name = "go_default_library",
    actual = ":example_protos_go_proto",
    visibility = ["//visibility:public"],
)
