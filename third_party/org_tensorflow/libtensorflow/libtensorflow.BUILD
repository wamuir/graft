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

C_HEADERS = glob(["include/**/*.h"])

filegroup(
    name = "all_files",
    srcs = ALL_FILES,
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

cc_import(
    name = "libtensorflow_import_lib_win",
    interface_library = "lib/tensorflow.lib",
    shared_library = "lib/tensorflow.dll",
    visibility = ["@com_github_wamuir_graft//third_party/org_tensorflow:__pkg__"],
)

cc_library(
    name = "libtensorflow",
    hdrs = C_HEADERS,
    includes = ["include"],
    visibility = ["@com_github_wamuir_graft//third_party/org_tensorflow:__pkg__"],
    deps = select({
        "@platforms//os:linux": [
            ":libtensorflow_framework_import_lib",
            ":libtensorflow_import_lib",
        ],
        "@platforms//os:macos": [
            ":libtensorflow_framework_import_lib",
            ":libtensorflow_import_lib",
        ],
        "@platforms//os:windows": [
            ":libtensorflow_import_lib_win",
        ],
    }),
)
