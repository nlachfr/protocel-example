GOPATH := $(PWD)/.cache
PATH := $(PATH):$(GOPATH)/bin
SHELL := env PATH=$(PATH) /bin/bash

PROTOC_GEN_GO := $(GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC := $(GOPATH)/bin/protoc-gen-go-grpc
PROTOC_GEN_GO_CEL_AUTHORIZE := $(GOPATH)/bin/protoc-gen-go-cel-authorize
PROTOC_GEN_GO_CEL_VALIDATE := $(GOPATH)/bin/protoc-gen-go-cel-validate

PROTO := $(shell find proto -name '*.proto')
PROTO_OPTS := -I. -Isubmodules/protocel -Isubmodules/googleapis -I/include
GENPROTO_GO := $(PROTO:.proto=.pb.go) $(PROTO:.proto=_grpc.pb.go) $(PROTO:.proto=.pb.cel.authorize.go) $(PROTO:.proto=.pb.cel.validate.go)
GENPROTO_GO_IMPS := $(foreach i,$(PROTO),$(shell echo "M$(i)=$$(dirname $(i));$$(basename $$(dirname $(i)))"))

.PHONY: all
all: go-genproto

$(PROTOC_GEN_GO):
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

$(PROTOC_GEN_GO_GRPC):
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

$(PROTOC_GEN_GO_CEL_AUTHORIZE):
	go install github.com/Neakxs/protocel/cmd/protoc-gen-go-cel-authorize@v0.1.0

$(PROTOC_GEN_GO_CEL_VALIDATE):
	go install github.com/Neakxs/protocel/cmd/protoc-gen-go-cel-validate@v0.1.0


%.pb.go: %.proto
	protoc $(PROTO_OPTS) --go_out=. --go_opt=paths=source_relative $(addsuffix ",$(addprefix --go_opt=",$(GENPROTO_GO_IMPS))) $<

%_grpc.pb.go: %.proto
	protoc $(PROTO_OPTS) --go-grpc_out=. --go-grpc_opt=paths=source_relative $(addsuffix ",$(addprefix --go-grpc_opt=",$(GENPROTO_GO_IMPS))) $<

%.pb.cel.authorize.go: %.proto
	protoc $(PROTO_OPTS) --go-cel-authorize_out=. --go-cel-authorize_opt=paths=source_relative $(addsuffix ",$(addprefix --go-cel-authorize_opt=",$(GENPROTO_GO_IMPS))) $<

%.pb.cel.validate.go: %.proto
	protoc $(PROTO_OPTS) --go-cel-validate_out=. --go-cel-validate_opt=paths=source_relative $(addsuffix ",$(addprefix --go-cel-validate_opt=",$(GENPROTO_GO_IMPS))) $<


.PHONY: go-genproto
go-genproto: $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC) $(PROTOC_GEN_GO_CEL_AUTHORIZE) $(PROTOC_GEN_GO_CEL_VALIDATE) $(GENPROTO_GO)