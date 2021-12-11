# Description:
# Generation of Go source files for TensorFlow ops.

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

package(default_visibility = ["//visibility:private"])

go_library(
    name = "lib",
    srcs = ["main.go"],
    importpath = "github.com/wamuir/graft/tools/genop",
    deps = ["//tools/genop/internal"],
)

go_binary(
    name = "genop",
    embed = [":lib"],
    visibility = ["//:__subpackages__"],
)
