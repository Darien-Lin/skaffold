// Copyright 2022 The Sigstore Authors.
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
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.4
// source: sigstore_rekor.proto

package v1

import (
	v1 "github.com/sigstore/protobuf-specs/gen/pb-go/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// KindVersion contains the entry's kind and api version.
type KindVersion struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Kind is the type of entry being stored in the log.
	// See here for a list: https://github.com/sigstore/rekor/tree/main/pkg/types
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// The specific api version of the type.
	Version       string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KindVersion) Reset() {
	*x = KindVersion{}
	mi := &file_sigstore_rekor_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KindVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KindVersion) ProtoMessage() {}

func (x *KindVersion) ProtoReflect() protoreflect.Message {
	mi := &file_sigstore_rekor_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KindVersion.ProtoReflect.Descriptor instead.
func (*KindVersion) Descriptor() ([]byte, []int) {
	return file_sigstore_rekor_proto_rawDescGZIP(), []int{0}
}

func (x *KindVersion) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *KindVersion) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// The checkpoint MUST contain an origin string as a unique log identifier,
// the tree size, and the root hash. It MAY also be followed by optional data,
// and clients MUST NOT assume optional data. The checkpoint MUST also contain
// a signature over the root hash (tree head). The checkpoint MAY contain additional
// signatures, but the first SHOULD be the signature from the log. Checkpoint contents
// are concatenated with newlines into a single string.
// The checkpoint format is described in
// https://github.com/transparency-dev/formats/blob/main/log/README.md
// and https://github.com/C2SP/C2SP/blob/main/tlog-checkpoint.md.
// An example implementation can be found in https://github.com/sigstore/rekor/blob/main/pkg/util/signed_note.go
type Checkpoint struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Envelope      string                 `protobuf:"bytes,1,opt,name=envelope,proto3" json:"envelope,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Checkpoint) Reset() {
	*x = Checkpoint{}
	mi := &file_sigstore_rekor_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Checkpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Checkpoint) ProtoMessage() {}

func (x *Checkpoint) ProtoReflect() protoreflect.Message {
	mi := &file_sigstore_rekor_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Checkpoint.ProtoReflect.Descriptor instead.
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return file_sigstore_rekor_proto_rawDescGZIP(), []int{1}
}

func (x *Checkpoint) GetEnvelope() string {
	if x != nil {
		return x.Envelope
	}
	return ""
}

// InclusionProof is the proof returned from the transparency log. Can
// be used for offline or online verification against the log.
type InclusionProof struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The index of the entry in the tree it was written to.
	LogIndex int64 `protobuf:"varint,1,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty"`
	// The hash digest stored at the root of the merkle tree at the time
	// the proof was generated.
	RootHash []byte `protobuf:"bytes,2,opt,name=root_hash,json=rootHash,proto3" json:"root_hash,omitempty"`
	// The size of the merkle tree at the time the proof was generated.
	TreeSize int64 `protobuf:"varint,3,opt,name=tree_size,json=treeSize,proto3" json:"tree_size,omitempty"`
	// A list of hashes required to compute the inclusion proof, sorted
	// in order from leaf to root.
	// Note that leaf and root hashes are not included.
	// The root hash is available separately in this message, and the
	// leaf hash should be calculated by the client.
	Hashes [][]byte `protobuf:"bytes,4,rep,name=hashes,proto3" json:"hashes,omitempty"`
	// Signature of the tree head, as of the time of this proof was
	// generated. See above info on 'Checkpoint' for more details.
	Checkpoint    *Checkpoint `protobuf:"bytes,5,opt,name=checkpoint,proto3" json:"checkpoint,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InclusionProof) Reset() {
	*x = InclusionProof{}
	mi := &file_sigstore_rekor_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InclusionProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InclusionProof) ProtoMessage() {}

func (x *InclusionProof) ProtoReflect() protoreflect.Message {
	mi := &file_sigstore_rekor_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InclusionProof.ProtoReflect.Descriptor instead.
func (*InclusionProof) Descriptor() ([]byte, []int) {
	return file_sigstore_rekor_proto_rawDescGZIP(), []int{2}
}

func (x *InclusionProof) GetLogIndex() int64 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

func (x *InclusionProof) GetRootHash() []byte {
	if x != nil {
		return x.RootHash
	}
	return nil
}

func (x *InclusionProof) GetTreeSize() int64 {
	if x != nil {
		return x.TreeSize
	}
	return 0
}

func (x *InclusionProof) GetHashes() [][]byte {
	if x != nil {
		return x.Hashes
	}
	return nil
}

func (x *InclusionProof) GetCheckpoint() *Checkpoint {
	if x != nil {
		return x.Checkpoint
	}
	return nil
}

// The inclusion promise is calculated by Rekor. It's calculated as a
// signature over a canonical JSON serialization of the persisted entry, the
// log ID, log index and the integration timestamp.
// See https://github.com/sigstore/rekor/blob/a6e58f72b6b18cc06cefe61808efd562b9726330/pkg/api/entries.go#L54
// The format of the signature depends on the transparency log's public key.
// If the signature algorithm requires a hash function and/or a signature
// scheme (e.g. RSA) those has to be retrieved out-of-band from the log's
// operators, together with the public key.
// This is used to verify the integration timestamp's value and that the log
// has promised to include the entry.
type InclusionPromise struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	SignedEntryTimestamp []byte                 `protobuf:"bytes,1,opt,name=signed_entry_timestamp,json=signedEntryTimestamp,proto3" json:"signed_entry_timestamp,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *InclusionPromise) Reset() {
	*x = InclusionPromise{}
	mi := &file_sigstore_rekor_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InclusionPromise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InclusionPromise) ProtoMessage() {}

func (x *InclusionPromise) ProtoReflect() protoreflect.Message {
	mi := &file_sigstore_rekor_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InclusionPromise.ProtoReflect.Descriptor instead.
func (*InclusionPromise) Descriptor() ([]byte, []int) {
	return file_sigstore_rekor_proto_rawDescGZIP(), []int{3}
}

func (x *InclusionPromise) GetSignedEntryTimestamp() []byte {
	if x != nil {
		return x.SignedEntryTimestamp
	}
	return nil
}

// TransparencyLogEntry captures all the details required from Rekor to
// reconstruct an entry, given that the payload is provided via other means.
// This type can easily be created from the existing response from Rekor.
// Future iterations could rely on Rekor returning the minimal set of
// attributes (excluding the payload) that are required for verifying the
// inclusion promise. The inclusion promise (called SignedEntryTimestamp in
// the response from Rekor) is similar to a Signed Certificate Timestamp
// as described here https://www.rfc-editor.org/rfc/rfc6962.html#section-3.2.
type TransparencyLogEntry struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The global index of the entry, used when querying the log by index.
	LogIndex int64 `protobuf:"varint,1,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty"`
	// The unique identifier of the log.
	LogId *v1.LogId `protobuf:"bytes,2,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	// The kind (type) and version of the object associated with this
	// entry. These values are required to construct the entry during
	// verification.
	KindVersion *KindVersion `protobuf:"bytes,3,opt,name=kind_version,json=kindVersion,proto3" json:"kind_version,omitempty"`
	// The UNIX timestamp from the log when the entry was persisted.
	// The integration time MUST NOT be trusted if inclusion_promise
	// is omitted.
	IntegratedTime int64 `protobuf:"varint,4,opt,name=integrated_time,json=integratedTime,proto3" json:"integrated_time,omitempty"`
	// The inclusion promise/signed entry timestamp from the log.
	// Required for v0.1 bundles, and MUST be verified.
	// Optional for >= v0.2 bundles if another suitable source of
	// time is present (such as another source of signed time,
	// or the current system time for long-lived certificates).
	// MUST be verified if no other suitable source of time is present,
	// and SHOULD be verified otherwise.
	InclusionPromise *InclusionPromise `protobuf:"bytes,5,opt,name=inclusion_promise,json=inclusionPromise,proto3" json:"inclusion_promise,omitempty"`
	// The inclusion proof can be used for offline or online verification
	// that the entry was appended to the log, and that the log has not been
	// altered.
	InclusionProof *InclusionProof `protobuf:"bytes,6,opt,name=inclusion_proof,json=inclusionProof,proto3" json:"inclusion_proof,omitempty"`
	// Optional. The canonicalized transparency log entry, used to
	// reconstruct the Signed Entry Timestamp (SET) during verification.
	// The contents of this field are the same as the `body` field in
	// a Rekor response, meaning that it does **not** include the "full"
	// canonicalized form (of log index, ID, etc.) which are
	// exposed as separate fields. The verifier is responsible for
	// combining the `canonicalized_body`, `log_index`, `log_id`,
	// and `integrated_time` into the payload that the SET's signature
	// is generated over.
	// This field is intended to be used in cases where the SET cannot be
	// produced determinisitically (e.g. inconsistent JSON field ordering,
	// differing whitespace, etc).
	//
	// If set, clients MUST verify that the signature referenced in the
	// `canonicalized_body` matches the signature provided in the
	// `Bundle.content`.
	// If not set, clients are responsible for constructing an equivalent
	// payload from other sources to verify the signature.
	CanonicalizedBody []byte `protobuf:"bytes,7,opt,name=canonicalized_body,json=canonicalizedBody,proto3" json:"canonicalized_body,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *TransparencyLogEntry) Reset() {
	*x = TransparencyLogEntry{}
	mi := &file_sigstore_rekor_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransparencyLogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransparencyLogEntry) ProtoMessage() {}

func (x *TransparencyLogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_sigstore_rekor_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransparencyLogEntry.ProtoReflect.Descriptor instead.
func (*TransparencyLogEntry) Descriptor() ([]byte, []int) {
	return file_sigstore_rekor_proto_rawDescGZIP(), []int{4}
}

func (x *TransparencyLogEntry) GetLogIndex() int64 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

func (x *TransparencyLogEntry) GetLogId() *v1.LogId {
	if x != nil {
		return x.LogId
	}
	return nil
}

func (x *TransparencyLogEntry) GetKindVersion() *KindVersion {
	if x != nil {
		return x.KindVersion
	}
	return nil
}

func (x *TransparencyLogEntry) GetIntegratedTime() int64 {
	if x != nil {
		return x.IntegratedTime
	}
	return 0
}

func (x *TransparencyLogEntry) GetInclusionPromise() *InclusionPromise {
	if x != nil {
		return x.InclusionPromise
	}
	return nil
}

func (x *TransparencyLogEntry) GetInclusionProof() *InclusionProof {
	if x != nil {
		return x.InclusionProof
	}
	return nil
}

func (x *TransparencyLogEntry) GetCanonicalizedBody() []byte {
	if x != nil {
		return x.CanonicalizedBody
	}
	return nil
}

var File_sigstore_rekor_proto protoreflect.FileDescriptor

var file_sigstore_rekor_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x6b, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x6b, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x0b, 0x4b, 0x69, 0x6e, 0x64, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1d, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x2d, 0x0a, 0x0a,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x08, 0x65, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x08, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x22, 0xdb, 0x01, 0x0a, 0x0e,
	0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x20,
	0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x20, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x20, 0x0a, 0x09, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x74, 0x72, 0x65, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0c, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65,
	0x73, 0x12, 0x46, 0x0a, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x6b, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x4d, 0x0a, 0x10, 0x49, 0x6e, 0x63,
	0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x16, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x14, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xc7, 0x03, 0x0a, 0x14, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x20, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x39, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67,
	0x49, 0x64, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x4a,
	0x0a, 0x0c, 0x6b, 0x69, 0x6e, 0x64, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x6b, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x69, 0x6e,
	0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x6b,
	0x69, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x0f, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x54, 0x0a, 0x11, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x72, 0x65, 0x6b, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x6c,
	0x75, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x52, 0x10, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x12, 0x53,
	0x0a, 0x0f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69,
	0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x6b, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x0e, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x6f, 0x66, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x64, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x11, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x42, 0x6f,
	0x64, 0x79, 0x42, 0x78, 0x0a, 0x1b, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6b, 0x6f, 0x72, 0x2e, 0x76,
	0x31, 0x42, 0x0a, 0x52, 0x65, 0x6b, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x67, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x73, 0x70,
	0x65, 0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x72, 0x65,
	0x6b, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0xea, 0x02, 0x13, 0x53, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x3a, 0x3a, 0x52, 0x65, 0x6b, 0x6f, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_sigstore_rekor_proto_rawDescOnce sync.Once
	file_sigstore_rekor_proto_rawDescData []byte
)

func file_sigstore_rekor_proto_rawDescGZIP() []byte {
	file_sigstore_rekor_proto_rawDescOnce.Do(func() {
		file_sigstore_rekor_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_sigstore_rekor_proto_rawDesc), len(file_sigstore_rekor_proto_rawDesc)))
	})
	return file_sigstore_rekor_proto_rawDescData
}

var file_sigstore_rekor_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sigstore_rekor_proto_goTypes = []any{
	(*KindVersion)(nil),          // 0: dev.sigstore.rekor.v1.KindVersion
	(*Checkpoint)(nil),           // 1: dev.sigstore.rekor.v1.Checkpoint
	(*InclusionProof)(nil),       // 2: dev.sigstore.rekor.v1.InclusionProof
	(*InclusionPromise)(nil),     // 3: dev.sigstore.rekor.v1.InclusionPromise
	(*TransparencyLogEntry)(nil), // 4: dev.sigstore.rekor.v1.TransparencyLogEntry
	(*v1.LogId)(nil),             // 5: dev.sigstore.common.v1.LogId
}
var file_sigstore_rekor_proto_depIdxs = []int32{
	1, // 0: dev.sigstore.rekor.v1.InclusionProof.checkpoint:type_name -> dev.sigstore.rekor.v1.Checkpoint
	5, // 1: dev.sigstore.rekor.v1.TransparencyLogEntry.log_id:type_name -> dev.sigstore.common.v1.LogId
	0, // 2: dev.sigstore.rekor.v1.TransparencyLogEntry.kind_version:type_name -> dev.sigstore.rekor.v1.KindVersion
	3, // 3: dev.sigstore.rekor.v1.TransparencyLogEntry.inclusion_promise:type_name -> dev.sigstore.rekor.v1.InclusionPromise
	2, // 4: dev.sigstore.rekor.v1.TransparencyLogEntry.inclusion_proof:type_name -> dev.sigstore.rekor.v1.InclusionProof
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_sigstore_rekor_proto_init() }
func file_sigstore_rekor_proto_init() {
	if File_sigstore_rekor_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sigstore_rekor_proto_rawDesc), len(file_sigstore_rekor_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sigstore_rekor_proto_goTypes,
		DependencyIndexes: file_sigstore_rekor_proto_depIdxs,
		MessageInfos:      file_sigstore_rekor_proto_msgTypes,
	}.Build()
	File_sigstore_rekor_proto = out.File
	file_sigstore_rekor_proto_goTypes = nil
	file_sigstore_rekor_proto_depIdxs = nil
}
