## Go library

The Go library package consists of the TensorFlow [Go
bindings](https://www.tensorflow.org/code/tensorflow/go/) (module
`github.com/tensorflow/tensorflow/tensorflow/go`) and the TensorFlow [C
API](https://www.tensorflow.org/code/tensorflow/c/c_api.h) that it requires.
Dependent Go modules should use the [Go module replace
directive](https://go.dev/ref/mod#go-mod-file-replace) to point to the local
installation of the TensorFlow Go bindings.  The following commands:

```sh
bazel test //tools/lib_package:libtensorflow_go_test
bazel build //tools/lib_package:libtensorflow_go
```

test and produce the archive at
`bazel-bin/tools/lib_package/libtensorflow_go.tar.gz`, which can be distributed
and installed using something like:

```sh
tar -C /usr/local/ -xzf libtensorflow_go.tar.gz
```
