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
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-2.9.2.tar.gz"],
        sha256 = "69acedca22272f45f79bfd0738aca19a06ebe2c95f1080d63d1e7ccb0186ca35",
    )

    http_archive(
        name = "libtensorflow_linux_x86_64_gpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        patch_args = ["-p1"],
        patches = [],
        strip_prefix = "",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-gpu-linux-x86_64-2.9.2.tar.gz"],
        sha256 = "09fd0f9ce7e6a283bfba732887459a49aa05a217f92add911b3e6de20f087275",
    )

    http_archive(
        name = "libtensorflow_macos_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        patch_args = ["-p1"],
        patches = [],
        strip_prefix = "",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-darwin-x86_64-2.9.2.tar.gz"],
        sha256 = "5d9dc806fe6b1d6f1bacfaee7f438e3e2c6457b73bbc899d49d2cbca66a1fb38",
    )

    ###########################################################################
    # protos
    ###########################################################################

    http_archive(
        name = "libtensorflow_proto_linux_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.BUILD",
        patch_args = ["-p1"],
        patches = ["@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.patch"],
        strip_prefix = "",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow_proto-2.9.2.zip"],
        sha256 = "c03af8326bc767fdb79f8e796e97a221bee0abb8d5528f057c5ebe40a64d7542",
    )

    http_archive(
        name = "libtensorflow_proto_linux_x86_64_gpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.BUILD",
        patch_args = ["-p1"],
        patches = ["@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.patch"],
        strip_prefix = "",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow_proto-2.9.2.zip"],
        sha256 = "c03af8326bc767fdb79f8e796e97a221bee0abb8d5528f057c5ebe40a64d7542",
    )

    http_archive(
        name = "libtensorflow_proto_macos_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.BUILD",
        patch_args = ["-p1"],
        patches = ["@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.patch"],
        strip_prefix = "",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow_proto-2.9.2.zip"],
        sha256 = "c03af8326bc767fdb79f8e796e97a221bee0abb8d5528f057c5ebe40a64d7542",
    )

    ###########################################################################
    # base api def
    ###########################################################################

    http_archive(
        name = "tensorflow_base_api_def",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/tensorflow_base_api_def:tensorflow_base_api_def.BUILD",
        patch_args = ["-p1"],
        patches = [],
        strip_prefix = "tensorflow-2.9.2/tensorflow/core/api_def/base_api",
        urls = ["https://github.com/tensorflow/tensorflow/archive/refs/tags/v2.9.2.tar.gz"],
        sha256 = "8cd7ed82b096dc349764c3369331751e870d39c86e73bbb5374e1664a59dcdf7",
    )

    ###########################################################################
    # go bindings
    ###########################################################################

    http_archive(
        name = "tensorflow_go",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/tensorflow_go:tensorflow_go.BUILD",
        patch_args = ["-p1"],
        patches = [],
        strip_prefix = "tensorflow-2.9.2/tensorflow/go",
        urls = ["https://github.com/tensorflow/tensorflow/archive/refs/tags/v2.9.2.tar.gz"],
        sha256 = "8cd7ed82b096dc349764c3369331751e870d39c86e73bbb5374e1664a59dcdf7",
    )
