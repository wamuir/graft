load("//third_party/org_tensorflow:workspace.bzl", _tf_repositories = "tf_repositories")
load("@bazel_gazelle//:deps.bzl", "go_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")

def _go_dependencies(ctx):
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
        sum = "h1:kPPoIgf3TsEvrm0PFe15JQ+570QVxYzEvvHqChK+cng=",
        version = "v1.30.0",
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
            sha256 = "8f9ee2dc10c1ae514ee599a8b42ed99fa262b757058f65ad3c384289ff70c4b8",
            urls = [
                "https://mirror.bazel.build/github.com/bazelbuild/rules_pkg/releases/download/0.9.1/rules_pkg-0.9.1.tar.gz",
                "https://github.com/bazelbuild/rules_pkg/releases/download/0.9.1/rules_pkg-0.9.1.tar.gz",
            ],
        ),
    )

    maybe(
        name = "com_google_protobuf",
        repo_rule = http_archive(
            name = "com_google_protobuf",
            sha256 = "0b0395d34e000f1229679e10d984ed7913078f3dd7f26cf0476467f5e65716f4",
            strip_prefix = "protobuf-23.2",
            urls = [
                "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v23.2.tar.gz",
                "https://github.com/protocolbuffers/protobuf/archive/v23.2.tar.gz",
            ],
        ),
    )

    _go_dependencies(ctx = None)
    _tf_repositories(ctx = None)

go_dependencies = module_extension(
    implementation = _go_dependencies,
)
