//go:build tools
// +build tools

package tools

import (
	_ "github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto"
	_ "golang.org/x/perf/cmd/benchstat"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
	_ "karmem.org/cmd/karmem"
)
