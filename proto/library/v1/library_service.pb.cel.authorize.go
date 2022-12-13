// Code generated by protoc-gen-go-cel-authorize. DO NOT EDIT.
// versions:
//  protoc-gen-go-cel-authorize	v0.1.0
//  protoc 				        v3.21.9
// source: proto/library/v1/library_service.proto

package v1

import (
	authorize "github.com/Neakxs/protocel/authorize"
	options "github.com/Neakxs/protocel/options"
	cel "github.com/google/cel-go/cel"
	proto "google.golang.org/protobuf/proto"
)

var _File_proto_library_v1_library_service_proto_rawAuthorizeOptions = []byte{}

func NewLibraryServiceAuthzInterceptor(opts ...options.RuntimeOptions) (authorize.AuthzInterceptor, error) {
	cfg := &authorize.AuthorizeOptions{}
	if err := proto.Unmarshal(_File_proto_library_v1_library_service_proto_rawAuthorizeOptions, cfg); err != nil {
		return nil, err
	}
	lib := options.BuildRuntimeLibrary(cfg.Options, opts...)
	m := map[string]cel.Program{}
	for k, v := range map[string]struct {
		expr string
		req  proto.Message
	}{} {
		if pgr, err := authorize.BuildAuthzProgram(v.expr, v.req, cfg, lib); err != nil {
			return nil, err
		} else {
			m[k] = pgr
		}
	}
	return authorize.NewAuthzInterceptor(m), nil
}
