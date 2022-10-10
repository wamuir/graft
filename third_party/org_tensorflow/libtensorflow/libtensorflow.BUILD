load("@rules_pkg//:pkg.bzl", "pkg_tar")

package(
    default_visibility = ["//visibility:private"],
)

VERSION_MAJOR = 2

ALL_FILES = glob(
    ["**/*"],
    exclude = [
        "BUILD.bazel",
        "WORKSPACE",
    ],
)

SYMLINKS = select({
        "@platforms//os:linux": [
		"BUILD.bazel",
		"WORKSPACE",
		"libtensorflow.so",
		"libtensorflow.so.%s" % VERSION_MAJOR,
		"libtensorflow_framework.so",
		"libtensorflow_framework.so.%s" % VERSION_MAJOR,
	],
        "@platforms//os:macos": [
		"BUILD.bazel",
		"WORKSPACE",
	],
    })

# exclude symlinks (symlinks not preserved by rule_pkg)
# github.com/bazelbuild/rules_pkg/issues/115
PKG_FILES = glob(
    ["**/*"],
    exclude = [
		"BUILD.bazel",
		"WORKSPACE",
		"lib/libtensorflow.dylib",
		"lib/libtensorflow.%s.dylib" % VERSION_MAJOR,
		"lib/libtensorflow_framework.dylib",
		"lib/libtensorflow_framework.%s.dylib" % VERSION_MAJOR,
		"lib/libtensorflow.so",
		"lib/libtensorflow.so.%s" % VERSION_MAJOR,
		"lib/libtensorflow_framework.so",
		"lib/libtensorflow_framework.so.%s" % VERSION_MAJOR,
	],
)

C_HEADERS = glob(["include/**/*.h"])

filegroup(
    name = "all_files",
    srcs = ALL_FILES,
    visibility = ["@com_github_wamuir_graft//third_party/org_tensorflow:__pkg__"],
)

filegroup(
    name = "pkg_files",
    srcs = PKG_FILES,
    visibility = ["@com_github_wamuir_graft//third_party/org_tensorflow:__pkg__"],
)

filegroup(
    name = "c_headers",
    srcs = C_HEADERS,
    visibility = ["@com_github_wamuir_graft//third_party/org_tensorflow:__pkg__"],
)

cc_import(
    name = "libtensorflow_import_lib",
    shared_library = select({
        "@platforms//os:linux": "lib/libtensorflow.so.%s" % VERSION_MAJOR,
        "@platforms//os:macos": "lib/libtensorflow.%s.dylib" % VERSION_MAJOR,
    }),
    visibility = ["@com_github_wamuir_graft//third_party/org_tensorflow:__pkg__"],
)

cc_import(
    name = "libtensorflow_framework_import_lib",
    shared_library = select({
        "@platforms//os:linux": "lib/libtensorflow_framework.so.%s" % VERSION_MAJOR,
        "@platforms//os:macos": "lib/libtensorflow_framework.%s.dylib" % VERSION_MAJOR,
    }),
    visibility = ["@com_github_wamuir_graft//third_party/org_tensorflow:__pkg__"],
)

cc_library(
    name = "libtensorflow",
    hdrs = C_HEADERS,
    includes = ["include"],
    visibility = ["@com_github_wamuir_graft//third_party/org_tensorflow:__pkg__"],
    deps = [
        ":libtensorflow_framework_import_lib",
        ":libtensorflow_import_lib",
    ],
)
