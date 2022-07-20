package models

import (
	"encoding"
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type DeliveryDateValue struct {
	LocationUid string
	WarehouseId int64
	SourceId    int32
	IsBulk      bool
}

func (*DeliveryDateValue) ProtoMessage() {}

func (v *DeliveryDateValue) ProtoReflect() protoreflect.Message {
	return nil
}

// DeliveryDateValue implemetns encoding.BinaryMarshaler
var _ encoding.BinaryMarshaler = (*DeliveryDateValue)(nil)

func (v *DeliveryDateValue) MarshalBinary() (data []byte, err error) {
	if v == nil {
		return nil, nil
	}
	bytes, err := proto.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (v DeliveryDateValue) Key() string {
	return fmt.Sprintf("%s-%d-%v-%d", v.LocationUid, v.WarehouseId, v.IsBulk, v.SourceId)
}
