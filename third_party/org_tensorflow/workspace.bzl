load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@com_github_wamuir_graft//tools/build_defs/repo:http.bzl", "http_embedded_archive")

def tf_repositories():
    ###########################################################################
    # libtensorflow
    ###########################################################################

    http_archive(
        name = "libtensorflow_linux_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        urls = ["https://storage.googleapis.com/libtensorflow-nightly/latest/libtensorflow-cpu-linux-x86_64.tar.gz"],
    )

    http_archive(
        name = "libtensorflow_linux_x86_64_gpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        urls = ["https://storage.googleapis.com/libtensorflow-nightly/latest/libtensorflow-gpu-linux-x86_64.tar.gz"],
    )

    http_embedded_archive(
        name = "libtensorflow_macos_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        urls = ["https://storage.googleapis.com/libtensorflow-nightly/prod/tensorflow/release/macos/latest/macos_cpu_libtensorflow_binaries.tar.gz"],
        inner_archive = "lib_package/libtensorflow-cpu-darwin-x86_64.tar.gz",
    )

    http_embedded_archive(
        name = "libtensorflow_windows_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        urls = ["https://storage.googleapis.com/libtensorflow-nightly/prod/tensorflow/release/windows/latest/cpu/windows_cpu_libtensorflow_binaries.tar.gz"],
        inner_archive = "lib_package/libtensorflow-cpu-windows-x86_64.zip",
    )

    http_embedded_archive(
        name = "libtensorflow_windows_x86_64_gpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        urls = ["https://storage.googleapis.com/libtensorflow-nightly/prod/tensorflow/release/windows/latest/gpu/windows_gpu_libtensorflow_binaries.tar.gz"],
        inner_archive = "lib_package/libtensorflow-gpu-linux-x86_64.tar.gz",
    )

    ###########################################################################
    # protos
    ###########################################################################

    http_archive(
        name = "libtensorflow_proto_linux_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.BUILD",
        patch_args = ["-p1"],
        patches = ["@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.patch"],
        urls = ["https://storage.googleapis.com/libtensorflow-nightly/latest/libtensorflow_proto.zip"],
    )

    http_archive(
        name = "libtensorflow_proto_linux_x86_64_gpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.BUILD",
        patch_args = ["-p1"],
        patches = ["@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.patch"],
        urls = ["https://storage.googleapis.com/libtensorflow-nightly/latest/libtensorflow_proto.zip"],
    )

    http_archive(
        name = "libtensorflow_proto_macos_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.BUILD",
        patch_args = ["-p1"],
        patches = ["@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.patch"],
        urls = ["https://storage.googleapis.com/libtensorflow-nightly/latest/libtensorflow_proto.zip"],
    )

    # NOTE: Windows libtensorflow nightly archives do not include protos (rely instead on Linux CPU nightly archive)

    http_embedded_archive(
        name = "libtensorflow_proto_windows_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.BUILD",
        patch_args = ["-p1"],
        patches = ["@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.patch"],
        urls = ["https://storage.googleapis.com/libtensorflow-nightly/prod/tensorflow/release/ubuntu_16/latest/cpu/ubuntu_cpu_libtensorflow_binaries.tar.gz"],
        inner_archive = "lib_package/libtensorflow_proto.zip",
    )

    http_embedded_archive(
        name = "libtensorflow_proto_windows_x86_64_gpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.BUILD",
        patch_args = ["-p1"],
        patches = ["@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.patch"],
        urls = ["https://storage.googleapis.com/libtensorflow-nightly/prod/tensorflow/release/ubuntu_16/latest/cpu/ubuntu_cpu_libtensorflow_binaries.tar.gz"],
        inner_archive = "lib_package/libtensorflow_proto.zip",
    )

    ###########################################################################
    # base api def
    ###########################################################################

    http_archive(
        name = "tensorflow_base_api_def",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/tensorflow_base_api_def:tensorflow_base_api_def.BUILD",
        strip_prefix = "tensorflow-nightly/tensorflow/core/api_def/base_api",
        urls = ["https://github.com/tensorflow/tensorflow/archive/refs/heads/nightly.tar.gz"],
    )

    ###########################################################################
    # go bindings
    ###########################################################################

    http_archive(
        name = "tensorflow_go",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/tensorflow_go:tensorflow_go.BUILD",
        strip_prefix = "tensorflow-nightly/tensorflow/go",
        urls = ["https://github.com/tensorflow/tensorflow/archive/refs/heads/nightly.tar.gz"],
    )
