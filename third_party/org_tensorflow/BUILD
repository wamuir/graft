load("@rules_pkg//:pkg.bzl", "pkg_tar")
load("//third_party/org_tensorflow:version.bzl", "VERSION_MAJOR", "VERSION_MINOR", "VERSION_PATCH")

package(
    default_visibility = ["//visibility:private"],
)

alias(
    name = "libtensorflow",
    actual = select({
        "//tools/conditions:linux_x86_64_cpu": "@libtensorflow_linux_x86_64_cpu//:libtensorflow",
        "//tools/conditions:linux_x86_64_gpu": "@libtensorflow_linux_x86_64_gpu//:libtensorflow",
        "//tools/conditions:macos_x86_64_cpu": "@libtensorflow_macos_x86_64_cpu//:libtensorflow",
        "//conditions:default": "@libtensorflow_other_build//:libtensorflow",
    }),
    visibility = ["//:__subpackages__"],
)

pkg_tar(
    name = "libtensorflow_proto",
    srcs = select({
        "//tools/conditions:linux_x86_64_cpu": ["@libtensorflow_proto_linux_x86_64_cpu//:protos_all_proto_srcs"],
        "//tools/conditions:linux_x86_64_gpu": ["@libtensorflow_proto_linux_x86_64_gpu//:protos_all_proto_srcs"],
        "//tools/conditions:macos_x86_64_cpu": ["@libtensorflow_proto_macos_x86_64_cpu//:protos_all_proto_srcs"],
        "//conditions:default": ["@libtensorflow_proto_linux_x86_64_cpu//:protos_all_proto_srcs"],
    }),
    remap_paths = {
        "external/_main~download_tf_repositories~libtensorflow_proto_linux_x86_64_cpu": "",
    },
    strip_prefix = select({
        "//tools/conditions:linux_x86_64_cpu": "/external/libtensorflow_proto_linux_x86_64_cpu",
        "//tools/conditions:linux_x86_64_gpu": "/external/libtensorflow_proto_linux_x86_64_gpu",
        "//tools/conditions:macos_x86_64_cpu": "/external/libtensorflow_proto_macos_x86_64_cpu",
        "//conditions:default": "/external/libtensorflow_proto_linux_x86_64_cpu",
    }),
    visibility = ["//:__subpackages__"],
)

pkg_tar(
    name = "libtensorflow_pkg",
    srcs = select({
        "//tools/conditions:linux_x86_64_cpu": ["@libtensorflow_linux_x86_64_cpu//:pkg_files"],
        "//tools/conditions:linux_x86_64_gpu": ["@libtensorflow_linux_x86_64_gpu//:pkg_files"],
        "//tools/conditions:macos_x86_64_cpu": ["@libtensorflow_macos_x86_64_cpu//:pkg_files"],
        "//conditions:default": ["@libtensorflow_other_build//:pkg_files"],
    }),
    remap_paths = {
        "external/_main~download_tf_repositories~libtensorflow_linux_x86_64_cpu": "",
        "external/_main~download_tf_repositories~libtensorflow_linux_x86_64_gpu": "",
        "external/_main~download_tf_repositories~libtensorflow_macos_x86_64_cpu": "",
        "external/_main~download_tf_repositories~libtensorflow_other_build": "",
    },
    strip_prefix = select({
        "//tools/conditions:linux_x86_64_cpu": "/external/libtensorflow_linux_x86_64_cpu",
        "//tools/conditions:linux_x86_64_gpu": "/external/libtensorflow_linux_x86_64_gpu",
        "//tools/conditions:macos_x86_64_cpu": "/external/libtensorflow_macos_x86_64_cpu",
        "//conditions:default": "/external/libtensorflow_other_build",
    }),
    symlinks = select({
        "@platforms//os:linux": {
            "lib/libtensorflow.so": "libtensorflow.so.%s" % VERSION_MAJOR,
            "lib/libtensorflow_framework.so": "libtensorflow_framework.so.%s" % VERSION_MAJOR,
            "lib/libtensorflow.so.%s" % VERSION_MAJOR: "libtensorflow.so.%s.%s.%s" % (VERSION_MAJOR, VERSION_MINOR, VERSION_PATCH),
            "lib/libtensorflow_framework.so.%s" % VERSION_MAJOR: "libtensorflow_framework.so.%s.%s.%s" % (VERSION_MAJOR, VERSION_MINOR, VERSION_PATCH),
        },
        "@platforms//os:macos": {
            "lib/libtensorflow.dylib": "libtensorflow.%s.dylib" % VERSION_MAJOR,
            "lib/libtensorflow_framework.dylib": "libtensorflow_framework.%s.dylib" % VERSION_MAJOR,
            "lib/libtensorflow.%s.dylib" % VERSION_MAJOR: "libtensorflow.%s.%s.%s.dylib" % (VERSION_MAJOR, VERSION_MINOR, VERSION_PATCH),
            "lib/libtensorflow_framework.%s.dylib" % VERSION_MAJOR: "libtensorflow_framework.%s.%s.%s.dylib" % (VERSION_MAJOR, VERSION_MINOR, VERSION_PATCH),
        },
    }),
    visibility = ["//:__subpackages__"],
)

pkg_tar(
    name = "go_bindings",
    srcs = ["@tensorflow_go//:go_source"],
    mode = "0644",
    package_dir = "src/github.com/tensorflow/tensorflow/tensorflow/go",
    remap_paths = {
        "external/_main~download_tf_repositories~tensorflow_go": "",
    },
    strip_prefix = "/external/tensorflow_go",
    visibility = ["//:__subpackages__"],
)

alias(
    name = "base_api_def",
    actual = "@tensorflow_base_api_def//:all_pbtxt",
    visibility = ["//:__subpackages__"],
)
