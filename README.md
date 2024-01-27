# Graft
[![tensorflow version](https://img.shields.io/badge/tf-v2.15.0-FF6F00?logo=tensorflow&logoColor=FF6F00)](https://github.com/tensorflow/tensorflow/tree/v2.15.0)
[![build](https://img.shields.io/github/actions/workflow/status/wamuir/graft/build-and-test-bindings.yml?branch=bzl-test-other-platforms&label=build&logo=github)](https://github.com/wamuir/graft/actions/workflows/test.yml?query=branch%3Abzl-test-other-platforms)
[![go.dev reference](https://pkg.go.dev/badge/wamuir/graft)](https://pkg.go.dev/github.com/wamuir/graft/tensorflow)

## About

**Go language bindings to the TensorFlow C API**

Graft contains nightly and release builds of the Go language bindings to the
TensorFlow C API, including Go-compiled TensorFlow protocol buffers and
generated Go wrappers for TensorFlow operations.

Use Graft exactly as you would use the Go bindings found in the main TensorFlow
repo, and with the following import statement: `github.com/wamuir/graft/tensorflow`

## Getting Started

> Note: the Go bindings depend on
> [libtensorflow](https://www.tensorflow.org/install/lang_c), which should be
> downloaded (or compiled) and installed first.

Installation is performed using `go get`:

```sh
go get -u github.com/wamuir/graft/tensorflow/...
```

<details><summary><b>Hello TensorFlow</b></summary>

```go
package main

import (
	tf "github.com/wamuir/graft/tensorflow"
	"github.com/wamuir/graft/tensorflow/op"
	"fmt"
)

func main() {
	// Construct a graph with an operation that produces a string constant.
	s := op.NewScope()
	c := op.Const(s, "Hello from TensorFlow version " + tf.Version())
	graph, err := s.Finalize()
	if err != nil {
		panic(err)
	}

	// Execute the graph in a session.
	sess, err := tf.NewSession(graph, nil)
	if err != nil {
		panic(err)
	}
	output, err := sess.Run(nil, []tf.Output{c}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(output[0].Value())
}
```

</details>

## Credits

Graft is a compilation of [TensorFlow](https://tensorflow.org/code) source
code.
