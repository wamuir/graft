load("@rules_pkg//:pkg.bzl", "pkg_tar")

package(
    default_visibility = ["//visibility:private"],
)

alias(
    name = "libtensorflow",
    actual = select({
        "//tools/conditions:linux_x86_64_cpu": "@libtensorflow_linux_x86_64_cpu//:libtensorflow",
        "//tools/conditions:linux_x86_64_gpu": "@libtensorflow_linux_x86_64_gpu//:libtensorflow",
        "//tools/conditions:macos_x86_64_cpu": "@libtensorflow_macos_x86_64_cpu//:libtensorflow",
    }),
    visibility = ["//:__subpackages__"],
)

pkg_tar(
    name = "libtensorflow_proto",
    srcs = select({
        "//tools/conditions:linux_x86_64_cpu": ["@libtensorflow_proto_linux_x86_64_cpu//:protos_all_proto_srcs"],
        "//tools/conditions:linux_x86_64_gpu": ["@libtensorflow_proto_linux_x86_64_gpu//:protos_all_proto_srcs"],
        "//tools/conditions:macos_x86_64_cpu": ["@libtensorflow_proto_macos_x86_64_cpu//:protos_all_proto_srcs"],
    }),
    extension = ".tar.gz",
    strip_prefix = select({
        "//tools/conditions:linux_x86_64_cpu": "/external/libtensorflow_proto_linux_x86_64_cpu",
        "//tools/conditions:linux_x86_64_gpu": "/external/libtensorflow_proto_linux_x86_64_gpu",
        "//tools/conditions:macos_x86_64_cpu": "/external/libtensorflow_proto_macos_x86_64_cpu",
    }),
    visibility = ["//:__subpackages__"],
)

pkg_tar(
    name = "libtensorflow_pkg",
    srcs = select({
        "//tools/conditions:linux_x86_64_cpu": ["@libtensorflow_linux_x86_64_cpu//:all_files"],
        "//tools/conditions:linux_x86_64_gpu": ["@libtensorflow_linux_x86_64_gpu//:all_files"],
        "//tools/conditions:macos_x86_64_cpu": ["@libtensorflow_macos_x86_64_cpu//:all_files"],
    }),
    extension = ".tar.gz",
    strip_prefix = select({
        "//tools/conditions:linux_x86_64_cpu": "/external/libtensorflow_linux_x86_64_cpu",
        "//tools/conditions:linux_x86_64_gpu": "/external/libtensorflow_linux_x86_64_gpu",
        "//tools/conditions:macos_x86_64_cpu": "/external/libtensorflow_macos_x86_64_cpu",
    }),
    visibility = ["//:__subpackages__"],
)

pkg_tar(
    name = "go_bindings",
    srcs = ["@tensorflow_go//:go_source"],
    mode = "0644",
    package_dir = "src/github.com/tensorflow/tensorflow/tensorflow/go",
    strip_prefix = "/external/tensorflow_go",
    visibility = ["//:__subpackages__"],
)

alias(
    name = "base_api_def",
    actual = "@tensorflow_base_api_def//:all_pbtxt",
    visibility = ["//:__subpackages__"],
)
