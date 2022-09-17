GOPATH := $(shell go env GOPATH)

.DEFAULT_GOAL := bench

.PHONY: clean
clean:
	git clean -xdf

.PHONY: generate
generate: schema/proto.pb.go schema/proto_vtproto.pb.go schema/karmem_generated.go

.PHONY: test
test:
	go test -race ./...

.PHONY: bench
bench: generate
	mkdir -p result
	go test -tags proto   -bench=. -benchmem -count=5 github.com/leki75/iceprotobench/schema > result/proto.out
	go test -tags protovt -bench=. -benchmem -count=5 github.com/leki75/iceprotobench/schema > result/protovt.out
	go test -tags karmem  -bench=. -benchmem -count=5 github.com/leki75/iceprotobench/schema > result/karmem.out
	go test -tags raw     -bench=. -benchmem -count=5 github.com/leki75/iceprotobench/schema > result/raw.out

	go install golang.org/x/perf/cmd/benchstat
	benchstat result/proto.out result/protovt.out result/karmem.out result/raw.out

schema/proto.pb.go: schema/proto.proto
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	protoc --go_out=. --go_opt=paths=source_relative ./schema/proto.proto

schema/proto_vtproto.pb.go: schema/proto.proto
	go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto
	protoc \
		--plugin protoc-gen-go-vtproto="$(GOPATH)/bin/protoc-gen-go-vtproto" \
		--go-vtproto_out=. \
		--go-vtproto_opt=paths=source_relative \
		--go-vtproto_opt=features=marshal+unmarshal+size \
		./schema/proto.proto

schema/karmem_generated.go: schema/karmem.km
	go run karmem.org/cmd/karmem build --golang -o "schema" schema/karmem.km
