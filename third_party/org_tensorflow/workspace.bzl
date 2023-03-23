load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def tf_repositories():
    ###########################################################################
    # libtensorflow
    ###########################################################################

    http_archive(
        name = "libtensorflow_linux_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        patch_args = ["-p1"],
        patches = [],
        strip_prefix = "",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-2.11.1.tar.gz"],
        sha256 = "c994bd70960a611fa42e308c0483838a17936157b59b1ebe1966f061a519fae2",
    )

    http_archive(
        name = "libtensorflow_linux_x86_64_gpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        patch_args = ["-p1"],
        patches = [],
        strip_prefix = "",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-gpu-linux-x86_64-2.11.1.tar.gz"],
        sha256 = "9c101cb4ad3c05a1a11b327941bafb3c41f4f70a3a93d0d290ec0949f088eca2",
    )

    http_archive(
        name = "libtensorflow_macos_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        patch_args = ["-p1"],
        patches = [],
        strip_prefix = "",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-darwin-x86_64-2.11.1.tar.gz"],
        sha256 = "f4083e06c703f740e6e81c9a09bdb713c5680b48530478bf3f6d740d66da82d1",
    )

    ###########################################################################
    # protos
    ###########################################################################

    http_archive(
        name = "libtensorflow_proto_linux_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.BUILD",
        patch_args = ["-p1"],
        patches = [
            "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.patch",
            "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto_tsl.patch",
        ],
        strip_prefix = "",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow_proto-2.11.1.zip"],
        sha256 = "14f1894966a6f775bb627b13728049b23d9a32e4ffde1f5edb20352d419de413",
    )

    http_archive(
        name = "libtensorflow_proto_linux_x86_64_gpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.BUILD",
        patch_args = ["-p1"],
        patches = [
            "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.patch",
            "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto_tsl.patch",
        ],
        strip_prefix = "",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow_proto-2.11.1.zip"],
        sha256 = "14f1894966a6f775bb627b13728049b23d9a32e4ffde1f5edb20352d419de413",
    )

    http_archive(
        name = "libtensorflow_proto_macos_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.BUILD",
        patch_args = ["-p1"],
        patches = [
            "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.patch",
            "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto_tsl.patch",
        ],
        strip_prefix = "",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow_proto-2.11.1.zip"],
        sha256 = "14f1894966a6f775bb627b13728049b23d9a32e4ffde1f5edb20352d419de413",
    )

    ###########################################################################
    # base api def
    ###########################################################################

    http_archive(
        name = "tensorflow_base_api_def",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/tensorflow_base_api_def:tensorflow_base_api_def.BUILD",
        patch_args = ["-p1"],
        patches = [],
        strip_prefix = "tensorflow-2.11.1/tensorflow/core/api_def/base_api",
        urls = ["https://github.com/tensorflow/tensorflow/archive/refs/tags/v2.11.1.tar.gz"],
        sha256 = "624ed1cc170cdcc19e8a15d8cdde989a9a1c6b0534c90b38a6b2f06fb2963e5f",
    )

    ###########################################################################
    # go bindings
    ###########################################################################

    http_archive(
        name = "tensorflow_go",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/tensorflow_go:tensorflow_go.BUILD",
        patch_args = ["-p1"],
        patches = [],
        strip_prefix = "tensorflow-2.11.1/tensorflow/go",
        urls = ["https://github.com/tensorflow/tensorflow/archive/refs/tags/v2.11.1.tar.gz"],
        sha256 = "624ed1cc170cdcc19e8a15d8cdde989a9a1c6b0534c90b38a6b2f06fb2963e5f",
    )
