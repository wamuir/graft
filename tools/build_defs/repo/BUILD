load("@bazel_skylib//:bzl_library.bzl", "bzl_library")
load("@io_bazel_stardoc//stardoc:stardoc.bzl", "stardoc")

bzl_library(
    name = "utils",
    srcs = ["@bazel_tools//tools/build_defs/repo:utils.bzl"],
)

stardoc(
    name = "http_md",
    out = "http.md",
    input = "http.bzl",
    deps = [":utils"],
)
