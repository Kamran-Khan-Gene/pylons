// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pylons/pylons/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the pylons module's genesis state.
type GenesisState struct {
	// this line is used by starport scaffolding # genesis/proto/state
	RedeemInfoList               []RedeemInfo               `protobuf:"bytes,16,rep,name=redeem_info_list,json=redeemInfoList,proto3" json:"redeem_info_list"`
	PaymentInfoList              []PaymentInfo              `protobuf:"bytes,15,rep,name=payment_info_list,json=paymentInfoList,proto3" json:"payment_info_list"`
	AccountList                  []UserMap                  `protobuf:"bytes,14,rep,name=account_list,json=accountList,proto3" json:"account_list"`
	TradeList                    []Trade                    `protobuf:"bytes,13,rep,name=trade_list,json=tradeList,proto3" json:"trade_list"`
	TradeCount                   uint64                     `protobuf:"varint,12,opt,name=trade_count,json=tradeCount,proto3" json:"trade_count,omitempty"`
	EntityCount                  uint64                     `protobuf:"varint,11,opt,name=entity_count,json=entityCount,proto3" json:"entity_count,omitempty"`
	Params                       Params                     `protobuf:"bytes,10,opt,name=params,proto3" json:"params"`
	GoogleInAppPurchaseOrderList []GoogleInAppPurchaseOrder `protobuf:"bytes,8,rep,name=google_in_app_purchase_order_list,json=googleInAppPurchaseOrderList,proto3" json:"google_in_app_purchase_order_list"`
	GoogleIapOrderCount          uint64                     `protobuf:"varint,9,opt,name=google_iap_order_count,json=googleIapOrderCount,proto3" json:"google_iap_order_count,omitempty"`
	ExecutionList                []Execution                `protobuf:"bytes,7,rep,name=execution_list,json=executionList,proto3" json:"execution_list"`
	ExecutionCount               uint64                     `protobuf:"varint,6,opt,name=execution_count,json=executionCount,proto3" json:"execution_count,omitempty"`
	PendingExecutionList         []Execution                `protobuf:"bytes,5,rep,name=pending_execution_list,json=pendingExecutionList,proto3" json:"pending_execution_list"`
	PendingExecutionCount        uint64                     `protobuf:"varint,4,opt,name=pending_execution_count,json=pendingExecutionCount,proto3" json:"pending_execution_count,omitempty"`
	ItemList                     []Item                     `protobuf:"bytes,3,rep,name=item_list,json=itemList,proto3" json:"item_list"`
	RecipeList                   []Recipe                   `protobuf:"bytes,2,rep,name=recipe_list,json=recipeList,proto3" json:"recipe_list"`
	CookbookList                 []Cookbook                 `protobuf:"bytes,1,rep,name=cookbook_list,json=cookbookList,proto3" json:"cookbook_list"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_f41fd395b1a4953e, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetRedeemInfoList() []RedeemInfo {
	if m != nil {
		return m.RedeemInfoList
	}
	return nil
}

func (m *GenesisState) GetPaymentInfoList() []PaymentInfo {
	if m != nil {
		return m.PaymentInfoList
	}
	return nil
}

func (m *GenesisState) GetAccountList() []UserMap {
	if m != nil {
		return m.AccountList
	}
	return nil
}

func (m *GenesisState) GetTradeList() []Trade {
	if m != nil {
		return m.TradeList
	}
	return nil
}

func (m *GenesisState) GetTradeCount() uint64 {
	if m != nil {
		return m.TradeCount
	}
	return 0
}

func (m *GenesisState) GetEntityCount() uint64 {
	if m != nil {
		return m.EntityCount
	}
	return 0
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetGoogleInAppPurchaseOrderList() []GoogleInAppPurchaseOrder {
	if m != nil {
		return m.GoogleInAppPurchaseOrderList
	}
	return nil
}

func (m *GenesisState) GetGoogleIapOrderCount() uint64 {
	if m != nil {
		return m.GoogleIapOrderCount
	}
	return 0
}

func (m *GenesisState) GetExecutionList() []Execution {
	if m != nil {
		return m.ExecutionList
	}
	return nil
}

func (m *GenesisState) GetExecutionCount() uint64 {
	if m != nil {
		return m.ExecutionCount
	}
	return 0
}

func (m *GenesisState) GetPendingExecutionList() []Execution {
	if m != nil {
		return m.PendingExecutionList
	}
	return nil
}

func (m *GenesisState) GetPendingExecutionCount() uint64 {
	if m != nil {
		return m.PendingExecutionCount
	}
	return 0
}

func (m *GenesisState) GetItemList() []Item {
	if m != nil {
		return m.ItemList
	}
	return nil
}

func (m *GenesisState) GetRecipeList() []Recipe {
	if m != nil {
		return m.RecipeList
	}
	return nil
}

func (m *GenesisState) GetCookbookList() []Cookbook {
	if m != nil {
		return m.CookbookList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "pylons.pylons.GenesisState")
}

func init() { proto.RegisterFile("pylons/pylons/genesis.proto", fileDescriptor_f41fd395b1a4953e) }

var fileDescriptor_f41fd395b1a4953e = []byte{
	// 630 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x4f, 0x6f, 0xda, 0x4c,
	0x10, 0xc6, 0xf1, 0x9b, 0xbc, 0x69, 0xb2, 0x86, 0x24, 0x25, 0x40, 0x08, 0x4d, 0x0d, 0xa9, 0x2a,
	0x85, 0x43, 0x0b, 0x52, 0x90, 0x22, 0x55, 0xaa, 0x54, 0x95, 0x88, 0x46, 0x48, 0xa9, 0x8a, 0x68,
	0x7a, 0xe9, 0xc5, 0x32, 0x66, 0x63, 0xac, 0xe0, 0xdd, 0x95, 0xbd, 0x54, 0xe1, 0xd8, 0x6f, 0xd0,
	0x8f, 0x95, 0x63, 0x8e, 0x3d, 0x55, 0x15, 0x7c, 0x91, 0xca, 0x33, 0x6b, 0xc0, 0x4b, 0x2a, 0xf5,
	0x64, 0x6b, 0xe6, 0xf9, 0x3d, 0x33, 0xb3, 0xff, 0xc8, 0x33, 0x31, 0x1d, 0x73, 0x16, 0x35, 0xd5,
	0xc7, 0xa3, 0x8c, 0x46, 0x7e, 0xd4, 0x10, 0x21, 0x97, 0x3c, 0x9f, 0xc3, 0x68, 0x03, 0x3f, 0x95,
	0x6a, 0x5a, 0x1b, 0xd2, 0x21, 0xa5, 0x81, 0xed, 0xb3, 0x1b, 0x8e, 0xfa, 0x4a, 0x2d, 0x2d, 0x10,
	0xce, 0x34, 0xa0, 0x4c, 0xae, 0x2a, 0x8e, 0xd3, 0x0a, 0xc7, 0x75, 0xf9, 0x84, 0x49, 0x55, 0xaf,
	0x72, 0x94, 0xce, 0xca, 0xd0, 0x19, 0x52, 0x95, 0x7a, 0xa9, 0xf5, 0xc9, 0xb9, 0x37, 0xa6, 0xb6,
	0xef, 0x08, 0x9b, 0x87, 0x43, 0x1a, 0x2a, 0xd5, 0xf3, 0xb4, 0x8a, 0xde, 0x51, 0x77, 0x22, 0x7d,
	0xce, 0x54, 0xba, 0x9c, 0x4e, 0xfb, 0x92, 0x06, 0x2a, 0x53, 0xd1, 0x47, 0x73, 0x7d, 0x41, 0x1f,
	0xef, 0xd9, 0xe5, 0xfc, 0x76, 0xc0, 0xf9, 0xed, 0xe3, 0xa4, 0x70, 0x42, 0x27, 0x48, 0xe6, 0x29,
	0x78, 0xdc, 0xe3, 0xf0, 0xdb, 0x8c, 0xff, 0x30, 0xfa, 0xe2, 0xfb, 0x36, 0xc9, 0x5e, 0xe2, 0x3a,
	0x7f, 0x96, 0x8e, 0xa4, 0xf9, 0x2e, 0xd9, 0x5f, 0x59, 0x4b, 0x7b, 0xec, 0x47, 0xb2, 0xbc, 0x5f,
	0xdb, 0xa8, 0x9b, 0x67, 0x47, 0x8d, 0xd4, 0x0e, 0x34, 0xfa, 0x20, 0xeb, 0xb2, 0x1b, 0xde, 0xde,
	0xbc, 0xff, 0x55, 0xcd, 0xf4, 0x77, 0xc3, 0x45, 0xe4, 0xca, 0x8f, 0x64, 0xfe, 0x8a, 0x3c, 0x5d,
	0x5d, 0x75, 0xf4, 0xda, 0x03, 0xaf, 0x8a, 0xe6, 0xd5, 0x43, 0xdd, 0x8a, 0xd9, 0x9e, 0x58, 0x86,
	0xc0, 0xed, 0x1d, 0xc9, 0xaa, 0x1d, 0x42, 0xa3, 0x5d, 0x30, 0x2a, 0x69, 0x46, 0x5f, 0x22, 0x1a,
	0x7e, 0x74, 0x84, 0x32, 0x31, 0x15, 0x01, 0x06, 0x6f, 0x08, 0x81, 0x4d, 0x44, 0x3c, 0x07, 0x78,
	0x41, 0xc3, 0xaf, 0x63, 0x81, 0x82, 0x77, 0x40, 0x0d, 0x68, 0x95, 0x98, 0x88, 0x82, 0x5b, 0x39,
	0x5b, 0x33, 0xea, 0x9b, 0x7d, 0x74, 0xbb, 0x88, 0x23, 0xf9, 0x13, 0x92, 0xa5, 0x4c, 0xfa, 0x72,
	0xaa, 0x14, 0x26, 0x28, 0x4c, 0x8c, 0xa1, 0xa4, 0x45, 0xb6, 0x70, 0x3f, 0xca, 0xa4, 0x66, 0xd4,
	0xcd, 0xb3, 0xe2, 0xda, 0x12, 0xc4, 0x49, 0x55, 0x5b, 0x49, 0xf3, 0xdf, 0xc8, 0x49, 0x72, 0xba,
	0x98, 0xed, 0x08, 0x61, 0x8b, 0x49, 0xe8, 0x8e, 0x9c, 0x88, 0xe2, 0x49, 0xc3, 0x51, 0xb6, 0x61,
	0x94, 0x53, 0xcd, 0xef, 0x12, 0xb8, 0x2e, 0x7b, 0x2f, 0x44, 0x4f, 0x41, 0x9f, 0x62, 0x46, 0x55,
	0x38, 0xf6, 0xfe, 0x92, 0x87, 0x81, 0x5b, 0xa4, 0xa4, 0x9f, 0x6a, 0x35, 0xd9, 0x0e, 0x4c, 0x76,
	0xa0, 0x68, 0x47, 0x00, 0x83, 0x13, 0x76, 0xc8, 0xee, 0xe2, 0x90, 0x63, 0x67, 0x4f, 0xa0, 0xb3,
	0xb2, 0xd6, 0x59, 0x27, 0x11, 0xa9, 0x56, 0x72, 0x0b, 0x0a, 0x6a, 0x9f, 0x92, 0xbd, 0xa5, 0x0d,
	0x16, 0xdd, 0x82, 0xa2, 0x4b, 0x77, 0xac, 0x77, 0x4d, 0x4a, 0x82, 0xb2, 0xa1, 0xcf, 0x3c, 0x5b,
	0xab, 0xfb, 0xff, 0x3f, 0xd5, 0x2d, 0x28, 0xba, 0x93, 0x2a, 0x7f, 0x4e, 0x0e, 0xd7, 0x5d, 0xb1,
	0x8d, 0x4d, 0x68, 0xa3, 0xa8, 0x63, 0xd8, 0xcd, 0x39, 0xd9, 0x89, 0xef, 0x30, 0x36, 0xb0, 0x01,
	0x0d, 0x1c, 0x68, 0x0d, 0x74, 0x25, 0x0d, 0x54, 0xed, 0xed, 0x58, 0x0b, 0xf5, 0xde, 0x12, 0x13,
	0x6f, 0x38, 0x92, 0xff, 0x01, 0x59, 0x5c, 0xbb, 0x6b, 0xb1, 0x42, 0xb1, 0x04, 0xf5, 0x40, 0xb7,
	0x49, 0x2e, 0x79, 0x03, 0x90, 0x37, 0x80, 0x3f, 0xd4, 0xf8, 0x0b, 0xa5, 0x51, 0x0e, 0xd9, 0x84,
	0x89, 0x3d, 0xda, 0x1f, 0xee, 0x67, 0x96, 0xf1, 0x30, 0xb3, 0x8c, 0xdf, 0x33, 0xcb, 0xf8, 0x31,
	0xb7, 0x32, 0x0f, 0x73, 0x2b, 0xf3, 0x73, 0x6e, 0x65, 0xbe, 0xbe, 0xf2, 0x7c, 0x39, 0x9a, 0x0c,
	0x1a, 0x2e, 0x0f, 0x9a, 0x3d, 0x70, 0x7a, 0x2d, 0xa9, 0x3b, 0x4a, 0xde, 0x97, 0xbb, 0xc5, 0xe3,
	0x38, 0x15, 0x34, 0x1a, 0x6c, 0xc1, 0x93, 0xd2, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0x48, 0x5b,
	0xff, 0x37, 0xc7, 0x05, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RedeemInfoList) > 0 {
		for iNdEx := len(m.RedeemInfoList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RedeemInfoList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x82
		}
	}
	if len(m.PaymentInfoList) > 0 {
		for iNdEx := len(m.PaymentInfoList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PaymentInfoList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x7a
		}
	}
	if len(m.AccountList) > 0 {
		for iNdEx := len(m.AccountList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccountList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x72
		}
	}
	if len(m.TradeList) > 0 {
		for iNdEx := len(m.TradeList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TradeList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6a
		}
	}
	if m.TradeCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TradeCount))
		i--
		dAtA[i] = 0x60
	}
	if m.EntityCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EntityCount))
		i--
		dAtA[i] = 0x58
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if m.GoogleIapOrderCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.GoogleIapOrderCount))
		i--
		dAtA[i] = 0x48
	}
	if len(m.GoogleInAppPurchaseOrderList) > 0 {
		for iNdEx := len(m.GoogleInAppPurchaseOrderList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GoogleInAppPurchaseOrderList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.ExecutionList) > 0 {
		for iNdEx := len(m.ExecutionList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExecutionList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.ExecutionCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ExecutionCount))
		i--
		dAtA[i] = 0x30
	}
	if len(m.PendingExecutionList) > 0 {
		for iNdEx := len(m.PendingExecutionList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PendingExecutionList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.PendingExecutionCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PendingExecutionCount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ItemList) > 0 {
		for iNdEx := len(m.ItemList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ItemList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.RecipeList) > 0 {
		for iNdEx := len(m.RecipeList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RecipeList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.CookbookList) > 0 {
		for iNdEx := len(m.CookbookList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CookbookList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CookbookList) > 0 {
		for _, e := range m.CookbookList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RecipeList) > 0 {
		for _, e := range m.RecipeList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ItemList) > 0 {
		for _, e := range m.ItemList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.PendingExecutionCount != 0 {
		n += 1 + sovGenesis(uint64(m.PendingExecutionCount))
	}
	if len(m.PendingExecutionList) > 0 {
		for _, e := range m.PendingExecutionList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ExecutionCount != 0 {
		n += 1 + sovGenesis(uint64(m.ExecutionCount))
	}
	if len(m.ExecutionList) > 0 {
		for _, e := range m.ExecutionList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GoogleInAppPurchaseOrderList) > 0 {
		for _, e := range m.GoogleInAppPurchaseOrderList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.GoogleIapOrderCount != 0 {
		n += 1 + sovGenesis(uint64(m.GoogleIapOrderCount))
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.EntityCount != 0 {
		n += 1 + sovGenesis(uint64(m.EntityCount))
	}
	if m.TradeCount != 0 {
		n += 1 + sovGenesis(uint64(m.TradeCount))
	}
	if len(m.TradeList) > 0 {
		for _, e := range m.TradeList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AccountList) > 0 {
		for _, e := range m.AccountList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PaymentInfoList) > 0 {
		for _, e := range m.PaymentInfoList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RedeemInfoList) > 0 {
		for _, e := range m.RedeemInfoList {
			l = e.Size()
			n += 2 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CookbookList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CookbookList = append(m.CookbookList, Cookbook{})
			if err := m.CookbookList[len(m.CookbookList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipeList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipeList = append(m.RecipeList, Recipe{})
			if err := m.RecipeList[len(m.RecipeList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ItemList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ItemList = append(m.ItemList, Item{})
			if err := m.ItemList[len(m.ItemList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingExecutionCount", wireType)
			}
			m.PendingExecutionCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PendingExecutionCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingExecutionList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PendingExecutionList = append(m.PendingExecutionList, Execution{})
			if err := m.PendingExecutionList[len(m.PendingExecutionList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionCount", wireType)
			}
			m.ExecutionCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExecutionCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecutionList = append(m.ExecutionList, Execution{})
			if err := m.ExecutionList[len(m.ExecutionList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoogleInAppPurchaseOrderList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoogleInAppPurchaseOrderList = append(m.GoogleInAppPurchaseOrderList, GoogleInAppPurchaseOrder{})
			if err := m.GoogleInAppPurchaseOrderList[len(m.GoogleInAppPurchaseOrderList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoogleIapOrderCount", wireType)
			}
			m.GoogleIapOrderCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GoogleIapOrderCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityCount", wireType)
			}
			m.EntityCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EntityCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeCount", wireType)
			}
			m.TradeCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TradeCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TradeList = append(m.TradeList, Trade{})
			if err := m.TradeList[len(m.TradeList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountList = append(m.AccountList, UserMap{})
			if err := m.AccountList[len(m.AccountList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentInfoList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PaymentInfoList = append(m.PaymentInfoList, PaymentInfo{})
			if err := m.PaymentInfoList[len(m.PaymentInfoList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedeemInfoList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RedeemInfoList = append(m.RedeemInfoList, RedeemInfo{})
			if err := m.RedeemInfoList[len(m.RedeemInfoList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
