load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "op",
    srcs = [
        "gradients.go",
        "op.go",
        "scope.go",
        ":wrappers",  # keep
    ],
    importpath = "github.com/wamuir/graft/tensorflow/op",
    visibility = ["//visibility:public"],
    deps = ["//tensorflow"],
)

go_test(
    name = "op_test",
    size = "small",  # keep
    srcs = [
        "gradients_test.go",
        "op_test.go",
        "scope_test.go",
    ],
    embed = [":op"],
    deps = ["//tensorflow"],
)

genrule(
    name = "wrappers",
    srcs = ["//third_party/org_tensorflow:base_api_def"],
    outs = ["wrappers.go"],
    cmd = "$(execpath //tools/genop) $(SRCS) > $@",
    tools = ["//tools/genop"],
)

alias(
    name = "go_default_library",
    actual = ":op",
    visibility = ["//visibility:public"],
)
