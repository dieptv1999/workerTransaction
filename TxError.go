package main

import (
	"encoding/json"
	"fmt"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"log"
	"time"
	"wokerTrans/helps"
)

type TxError struct {
	BlockHash         string `json:"blockHash" xml:"blockHash"`
	BlockNumber       string `json:"blockNumber" xml:"blockNumber"`
	Confirmations     string `json:"confirmations" xml:"confirmations"`
	ContractAddress   string `json:"contractAddress" xml:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed" xml:"cumulativeGasUsed"`
	From              string `json:"from" xml:"from"`
	Gas               string `json:"gas" xml:"gas"`
	GasPrice          string `json:"gasPrice" xml:"gasPrice"`
	GasUsed           string `json:"gasUsed" xml:"gasUsed"`
	Hash              string `json:"hash" xml:"hash"`
	Input             string `json:"input" xml:"input"`
	IsError           string `json:"isError" xml:"isError"`
	Nonce             string `json:"nonce" xml:"nonce"`
	TimeStamp         string `json:"timeStamp" xml:"timeStamp"`
	To                string `json:"to" xml:"to"`
	TransactionIndex  string `json:"transactionIndex" xml:"transactionIndex"`
	Txreceipt_status  string `json:"txreceipt_status" xml:"txreceipt_status"`
	Value             string `json:"value" xml:"value"`
	CreatedAt         int64  `json:"created_at" xml:"created_at"`
	UpdatedAt         int64  `json:"updated_at" xml:"updated_at"`
}

func (this *TxError) String() string {
	return fmt.Sprintf("BlockHash:%s Hash:%s From:%s To:%s TimeStamp:%s", this.BlockHash, this.Hash, this.From, this.To, this.TimeStamp)
}

func (this *TxError) GetBsKey() generic.TStringKey {
	return generic.TStringKey(fmt.Sprintf("%s:%s", BS_TX_ERROR, coinName))
}

func (this *TxError) GetAll() ([]TxError, int64, error) {
	var err error
	if totalCount, err := GetBsTransService().GetTotalCount(this.GetBsKey()); totalCount > 0 && (err == nil || !helps.IsError(err)) {
		slice, err := GetBsTransService().BsGetSliceR(this.GetBsKey(), 0, int32(totalCount))
		if helps.IsError(err) {
			return make([]TxError, 0), 0, err
		}
		mapTxHash2Tran, err := this.UnMarshalArrayTItem(slice)
		return mapTxHash2Tran, totalCount, err
	}

	return make([]TxError, 0), 0, err
}

func (this *TxError) GetPaginate(pos, count int32) ([]TxError, int64, error) {
	totalCount, err := GetBsTransService().GetTotalCount(this.GetBsKey())
	if helps.IsError(err) || totalCount < 1 {
		return nil, 0, err
	}

	setItems, err := GetBsTransService().BsGetSlice(this.GetBsKey(), pos, count)
	if helps.IsError(err) {
		return nil, 0, err
	}

	TxError, err := this.UnMarshalArrayTItem(setItems)
	if err != nil {
		return nil, 0, err
	}

	return TxError, totalCount, err
}

func (this *TxError) Create() error {
	now := time.Now().Unix()
	this.CreatedAt = now
	this.UpdatedAt = now

	bTxError, key, err := helps.MarshalBytes(this)
	if err != nil {
		return err
	}

	log.Println(string(key), "-- string(key)")
	log.Println(this, "-- this")

	return GetBsOriginal().BsPutItem(this.GetBsKey(), &generic.TItem{
		Key:   key,
		Value: bTxError,
	})
}

func (this *TxError) PutItem() error {
	this.UpdateTime()

	bTxError, key, err := helps.MarshalBytes(this)

	if err != nil {
		return err
	}

	return GetBsOriginal().BsPutItem(this.GetBsKey(), &generic.TItem{
		Key:   key,
		Value: bTxError,
	})
}

func (this *TxError) Delete() error {
	return GetBsTransService().BsRemoveItem(this.GetBsKey(), []byte(this.String()))
}

func (this *TxError) Get() (interface{}, error) {
	bytes, err := this.GetItemBytes()
	if err != nil {
		return nil, err
	}

	return helps.UnMarshalBytes(bytes)
}

func (this *TxError) GetItemBytes() ([]byte, error) {
	tBlockChain2Address, err := GetBsTransService().BsGetItem(this.GetBsKey(), generic.TItemKey(this.String()))
	if helps.IsError(err) {
		return nil, err
	}

	return tBlockChain2Address.GetValue(), nil
}

func (this *TxError) UnMarshalArrayTItem(objects []*generic.TItem) ([]TxError, error) {
	objs := make([]TxError, 0)

	for _, object := range objects {
		obj := TxError{}
		err := json.Unmarshal(object.GetValue(), &obj)

		if err != nil {
			return make([]TxError, 0), err
		}

		objs = append(objs, obj)
	}

	return objs, nil
}

func (this *TxError) GetFromKey(key string) (*TxError, error) {
	item, err := GetBsTransService().BsGetItem(this.GetBsKey(), generic.TItemKey(key))
	if helps.IsError(err) {
		return nil, err
	}
	obj := &TxError{}
	err = json.Unmarshal(item.GetValue(), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (this *TxError) UpdateTime() {
	this.UpdatedAt = time.Now().Unix()
}
