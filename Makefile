GOPATH := $(shell go env GOPATH)

.PHONY: clean
clean:
	git clean -xdf

.PHONY: install
install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@v0.2.0

.PHONY: generate
generate: install
	@# protobuf
	protoc --go_out=. --go_opt=paths=source_relative ./schema/proto.proto

	@# VT protobuf
	protoc \
		--plugin protoc-gen-go-vtproto="$(GOPATH)/bin/protoc-gen-go-vtproto" \
		--go-vtproto_out=. \
		--go-vtproto_opt=paths=source_relative \
		--go-vtproto_opt=features=marshal+unmarshal+size \
		./schema/proto.proto

.PHONY:
bench:
	go test -bench '^Benchmark.*_Marshal.*$$'   github.com/leki75/iceprotobench/schema
	go test -bench '^Benchmark.*_Unmarshal.*$$' github.com/leki75/iceprotobench/schema