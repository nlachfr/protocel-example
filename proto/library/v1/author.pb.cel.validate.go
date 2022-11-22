// Code generated by protoc-gen-go-cel-validate. DO NOT EDIT.
// versions:
//  protoc-gen-go-cel-validate	v0.0.0
//  protoc						v3.21.7
// source: proto/library/v1/author.proto

package v1

import (
	context "context"
	validate "github.com/Neakxs/protocel/validate"
	cel "github.com/google/cel-go/cel"
	proto "google.golang.org/protobuf/proto"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	sync "sync"
)

var _File_proto_library_v1_author_proto_rawValidateOptions = []byte{}

var (
	_File_proto_library_v1_author_proto_Author_celValidateMap     map[string]cel.Program = nil
	_File_proto_library_v1_author_proto_Author_celValidateMapOnce sync.Once
)

func (m *Author) Validate(ctx context.Context) error {
	return m.ValidateWithMask(ctx, &fieldmaskpb.FieldMask{
		Paths: []string{"*"},
	})
}

func (m *Author) ValidateWithMask(ctx context.Context, fm *fieldmaskpb.FieldMask) error {
	_File_proto_library_v1_author_proto_Author_celValidateMapOnce.Do(func() {
		cfg := &validate.ValidateOptions{}
		if err := proto.Unmarshal(_File_proto_library_v1_author_proto_rawValidateOptions, cfg); err != nil {
			return
		}
		tmp := map[string]cel.Program{}
		for k, v := range map[string]struct {
			expr string
			req  proto.Message
		}{

			"death_date": {expr: `death_date > birth_date`, req: m},
		} {
			if pgr, err := validate.BuildValidateProgram(v.expr, v.req, cfg); err != nil {
				return
			} else {
				tmp[k] = pgr
			}
		}
		_File_proto_library_v1_author_proto_Author_celValidateMap = tmp
	})
	return validate.ValidateWithMask(ctx, m, fm, _File_proto_library_v1_author_proto_Author_celValidateMap)
}
