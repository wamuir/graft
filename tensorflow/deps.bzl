load("//third_party/org_tensorflow:workspace.bzl", _tf_repositories = "tf_repositories")
load("@bazel_gazelle//:deps.bzl", "go_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")

def _go_dependencies():
    go_repository(
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        sum = "h1:LUVKkCeviFUMKqHa4tXIIij/lbhnMbP7Fn5wKdKkRh4=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        sum = "h1:Khx7svrCpmxxtHBq5j2mp/xVjsi8hQMfNLvJFAlrGgU=",
        version = "v0.5.5",
    )
    go_repository(
        name = "org_golang_google_protobuf",
        importpath = "google.golang.org/protobuf",
        sum = "h1:d0NfwRgPtno5B1Wa6L2DAG+KivqkdutMf1UhdNx175w=",
        version = "v1.28.1",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        sum = "h1:E7g+9GITq07hpfrRu66IVDexMakfv52eLZ2CXBWiKr4=",
        version = "v0.0.0-20191204190536-9bdfabe68543",
    )

def graft_dependencies():
    """Call in WORKSPACE file to load @com_github_wamuir_graft dependencies."""

    maybe(
        name = "rules_pkg",
        repo_rule = http_archive(
            name = "rules_pkg",
            sha256 = "eea0f59c28a9241156a47d7a8e32db9122f3d50b505fae0f33de6ce4d9b61834",
            urls = [
                "https://mirror.bazel.build/github.com/bazelbuild/rules_pkg/releases/download/0.8.0/rules_pkg-0.8.0.tar.gz",
                "https://github.com/bazelbuild/rules_pkg/releases/download/0.8.0/rules_pkg-0.8.0.tar.gz",
            ],
        ),
    )

    maybe(
        name = "com_google_protobuf",
        repo_rule = http_archive(
            name = "com_google_protobuf",
            sha256 = "b1d6dd2cbb5d87e17af41cadb720322ce7e13af826268707bd8db47e5654770b",
            strip_prefix = "protobuf-21.11",
            urls = [
                "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v21.11.tar.gz",
                "https://github.com/protocolbuffers/protobuf/archive/v21.11.tar.gz",
            ],
        ),
    )

    _go_dependencies()
    _tf_repositories()
