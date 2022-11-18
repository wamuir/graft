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
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-2.10.1.tar.gz"],
        sha256 = "201fb4f0e612c3e6f0ef9b6415efa23072eadda37de1c8a901fd196042f05a3e",
    )

    http_archive(
        name = "libtensorflow_linux_x86_64_gpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        patch_args = ["-p1"],
        patches = [],
        strip_prefix = "",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-gpu-linux-x86_64-2.10.1.tar.gz"],
        sha256 = "5e8efbd334526d922413018b8890aeeb78470abd3612ee248aa8eb34014aad05",
    )

    http_archive(
        name = "libtensorflow_macos_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        patch_args = ["-p1"],
        patches = [],
        strip_prefix = "",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-darwin-x86_64-2.10.1.tar.gz"],
        sha256 = "f4a3c7f9fe861edc9e628e7bcf9ac29307885ac19bc562c93889decc2e7e144c",
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
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow_proto-2.10.1.zip"],
        sha256 = "a4991bf55f5ef4d06b8004a0c446c4314911548da597c64f054085e59ef2677f",
    )

    http_archive(
        name = "libtensorflow_proto_linux_x86_64_gpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.BUILD",
        patch_args = ["-p1"],
        patches = ["@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.patch"],
        strip_prefix = "",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow_proto-2.10.1.zip"],
        sha256 = "a4991bf55f5ef4d06b8004a0c446c4314911548da597c64f054085e59ef2677f",
    )

    http_archive(
        name = "libtensorflow_proto_macos_x86_64_cpu",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.BUILD",
        patch_args = ["-p1"],
        patches = ["@com_github_wamuir_graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.patch"],
        strip_prefix = "",
        urls = ["https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow_proto-2.10.1.zip"],
        sha256 = "a4991bf55f5ef4d06b8004a0c446c4314911548da597c64f054085e59ef2677f",
    )

    ###########################################################################
    # base api def
    ###########################################################################

    http_archive(
        name = "tensorflow_base_api_def",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/tensorflow_base_api_def:tensorflow_base_api_def.BUILD",
        patch_args = ["-p1"],
        patches = [],
        strip_prefix = "tensorflow-2.10.1/tensorflow/core/api_def/base_api",
        urls = ["https://github.com/tensorflow/tensorflow/archive/refs/tags/v2.10.1.tar.gz"],
        sha256 = "622a92e22e6f3f4300ea43b3025a0b6122f1cc0e2d9233235e4c628c331a94a3",
    )

    ###########################################################################
    # go bindings
    ###########################################################################

    http_archive(
        name = "tensorflow_go",
        build_file = "@com_github_wamuir_graft//third_party/org_tensorflow/tensorflow_go:tensorflow_go.BUILD",
        patch_args = ["-p1"],
        patches = [],
        strip_prefix = "tensorflow-2.10.1/tensorflow/go",
        urls = ["https://github.com/tensorflow/tensorflow/archive/refs/tags/v2.10.1.tar.gz"],
        sha256 = "622a92e22e6f3f4300ea43b3025a0b6122f1cc0e2d9233235e4c628c331a94a3",
    )
