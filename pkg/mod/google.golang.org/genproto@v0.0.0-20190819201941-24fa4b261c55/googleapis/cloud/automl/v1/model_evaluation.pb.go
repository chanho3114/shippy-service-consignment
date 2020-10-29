// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1/model_evaluation.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Evaluation results of a model.
type ModelEvaluation struct {
	// Output only. Problem type specific evaluation metrics.
	//
	// Types that are valid to be assigned to Metrics:
	//	*ModelEvaluation_TranslationEvaluationMetrics
	Metrics isModelEvaluation_Metrics `protobuf_oneof:"metrics"`
	// Output only. Resource name of the model evaluation.
	// Format:
	//
	// `projects/{project_id}/locations/{location_id}/models/{model_id}/modelEvaluations/{model_evaluation_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The ID of the annotation spec that the model evaluation applies to. The
	// The ID is empty for the overall model evaluation.
	AnnotationSpecId string `protobuf:"bytes,2,opt,name=annotation_spec_id,json=annotationSpecId,proto3" json:"annotation_spec_id,omitempty"`
	// Output only. Timestamp when this model evaluation was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The number of examples used for model evaluation, i.e. for
	// which ground truth from time of model creation is compared against the
	// predicted annotations created by the model.
	// For overall ModelEvaluation (i.e. with annotation_spec_id not set) this is
	// the total number of all examples used for evaluation.
	// Otherwise, this is the count of examples that according to the ground
	// truth were annotated by the
	//
	// [annotation_spec_id][google.cloud.automl.v1beta1.ModelEvaluation.annotation_spec_id].
	EvaluatedExampleCount int32    `protobuf:"varint,6,opt,name=evaluated_example_count,json=evaluatedExampleCount,proto3" json:"evaluated_example_count,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *ModelEvaluation) Reset()         { *m = ModelEvaluation{} }
func (m *ModelEvaluation) String() string { return proto.CompactTextString(m) }
func (*ModelEvaluation) ProtoMessage()    {}
func (*ModelEvaluation) Descriptor() ([]byte, []int) {
	return fileDescriptor_008481175b84a2ca, []int{0}
}

func (m *ModelEvaluation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelEvaluation.Unmarshal(m, b)
}
func (m *ModelEvaluation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelEvaluation.Marshal(b, m, deterministic)
}
func (m *ModelEvaluation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelEvaluation.Merge(m, src)
}
func (m *ModelEvaluation) XXX_Size() int {
	return xxx_messageInfo_ModelEvaluation.Size(m)
}
func (m *ModelEvaluation) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelEvaluation.DiscardUnknown(m)
}

var xxx_messageInfo_ModelEvaluation proto.InternalMessageInfo

type isModelEvaluation_Metrics interface {
	isModelEvaluation_Metrics()
}

type ModelEvaluation_TranslationEvaluationMetrics struct {
	TranslationEvaluationMetrics *TranslationEvaluationMetrics `protobuf:"bytes,9,opt,name=translation_evaluation_metrics,json=translationEvaluationMetrics,proto3,oneof"`
}

func (*ModelEvaluation_TranslationEvaluationMetrics) isModelEvaluation_Metrics() {}

func (m *ModelEvaluation) GetMetrics() isModelEvaluation_Metrics {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *ModelEvaluation) GetTranslationEvaluationMetrics() *TranslationEvaluationMetrics {
	if x, ok := m.GetMetrics().(*ModelEvaluation_TranslationEvaluationMetrics); ok {
		return x.TranslationEvaluationMetrics
	}
	return nil
}

func (m *ModelEvaluation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModelEvaluation) GetAnnotationSpecId() string {
	if m != nil {
		return m.AnnotationSpecId
	}
	return ""
}

func (m *ModelEvaluation) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ModelEvaluation) GetEvaluatedExampleCount() int32 {
	if m != nil {
		return m.EvaluatedExampleCount
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ModelEvaluation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ModelEvaluation_TranslationEvaluationMetrics)(nil),
	}
}

func init() {
	proto.RegisterType((*ModelEvaluation)(nil), "google.cloud.automl.v1.ModelEvaluation")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1/model_evaluation.proto", fileDescriptor_008481175b84a2ca)
}

var fileDescriptor_008481175b84a2ca = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x6a, 0xdb, 0x30,
	0x1c, 0xc7, 0x67, 0xb3, 0x64, 0x44, 0x39, 0x6c, 0x08, 0x96, 0x79, 0x26, 0x6c, 0x61, 0x27, 0x1f,
	0x36, 0x09, 0xaf, 0xa5, 0x07, 0xa7, 0x97, 0x26, 0x84, 0xb6, 0xd0, 0x40, 0x71, 0x43, 0x0e, 0x25,
	0x60, 0x14, 0x5b, 0x35, 0x06, 0xfd, 0x31, 0xb6, 0x1c, 0x7a, 0xe8, 0x4b, 0xf4, 0xb5, 0x7a, 0xeb,
	0x6b, 0xf4, 0x29, 0x8a, 0x25, 0x27, 0x0e, 0xa5, 0xe9, 0x4d, 0xf6, 0xe7, 0xf3, 0xfb, 0xf9, 0xeb,
	0xaf, 0xc0, 0xbf, 0x54, 0xca, 0x94, 0x51, 0x1c, 0x33, 0x59, 0x25, 0x98, 0x54, 0x4a, 0x72, 0x86,
	0x37, 0x3e, 0xe6, 0x32, 0xa1, 0x2c, 0xa2, 0x1b, 0xc2, 0x2a, 0xa2, 0x32, 0x29, 0x50, 0x5e, 0x48,
	0x25, 0xe1, 0xc0, 0xe8, 0x48, 0xeb, 0xc8, 0xe8, 0x68, 0xe3, 0xbb, 0xde, 0x81, 0x35, 0xaa, 0x20,
	0xa2, 0x64, 0x7b, 0x1b, 0xdc, 0xdf, 0x8d, 0xa9, 0x9f, 0xd6, 0xd5, 0x1d, 0x56, 0x19, 0xa7, 0xa5,
	0x22, 0x3c, 0x6f, 0x84, 0x61, 0x23, 0x90, 0x3c, 0xc3, 0x44, 0x08, 0xa9, 0xf4, 0x74, 0x69, 0xe8,
	0x9f, 0x67, 0x1b, 0x7c, 0x9d, 0xd7, 0xd9, 0x66, 0xbb, 0x68, 0xf0, 0x01, 0xfc, 0xda, 0xfb, 0xce,
	0x5e, 0xe8, 0x88, 0x53, 0x55, 0x64, 0x71, 0xe9, 0xf4, 0x46, 0x96, 0xd7, 0xff, 0x7f, 0x8c, 0xde,
	0x4f, 0x8f, 0x16, 0xed, 0x74, 0xbb, 0x76, 0x6e, 0x66, 0x2f, 0x3e, 0x85, 0x43, 0xf5, 0x01, 0x87,
	0x10, 0x7c, 0x16, 0x84, 0x53, 0xc7, 0x1a, 0x59, 0x5e, 0x2f, 0xd4, 0x67, 0xf8, 0x17, 0xc0, 0x36,
	0x7a, 0x54, 0xe6, 0x34, 0x8e, 0xb2, 0xc4, 0xb1, 0xb5, 0xf1, 0xad, 0x25, 0x37, 0x39, 0x8d, 0x2f,
	0x13, 0x38, 0x06, 0xfd, 0xb8, 0xa0, 0x44, 0xd1, 0xa8, 0xee, 0xc2, 0xe9, 0xe8, 0xb0, 0xee, 0x36,
	0xec, 0xb6, 0x28, 0xb4, 0xd8, 0x16, 0x15, 0x02, 0xa3, 0xd7, 0x2f, 0xe0, 0x09, 0xf8, 0xd1, 0xfc,
	0x30, 0x4d, 0x22, 0x7a, 0x4f, 0x78, 0xce, 0x68, 0x14, 0xcb, 0x4a, 0x28, 0xa7, 0x3b, 0xb2, 0xbc,
	0x4e, 0xf8, 0x7d, 0x87, 0x67, 0x86, 0x4e, 0x6b, 0x38, 0xe9, 0x81, 0x2f, 0x4d, 0x3b, 0x93, 0x47,
	0x0b, 0xb8, 0xb1, 0xe4, 0x07, 0xda, 0xb9, 0xb6, 0x6e, 0x4f, 0x1b, 0x92, 0x4a, 0x46, 0x44, 0x8a,
	0x64, 0x91, 0xe2, 0x94, 0x0a, 0x1d, 0x0c, 0x1b, 0x44, 0xf2, 0xac, 0x7c, 0x7b, 0xf9, 0x63, 0x73,
	0x7a, 0xb2, 0x07, 0xe7, 0xda, 0x59, 0x4d, 0x6b, 0xbe, 0x3a, 0xab, 0x94, 0x9c, 0xb3, 0xd5, 0xd2,
	0x7f, 0xb1, 0x7f, 0x1a, 0x10, 0x04, 0x9a, 0x04, 0x81, 0x46, 0x57, 0x41, 0xb0, 0xf4, 0xd7, 0x5d,
	0xbd, 0xfd, 0xe8, 0x35, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x27, 0xb3, 0x8d, 0xa0, 0x02, 0x00, 0x00,
}
