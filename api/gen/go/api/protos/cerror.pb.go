// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/protos/cerror.proto

package genpb

import (
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Error struct {
	Type         uint32            `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	InternalCode string            `protobuf:"bytes,2,opt,name=internal_code,json=internalCode,proto3" json:"internal_code,omitempty"`
	Text         string            `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Description  string            `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Reason       string            `protobuf:"bytes,5,opt,name=reason,proto3" json:"reason,omitempty"`
	IsSensible   string            `protobuf:"bytes,6,opt,name=is_sensible,json=isSensible,proto3" json:"is_sensible,omitempty"`
	NeedReport   string            `protobuf:"bytes,7,opt,name=need_report,json=needReport,proto3" json:"need_report,omitempty"`
	Meta         map[string]string `protobuf:"bytes,8,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ComesFrom    *Error            `protobuf:"bytes,9,opt,name=comes_from,json=comesFrom,proto3" json:"comes_from,omitempty"`
}

func (m *Error) Reset()      { *m = Error{} }
func (*Error) ProtoMessage() {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_99114d4472a2cf7d, []int{0}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Error.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return m.Size()
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Error) GetInternalCode() string {
	if m != nil {
		return m.InternalCode
	}
	return ""
}

func (m *Error) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Error) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *Error) GetIsSensible() string {
	if m != nil {
		return m.IsSensible
	}
	return ""
}

func (m *Error) GetNeedReport() string {
	if m != nil {
		return m.NeedReport
	}
	return ""
}

func (m *Error) GetMeta() map[string]string {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Error) GetComesFrom() *Error {
	if m != nil {
		return m.ComesFrom
	}
	return nil
}

func init() {
	proto.RegisterType((*Error)(nil), "api.gen.Error")
	proto.RegisterMapType((map[string]string)(nil), "api.gen.Error.MetaEntry")
}

func init() { proto.RegisterFile("api/protos/cerror.proto", fileDescriptor_99114d4472a2cf7d) }

var fileDescriptor_99114d4472a2cf7d = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x33, 0x6d, 0xd3, 0x9a, 0x89, 0x15, 0x19, 0x44, 0x87, 0x22, 0x63, 0xd0, 0x4d, 0x16,
	0x9a, 0x42, 0x5d, 0x28, 0x2e, 0x5b, 0xea, 0x4e, 0x28, 0x71, 0x21, 0xb8, 0x09, 0xd3, 0xe4, 0x18,
	0x06, 0x93, 0x99, 0x30, 0x33, 0x8a, 0xdd, 0xf9, 0x08, 0x3e, 0x86, 0x8f, 0xe2, 0xb2, 0xcb, 0x2e,
	0x6d, 0x0a, 0x97, 0xbb, 0xec, 0x23, 0x5c, 0x32, 0x6d, 0x2f, 0xf7, 0xee, 0xce, 0xf9, 0xbe, 0xdf,
	0x99, 0x3f, 0xe7, 0xc3, 0xcf, 0x78, 0x23, 0xa6, 0x8d, 0x56, 0x56, 0x99, 0x69, 0x0e, 0x5a, 0x2b,
	0x9d, 0xb8, 0x8e, 0x8c, 0x78, 0x23, 0x92, 0x12, 0xe4, 0xe4, 0x79, 0xa9, 0x54, 0x59, 0xc1, 0xb4,
	0x03, 0xb9, 0x94, 0xca, 0x72, 0x2b, 0x94, 0x34, 0x27, 0xec, 0xe5, 0x55, 0x0f, 0xfb, 0xcb, 0x6e,
	0x8c, 0x10, 0x3c, 0xb0, 0x9b, 0x06, 0x28, 0x8a, 0x50, 0x3c, 0x4e, 0x5d, 0x4d, 0x5e, 0xe1, 0xb1,
	0x90, 0x16, 0xb4, 0xe4, 0x55, 0x96, 0xab, 0x02, 0x68, 0x2f, 0x42, 0x71, 0x90, 0x3e, 0xbc, 0x88,
	0x0b, 0x55, 0x80, 0x1b, 0x84, 0x5f, 0x96, 0xf6, 0x9d, 0xe7, 0x6a, 0x12, 0xe1, 0xb0, 0x00, 0x93,
	0x6b, 0xd1, 0x74, 0x97, 0xd1, 0x81, 0xb3, 0xee, 0x4a, 0xe4, 0x29, 0x1e, 0x6a, 0xe0, 0x46, 0x49,
	0xea, 0x3b, 0xf3, 0xdc, 0x91, 0x17, 0x38, 0x14, 0x26, 0x33, 0x20, 0x8d, 0x58, 0x57, 0x40, 0x87,
	0xce, 0xc4, 0xc2, 0x7c, 0x3e, 0x2b, 0x1d, 0x20, 0x01, 0x8a, 0x4c, 0x43, 0xa3, 0xb4, 0xa5, 0xa3,
	0x13, 0xd0, 0x49, 0xa9, 0x53, 0xc8, 0x6b, 0x3c, 0xa8, 0xc1, 0x72, 0xfa, 0x20, 0xea, 0xc7, 0xe1,
	0x8c, 0x26, 0xe7, 0x45, 0x24, 0xee, 0x9b, 0xc9, 0x27, 0xb0, 0x7c, 0x29, 0xad, 0xde, 0xa4, 0x8e,
	0x22, 0x6f, 0x30, 0xce, 0x55, 0x0d, 0x26, 0xfb, 0xa6, 0x55, 0x4d, 0x83, 0x08, 0xc5, 0xe1, 0xec,
	0xd1, 0xfd, 0x99, 0x34, 0x70, 0xc4, 0x47, 0xad, 0xea, 0xc9, 0x3b, 0x1c, 0xdc, 0x9e, 0x40, 0x1e,
	0xe3, 0xfe, 0x77, 0xd8, 0xb8, 0x8d, 0x05, 0x69, 0x57, 0x92, 0x27, 0xd8, 0xff, 0xc9, 0xab, 0x1f,
	0x97, 0x45, 0x9d, 0x9a, 0x0f, 0xbd, 0xf7, 0x68, 0xfe, 0x65, 0xbb, 0x67, 0xde, 0x6e, 0xcf, 0xbc,
	0xe3, 0x9e, 0xa1, 0xdf, 0x2d, 0x43, 0x7f, 0x5b, 0x86, 0xfe, 0xb5, 0x0c, 0x6d, 0x5b, 0x86, 0xfe,
	0xb7, 0x0c, 0x5d, 0xb7, 0xcc, 0x3b, 0xb6, 0x0c, 0xfd, 0x39, 0x30, 0x6f, 0x7b, 0x60, 0xde, 0xee,
	0xc0, 0x3c, 0x1c, 0xe6, 0xaa, 0xbe, 0x3c, 0x66, 0x1e, 0x2e, 0x5c, 0xc0, 0xab, 0x2e, 0xb8, 0x15,
	0xfa, 0xea, 0x97, 0x20, 0x9b, 0xf5, 0x7a, 0xe8, 0x82, 0x7c, 0x7b, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x43, 0xea, 0x4e, 0x9e, 0x0a, 0x02, 0x00, 0x00,
}

func (this *Error) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Error)
	if !ok {
		that2, ok := that.(Error)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.InternalCode != that1.InternalCode {
		return false
	}
	if this.Text != that1.Text {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Reason != that1.Reason {
		return false
	}
	if this.IsSensible != that1.IsSensible {
		return false
	}
	if this.NeedReport != that1.NeedReport {
		return false
	}
	if len(this.Meta) != len(that1.Meta) {
		return false
	}
	for i := range this.Meta {
		if this.Meta[i] != that1.Meta[i] {
			return false
		}
	}
	if !this.ComesFrom.Equal(that1.ComesFrom) {
		return false
	}
	return true
}
func (this *Error) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 13)
	s = append(s, "&genpb.Error{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "InternalCode: "+fmt.Sprintf("%#v", this.InternalCode)+",\n")
	s = append(s, "Text: "+fmt.Sprintf("%#v", this.Text)+",\n")
	s = append(s, "Description: "+fmt.Sprintf("%#v", this.Description)+",\n")
	s = append(s, "Reason: "+fmt.Sprintf("%#v", this.Reason)+",\n")
	s = append(s, "IsSensible: "+fmt.Sprintf("%#v", this.IsSensible)+",\n")
	s = append(s, "NeedReport: "+fmt.Sprintf("%#v", this.NeedReport)+",\n")
	keysForMeta := make([]string, 0, len(this.Meta))
	for k, _ := range this.Meta {
		keysForMeta = append(keysForMeta, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForMeta)
	mapStringForMeta := "map[string]string{"
	for _, k := range keysForMeta {
		mapStringForMeta += fmt.Sprintf("%#v: %#v,", k, this.Meta[k])
	}
	mapStringForMeta += "}"
	if this.Meta != nil {
		s = append(s, "Meta: "+mapStringForMeta+",\n")
	}
	if this.ComesFrom != nil {
		s = append(s, "ComesFrom: "+fmt.Sprintf("%#v", this.ComesFrom)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringCerror(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Error) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Error) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCerror(dAtA, i, uint64(m.Type))
	}
	if len(m.InternalCode) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCerror(dAtA, i, uint64(len(m.InternalCode)))
		i += copy(dAtA[i:], m.InternalCode)
	}
	if len(m.Text) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCerror(dAtA, i, uint64(len(m.Text)))
		i += copy(dAtA[i:], m.Text)
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCerror(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if len(m.Reason) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintCerror(dAtA, i, uint64(len(m.Reason)))
		i += copy(dAtA[i:], m.Reason)
	}
	if len(m.IsSensible) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintCerror(dAtA, i, uint64(len(m.IsSensible)))
		i += copy(dAtA[i:], m.IsSensible)
	}
	if len(m.NeedReport) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintCerror(dAtA, i, uint64(len(m.NeedReport)))
		i += copy(dAtA[i:], m.NeedReport)
	}
	if len(m.Meta) > 0 {
		for k, _ := range m.Meta {
			dAtA[i] = 0x42
			i++
			v := m.Meta[k]
			mapSize := 1 + len(k) + sovCerror(uint64(len(k))) + 1 + len(v) + sovCerror(uint64(len(v)))
			i = encodeVarintCerror(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintCerror(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintCerror(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.ComesFrom != nil {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintCerror(dAtA, i, uint64(m.ComesFrom.Size()))
		n1, err1 := m.ComesFrom.MarshalTo(dAtA[i:])
		if err1 != nil {
			return 0, err1
		}
		i += n1
	}
	return i, nil
}

func encodeVarintCerror(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Error) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovCerror(uint64(m.Type))
	}
	l = len(m.InternalCode)
	if l > 0 {
		n += 1 + l + sovCerror(uint64(l))
	}
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovCerror(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovCerror(uint64(l))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovCerror(uint64(l))
	}
	l = len(m.IsSensible)
	if l > 0 {
		n += 1 + l + sovCerror(uint64(l))
	}
	l = len(m.NeedReport)
	if l > 0 {
		n += 1 + l + sovCerror(uint64(l))
	}
	if len(m.Meta) > 0 {
		for k, v := range m.Meta {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovCerror(uint64(len(k))) + 1 + len(v) + sovCerror(uint64(len(v)))
			n += mapEntrySize + 1 + sovCerror(uint64(mapEntrySize))
		}
	}
	if m.ComesFrom != nil {
		l = m.ComesFrom.Size()
		n += 1 + l + sovCerror(uint64(l))
	}
	return n
}

func sovCerror(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCerror(x uint64) (n int) {
	return sovCerror(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Error) String() string {
	if this == nil {
		return "nil"
	}
	keysForMeta := make([]string, 0, len(this.Meta))
	for k, _ := range this.Meta {
		keysForMeta = append(keysForMeta, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForMeta)
	mapStringForMeta := "map[string]string{"
	for _, k := range keysForMeta {
		mapStringForMeta += fmt.Sprintf("%v: %v,", k, this.Meta[k])
	}
	mapStringForMeta += "}"
	s := strings.Join([]string{`&Error{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`InternalCode:` + fmt.Sprintf("%v", this.InternalCode) + `,`,
		`Text:` + fmt.Sprintf("%v", this.Text) + `,`,
		`Description:` + fmt.Sprintf("%v", this.Description) + `,`,
		`Reason:` + fmt.Sprintf("%v", this.Reason) + `,`,
		`IsSensible:` + fmt.Sprintf("%v", this.IsSensible) + `,`,
		`NeedReport:` + fmt.Sprintf("%v", this.NeedReport) + `,`,
		`Meta:` + mapStringForMeta + `,`,
		`ComesFrom:` + strings.Replace(this.ComesFrom.String(), "Error", "Error", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCerror(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Error) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCerror
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Error: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Error: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerror
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InternalCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerror
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCerror
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCerror
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InternalCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerror
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCerror
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCerror
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerror
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCerror
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCerror
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerror
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCerror
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCerror
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsSensible", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerror
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCerror
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCerror
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IsSensible = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NeedReport", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerror
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCerror
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCerror
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NeedReport = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerror
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCerror
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCerror
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Meta == nil {
				m.Meta = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCerror
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCerror
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthCerror
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthCerror
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCerror
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthCerror
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthCerror
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipCerror(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthCerror
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Meta[mapkey] = mapvalue
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComesFrom", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerror
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCerror
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCerror
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ComesFrom == nil {
				m.ComesFrom = &Error{}
			}
			if err := m.ComesFrom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCerror(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCerror
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCerror
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCerror(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCerror
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCerror
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCerror
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCerror
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCerror
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCerror
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCerror(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCerror
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCerror = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCerror   = fmt.Errorf("proto: integer overflow")
)