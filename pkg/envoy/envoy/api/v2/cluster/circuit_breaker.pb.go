// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/cluster/circuit_breaker.proto

/*
Package cluster is a generated protocol buffer package.

It is generated from these files:
	envoy/api/v2/cluster/circuit_breaker.proto
	envoy/api/v2/cluster/outlier_detection.proto

It has these top-level messages:
	CircuitBreakers
	OutlierDetection
*/
package cluster

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import google_protobuf1 "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// :ref:`Circuit breaking<arch_overview_circuit_break>` settings can be
// specified individually for each defined priority.
type CircuitBreakers struct {
	// If multiple :ref:`Thresholds<envoy_api_msg_cluster.CircuitBreakers.Thresholds>`
	// are defined with the same :ref:`RoutingPriority<envoy_api_enum_core.RoutingPriority>`,
	// the first one in the list is used. If no Thresholds is defined for a given
	// :ref:`RoutingPriority<envoy_api_enum_core.RoutingPriority>`, the default values
	// are used.
	Thresholds []*CircuitBreakers_Thresholds `protobuf:"bytes,1,rep,name=thresholds" json:"thresholds,omitempty"`
}

func (m *CircuitBreakers) Reset()                    { *m = CircuitBreakers{} }
func (m *CircuitBreakers) String() string            { return proto.CompactTextString(m) }
func (*CircuitBreakers) ProtoMessage()               {}
func (*CircuitBreakers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CircuitBreakers) GetThresholds() []*CircuitBreakers_Thresholds {
	if m != nil {
		return m.Thresholds
	}
	return nil
}

// A Thresholds defines CircuitBreaker settings for a
// :ref:`RoutingPriority<envoy_api_enum_core.RoutingPriority>`.
type CircuitBreakers_Thresholds struct {
	// The :ref:`RoutingPriority<envoy_api_enum_core.RoutingPriority>`
	// the specified CircuitBreaker settings apply to.
	// [#comment:TODO(htuch): add (validate.rules).enum.defined_only = true once
	// https://github.com/lyft/protoc-gen-validate/issues/42 is resolved.]
	Priority envoy_api_v2_core.RoutingPriority `protobuf:"varint,1,opt,name=priority,enum=envoy.api.v2.core.RoutingPriority" json:"priority,omitempty"`
	// The maximum number of connections that Envoy will make to the upstream
	// cluster. If not specified, the default is 1024.
	MaxConnections *google_protobuf1.UInt32Value `protobuf:"bytes,2,opt,name=max_connections,json=maxConnections" json:"max_connections,omitempty"`
	// The maximum number of pending requests that Envoy will allow to the
	// upstream cluster. If not specified, the default is 1024.
	MaxPendingRequests *google_protobuf1.UInt32Value `protobuf:"bytes,3,opt,name=max_pending_requests,json=maxPendingRequests" json:"max_pending_requests,omitempty"`
	// The maximum number of parallel requests that Envoy will make to the
	// upstream cluster. If not specified, the default is 1024.
	MaxRequests *google_protobuf1.UInt32Value `protobuf:"bytes,4,opt,name=max_requests,json=maxRequests" json:"max_requests,omitempty"`
	// The maximum number of parallel retries that Envoy will allow to the
	// upstream cluster. If not specified, the default is 3.
	MaxRetries *google_protobuf1.UInt32Value `protobuf:"bytes,5,opt,name=max_retries,json=maxRetries" json:"max_retries,omitempty"`
}

func (m *CircuitBreakers_Thresholds) Reset()                    { *m = CircuitBreakers_Thresholds{} }
func (m *CircuitBreakers_Thresholds) String() string            { return proto.CompactTextString(m) }
func (*CircuitBreakers_Thresholds) ProtoMessage()               {}
func (*CircuitBreakers_Thresholds) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *CircuitBreakers_Thresholds) GetPriority() envoy_api_v2_core.RoutingPriority {
	if m != nil {
		return m.Priority
	}
	return envoy_api_v2_core.RoutingPriority_DEFAULT
}

func (m *CircuitBreakers_Thresholds) GetMaxConnections() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.MaxConnections
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxPendingRequests() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.MaxPendingRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRequests() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.MaxRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRetries() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.MaxRetries
	}
	return nil
}

func init() {
	proto.RegisterType((*CircuitBreakers)(nil), "envoy.api.v2.cluster.CircuitBreakers")
	proto.RegisterType((*CircuitBreakers_Thresholds)(nil), "envoy.api.v2.cluster.CircuitBreakers.Thresholds")
}

func init() { proto.RegisterFile("envoy/api/v2/cluster/circuit_breaker.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x53, 0xf0, 0x5f, 0xb6, 0x0a, 0x49, 0xc3, 0xa1, 0x21, 0x84, 0x10, 0x4e, 0xc4, 0xc3,
	0xae, 0x29, 0x67, 0x35, 0x81, 0x78, 0xf0, 0x62, 0x48, 0xa3, 0x1e, 0xbc, 0x90, 0x6d, 0x19, 0xcb,
	0xc6, 0x76, 0xb7, 0xce, 0x6e, 0xb1, 0xbc, 0x91, 0xaf, 0xe2, 0x6b, 0xe8, 0x8b, 0x18, 0xda, 0x5a,
	0x94, 0x78, 0xe0, 0x36, 0xed, 0x7c, 0xbf, 0x5f, 0xbe, 0xcc, 0x92, 0x73, 0x90, 0x2b, 0xb5, 0x66,
	0x3c, 0x15, 0x6c, 0xe5, 0xb1, 0x30, 0xce, 0xb4, 0x01, 0x64, 0xa1, 0xc0, 0x30, 0x13, 0x66, 0x1e,
	0x20, 0xf0, 0x17, 0x40, 0x9a, 0xa2, 0x32, 0xca, 0xe9, 0x14, 0x59, 0xca, 0x53, 0x41, 0x57, 0x1e,
	0xad, 0xb2, 0xdd, 0xde, 0x5f, 0x83, 0x42, 0x60, 0x01, 0xd7, 0x50, 0x32, 0xdd, 0x7e, 0xa4, 0x54,
	0x14, 0x03, 0x2b, 0xbe, 0x82, 0xec, 0x99, 0xbd, 0x21, 0x4f, 0x53, 0x40, 0x5d, 0xed, 0x3b, 0x91,
	0x8a, 0x54, 0x31, 0xb2, 0xcd, 0x54, 0xfe, 0x1d, 0x7e, 0x34, 0x49, 0x7b, 0x5a, 0x76, 0x98, 0x94,
	0x15, 0xb4, 0x33, 0x23, 0xc4, 0x2c, 0x11, 0xf4, 0x52, 0xc5, 0x0b, 0xed, 0x5a, 0x83, 0xe6, 0xc8,
	0xf6, 0x2e, 0xe8, 0x7f, 0x95, 0xe8, 0x0e, 0x4a, 0xef, 0x6b, 0xce, 0xff, 0xe5, 0xe8, 0x7e, 0x35,
	0x08, 0xd9, 0xae, 0x9c, 0x2b, 0x72, 0x92, 0xa2, 0x50, 0x28, 0xcc, 0xda, 0xb5, 0x06, 0xd6, 0xa8,
	0xe5, 0x0d, 0x77, 0xf4, 0x0a, 0x81, 0xfa, 0x2a, 0x33, 0x42, 0x46, 0xb3, 0x2a, 0xe9, 0xd7, 0x8c,
	0x73, 0x43, 0xda, 0x09, 0xcf, 0xe7, 0xa1, 0x92, 0x12, 0x42, 0x23, 0x94, 0xd4, 0x6e, 0x63, 0x60,
	0x8d, 0x6c, 0xaf, 0x47, 0xcb, 0x23, 0xd0, 0x9f, 0x23, 0xd0, 0x87, 0x5b, 0x69, 0xc6, 0xde, 0x23,
	0x8f, 0x33, 0xf0, 0x5b, 0x09, 0xcf, 0xa7, 0x5b, 0xc6, 0xb9, 0x23, 0x9d, 0x8d, 0x26, 0x05, 0xb9,
	0x10, 0x32, 0x9a, 0x23, 0xbc, 0x66, 0xa0, 0x8d, 0x76, 0x9b, 0x7b, 0xb8, 0x9c, 0x84, 0xe7, 0xb3,
	0x12, 0xf4, 0x2b, 0xce, 0xb9, 0x26, 0xa7, 0x1b, 0x5f, 0xed, 0x39, 0xd8, 0xc3, 0x63, 0x27, 0x3c,
	0xaf, 0x05, 0x97, 0xc4, 0x2e, 0x05, 0x06, 0x05, 0x68, 0xf7, 0x70, 0x0f, 0x9e, 0x14, 0x7c, 0x91,
	0x9f, 0x9c, 0xbd, 0x7f, 0xf6, 0xad, 0xa7, 0xe3, 0xea, 0x6d, 0x82, 0xa3, 0x02, 0x18, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xcf, 0x8a, 0x0b, 0x1f, 0x79, 0x02, 0x00, 0x00,
}
