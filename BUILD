load("@bazel_gazelle//:def.bzl", "gazelle")

package(
    default_visibility = ["//visibility:private"],
)

# gazelle:prefix github.com/wamuir/graft
# gazelle:build_file_name BUILD
# gazelle:exclude tools/graft_package
# gazelle:exclude tools/lib_package
# gazelle:go_naming_convention import_alias
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=tensorflow/deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)
