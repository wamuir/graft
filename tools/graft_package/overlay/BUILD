load("@bazel_skylib//rules:copy_file.bzl", "copy_file")
load("@rules_pkg//:pkg.bzl", "pkg_tar")

package(
    default_visibility = ["//:__subpackages__"],
)

licenses(["notice"])  # Apache 2.0

IN_FILES = glob(
    ["**/*.in"],
)

[
    copy_file(
        name = f.replace("/", "_"),
        src = f,
        out = f.replace(".in", ""),
    )
    for f in IN_FILES
]

pkg_tar(
    name = "overlay",
    srcs = glob(
        ["**/*"],
        exclude = ["BUILD"] + IN_FILES,
    ) + [f.replace(".in", "") for f in IN_FILES],
    extension = ".tar.gz",
    mode = "0644",
    package_dir = "src/github.com/wamuir/graft",
    strip_prefix = "/tools/graft_package/overlay",
)
