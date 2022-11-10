load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "test_log_proto_go_proto",
    srcs = ["test_log.pb.go"],
    importpath = "github.com/wamuir/graft/tensorflow/core/util/test_log_proto_go_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//tensorflow/tsl/protobuf/for_core_protos_go_proto",
        "@org_golang_google_protobuf//reflect/protoreflect:go_default_library",
        "@org_golang_google_protobuf//runtime/protoimpl:go_default_library",
    ],
)

alias(
    name = "go_default_library",
    actual = ":test_log_proto_go_proto",
    visibility = ["//visibility:public"],
)
