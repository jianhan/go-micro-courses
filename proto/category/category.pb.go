// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/category/category.proto

/*
Package category is a generated protocol buffer package.

It is generated from these files:
	proto/category/category.proto

It has these top-level messages:
	Category
	CategorySlice
	InsertCategoriesResponse
	UpdateCategoriesResponse
	FindCategoriesRequest
	DeleteCategoriesByIDsRequest
*/
package category

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/empty"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Category struct {
	// @inject_tag: bson:"_id" valid:"uuid,required~ID is required"
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty" bson:"_id" valid:"uuid,required~ID is required"`
	// @inject_tag: bson:"name" valid:"required~name is required,length(1|256)~name must be max 256 characters"
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty" bson:"name" valid:"required~name is required,length(1|256)~name must be max 256 characters"`
	// @inject_tag: valid:"required~Slug is required"
	Slug string `protobuf:"bytes,3,opt,name=slug" json:"slug,omitempty" valid:"required~Slug is required"`
	// @inject_tag: bson:"display_order"
	DisplayOrder uint64 `protobuf:"varint,4,opt,name=display_order,json=displayOrder" json:"display_order,omitempty" bson:"display_order"`
	// @inject_tag: valid:"required~Description is required"
	Description string `protobuf:"bytes,5,opt,name=description" json:"description,omitempty" valid:"required~Description is required"`
	Hidden      bool   `protobuf:"varint,6,opt,name=hidden" json:"hidden,omitempty"`
	// @inject_tag: bson:"created_at"
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt" json:"created_at,omitempty" bson:"created_at"`
	// @inject_tag: bson:"updated_at"
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty" bson:"updated_at"`
}

func (m *Category) Reset()                    { *m = Category{} }
func (m *Category) String() string            { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()               {}
func (*Category) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Category) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Category) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Category) GetDisplayOrder() uint64 {
	if m != nil {
		return m.DisplayOrder
	}
	return 0
}

func (m *Category) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Category) GetHidden() bool {
	if m != nil {
		return m.Hidden
	}
	return false
}

func (m *Category) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Category) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type CategorySlice struct {
	Categories []*Category `protobuf:"bytes,1,rep,name=categories" json:"categories,omitempty"`
}

func (m *CategorySlice) Reset()                    { *m = CategorySlice{} }
func (m *CategorySlice) String() string            { return proto.CompactTextString(m) }
func (*CategorySlice) ProtoMessage()               {}
func (*CategorySlice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CategorySlice) GetCategories() []*Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

type InsertCategoriesResponse struct {
	Inserted int64 `protobuf:"varint,1,opt,name=inserted" json:"inserted,omitempty"`
}

func (m *InsertCategoriesResponse) Reset()                    { *m = InsertCategoriesResponse{} }
func (m *InsertCategoriesResponse) String() string            { return proto.CompactTextString(m) }
func (*InsertCategoriesResponse) ProtoMessage()               {}
func (*InsertCategoriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *InsertCategoriesResponse) GetInserted() int64 {
	if m != nil {
		return m.Inserted
	}
	return 0
}

type UpdateCategoriesResponse struct {
	Updated int64 `protobuf:"varint,1,opt,name=updated" json:"updated,omitempty"`
}

func (m *UpdateCategoriesResponse) Reset()                    { *m = UpdateCategoriesResponse{} }
func (m *UpdateCategoriesResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateCategoriesResponse) ProtoMessage()               {}
func (*UpdateCategoriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateCategoriesResponse) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

type FindCategoriesRequest struct {
	Ids         []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
	Query       string   `protobuf:"bytes,2,opt,name=query" json:"query,omitempty"`
	Hidden      bool     `protobuf:"varint,3,opt,name=hidden" json:"hidden,omitempty"`
	Sort        []string `protobuf:"bytes,4,rep,name=sort" json:"sort,omitempty"`
	PerPage     int64    `protobuf:"varint,5,opt,name=per_page,json=perPage" json:"per_page,omitempty"`
	CurrentPage int64    `protobuf:"varint,6,opt,name=current_page,json=currentPage" json:"current_page,omitempty"`
}

func (m *FindCategoriesRequest) Reset()                    { *m = FindCategoriesRequest{} }
func (m *FindCategoriesRequest) String() string            { return proto.CompactTextString(m) }
func (*FindCategoriesRequest) ProtoMessage()               {}
func (*FindCategoriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FindCategoriesRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *FindCategoriesRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *FindCategoriesRequest) GetHidden() bool {
	if m != nil {
		return m.Hidden
	}
	return false
}

func (m *FindCategoriesRequest) GetSort() []string {
	if m != nil {
		return m.Sort
	}
	return nil
}

func (m *FindCategoriesRequest) GetPerPage() int64 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

func (m *FindCategoriesRequest) GetCurrentPage() int64 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

type DeleteCategoriesByIDsRequest struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *DeleteCategoriesByIDsRequest) Reset()                    { *m = DeleteCategoriesByIDsRequest{} }
func (m *DeleteCategoriesByIDsRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteCategoriesByIDsRequest) ProtoMessage()               {}
func (*DeleteCategoriesByIDsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteCategoriesByIDsRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto.RegisterType((*Category)(nil), "go.micro.srv.courses.Category")
	proto.RegisterType((*CategorySlice)(nil), "go.micro.srv.courses.CategorySlice")
	proto.RegisterType((*InsertCategoriesResponse)(nil), "go.micro.srv.courses.InsertCategoriesResponse")
	proto.RegisterType((*UpdateCategoriesResponse)(nil), "go.micro.srv.courses.UpdateCategoriesResponse")
	proto.RegisterType((*FindCategoriesRequest)(nil), "go.micro.srv.courses.FindCategoriesRequest")
	proto.RegisterType((*DeleteCategoriesByIDsRequest)(nil), "go.micro.srv.courses.DeleteCategoriesByIDsRequest")
}

func init() { proto.RegisterFile("proto/category/category.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0xe3, 0x34, 0x71, 0x26, 0x6d, 0x15, 0xad, 0xda, 0x6a, 0x31, 0x5f, 0xc6, 0xbd, 0x44,
	0x42, 0x72, 0x50, 0x40, 0x48, 0x5c, 0x90, 0xda, 0x06, 0xa4, 0x9c, 0x8a, 0x0c, 0x5c, 0xb8, 0x44,
	0x8e, 0x3d, 0x98, 0x95, 0x12, 0xaf, 0xbb, 0xbb, 0x46, 0xca, 0x2f, 0xe1, 0x27, 0x20, 0xfe, 0x25,
	0xf2, 0x7a, 0x1d, 0xdc, 0xe0, 0x14, 0xa4, 0xde, 0x66, 0xde, 0xbc, 0xb7, 0x3b, 0x33, 0xcf, 0x5e,
	0x78, 0x9c, 0x0b, 0xae, 0xf8, 0x24, 0x8e, 0x14, 0xa6, 0x5c, 0x6c, 0xb6, 0x41, 0xa0, 0x71, 0x72,
	0x92, 0xf2, 0x60, 0xcd, 0x62, 0xc1, 0x03, 0x29, 0xbe, 0x07, 0x31, 0x2f, 0x84, 0x44, 0xe9, 0x3e,
	0x4d, 0x39, 0x4f, 0x57, 0x38, 0xd1, 0x9c, 0x65, 0xf1, 0x75, 0xa2, 0xd8, 0x1a, 0xa5, 0x8a, 0xd6,
	0x79, 0x25, 0x73, 0x1f, 0xee, 0x12, 0x70, 0x9d, 0x2b, 0x73, 0xa6, 0xff, 0xa3, 0x03, 0xce, 0x95,
	0xb9, 0x86, 0x1c, 0x43, 0x67, 0x3e, 0xa3, 0x96, 0x67, 0x8d, 0x07, 0x61, 0x67, 0x3e, 0x23, 0x04,
	0xba, 0x59, 0xb4, 0x46, 0xda, 0xd1, 0x88, 0x8e, 0x4b, 0x4c, 0xae, 0x8a, 0x94, 0xda, 0x15, 0x56,
	0xc6, 0xe4, 0x1c, 0x8e, 0x12, 0x26, 0xf3, 0x55, 0xb4, 0x59, 0x70, 0x91, 0xa0, 0xa0, 0x5d, 0xcf,
	0x1a, 0x77, 0xc3, 0x43, 0x03, 0x5e, 0x97, 0x18, 0xf1, 0x60, 0x98, 0xa0, 0x8c, 0x05, 0xcb, 0x15,
	0xe3, 0x19, 0x3d, 0xd0, 0xfa, 0x26, 0x44, 0xce, 0xa0, 0xf7, 0x8d, 0x25, 0x09, 0x66, 0xb4, 0xe7,
	0x59, 0x63, 0x27, 0x34, 0x19, 0x79, 0x03, 0x10, 0x0b, 0x8c, 0x14, 0x26, 0x8b, 0x48, 0xd1, 0xbe,
	0x67, 0x8d, 0x87, 0x53, 0x37, 0xa8, 0xa6, 0x0a, 0xea, 0xa9, 0x82, 0x4f, 0xf5, 0xd8, 0xe1, 0xc0,
	0xb0, 0x2f, 0x54, 0x29, 0x2d, 0xf2, 0xa4, 0x96, 0x3a, 0xff, 0x96, 0x1a, 0xf6, 0x85, 0xf2, 0xaf,
	0xe1, 0xa8, 0x5e, 0xcc, 0xc7, 0x15, 0x8b, 0x91, 0xbc, 0x05, 0x30, 0x86, 0x30, 0x94, 0xd4, 0xf2,
	0xec, 0xf1, 0x70, 0xfa, 0x24, 0x68, 0xf3, 0x24, 0xa8, 0x85, 0x61, 0x43, 0xe1, 0xbf, 0x06, 0x3a,
	0xcf, 0x24, 0x0a, 0x75, 0xb5, 0xc5, 0x42, 0x94, 0x39, 0xcf, 0x24, 0x12, 0x17, 0x1c, 0xa6, 0x6b,
	0x98, 0xe8, 0xfd, 0xdb, 0xe1, 0x36, 0xf7, 0x5f, 0x01, 0xfd, 0xac, 0xbb, 0x6a, 0xd1, 0x51, 0xe8,
	0x9b, 0x8e, 0x8d, 0xac, 0x4e, 0xfd, 0x5f, 0x16, 0x9c, 0xbe, 0x67, 0x59, 0xd2, 0x14, 0xdd, 0x14,
	0x28, 0x15, 0x19, 0x81, 0xcd, 0x92, 0x6a, 0x80, 0x41, 0x58, 0x86, 0xe4, 0x04, 0x0e, 0x6e, 0x0a,
	0x14, 0x1b, 0x63, 0x74, 0x95, 0x34, 0xec, 0xb0, 0x6f, 0xd9, 0x51, 0x7e, 0x01, 0x5c, 0x28, 0xda,
	0xd5, 0x07, 0xe8, 0x98, 0x3c, 0x00, 0x27, 0x47, 0xb1, 0xc8, 0xa3, 0x14, 0xb5, 0xb3, 0x76, 0xd8,
	0xcf, 0x51, 0x7c, 0x88, 0x52, 0x24, 0xcf, 0xe0, 0x30, 0x2e, 0x84, 0xc0, 0x4c, 0x55, 0xe5, 0x9e,
	0x2e, 0x0f, 0x0d, 0x56, 0x52, 0xfc, 0x17, 0xf0, 0x68, 0x86, 0x2b, 0x6c, 0x4e, 0x78, 0xb9, 0x99,
	0xcf, 0xf6, 0x77, 0x3c, 0xfd, 0x69, 0x03, 0xfc, 0x21, 0x93, 0x14, 0x46, 0xbb, 0xab, 0x25, 0xe7,
	0x77, 0x5b, 0xa3, 0x3d, 0x75, 0x83, 0x76, 0xd2, 0x5e, 0x9f, 0x52, 0x18, 0xed, 0x7a, 0x71, 0xaf,
	0x8b, 0xf6, 0x1a, 0xbb, 0x84, 0xe3, 0xdb, 0xee, 0x91, 0xe7, 0xed, 0x27, 0xb4, 0x7a, 0xec, 0xfe,
	0x4f, 0x4f, 0x24, 0x86, 0xd3, 0xd6, 0xb5, 0x93, 0x69, 0xbb, 0xfa, 0x2e, 0x8f, 0xdc, 0xb3, 0xbf,
	0xfe, 0xaa, 0x77, 0xe5, 0x33, 0x73, 0x09, 0x5f, 0x9c, 0xfa, 0x19, 0x5b, 0xf6, 0x74, 0xed, 0xe5,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xe4, 0x75, 0x11, 0xe8, 0x04, 0x00, 0x00,
}
