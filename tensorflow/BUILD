load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(
    default_visibility = ["//visibility:public"],
)

licenses(["notice"])  # Apache 2.0

go_library(
    name = "tensorflow",
    srcs = [
        "android.go",
        "attrs.go",
        "context.go",
        "doc.go",
        "graph.go",
        "lib.go",
        "operation.go",
        "saved_model.go",
        "session.go",
        "shape.go",
        "signature.go",
        "status.go",
        "tensor.go",
        "tensor_handle.go",
        "version.go",
    ],
    cdeps = ["//third_party/org_tensorflow:libtensorflow"],  # keep
    cgo = True,
    clinkopts = select({
        "@io_bazel_rules_go//go/platform:android": [
            "-landroid -llog -lm -lz -ldl",
        ],
        "//conditions:default": [],
    }),  # keep
    copts = [],  # keep
    importpath = "github.com/wamuir/graft/tensorflow",
    deps = [
        "//tensorflow/core/protobuf/for_core_protos_go_proto",
        "@org_golang_google_protobuf//proto:go_default_library",
    ],
)

alias(
    name = "go_default_library",
    actual = ":tensorflow",
    visibility = ["//visibility:public"],
)

go_test(
    name = "tensorflow_test",
    size = "small",  # keep
    srcs = [
        "attrs_test.go",
        "context_test.go",
        "example_inception_inference_test.go",
        "graph_test.go",
        "operation_test.go",
        "saved_model_test.go",
        "session_test.go",
        "shape_test.go",
        "signature_test.go",
        "tensor_handle_test.go",
        "tensor_test.go",
        "util_test.go",
    ],
    data = glob(["testdata/**"]),
    embed = [":tensorflow"],
    deps = [
        "//tensorflow/core/framework/tensor_shape_go_proto",
        "//tensorflow/core/framework/types_go_proto",
        "//tensorflow/core/protobuf/for_core_protos_go_proto",
        "//tensorflow/op",
    ],
)
