load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "graph_go_proto",
    srcs = ["graph.pb.go"],
    importpath = "github.com/wamuir/graft/tensorflow/core/framework/graph_go_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//tensorflow/core/framework/function_go_proto",
        "//tensorflow/core/framework/node_def_go_proto",
        "//tensorflow/core/framework/versions_go_proto",
        "@org_golang_google_protobuf//reflect/protoreflect:go_default_library",
        "@org_golang_google_protobuf//runtime/protoimpl:go_default_library",
    ],
)

alias(
    name = "go_default_library",
    actual = ":graph_go_proto",
    visibility = ["//visibility:public"],
)
