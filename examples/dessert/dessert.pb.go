// Code generated by protoc-gen-gogo.
// source: dessert.proto
// DO NOT EDIT!

/*
Package flavortown_dessert is a generated protocol buffer package.

It is generated from these files:
	dessert.proto

It has these top-level messages:
	Dessert
*/
package flavortown_dessert

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/opsee/protobuf/opseeproto"

import github_com_graphql_go_graphql "github.com/graphql-go/graphql"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

// A delicious dessert dish on the menu
type Dessert struct {
	// The name of the dish
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// How sweet is the dish, an integer between 0 and 10
	Sweetness int32 `protobuf:"varint,2,opt,name=sweetness,proto3" json:"sweetness,omitempty"`
}

func (m *Dessert) Reset()                    { *m = Dessert{} }
func (m *Dessert) String() string            { return proto.CompactTextString(m) }
func (*Dessert) ProtoMessage()               {}
func (*Dessert) Descriptor() ([]byte, []int) { return fileDescriptorDessert, []int{0} }

func init() {
	proto.RegisterType((*Dessert)(nil), "flavortown.dessert.Dessert")
}
func (this *Dessert) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Dessert)
	if !ok {
		that2, ok := that.(Dessert)
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
	if this.Sweetness != that1.Sweetness {
		return false
	}
	return true
}

type DessertGetter interface {
	GetDessert() *Dessert
}

var GraphQLDessertType *github_com_graphql_go_graphql.Object

func init() {
	GraphQLDessertType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "flavortown_dessertDessert",
		Description: "A delicious dessert dish on the menu",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"name": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.String,
					Description: "The name of the dish",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Dessert)
						if ok {
							return obj.Name, nil
						}
						inter, ok := p.Source.(DessertGetter)
						if ok {
							face := inter.GetDessert()
							if face == nil {
								return nil, nil
							}
							return face.Name, nil
						}
						return nil, fmt.Errorf("field name not resolved")
					},
				},
				"sweetness": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Int,
					Description: "How sweet is the dish, an integer between 0 and 10",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Dessert)
						if ok {
							return obj.Sweetness, nil
						}
						inter, ok := p.Source.(DessertGetter)
						if ok {
							face := inter.GetDessert()
							if face == nil {
								return nil, nil
							}
							return face.Sweetness, nil
						}
						return nil, fmt.Errorf("field sweetness not resolved")
					},
				},
			}
		}),
	})
}
func NewPopulatedDessert(r randyDessert, easy bool) *Dessert {
	this := &Dessert{}
	this.Name = randStringDessert(r)
	this.Sweetness = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Sweetness *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyDessert interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneDessert(r randyDessert) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringDessert(r randyDessert) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneDessert(r)
	}
	return string(tmps)
}
func randUnrecognizedDessert(r randyDessert, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldDessert(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldDessert(data []byte, r randyDessert, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateDessert(data, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		data = encodeVarintPopulateDessert(data, uint64(v2))
	case 1:
		data = encodeVarintPopulateDessert(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateDessert(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateDessert(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateDessert(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateDessert(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}

var fileDescriptorDessert = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x49, 0x2d, 0x2e,
	0x4e, 0x2d, 0x2a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4a, 0xcb, 0x49, 0x2c, 0xcb,
	0x2f, 0x2a, 0xc9, 0x2f, 0xcf, 0xd3, 0x83, 0xca, 0x48, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0x95, 0x26, 0x95, 0xa6, 0x81,
	0x79, 0x60, 0x0e, 0x98, 0x05, 0x31, 0x42, 0xca, 0x00, 0x49, 0x79, 0x7e, 0x41, 0x71, 0x6a, 0x2a,
	0x42, 0x3d, 0x98, 0x0b, 0xd1, 0x00, 0x66, 0x42, 0x74, 0x28, 0x59, 0x73, 0xb1, 0xbb, 0x40, 0xec,
	0x12, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02,
	0xb3, 0x85, 0x64, 0xb8, 0x38, 0x8b, 0xcb, 0x53, 0x53, 0x4b, 0xf2, 0x80, 0x6a, 0x24, 0x98, 0x80,
	0x12, 0xac, 0x41, 0x08, 0x01, 0x27, 0x81, 0x1f, 0x0f, 0xe5, 0x18, 0x57, 0x3c, 0x92, 0x63, 0xdc,
	0x01, 0xc4, 0x07, 0x16, 0xc9, 0x33, 0x26, 0xb1, 0x81, 0x4d, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x52, 0x6a, 0x5f, 0x25, 0xdb, 0x00, 0x00, 0x00,
}
