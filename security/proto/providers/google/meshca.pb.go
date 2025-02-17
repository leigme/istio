// Copyright 2019 Istio Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: security/proto/providers/google/meshca.proto

package google

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Certificate request message.
type MeshCertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The request ID must be a valid UUID with the exception that zero UUID is
	// not supported (00000000-0000-0000-0000-000000000000).
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// PEM-encoded certificate request.
	Csr string `protobuf:"bytes,2,opt,name=csr,proto3" json:"csr,omitempty"`
	// Optional: requested certificate validity period.
	Validity *duration.Duration `protobuf:"bytes,3,opt,name=validity,proto3" json:"validity,omitempty"` // Reserved 4
}

func (x *MeshCertificateRequest) Reset() {
	*x = MeshCertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_proto_providers_google_meshca_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshCertificateRequest) ProtoMessage() {}

func (x *MeshCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_security_proto_providers_google_meshca_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshCertificateRequest.ProtoReflect.Descriptor instead.
func (*MeshCertificateRequest) Descriptor() ([]byte, []int) {
	return file_security_proto_providers_google_meshca_proto_rawDescGZIP(), []int{0}
}

func (x *MeshCertificateRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *MeshCertificateRequest) GetCsr() string {
	if x != nil {
		return x.Csr
	}
	return ""
}

func (x *MeshCertificateRequest) GetValidity() *duration.Duration {
	if x != nil {
		return x.Validity
	}
	return nil
}

// Certificate response message.
type MeshCertificateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PEM-encoded certificate chain.
	// Leaf cert is element '0'. Root cert is element 'n'.
	CertChain []string `protobuf:"bytes,1,rep,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
}

func (x *MeshCertificateResponse) Reset() {
	*x = MeshCertificateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_proto_providers_google_meshca_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshCertificateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshCertificateResponse) ProtoMessage() {}

func (x *MeshCertificateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_security_proto_providers_google_meshca_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshCertificateResponse.ProtoReflect.Descriptor instead.
func (*MeshCertificateResponse) Descriptor() ([]byte, []int) {
	return file_security_proto_providers_google_meshca_proto_rawDescGZIP(), []int{1}
}

func (x *MeshCertificateResponse) GetCertChain() []string {
	if x != nil {
		return x.CertChain
	}
	return nil
}

var File_security_proto_providers_google_meshca_proto protoreflect.FileDescriptor

var file_security_proto_providers_google_meshca_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x63, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x16, 0x4d, 0x65,
	0x73, 0x68, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x73, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x63, 0x73, 0x72, 0x12, 0x35, 0x0a, 0x08, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x22, 0x38, 0x0a, 0x17,
	0x4d, 0x65, 0x73, 0x68, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x65, 0x72, 0x74, 0x5f,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x63, 0x65, 0x72,
	0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x32, 0x96, 0x01, 0x0a, 0x16, 0x4d, 0x65, 0x73, 0x68, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x7c, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x63, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x65, 0x73, 0x68,
	0x63, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x5e, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x63, 0x61, 0x2e, 0x76, 0x31,
	0x42, 0x0b, 0x4d, 0x65, 0x73, 0x68, 0x43, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2f,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_security_proto_providers_google_meshca_proto_rawDescOnce sync.Once
	file_security_proto_providers_google_meshca_proto_rawDescData = file_security_proto_providers_google_meshca_proto_rawDesc
)

func file_security_proto_providers_google_meshca_proto_rawDescGZIP() []byte {
	file_security_proto_providers_google_meshca_proto_rawDescOnce.Do(func() {
		file_security_proto_providers_google_meshca_proto_rawDescData = protoimpl.X.CompressGZIP(file_security_proto_providers_google_meshca_proto_rawDescData)
	})
	return file_security_proto_providers_google_meshca_proto_rawDescData
}

var file_security_proto_providers_google_meshca_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_security_proto_providers_google_meshca_proto_goTypes = []any{
	(*MeshCertificateRequest)(nil),  // 0: google.security.meshca.v1.MeshCertificateRequest
	(*MeshCertificateResponse)(nil), // 1: google.security.meshca.v1.MeshCertificateResponse
	(*duration.Duration)(nil),       // 2: google.protobuf.Duration
}
var file_security_proto_providers_google_meshca_proto_depIdxs = []int32{
	2, // 0: google.security.meshca.v1.MeshCertificateRequest.validity:type_name -> google.protobuf.Duration
	0, // 1: google.security.meshca.v1.MeshCertificateService.CreateCertificate:input_type -> google.security.meshca.v1.MeshCertificateRequest
	1, // 2: google.security.meshca.v1.MeshCertificateService.CreateCertificate:output_type -> google.security.meshca.v1.MeshCertificateResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_security_proto_providers_google_meshca_proto_init() }
func file_security_proto_providers_google_meshca_proto_init() {
	if File_security_proto_providers_google_meshca_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_security_proto_providers_google_meshca_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MeshCertificateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_security_proto_providers_google_meshca_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MeshCertificateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_security_proto_providers_google_meshca_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_security_proto_providers_google_meshca_proto_goTypes,
		DependencyIndexes: file_security_proto_providers_google_meshca_proto_depIdxs,
		MessageInfos:      file_security_proto_providers_google_meshca_proto_msgTypes,
	}.Build()
	File_security_proto_providers_google_meshca_proto = out.File
	file_security_proto_providers_google_meshca_proto_rawDesc = nil
	file_security_proto_providers_google_meshca_proto_goTypes = nil
	file_security_proto_providers_google_meshca_proto_depIdxs = nil
}
