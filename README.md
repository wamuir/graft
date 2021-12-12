# Graft
[![tensorflow version](https://img.shields.io/badge/tf-v2.8.0--nightly-FF6F00?logo=tensorflow&logoColor=FF6F00)](https://github.com/tensorflow/tensorflow/tree/nightly)
[![nightly](https://img.shields.io/github/workflow/status/wamuir/graft/nightly?label=nightly%20ci&logo=github&event=schedule)](https://github.com/wamuir/graft/actions/workflows/nightly.yml)
[![go.dev reference](https://pkg.go.dev/badge/wamuir/graft)](https://pkg.go.dev/github.com/wamuir/graft/tensorflow)

## About

**Go language bindings to the TensorFlow C API**

Graft contains nightly and release builds of the Go language bindings to the
TensorFlow C API, including builds of generated Go code for TensorFlow protobuf
definitions and API wrappers.

Use Graft exactly as you would use the Go bindings found in the main TensorFlow
repo, except use the import statement `github.com/wamuir/graft/tensorflow`.

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
