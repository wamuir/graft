load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def tf_repositories(ctx):
    ###########################################################################
    # libtensorflow
    ###########################################################################

    http_archive(
        name = "libtensorflow_linux_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        patch_args = ["-p1"],
        patches = [],
        strip_prefix = "",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-2.14.0.tar.gz"],
        sha256 = "07b22de022eae9121618b47520bb2e68035b808d13fade907360b088b8807420",
    )

    http_archive(
        name = "libtensorflow_linux_x86_64_gpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        patch_args = ["-p1"],
        patches = [],
        strip_prefix = "",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-gpu-linux-x86_64-2.14.0.tar.gz"],
        sha256 = "d14e7f2130c94fc9f350a182cf818bb063d7b5aefd82474621df662044850d54",
    )

    http_archive(
        name = "libtensorflow_macos_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        patch_args = ["-p1"],
        patches = [],
        strip_prefix = "",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-darwin-x86_64-2.14.0.tar.gz"],
        sha256 = "1d2cce252d8d605c277aff2edc3aaef4f9da180711294b5ad782d0a25127a366",
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
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow_proto-2.14.0.zip"],
        sha256 = "14ee56e7e19098d98089007cf8a6740bfd0869fe5b53b49ff32ba1fe7e81de09",
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
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow_proto-2.14.0.zip"],
        sha256 = "14ee56e7e19098d98089007cf8a6740bfd0869fe5b53b49ff32ba1fe7e81de09",
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
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow_proto-2.14.0.zip"],
        sha256 = "14ee56e7e19098d98089007cf8a6740bfd0869fe5b53b49ff32ba1fe7e81de09",
    )

    ###########################################################################
    # base api def
    ###########################################################################

    http_archive(
        name = "tensorflow_base_api_def",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/tensorflow_base_api_def:tensorflow_base_api_def.BUILD",
        patch_args = ["-p1"],
        patches = [],
        strip_prefix = "tensorflow-2.14.0/tensorflow/core/api_def/base_api",
        urls = ["https://github.com/tensorflow/tensorflow/archive/refs/tags/v2.14.0.tar.gz"],
        sha256 = "ce357fd0728f0d1b0831d1653f475591662ec5bca736a94ff789e6b1944df19f",
    )

    ###########################################################################
    # go bindings
    ###########################################################################

    http_archive(
        name = "tensorflow_go",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/tensorflow_go:tensorflow_go.BUILD",
        patch_args = ["-p1"],
        patches = [
            "@com_github_wamuir_graft//third_party/org_tensorflow/tensorflow_go:tensorflow_go_op.patch",
        ],
        strip_prefix = "tensorflow-2.14.0/tensorflow/go",
        urls = ["https://github.com/tensorflow/tensorflow/archive/refs/tags/v2.14.0.tar.gz"],
        sha256 = "ce357fd0728f0d1b0831d1653f475591662ec5bca736a94ff789e6b1944df19f",
    )

download_tf_repositories = module_extension(
    implementation = tf_repositories,
)
