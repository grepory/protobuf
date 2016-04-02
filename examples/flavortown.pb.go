// Code generated by protoc-gen-gogo.
// source: flavortown.proto
// DO NOT EDIT!

/*
Package flavortown is a generated protocol buffer package.

It is generated from these files:
	flavortown.proto

It has these top-level messages:
	Menu
	LineItem
	Lunch
	Nothing
*/
package flavortown

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import opsee_types "github.com/opsee/protobuf/opseeproto/types"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/opsee/protobuf/opseeproto"
import flavortown_dessert "github.com/opsee/protobuf/examples/dessert"

import bytes "bytes"

import github_com_graphql_go_graphql "github.com/graphql-go/graphql"
import github_com_opsee_protobuf_plugin_graphql_scalars "github.com/opsee/protobuf/plugin/graphql/scalars"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

// The menu at Guy’s American Kitchen & Bar reflects his signature style of authentic and surprising flavors
type Menu struct {
	// These dishes are crafted with the heart and soul of hometown favorites and infused with Guy’s big, daring flavors
	Items []*LineItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *Menu) Reset()                    { *m = Menu{} }
func (m *Menu) String() string            { return proto.CompactTextString(m) }
func (*Menu) ProtoMessage()               {}
func (*Menu) Descriptor() ([]byte, []int) { return fileDescriptorFlavortown, []int{0} }

func (m *Menu) GetItems() []*LineItem {
	if m != nil {
		return m.Items
	}
	return nil
}

// A line item representing a dish and price
type LineItem struct {
	// The menu dish, can either be lunch or dessert
	//
	// Types that are valid to be assigned to Dish:
	//	*LineItem_Lunch
	//	*LineItem_TastyDessert
	Dish isLineItem_Dish `protobuf_oneof:"dish"`
	// The price of the dish in cents
	PriceCents int32 `protobuf:"varint,2,opt,name=price_cents,json=priceCents,proto3" json:"price_cents,omitempty"`
	// A timestamp representing when the dish was added to the menu
	CreatedAt *opsee_types.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	// A timestamp representing when the dish was updated
	UpdatedAt *opsee_types.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	// A list of nothing really
	Nothing *Nothing `protobuf:"bytes,5,opt,name=nothing" json:"nothing,omitempty"`
}

func (m *LineItem) Reset()                    { *m = LineItem{} }
func (m *LineItem) String() string            { return proto.CompactTextString(m) }
func (*LineItem) ProtoMessage()               {}
func (*LineItem) Descriptor() ([]byte, []int) { return fileDescriptorFlavortown, []int{1} }

type isLineItem_Dish interface {
	isLineItem_Dish()
	Equal(interface{}) bool
}

type LineItem_Lunch struct {
	Lunch *Lunch `protobuf:"bytes,100,opt,name=lunch,oneof"`
}
type LineItem_TastyDessert struct {
	TastyDessert *flavortown_dessert.Dessert `protobuf:"bytes,101,opt,name=tasty_dessert,json=tastyDessert,oneof"`
}

func (*LineItem_Lunch) isLineItem_Dish()        {}
func (*LineItem_TastyDessert) isLineItem_Dish() {}

func (m *LineItem) GetDish() isLineItem_Dish {
	if m != nil {
		return m.Dish
	}
	return nil
}

func (m *LineItem) GetLunch() *Lunch {
	if x, ok := m.GetDish().(*LineItem_Lunch); ok {
		return x.Lunch
	}
	return nil
}

func (m *LineItem) GetTastyDessert() *flavortown_dessert.Dessert {
	if x, ok := m.GetDish().(*LineItem_TastyDessert); ok {
		return x.TastyDessert
	}
	return nil
}

func (m *LineItem) GetCreatedAt() *opsee_types.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *LineItem) GetUpdatedAt() *opsee_types.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *LineItem) GetNothing() *Nothing {
	if m != nil {
		return m.Nothing
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LineItem) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LineItem_OneofMarshaler, _LineItem_OneofUnmarshaler, _LineItem_OneofSizer, []interface{}{
		(*LineItem_Lunch)(nil),
		(*LineItem_TastyDessert)(nil),
	}
}

func _LineItem_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LineItem)
	// dish
	switch x := m.Dish.(type) {
	case *LineItem_Lunch:
		_ = b.EncodeVarint(100<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Lunch); err != nil {
			return err
		}
	case *LineItem_TastyDessert:
		_ = b.EncodeVarint(101<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TastyDessert); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LineItem.Dish has unexpected type %T", x)
	}
	return nil
}

func _LineItem_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LineItem)
	switch tag {
	case 100: // dish.lunch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Lunch)
		err := b.DecodeMessage(msg)
		m.Dish = &LineItem_Lunch{msg}
		return true, err
	case 101: // dish.tasty_dessert
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(flavortown_dessert.Dessert)
		err := b.DecodeMessage(msg)
		m.Dish = &LineItem_TastyDessert{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LineItem_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LineItem)
	// dish
	switch x := m.Dish.(type) {
	case *LineItem_Lunch:
		s := proto.Size(x.Lunch)
		n += proto.SizeVarint(100<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LineItem_TastyDessert:
		s := proto.Size(x.TastyDessert)
		n += proto.SizeVarint(101<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A delicious lunch dish on the menu
type Lunch struct {
	// The name of the dish
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The description of the dish
	Description []byte `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *Lunch) Reset()                    { *m = Lunch{} }
func (m *Lunch) String() string            { return proto.CompactTextString(m) }
func (*Lunch) ProtoMessage()               {}
func (*Lunch) Descriptor() ([]byte, []int) { return fileDescriptorFlavortown, []int{2} }

// confusion
type Nothing struct {
	// the void
	Void string `protobuf:"bytes,1,opt,name=void,proto3" json:"void,omitempty"`
}

func (m *Nothing) Reset()                    { *m = Nothing{} }
func (m *Nothing) String() string            { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()               {}
func (*Nothing) Descriptor() ([]byte, []int) { return fileDescriptorFlavortown, []int{3} }

func init() {
	proto.RegisterType((*Menu)(nil), "flavortown.Menu")
	proto.RegisterType((*LineItem)(nil), "flavortown.LineItem")
	proto.RegisterType((*Lunch)(nil), "flavortown.Lunch")
	proto.RegisterType((*Nothing)(nil), "flavortown.Nothing")
}
func (this *Menu) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Menu)
	if !ok {
		that2, ok := that.(Menu)
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
	if len(this.Items) != len(that1.Items) {
		return false
	}
	for i := range this.Items {
		if !this.Items[i].Equal(that1.Items[i]) {
			return false
		}
	}
	return true
}
func (this *LineItem) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*LineItem)
	if !ok {
		that2, ok := that.(LineItem)
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
	if that1.Dish == nil {
		if this.Dish != nil {
			return false
		}
	} else if this.Dish == nil {
		return false
	} else if !this.Dish.Equal(that1.Dish) {
		return false
	}
	if this.PriceCents != that1.PriceCents {
		return false
	}
	if !this.CreatedAt.Equal(that1.CreatedAt) {
		return false
	}
	if !this.UpdatedAt.Equal(that1.UpdatedAt) {
		return false
	}
	if !this.Nothing.Equal(that1.Nothing) {
		return false
	}
	return true
}
func (this *LineItem_Lunch) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*LineItem_Lunch)
	if !ok {
		that2, ok := that.(LineItem_Lunch)
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
	if !this.Lunch.Equal(that1.Lunch) {
		return false
	}
	return true
}
func (this *LineItem_TastyDessert) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*LineItem_TastyDessert)
	if !ok {
		that2, ok := that.(LineItem_TastyDessert)
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
	if !this.TastyDessert.Equal(that1.TastyDessert) {
		return false
	}
	return true
}
func (this *Lunch) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Lunch)
	if !ok {
		that2, ok := that.(Lunch)
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
	if !bytes.Equal(this.Description, that1.Description) {
		return false
	}
	return true
}
func (this *Nothing) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Nothing)
	if !ok {
		that2, ok := that.(Nothing)
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
	if this.Void != that1.Void {
		return false
	}
	return true
}

type MenuGetter interface {
	GetMenu() *Menu
}

var GraphQLMenuType *github_com_graphql_go_graphql.Object

type LineItemGetter interface {
	GetLineItem() *LineItem
}

var GraphQLLineItemType *github_com_graphql_go_graphql.Object
var GraphQLLineItemDishUnion *github_com_graphql_go_graphql.Union

type LunchGetter interface {
	GetLunch() *Lunch
}

var GraphQLLunchType *github_com_graphql_go_graphql.Object

type NothingGetter interface {
	GetNothing() *Nothing
}

var GraphQLNothingType *github_com_graphql_go_graphql.Object

func (g *LineItem_Lunch) GetLunch() *Lunch {
	return g.Lunch
}
func (g *LineItem_TastyDessert) GetDessert() *flavortown_dessert.Dessert {
	return g.TastyDessert
}

func init() {
	GraphQLMenuType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "flavortownMenu",
		Description: "The menu at Guy’s American Kitchen & Bar reflects his signature style of authentic and surprising flavors",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"items": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.NewList(GraphQLLineItemType),
					Description: "These dishes are crafted with the heart and soul of hometown favorites and infused with Guy’s big, daring flavors",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Menu)
						if ok {
							return obj.Items, nil
						}
						inter, ok := p.Source.(MenuGetter)
						if ok {
							face := inter.GetMenu()
							if face == nil {
								return nil, nil
							}
							return face.Items, nil
						}
						return nil, fmt.Errorf("field items not resolved")
					},
				},
			}
		}),
	})
	GraphQLLineItemType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "flavortownLineItem",
		Description: "A line item representing a dish and price",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"price_cents": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Int,
					Description: "The price of the dish in cents",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*LineItem)
						if ok {
							return obj.PriceCents, nil
						}
						inter, ok := p.Source.(LineItemGetter)
						if ok {
							face := inter.GetLineItem()
							if face == nil {
								return nil, nil
							}
							return face.PriceCents, nil
						}
						return nil, fmt.Errorf("field price_cents not resolved")
					},
				},
				"created_at": &github_com_graphql_go_graphql.Field{
					Type:        github_com_opsee_protobuf_plugin_graphql_scalars.Timestamp,
					Description: "A timestamp representing when the dish was added to the menu",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*LineItem)
						if ok {
							if obj.CreatedAt == nil {
								return nil, nil
							}
							return obj.GetCreatedAt(), nil
						}
						inter, ok := p.Source.(LineItemGetter)
						if ok {
							face := inter.GetLineItem()
							if face == nil {
								return nil, nil
							}
							if face.CreatedAt == nil {
								return nil, nil
							}
							return face.GetCreatedAt(), nil
						}
						return nil, fmt.Errorf("field created_at not resolved")
					},
				},
				"updated_at": &github_com_graphql_go_graphql.Field{
					Type:        github_com_opsee_protobuf_plugin_graphql_scalars.Timestamp,
					Description: "A timestamp representing when the dish was updated",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*LineItem)
						if ok {
							if obj.UpdatedAt == nil {
								return nil, nil
							}
							return obj.GetUpdatedAt(), nil
						}
						inter, ok := p.Source.(LineItemGetter)
						if ok {
							face := inter.GetLineItem()
							if face == nil {
								return nil, nil
							}
							if face.UpdatedAt == nil {
								return nil, nil
							}
							return face.GetUpdatedAt(), nil
						}
						return nil, fmt.Errorf("field updated_at not resolved")
					},
				},
				"nothing": &github_com_graphql_go_graphql.Field{
					Type:        GraphQLNothingType,
					Description: "A list of nothing really",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*LineItem)
						if ok {
							if obj.Nothing == nil {
								return nil, nil
							}
							return obj.GetNothing(), nil
						}
						inter, ok := p.Source.(LineItemGetter)
						if ok {
							face := inter.GetLineItem()
							if face == nil {
								return nil, nil
							}
							if face.Nothing == nil {
								return nil, nil
							}
							return face.GetNothing(), nil
						}
						return nil, fmt.Errorf("field nothing not resolved")
					},
				},
				"dish": &github_com_graphql_go_graphql.Field{
					Type:        GraphQLLineItemDishUnion,
					Description: "The menu dish, can either be lunch or dessert",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*LineItem)
						if !ok {
							return nil, fmt.Errorf("field dish not resolved")
						}
						return obj.GetDish(), nil
					},
				},
			}
		}),
	})
	GraphQLLunchType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "flavortownLunch",
		Description: "A delicious lunch dish on the menu",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"name": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.String,
					Description: "The name of the dish",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Lunch)
						if ok {
							return obj.Name, nil
						}
						inter, ok := p.Source.(LunchGetter)
						if ok {
							face := inter.GetLunch()
							if face == nil {
								return nil, nil
							}
							return face.Name, nil
						}
						return nil, fmt.Errorf("field name not resolved")
					},
				},
				"description": &github_com_graphql_go_graphql.Field{
					Type:        github_com_opsee_protobuf_plugin_graphql_scalars.ByteString,
					Description: "The description of the dish",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Lunch)
						if ok {
							return obj.Description, nil
						}
						inter, ok := p.Source.(LunchGetter)
						if ok {
							face := inter.GetLunch()
							if face == nil {
								return nil, nil
							}
							return face.Description, nil
						}
						return nil, fmt.Errorf("field description not resolved")
					},
				},
			}
		}),
	})
	GraphQLNothingType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "flavortownNothing",
		Description: "confusion",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"void": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.String,
					Description: "the void",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Nothing)
						if ok {
							return obj.Void, nil
						}
						inter, ok := p.Source.(NothingGetter)
						if ok {
							face := inter.GetNothing()
							if face == nil {
								return nil, nil
							}
							return face.Void, nil
						}
						return nil, fmt.Errorf("field void not resolved")
					},
				},
			}
		}),
	})
	GraphQLLineItemDishUnion = github_com_graphql_go_graphql.NewUnion(github_com_graphql_go_graphql.UnionConfig{
		Name:        "LineItemDish",
		Description: "The menu dish, can either be lunch or dessert",
		Types: []*github_com_graphql_go_graphql.Object{
			GraphQLLunchType,
			flavortown_dessert.GraphQLDessertType,
		},
		ResolveType: func(value interface{}, info github_com_graphql_go_graphql.ResolveInfo) *github_com_graphql_go_graphql.Object {
			switch value.(type) {
			case *LineItem_Lunch:
				return GraphQLLunchType
			case *LineItem_TastyDessert:
				return flavortown_dessert.GraphQLDessertType
			}
			return nil
		},
	})
}
func NewPopulatedMenu(r randyFlavortown, easy bool) *Menu {
	this := &Menu{}
	if r.Intn(10) != 0 {
		v1 := r.Intn(5)
		this.Items = make([]*LineItem, v1)
		for i := 0; i < v1; i++ {
			this.Items[i] = NewPopulatedLineItem(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedLineItem(r randyFlavortown, easy bool) *LineItem {
	this := &LineItem{}
	oneofNumber_Dish := []int32{100, 101}[r.Intn(2)]
	switch oneofNumber_Dish {
	case 100:
		this.Dish = NewPopulatedLineItem_Lunch(r, easy)
	case 101:
		this.Dish = NewPopulatedLineItem_TastyDessert(r, easy)
	}
	this.PriceCents = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.PriceCents *= -1
	}
	if r.Intn(10) != 0 {
		this.CreatedAt = opsee_types.NewPopulatedTimestamp(r, easy)
	}
	if r.Intn(10) != 0 {
		this.UpdatedAt = opsee_types.NewPopulatedTimestamp(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Nothing = NewPopulatedNothing(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedLineItem_Lunch(r randyFlavortown, easy bool) *LineItem_Lunch {
	this := &LineItem_Lunch{}
	this.Lunch = NewPopulatedLunch(r, easy)
	return this
}
func NewPopulatedLineItem_TastyDessert(r randyFlavortown, easy bool) *LineItem_TastyDessert {
	this := &LineItem_TastyDessert{}
	this.TastyDessert = flavortown_dessert.NewPopulatedDessert(r, easy)
	return this
}
func NewPopulatedLunch(r randyFlavortown, easy bool) *Lunch {
	this := &Lunch{}
	this.Name = randStringFlavortown(r)
	v2 := r.Intn(100)
	this.Description = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.Description[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedNothing(r randyFlavortown, easy bool) *Nothing {
	this := &Nothing{}
	this.Void = randStringFlavortown(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyFlavortown interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneFlavortown(r randyFlavortown) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringFlavortown(r randyFlavortown) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneFlavortown(r)
	}
	return string(tmps)
}
func randUnrecognizedFlavortown(r randyFlavortown, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldFlavortown(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldFlavortown(data []byte, r randyFlavortown, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateFlavortown(data, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		data = encodeVarintPopulateFlavortown(data, uint64(v4))
	case 1:
		data = encodeVarintPopulateFlavortown(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateFlavortown(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateFlavortown(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateFlavortown(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateFlavortown(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}

var fileDescriptorFlavortown = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xc7, 0xc9, 0x6e, 0xb2, 0xcb, 0x4e, 0x17, 0x69, 0x31, 0x08, 0x45, 0x45, 0xb4, 0x55, 0x4e,
	0x05, 0xa9, 0x09, 0x2a, 0x42, 0x42, 0x48, 0x1c, 0x28, 0x1c, 0x40, 0x02, 0x0e, 0x11, 0xf7, 0x2a,
	0x4d, 0xa6, 0x89, 0xa5, 0xc6, 0x8e, 0xe2, 0x49, 0xa1, 0xaf, 0xc3, 0x89, 0x47, 0xe0, 0x06, 0xaf,
	0x02, 0x4f, 0xc1, 0x11, 0xc7, 0x4e, 0xd4, 0x70, 0x00, 0xed, 0xc1, 0xf2, 0x7c, 0xfc, 0x7f, 0xe3,
	0xf1, 0x0c, 0x5c, 0x6d, 0x77, 0xc9, 0x5e, 0xd6, 0x24, 0x3f, 0x89, 0xb0, 0xaa, 0x25, 0x49, 0x06,
	0xc7, 0xc8, 0xf8, 0x79, 0xce, 0xa9, 0x68, 0x36, 0x61, 0x2a, 0xcb, 0x48, 0x56, 0x0a, 0x31, 0x32,
	0x9a, 0x4d, 0xb3, 0xb5, 0xae, 0xf1, 0x22, 0x3a, 0x54, 0xa8, 0x22, 0xe2, 0x25, 0x2a, 0x4a, 0xca,
	0xca, 0xd6, 0x19, 0x2f, 0x06, 0x6c, 0x2e, 0x73, 0x79, 0x44, 0x5b, 0xcf, 0x92, 0xad, 0xd5, 0xc9,
	0x1f, 0x5f, 0xeb, 0x29, 0x63, 0x76, 0xc4, 0xb3, 0x7f, 0x13, 0xf8, 0x59, 0xf7, 0xb1, 0xd3, 0x4d,
	0x65, 0xa8, 0x14, 0xd6, 0xd4, 0xdf, 0x96, 0x0c, 0x96, 0xe0, 0xbe, 0x47, 0xd1, 0xb0, 0x47, 0xe0,
	0x71, 0xc2, 0x52, 0xf9, 0xce, 0xec, 0x74, 0x3e, 0x5a, 0xde, 0x0d, 0x07, 0xc3, 0x78, 0xc7, 0x05,
	0xbe, 0xd5, 0xc9, 0xd8, 0x4a, 0x82, 0xef, 0x27, 0x70, 0xb3, 0x8f, 0xb1, 0x87, 0xe0, 0xed, 0x1a,
	0x91, 0x16, 0x7e, 0x36, 0x73, 0x34, 0x78, 0xfb, 0x2f, 0xb0, 0x4d, 0xbc, 0xb9, 0x11, 0x5b, 0x05,
	0x5b, 0xc1, 0x2d, 0x4a, 0x14, 0x1d, 0xd6, 0x5d, 0x0b, 0x3e, 0x1a, 0xe4, 0xfe, 0x10, 0xe9, 0xbb,
	0x7b, 0x6d, 0x6f, 0x0d, 0x5f, 0x1a, 0xa6, 0xf3, 0xd9, 0x14, 0x46, 0x55, 0xcd, 0x53, 0x5c, 0xa7,
	0x28, 0x48, 0xf9, 0x27, 0xba, 0x82, 0x17, 0x83, 0x09, 0xbd, 0x6a, 0x23, 0xec, 0x29, 0x40, 0x5a,
	0x63, 0x42, 0x98, 0xad, 0x13, 0xf2, 0x4f, 0xcd, 0x0b, 0xf7, 0x42, 0x3b, 0x2c, 0xb3, 0x9d, 0xf0,
	0x63, 0xbf, 0x9d, 0xf8, 0xa2, 0x53, 0xbe, 0xa4, 0x16, 0x6b, 0xaa, 0xac, 0xc7, 0xdc, 0xff, 0x63,
	0x9d, 0x52, 0x63, 0x0b, 0x38, 0x17, 0x92, 0x0a, 0x2e, 0x72, 0xdf, 0x33, 0xcc, 0x9d, 0xe1, 0x67,
	0x3e, 0xd8, 0x54, 0xdc, 0x6b, 0x56, 0x67, 0xe0, 0x66, 0x5c, 0x15, 0xc1, 0x0b, 0xf0, 0xcc, 0x6c,
	0x18, 0x03, 0x57, 0x24, 0x25, 0xea, 0xa9, 0x3b, 0xf3, 0x8b, 0xd8, 0xd8, 0x6c, 0x06, 0x23, 0x3d,
	0x85, 0xb4, 0xe6, 0x15, 0x71, 0x29, 0xcc, 0x17, 0x2f, 0xe3, 0x61, 0x28, 0x78, 0x00, 0xe7, 0x5d,
	0xe9, 0xb6, 0xc0, 0x5e, 0xf2, 0xac, 0x2f, 0xd0, 0xda, 0xab, 0xab, 0xdf, 0x3f, 0x27, 0xce, 0xd7,
	0x5f, 0x13, 0xe7, 0x9b, 0x3e, 0x3f, 0xbe, 0x4c, 0x9d, 0xcd, 0x99, 0x59, 0xf6, 0x93, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xdd, 0xe8, 0x23, 0xcb, 0xe3, 0x02, 0x00, 0x00,
}
