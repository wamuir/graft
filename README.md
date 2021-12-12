# Graft
[![tensorflow version](https://img.shields.io/badge/tf-v2.8.0--nightly-FF6F00?logo=tensorflow&logoColor=FF6F00)](https://github.com/tensorflow/tensorflow/tree/nightly)
[![nightly](https://img.shields.io/github/workflow/status/wamuir/graft/nightly?label=nightly%20ci&logo=github&event=push)](https://github.com/wamuir/graft/actions/workflows/nightly.yml)
[![go.dev reference](https://pkg.go.dev/badge/wamuir/graft)](https://pkg.go.dev/github.com/wamuir/graft/tensorflow)

## About

**Graft contains nightly and release builds of the Go language bindings to the TensorFlow C API.**

## Getting Started

> The Go bindings depend on [libtensorflow](https://www.tensorflow.org/install/lang_c), which should be downloaded (or compiled) and installed first.

Installation is performed using `go get`:

```sh
go get -u github.com/wamuir/graft/tensorflow/...
```

## Credits

Graft is a compilation of [Tensorflow](https://tensorflow.org/code) source code.
