package main

import (
	"encoding/json"
	"fmt"
	"wokerTrans/helps"

	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
)

//go:generate easytags $GOFILE json,xml
type TransactionType int

type MapTxHash2Trans struct {
	Txid      int64           `json:"txid" xml:"txid"`
	TxHash    string          `json:"tran_id" xml:"tran_id"`
	Coin      string          `json:"coin" xml:"coin"`
	Type      TransactionType `json:"type" xml:"type"`
	CreatedAt int64           `json:"created_at" xml:"created_at"`
}

func (this *MapTxHash2Trans) String() string {
	return this.TxHash
}

func (this *MapTxHash2Trans) GetBsKey() generic.TStringKey {
	return generic.TStringKey(fmt.Sprintf("%s:%s", BS_TX_HASH_2_TRANSACTION, this.Coin))
}

func (this *MapTxHash2Trans) GetAll() ([]MapTxHash2Trans, int64, error) {
	var err error
	if totalCount, err := GetBsTransService().GetTotalCount(this.GetBsKey()); totalCount > 0 && (err == nil || !helps.IsError(err)) {
		slice, err := GetBsTransService().BsGetSliceR(this.GetBsKey(), 0, int32(totalCount))
		if helps.IsError(err) {
			return make([]MapTxHash2Trans, 0), 0, err
		}
		mapTxHash2Tran, err := this.UnMarshalArrayTItem(slice)
		return mapTxHash2Tran, totalCount, err
	}

	return make([]MapTxHash2Trans, 0), 0, err
}

func (this *MapTxHash2Trans) GetPaginate(pos, count int32) ([]MapTxHash2Trans, int64, error) {
	totalCount, err := GetBsTransService().GetTotalCount(this.GetBsKey())
	if helps.IsError(err) || totalCount < 1 {
		return nil, 0, err
	}

	setItems, err := GetBsTransService().BsGetSlice(this.GetBsKey(), pos, count)
	if helps.IsError(err) {
		return nil, 0, err
	}

	mapTxHash2Trans, err := this.UnMarshalArrayTItem(setItems)
	if err != nil {
		return nil, 0, err
	}

	return mapTxHash2Trans, totalCount, err
}

func (this *MapTxHash2Trans) Delete() error {
	return GetBsTransService().BsRemoveItem(this.GetBsKey(), []byte(this.String()))
}

//func (this *MapTxHash2Trans) Get() (interface{}, error) {
//	bytes, err := this.GetItemBytes()
//	if err != nil {
//		return nil, err
//	}
//
//	return helps.UnMarshalBytes(bytes)
//}

func (this *MapTxHash2Trans) GetItemBytes() ([]byte, error) {
	tBlockChain2Address, err := GetBsTransService().BsGetItem(this.GetBsKey(), generic.TItemKey(this.String()))
	if helps.IsError(err) {
		return nil, err
	}

	return tBlockChain2Address.GetValue(), nil
}

func (this *MapTxHash2Trans) UnMarshalArrayTItem(objects []*generic.TItem) ([]MapTxHash2Trans, error) {
	objs := make([]MapTxHash2Trans, 0)

	for _, object := range objects {
		obj := MapTxHash2Trans{}
		err := json.Unmarshal(object.GetValue(), &obj)

		if err != nil {
			return make([]MapTxHash2Trans, 0), err
		}

		objs = append(objs, obj)
	}

	return objs, nil
}

func (this *MapTxHash2Trans) GetFromKey(key string) (*MapTxHash2Trans, error) {
	item, err := GetBsTransService().BsGetItem(this.GetBsKey(), generic.TItemKey(key))
	if helps.IsError(err) {
		return nil, err
	}
	obj := &MapTxHash2Trans{}
	err = json.Unmarshal(item.GetValue(), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
