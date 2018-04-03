// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/category/category.proto

/*
Package category is a generated protocol buffer package.

It is generated from these files:
	proto/category/category.proto

It has these top-level messages:
	Category
	UpsertCategoryReq
	UpsertCategoryRsp
	Categories
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

type UpsertCategoryReq struct {
	Catetory *Category `protobuf:"bytes,1,opt,name=catetory" json:"catetory,omitempty"`
	IsNew    bool      `protobuf:"varint,2,opt,name=is_new,json=isNew" json:"is_new,omitempty"`
}

func (m *UpsertCategoryReq) Reset()                    { *m = UpsertCategoryReq{} }
func (m *UpsertCategoryReq) String() string            { return proto.CompactTextString(m) }
func (*UpsertCategoryReq) ProtoMessage()               {}
func (*UpsertCategoryReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UpsertCategoryReq) GetCatetory() *Category {
	if m != nil {
		return m.Catetory
	}
	return nil
}

func (m *UpsertCategoryReq) GetIsNew() bool {
	if m != nil {
		return m.IsNew
	}
	return false
}

type UpsertCategoryRsp struct {
	Category *Category `protobuf:"bytes,1,opt,name=category" json:"category,omitempty"`
}

func (m *UpsertCategoryRsp) Reset()                    { *m = UpsertCategoryRsp{} }
func (m *UpsertCategoryRsp) String() string            { return proto.CompactTextString(m) }
func (*UpsertCategoryRsp) ProtoMessage()               {}
func (*UpsertCategoryRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpsertCategoryRsp) GetCategory() *Category {
	if m != nil {
		return m.Category
	}
	return nil
}

type Categories struct {
	Categories []*Category `protobuf:"bytes,1,rep,name=categories" json:"categories,omitempty"`
}

func (m *Categories) Reset()                    { *m = Categories{} }
func (m *Categories) String() string            { return proto.CompactTextString(m) }
func (*Categories) ProtoMessage()               {}
func (*Categories) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Categories) GetCategories() []*Category {
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
func (*InsertCategoriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
func (*UpdateCategoriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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
func (*FindCategoriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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
func (*DeleteCategoriesByIDsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeleteCategoriesByIDsRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto.RegisterType((*Category)(nil), "go.micro.srv.courses.Category")
	proto.RegisterType((*UpsertCategoryReq)(nil), "go.micro.srv.courses.UpsertCategoryReq")
	proto.RegisterType((*UpsertCategoryRsp)(nil), "go.micro.srv.courses.UpsertCategoryRsp")
	proto.RegisterType((*Categories)(nil), "go.micro.srv.courses.Categories")
	proto.RegisterType((*InsertCategoriesResponse)(nil), "go.micro.srv.courses.InsertCategoriesResponse")
	proto.RegisterType((*UpdateCategoriesResponse)(nil), "go.micro.srv.courses.UpdateCategoriesResponse")
	proto.RegisterType((*FindCategoriesRequest)(nil), "go.micro.srv.courses.FindCategoriesRequest")
	proto.RegisterType((*DeleteCategoriesByIDsRequest)(nil), "go.micro.srv.courses.DeleteCategoriesByIDsRequest")
}

func init() { proto.RegisterFile("proto/category/category.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x93, 0x34, 0x75, 0x27, 0xa5, 0x94, 0x55, 0x5b, 0x2d, 0xe6, 0x66, 0xcc, 0x03, 0x91,
	0x90, 0x5c, 0x14, 0x10, 0x12, 0x3c, 0x20, 0xb5, 0x04, 0xa4, 0x48, 0x88, 0x22, 0x43, 0x5f, 0x78,
	0x89, 0x5c, 0xef, 0xd4, 0xac, 0x94, 0x78, 0xb7, 0xbb, 0xeb, 0x56, 0xf9, 0x02, 0x3e, 0x81, 0x6f,
	0xe0, 0x2f, 0x91, 0x37, 0xb6, 0xe5, 0xb6, 0xee, 0x05, 0xde, 0x66, 0xce, 0x9e, 0xb3, 0xb3, 0x33,
	0x73, 0x6c, 0x78, 0x24, 0x95, 0x30, 0x62, 0x37, 0x89, 0x0d, 0xa6, 0x42, 0x2d, 0xea, 0x20, 0xb4,
	0x38, 0xd9, 0x4a, 0x45, 0x38, 0xe7, 0x89, 0x12, 0xa1, 0x56, 0xa7, 0x61, 0x22, 0x72, 0xa5, 0x51,
	0x7b, 0x4f, 0x52, 0x21, 0xd2, 0x19, 0xee, 0x5a, 0xce, 0x51, 0x7e, 0xbc, 0x6b, 0xf8, 0x1c, 0xb5,
	0x89, 0xe7, 0x72, 0x29, 0xf3, 0x1e, 0x5c, 0x24, 0xe0, 0x5c, 0x9a, 0xf2, 0xce, 0xe0, 0x77, 0x07,
	0xdc, 0x0f, 0x65, 0x19, 0xb2, 0x01, 0x9d, 0xc9, 0x98, 0x3a, 0xbe, 0x33, 0x5c, 0x8b, 0x3a, 0x93,
	0x31, 0x21, 0xd0, 0xcb, 0xe2, 0x39, 0xd2, 0x8e, 0x45, 0x6c, 0x5c, 0x60, 0x7a, 0x96, 0xa7, 0xb4,
	0xbb, 0xc4, 0x8a, 0x98, 0x3c, 0x83, 0x3b, 0x8c, 0x6b, 0x39, 0x8b, 0x17, 0x53, 0xa1, 0x18, 0x2a,
	0xda, 0xf3, 0x9d, 0x61, 0x2f, 0x5a, 0x2f, 0xc1, 0x83, 0x02, 0x23, 0x3e, 0x0c, 0x18, 0xea, 0x44,
	0x71, 0x69, 0xb8, 0xc8, 0xe8, 0x8a, 0xd5, 0x37, 0x21, 0xb2, 0x03, 0xfd, 0x9f, 0x9c, 0x31, 0xcc,
	0x68, 0xdf, 0x77, 0x86, 0x6e, 0x54, 0x66, 0xe4, 0x2d, 0x40, 0xa2, 0x30, 0x36, 0xc8, 0xa6, 0xb1,
	0xa1, 0xab, 0xbe, 0x33, 0x1c, 0x8c, 0xbc, 0x70, 0xd9, 0x55, 0x58, 0x75, 0x15, 0x7e, 0xaf, 0xda,
	0x8e, 0xd6, 0x4a, 0xf6, 0x9e, 0x29, 0xa4, 0xb9, 0x64, 0x95, 0xd4, 0xbd, 0x59, 0x5a, 0xb2, 0xf7,
	0x4c, 0x70, 0x0c, 0xf7, 0x0e, 0xa5, 0x46, 0x65, 0xaa, 0xf1, 0x44, 0x78, 0x42, 0xde, 0x81, 0x5b,
	0x2c, 0xc5, 0x08, 0xb5, 0xb0, 0x73, 0x1a, 0x8c, 0x1e, 0x87, 0x6d, 0x5b, 0x09, 0x6b, 0x51, 0xcd,
	0x27, 0xdb, 0xd0, 0xe7, 0x7a, 0x9a, 0xe1, 0x99, 0x9d, 0xa7, 0x1b, 0xad, 0x70, 0xfd, 0x05, 0xcf,
	0x82, 0x83, 0x4b, 0x75, 0xb4, 0xac, 0xea, 0xa4, 0xff, 0x58, 0xa7, 0x88, 0x82, 0xcf, 0x00, 0x25,
	0xca, 0x51, 0x93, 0xf7, 0x00, 0x49, 0x9d, 0x51, 0xc7, 0xef, 0xde, 0xe2, 0xae, 0x86, 0x22, 0x78,
	0x03, 0x74, 0x92, 0x35, 0x9e, 0xc7, 0x51, 0x47, 0xa8, 0xa5, 0xc8, 0x34, 0x12, 0x0f, 0x5c, 0x6e,
	0xcf, 0x90, 0xd9, 0x57, 0x76, 0xa3, 0x3a, 0x0f, 0x5e, 0x03, 0x3d, 0xb4, 0xb3, 0x6c, 0xd1, 0x51,
	0x58, 0x2d, 0xe7, 0x5c, 0xca, 0xaa, 0x34, 0xf8, 0xe3, 0xc0, 0xf6, 0x27, 0x9e, 0xb1, 0xa6, 0xe8,
	0x24, 0x47, 0x6d, 0xc8, 0x26, 0x74, 0x39, 0x5b, 0x36, 0xb0, 0x16, 0x15, 0x21, 0xd9, 0x82, 0x95,
	0x93, 0x1c, 0xd5, 0xa2, 0xb4, 0xe7, 0x32, 0x69, 0x98, 0xa8, 0x7b, 0xce, 0x44, 0x85, 0x6f, 0x85,
	0x32, 0xb4, 0x67, 0x2f, 0xb0, 0x31, 0xb9, 0x0f, 0xae, 0x44, 0x35, 0x95, 0x71, 0x8a, 0xd6, 0x8f,
	0xdd, 0x68, 0x55, 0xa2, 0xfa, 0x1a, 0xa7, 0x48, 0x9e, 0xc2, 0x7a, 0x92, 0x2b, 0x85, 0x99, 0x59,
	0x1e, 0xf7, 0xed, 0xf1, 0xa0, 0xc4, 0x0a, 0x4a, 0xf0, 0x12, 0x1e, 0x8e, 0x71, 0x86, 0xcd, 0x0e,
	0xf7, 0x17, 0x93, 0xf1, 0xd5, 0x2f, 0x1e, 0xfd, 0xea, 0xc1, 0xdd, 0x6a, 0xc8, 0xdf, 0x50, 0x9d,
	0xf2, 0x04, 0xc9, 0x11, 0x6c, 0x9c, 0x5f, 0x3f, 0x79, 0xde, 0xbe, 0x9d, 0x4b, 0x66, 0xf4, 0x6e,
	0x47, 0xd4, 0x92, 0x30, 0xd8, 0xbc, 0xb8, 0x43, 0xe2, 0x5f, 0xeb, 0x01, 0x8e, 0xda, 0x0b, 0xdb,
	0x19, 0x57, 0xba, 0x81, 0xc1, 0xe6, 0xc5, 0x8d, 0xff, 0x7f, 0x95, 0x2b, 0xbd, 0x33, 0x85, 0x8d,
	0xf3, 0x06, 0x21, 0x2f, 0xda, 0x6f, 0x68, 0xb5, 0x91, 0x77, 0xe3, 0x83, 0x48, 0x02, 0xdb, 0xad,
	0x6b, 0x25, 0xa3, 0x76, 0xe9, 0x75, 0x1e, 0xf0, 0x76, 0x2e, 0xfd, 0x6b, 0x3e, 0x16, 0x3f, 0xdf,
	0x7d, 0xf8, 0x51, 0x7f, 0xaf, 0x47, 0x7d, 0x7b, 0xf6, 0xea, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x82, 0x51, 0x99, 0x16, 0xfe, 0x05, 0x00, 0x00,
}
