load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@com_github_wamuir_graft//tools/build_defs/repo:http.bzl", "http_embedded_archive")

def tf_repositories():
    ###########################################################################
    # libtensorflow
    ###########################################################################

    http_archive(
        name = "libtensorflow_linux_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-2.9.1.tar.gz"],
        sha256 = "7f510d2546e3db5e10b1d119463689003d587842892ad3c95bca51cf82820173",
    )

    http_archive(
        name = "libtensorflow_linux_x86_64_gpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-gpu-linux-x86_64-2.9.1.tar.gz"],
        sha256 = "20c46f761d5905f592c4eac0812e8e50d9cb439f9440a9391912cf4c02ed8be8",
    )

    http_archive(
        name = "libtensorflow_macos_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-darwin-x86_64-2.9.1.tar.gz"],
        sha256 = "56f83d3320ba0f69aaafd6c058645e1b7fbc75fb9a3fc36d6be98de1291ca226",
    )

    ###########################################################################
    # protos
    ###########################################################################

    http_archive(
        name = "libtensorflow_proto_linux_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.BUILD",
        patch_args = ["-p1"],
        patches = ["@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.patch"],
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow_proto-2.9.1.zip"],
        sha256 = "9feb525216adc6848a42359c8157185a21421e27d9a9d7c57ed2646a8bef38cc",
    )

    http_archive(
        name = "libtensorflow_proto_linux_x86_64_gpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.BUILD",
        patch_args = ["-p1"],
        patches = ["@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.patch"],
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow_proto-2.9.1.zip"],
        sha256 = "9feb525216adc6848a42359c8157185a21421e27d9a9d7c57ed2646a8bef38cc",
    )

    http_archive(
        name = "libtensorflow_proto_macos_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.BUILD",
        patch_args = ["-p1"],
        patches = ["@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.patch"],
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow_proto-2.9.1.zip"],
        sha256 = "9feb525216adc6848a42359c8157185a21421e27d9a9d7c57ed2646a8bef38cc",
    )

    ###########################################################################
    # base api def
    ###########################################################################

    http_archive(
        name = "tensorflow_base_api_def",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/tensorflow_base_api_def:tensorflow_base_api_def.BUILD",
        sha256 = "6eaf86ead73e23988fe192da1db68f4d3828bcdd0f3a9dc195935e339c95dbdc",
        strip_prefix = "tensorflow-2.9.1/tensorflow/core/api_def/base_api",
        urls = ["https://github.com/tensorflow/tensorflow/archive/refs/tags/v2.9.1.tar.gz"],
    )

    ###########################################################################
    # go bindings
    ###########################################################################

    http_archive(
        name = "tensorflow_go",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/tensorflow_go:tensorflow_go.BUILD",
        sha256 = "6eaf86ead73e23988fe192da1db68f4d3828bcdd0f3a9dc195935e339c95dbdc",
        strip_prefix = "tensorflow-2.9.1/tensorflow/go",
        urls = ["https://github.com/tensorflow/tensorflow/archive/refs/tags/v2.9.1.tar.gz"],
    )
