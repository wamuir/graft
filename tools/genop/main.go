/*
Copyright 2016 The TensorFlow Authors. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Command genop generates a Go source file with functions for TensorFlow ops.
package main

import (
	"bytes"
	"flag"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/wamuir/graft/tools/genop/internal"
)

func main() {
	var (
		apidefs                      []string
		apiDefDirs, filename, header string
		buf                          bytes.Buffer
	)

	flag.StringVar(&filename, "outfile", "", "File to write generated source code to. If empty, source is written to stdout.")
	flag.StringVar(&header, "header", "", "Path to a file whose contents will be copied into the generated file. Can be empty.")
	flag.StringVar(&apiDefDirs, "apiDefDirs", "", "Comma-separated directories containing api_def_*.pbtxt files.")
	flag.Parse()

	apidefs = flag.Args() // non-flag command-line arguments, representing a list of api_def_*pbtxt files

	// Write header to the buffer.
	if len(header) > 0 {
		hdr, err := ioutil.ReadFile(header)
		if err != nil {
			log.Fatalf("Unable to read %s: %v", header, err)
		}
		buf.Write(hdr)
		buf.WriteString("\n\n")
	}

	// Add pbtxt files from any given api def directories to the list of api defs.
	if len(apiDefDirs) > 0 {
		apiDefDirsList := strings.Split(apiDefDirs, ",")
		for _, dir := range apiDefDirsList {
			files, err := ioutil.ReadDir(dir)
			if err != nil {
				log.Fatalf("Failed to read directory %s: %v", dir, err)
			}
			for _, file := range files {
				if file.IsDir() || !strings.HasSuffix(file.Name(), ".pbtxt") {
					continue
				}
				apidefs = append(apidefs, path.Join(dir, file.Name()))
			}
		}
	}

	// Generate wrapper functions.
	if err := internal.GenerateFunctionsForRegisteredOps(&buf, apidefs); err != nil {
		log.Fatal(err)
	}

	// Format generated source code.
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("Failed to generate valid source? 'go fmt' failed: %v", err)
	}

	// Write the source code.
	switch filename {
	case "-", "":
		if _, err := os.Stdout.Write(formatted); err != nil {
			log.Fatalf("Failed to write to stdout: %v", err)
		}
	default:
		if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
			log.Fatalf("Failed to write to %q: %v", filename, err)
		}
		if err := ioutil.WriteFile(filename, formatted, 0644); err != nil {
			log.Fatalf("Failed to write to %q: %v", filename, err)
		}
	}
}
