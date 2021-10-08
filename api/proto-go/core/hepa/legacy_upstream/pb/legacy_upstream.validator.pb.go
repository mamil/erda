// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: legacy_upstream.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *AsyncRegisterResponse) Validate() error {
	return nil
}
func (this *AsyncRegisterRequest) Validate() error {
	if this.Upstream != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Upstream); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Upstream", err)
		}
	}
	return nil
}
func (this *RegisterResponse) Validate() error {
	return nil
}
func (this *RegisterRequest) Validate() error {
	if this.Upstream != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Upstream); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Upstream", err)
		}
	}
	return nil
}
func (this *UpstreamApi) Validate() error {
	if this.Doc != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Doc); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Doc", err)
		}
	}
	return nil
}
func (this *Upstream) Validate() error {
	for _, item := range this.ApiList {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ApiList", err)
			}
		}
	}
	return nil
}