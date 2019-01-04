// Code generated by protoc-gen-gogo.
// source: image_layer.proto
// DO NOT EDIT!

package models

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strconv "strconv"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ImageLayer_DigestAlgorithm int32

const (
	DigestAlgorithmInvalid ImageLayer_DigestAlgorithm = 0
	DigestAlgorithmSha256  ImageLayer_DigestAlgorithm = 1
	DigestAlgorithmSha512  ImageLayer_DigestAlgorithm = 2
)

var ImageLayer_DigestAlgorithm_name = map[int32]string{
	0: "DigestAlgorithmInvalid",
	1: "SHA256",
	2: "SHA512",
}
var ImageLayer_DigestAlgorithm_value = map[string]int32{
	"DigestAlgorithmInvalid": 0,
	"SHA256":                 1,
	"SHA512":                 2,
}

func (x ImageLayer_DigestAlgorithm) Enum() *ImageLayer_DigestAlgorithm {
	p := new(ImageLayer_DigestAlgorithm)
	*p = x
	return p
}
func (x ImageLayer_DigestAlgorithm) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(ImageLayer_DigestAlgorithm_name, int32(x))
}
func (x *ImageLayer_DigestAlgorithm) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ImageLayer_DigestAlgorithm_value, data, "ImageLayer_DigestAlgorithm")
	if err != nil {
		return err
	}
	*x = ImageLayer_DigestAlgorithm(value)
	return nil
}
func (ImageLayer_DigestAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorImageLayer, []int{0, 0}
}

type ImageLayer_MediaType int32

const (
	MediaTypeInvalid ImageLayer_MediaType = 0
	MediaTypeTgz     ImageLayer_MediaType = 1
	MediaTypeTar     ImageLayer_MediaType = 2
	MediaTypeZip     ImageLayer_MediaType = 3
)

var ImageLayer_MediaType_name = map[int32]string{
	0: "MediaTypeInvalid",
	1: "TGZ",
	2: "TAR",
	3: "ZIP",
}
var ImageLayer_MediaType_value = map[string]int32{
	"MediaTypeInvalid": 0,
	"TGZ":              1,
	"TAR":              2,
	"ZIP":              3,
}

func (x ImageLayer_MediaType) Enum() *ImageLayer_MediaType {
	p := new(ImageLayer_MediaType)
	*p = x
	return p
}
func (x ImageLayer_MediaType) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(ImageLayer_MediaType_name, int32(x))
}
func (x *ImageLayer_MediaType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ImageLayer_MediaType_value, data, "ImageLayer_MediaType")
	if err != nil {
		return err
	}
	*x = ImageLayer_MediaType(value)
	return nil
}
func (ImageLayer_MediaType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorImageLayer, []int{0, 1}
}

type ImageLayer_Type int32

const (
	LayerTypeInvalid   ImageLayer_Type = 0
	LayerTypeShared    ImageLayer_Type = 1
	LayerTypeExclusive ImageLayer_Type = 2
)

var ImageLayer_Type_name = map[int32]string{
	0: "LayerTypeInvalid",
	1: "SHARED",
	2: "EXCLUSIVE",
}
var ImageLayer_Type_value = map[string]int32{
	"LayerTypeInvalid": 0,
	"SHARED":           1,
	"EXCLUSIVE":        2,
}

func (x ImageLayer_Type) Enum() *ImageLayer_Type {
	p := new(ImageLayer_Type)
	*p = x
	return p
}
func (x ImageLayer_Type) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(ImageLayer_Type_name, int32(x))
}
func (x *ImageLayer_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ImageLayer_Type_value, data, "ImageLayer_Type")
	if err != nil {
		return err
	}
	*x = ImageLayer_Type(value)
	return nil
}
func (ImageLayer_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptorImageLayer, []int{0, 2} }

type ImageLayer struct {
	Name            string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Url             string                     `protobuf:"bytes,2,opt,name=url" json:"url"`
	DestinationPath string                     `protobuf:"bytes,3,opt,name=destination_path,json=destinationPath" json:"destination_path"`
	LayerType       ImageLayer_Type            `protobuf:"varint,4,opt,name=layer_type,json=layerType,enum=models.ImageLayer_Type" json:"layer_type"`
	MediaType       ImageLayer_MediaType       `protobuf:"varint,5,opt,name=media_type,json=mediaType,enum=models.ImageLayer_MediaType" json:"media_type"`
	DigestAlgorithm ImageLayer_DigestAlgorithm `protobuf:"varint,6,opt,name=digest_algorithm,json=digestAlgorithm,enum=models.ImageLayer_DigestAlgorithm" json:"digest_algorithm,omitempty"`
	DigestValue     string                     `protobuf:"bytes,7,opt,name=digest_value,json=digestValue" json:"digest_value,omitempty"`
}

func (m *ImageLayer) Reset()                    { *m = ImageLayer{} }
func (*ImageLayer) ProtoMessage()               {}
func (*ImageLayer) Descriptor() ([]byte, []int) { return fileDescriptorImageLayer, []int{0} }

func (m *ImageLayer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ImageLayer) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ImageLayer) GetDestinationPath() string {
	if m != nil {
		return m.DestinationPath
	}
	return ""
}

func (m *ImageLayer) GetLayerType() ImageLayer_Type {
	if m != nil {
		return m.LayerType
	}
	return LayerTypeInvalid
}

func (m *ImageLayer) GetMediaType() ImageLayer_MediaType {
	if m != nil {
		return m.MediaType
	}
	return MediaTypeInvalid
}

func (m *ImageLayer) GetDigestAlgorithm() ImageLayer_DigestAlgorithm {
	if m != nil {
		return m.DigestAlgorithm
	}
	return DigestAlgorithmInvalid
}

func (m *ImageLayer) GetDigestValue() string {
	if m != nil {
		return m.DigestValue
	}
	return ""
}

func init() {
	proto.RegisterType((*ImageLayer)(nil), "models.ImageLayer")
	proto.RegisterEnum("models.ImageLayer_DigestAlgorithm", ImageLayer_DigestAlgorithm_name, ImageLayer_DigestAlgorithm_value)
	proto.RegisterEnum("models.ImageLayer_MediaType", ImageLayer_MediaType_name, ImageLayer_MediaType_value)
	proto.RegisterEnum("models.ImageLayer_Type", ImageLayer_Type_name, ImageLayer_Type_value)
}
func (x ImageLayer_DigestAlgorithm) String() string {
	s, ok := ImageLayer_DigestAlgorithm_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ImageLayer_MediaType) String() string {
	s, ok := ImageLayer_MediaType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ImageLayer_Type) String() string {
	s, ok := ImageLayer_Type_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *ImageLayer) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ImageLayer)
	if !ok {
		that2, ok := that.(ImageLayer)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Url != that1.Url {
		return false
	}
	if this.DestinationPath != that1.DestinationPath {
		return false
	}
	if this.LayerType != that1.LayerType {
		return false
	}
	if this.MediaType != that1.MediaType {
		return false
	}
	if this.DigestAlgorithm != that1.DigestAlgorithm {
		return false
	}
	if this.DigestValue != that1.DigestValue {
		return false
	}
	return true
}
func (this *ImageLayer) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&models.ImageLayer{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Url: "+fmt.Sprintf("%#v", this.Url)+",\n")
	s = append(s, "DestinationPath: "+fmt.Sprintf("%#v", this.DestinationPath)+",\n")
	s = append(s, "LayerType: "+fmt.Sprintf("%#v", this.LayerType)+",\n")
	s = append(s, "MediaType: "+fmt.Sprintf("%#v", this.MediaType)+",\n")
	s = append(s, "DigestAlgorithm: "+fmt.Sprintf("%#v", this.DigestAlgorithm)+",\n")
	s = append(s, "DigestValue: "+fmt.Sprintf("%#v", this.DigestValue)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringImageLayer(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ImageLayer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ImageLayer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintImageLayer(dAtA, i, uint64(len(m.Name)))
	i += copy(dAtA[i:], m.Name)
	dAtA[i] = 0x12
	i++
	i = encodeVarintImageLayer(dAtA, i, uint64(len(m.Url)))
	i += copy(dAtA[i:], m.Url)
	dAtA[i] = 0x1a
	i++
	i = encodeVarintImageLayer(dAtA, i, uint64(len(m.DestinationPath)))
	i += copy(dAtA[i:], m.DestinationPath)
	dAtA[i] = 0x20
	i++
	i = encodeVarintImageLayer(dAtA, i, uint64(m.LayerType))
	dAtA[i] = 0x28
	i++
	i = encodeVarintImageLayer(dAtA, i, uint64(m.MediaType))
	dAtA[i] = 0x30
	i++
	i = encodeVarintImageLayer(dAtA, i, uint64(m.DigestAlgorithm))
	dAtA[i] = 0x3a
	i++
	i = encodeVarintImageLayer(dAtA, i, uint64(len(m.DigestValue)))
	i += copy(dAtA[i:], m.DigestValue)
	return i, nil
}

func encodeFixed64ImageLayer(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32ImageLayer(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintImageLayer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ImageLayer) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovImageLayer(uint64(l))
	l = len(m.Url)
	n += 1 + l + sovImageLayer(uint64(l))
	l = len(m.DestinationPath)
	n += 1 + l + sovImageLayer(uint64(l))
	n += 1 + sovImageLayer(uint64(m.LayerType))
	n += 1 + sovImageLayer(uint64(m.MediaType))
	n += 1 + sovImageLayer(uint64(m.DigestAlgorithm))
	l = len(m.DigestValue)
	n += 1 + l + sovImageLayer(uint64(l))
	return n
}

func sovImageLayer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozImageLayer(x uint64) (n int) {
	return sovImageLayer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ImageLayer) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ImageLayer{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Url:` + fmt.Sprintf("%v", this.Url) + `,`,
		`DestinationPath:` + fmt.Sprintf("%v", this.DestinationPath) + `,`,
		`LayerType:` + fmt.Sprintf("%v", this.LayerType) + `,`,
		`MediaType:` + fmt.Sprintf("%v", this.MediaType) + `,`,
		`DigestAlgorithm:` + fmt.Sprintf("%v", this.DigestAlgorithm) + `,`,
		`DigestValue:` + fmt.Sprintf("%v", this.DigestValue) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringImageLayer(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ImageLayer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImageLayer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ImageLayer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImageLayer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageLayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthImageLayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageLayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthImageLayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageLayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthImageLayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LayerType", wireType)
			}
			m.LayerType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageLayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LayerType |= (ImageLayer_Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MediaType", wireType)
			}
			m.MediaType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageLayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MediaType |= (ImageLayer_MediaType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DigestAlgorithm", wireType)
			}
			m.DigestAlgorithm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageLayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DigestAlgorithm |= (ImageLayer_DigestAlgorithm(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DigestValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageLayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthImageLayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DigestValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipImageLayer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImageLayer
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
func skipImageLayer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowImageLayer
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
					return 0, ErrIntOverflowImageLayer
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
					return 0, ErrIntOverflowImageLayer
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthImageLayer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowImageLayer
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
				next, err := skipImageLayer(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthImageLayer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowImageLayer   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("image_layer.proto", fileDescriptorImageLayer) }

var fileDescriptorImageLayer = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x6f, 0xd3, 0x3e,
	0x18, 0xc7, 0xe3, 0xb6, 0xbf, 0xfe, 0x54, 0x33, 0xad, 0xc6, 0x40, 0xc9, 0x22, 0xe4, 0x46, 0x91,
	0x26, 0x4d, 0x68, 0xb4, 0x5a, 0xa5, 0xee, 0xde, 0xb2, 0x0a, 0x8a, 0x8a, 0x98, 0xd2, 0x32, 0xa1,
	0x5e, 0x2a, 0x6f, 0x31, 0x89, 0xa5, 0xa4, 0x09, 0xa9, 0x5b, 0x51, 0xb8, 0x70, 0xee, 0x89, 0x37,
	0xd0, 0x3b, 0x12, 0x6f, 0x64, 0xc7, 0x1d, 0x39, 0x4d, 0x34, 0x5c, 0x10, 0xa7, 0xbd, 0x04, 0x14,
	0xf7, 0x5f, 0x16, 0x8d, 0x5b, 0xfc, 0x7c, 0x3f, 0xcf, 0xc7, 0x8f, 0x1f, 0x05, 0xde, 0xe7, 0x1e,
	0xb5, 0xd9, 0xc0, 0xa5, 0x53, 0x16, 0x56, 0x82, 0xd0, 0x17, 0x3e, 0xce, 0x7b, 0xbe, 0xc5, 0xdc,
	0x91, 0xf6, 0xcc, 0xe6, 0xc2, 0x19, 0x9f, 0x57, 0x2e, 0x7c, 0xaf, 0x6a, 0xfb, 0xb6, 0x5f, 0x95,
	0xf1, 0xf9, 0xf8, 0xbd, 0x3c, 0xc9, 0x83, 0xfc, 0x5a, 0xb6, 0x19, 0xdf, 0xf3, 0x10, 0xb6, 0x63,
	0x59, 0x27, 0x76, 0xe1, 0xa7, 0x30, 0x37, 0xa4, 0x1e, 0x53, 0x81, 0x0e, 0x0e, 0x0a, 0xcd, 0xd2,
	0xe5, 0x75, 0x59, 0xf9, 0x73, 0x5d, 0xde, 0x8d, 0x6b, 0x87, 0xbe, 0xc7, 0x05, 0xf3, 0x02, 0x31,
	0x35, 0x25, 0x83, 0x4b, 0x30, 0x3b, 0x0e, 0x5d, 0x35, 0x23, 0xd1, 0x5c, 0x8c, 0x9a, 0x71, 0x01,
	0x57, 0x21, 0xb2, 0xd8, 0x48, 0xf0, 0x21, 0x15, 0xdc, 0x1f, 0x0e, 0x02, 0x2a, 0x1c, 0x35, 0x9b,
	0x80, 0x8a, 0x89, 0xf4, 0x94, 0x0a, 0x07, 0xbf, 0x82, 0x50, 0xbe, 0x64, 0x20, 0xa6, 0x01, 0x53,
	0x73, 0x3a, 0x38, 0xd8, 0xad, 0x3d, 0xae, 0x2c, 0xdf, 0x53, 0xd9, 0x0e, 0x57, 0xe9, 0x4d, 0x03,
	0xd6, 0xc4, 0xab, 0x99, 0x12, 0x2d, 0x66, 0x41, 0x7e, 0xc7, 0x31, 0x7e, 0x03, 0xa1, 0xc7, 0x2c,
	0x4e, 0x97, 0xae, 0xff, 0xa4, 0xeb, 0xc9, 0x1d, 0xae, 0xd7, 0x31, 0x74, 0x5b, 0xb8, 0xed, 0x33,
	0x0b, 0xde, 0x3a, 0xc6, 0x1f, 0x20, 0xb2, 0xb8, 0xcd, 0x46, 0x62, 0x40, 0x5d, 0xdb, 0x0f, 0xb9,
	0x70, 0x3c, 0x35, 0x2f, 0xb5, 0xc6, 0x1d, 0xda, 0x13, 0x89, 0x36, 0xd6, 0x64, 0xd3, 0x58, 0xc9,
	0xb5, 0xb4, 0x23, 0xb1, 0xcd, 0xa2, 0x75, 0xbb, 0x09, 0x37, 0xe0, 0xce, 0x0a, 0x9f, 0x50, 0x77,
	0xcc, 0xd4, 0xff, 0xe5, 0xf2, 0xc8, 0x4a, 0x55, 0x4a, 0x66, 0x09, 0xcd, 0xbd, 0x65, 0xfd, 0x2c,
	0x2e, 0x1b, 0x9f, 0x61, 0x31, 0x35, 0x0a, 0xd6, 0x60, 0x29, 0x55, 0x6a, 0x0f, 0x27, 0xd4, 0xe5,
	0x16, 0x52, 0xf0, 0x3e, 0xcc, 0x77, 0x5f, 0x36, 0x6a, 0xf5, 0x63, 0x04, 0xb4, 0xbd, 0xd9, 0x5c,
	0x7f, 0x94, 0x22, 0xbb, 0x0e, 0xad, 0xd5, 0x8f, 0x57, 0x58, 0xfd, 0xa8, 0x86, 0x32, 0xff, 0xc2,
	0xea, 0x47, 0x35, 0x23, 0x84, 0x85, 0xcd, 0x7a, 0xf1, 0x43, 0x88, 0x36, 0x87, 0xed, 0x85, 0x7b,
	0x30, 0xdb, 0x7b, 0xd1, 0x47, 0x40, 0x43, 0xb3, 0xb9, 0xbe, 0xb3, 0x01, 0x7a, 0xf6, 0x27, 0x19,
	0x35, 0x4c, 0x94, 0x49, 0x47, 0x34, 0x8c, 0xa3, 0x7e, 0xfb, 0x14, 0x65, 0x53, 0x51, 0x9f, 0x07,
	0x86, 0x05, 0x73, 0xeb, 0xeb, 0x3a, 0xeb, 0x9f, 0x61, 0x7b, 0x5d, 0x59, 0x0e, 0x6e, 0xb6, 0x4e,
	0x10, 0xd0, 0x1e, 0xcc, 0xe6, 0x7a, 0x71, 0xc3, 0x74, 0x1d, 0x1a, 0x32, 0x0b, 0xef, 0xc3, 0x42,
	0xeb, 0xdd, 0xf3, 0xce, 0xdb, 0x6e, 0xfb, 0xac, 0x85, 0x32, 0x5a, 0x69, 0x36, 0xd7, 0xf1, 0x86,
	0x69, 0x7d, 0xbc, 0x70, 0xc7, 0x23, 0x3e, 0x61, 0xcd, 0xc3, 0xab, 0x05, 0x51, 0x7e, 0x2c, 0x88,
	0x72, 0xb3, 0x20, 0xe0, 0x4b, 0x44, 0xc0, 0xb7, 0x88, 0x80, 0xcb, 0x88, 0x80, 0xab, 0x88, 0x80,
	0x9f, 0x11, 0x01, 0xbf, 0x23, 0xa2, 0xdc, 0x44, 0x04, 0x7c, 0xfd, 0x45, 0x94, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x42, 0x2b, 0x36, 0x0b, 0xa6, 0x03, 0x00, 0x00,
}