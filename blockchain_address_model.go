package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
	"wokerTrans/helps"

	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
)

// lưu địa chỉ blockchain

//go:generate easytags $GOFILE json,xml

type BlockchainAddress struct {
	BlockchainAddress string `json:"blockchain_address" xml:"blockchain_address"`
	Coin              string `json:"coin" xml:"coin"`
	//BlockchainTag     string `json:"blockchain_tag" xml:"blockchain_tag"`
	//DeviceId uint64
	//Path      string `json:"path" xml:"path"`
	Available bool   `json:"available" xml:"available"`
	Type      string `json:"type" xml:"type"`
	CreatedAt int64  `json:"created_at" xml:"created_at"`
	UpdatedAt int64  `json:"updated_at" xml:"updated_at"`
}


func (this *BlockchainAddress) String() string {
	return strings.ToLower(this.BlockchainAddress)
}

func (this *BlockchainAddress) GetBsKey() generic.TStringKey {
	return generic.TStringKey(fmt.Sprintf("%s:%s:%t", BS_BLOCKCHAIN_ADDRESS, this.Coin, this.Available))
}

func (this *BlockchainAddress) GetAll() ([]BlockchainAddress, int64, error) {
	var err error

	allAvailable, countAvailable, err := this.GetAllAvailable()
	if err != nil && countAvailable == 0 {
		log.Println(err.Error(), "-- err.Error() models/1_blockchain_address_model.go:43")
		return make([]BlockchainAddress, 0), 0, err
	}
	allUnAvailable, countUnavailable, err := this.GetAllUnAvailable()
	if err != nil && countUnavailable == 0 {
		log.Println(err.Error(), "-- err.Error() models/1_blockchain_address_model.go:43")
		return make([]BlockchainAddress, 0), 0, err
	}

	if (countAvailable + countUnavailable) == 0 {
		return make([]BlockchainAddress, 0), 0, err
	}
	return append(allAvailable, allUnAvailable...), countAvailable + countUnavailable, nil
}

func (this *BlockchainAddress) GetAllAvailable() ([]BlockchainAddress, int64, error) {
	var err error
	addresses := &BlockchainAddress{
		Available: true,
		Coin:      this.Coin,
	}

	if totalCount, err := GetBsOriginal().GetTotalCount(addresses.GetBsKey()); totalCount > 0 && (err == nil || err != nil) {
		slice, err := GetBsOriginal().BsGetSliceR(addresses.GetBsKey(), 0, int32(totalCount))
		if err != nil {
			return make([]BlockchainAddress, 0), 0, err
		}
		BlockchainAddress, err := addresses.UnMarshalArrayTItem(slice)
		return BlockchainAddress, totalCount, err
	}

	return make([]BlockchainAddress, 0), 0, err
}

func (this *BlockchainAddress) GetAllUnAvailable() ([]BlockchainAddress, int64, error) {
	var err error
	addresses := &BlockchainAddress{
		Available: false,
		Coin:      this.Coin,
	}

	if totalCount, err := GetBsOriginal().GetTotalCount(addresses.GetBsKey()); totalCount > 0 && (err == nil || err != nil) {
		slice, err := GetBsOriginal().BsGetSliceR(addresses.GetBsKey(), 0, int32(totalCount))
		if err != nil {
			return make([]BlockchainAddress, 0), 0, err
		}
		BlockchainAddress, err := addresses.UnMarshalArrayTItem(slice)
		return BlockchainAddress, totalCount, err
	}

	return make([]BlockchainAddress, 0), 0, err
}

func (this *BlockchainAddress) GetPaginate(pos, count int32) ([]BlockchainAddress, int64, error) {
	totalCount, err := GetBsOriginal().GetTotalCount(this.GetBsKey())
	if err != nil || totalCount < 1 {
		return nil, 0, err
	}

	setItems, err := GetBsOriginal().BsGetSlice(this.GetBsKey(), pos, count)
	if err != nil {
		return nil, 0, err
	}

	BlockchainAddresss, err := this.UnMarshalArrayTItem(setItems)
	if err != nil {
		return nil, 0, err
	}

	return BlockchainAddresss, totalCount, err
}

func (this *BlockchainAddress) GetTotalCount() (int64, error) {
	totalCount, err := GetBsOriginal().GetTotalCount(this.GetBsKey())
	if err != nil || totalCount < 1 {
		return 0, err
	}

	return totalCount, err
}

func (this *BlockchainAddress) Create() error {
	now := time.Now().Unix()
	this.CreatedAt = now
	this.UpdatedAt = now

	bBlockchainAddress, key, err := helps.MarshalBytes(this)
	if err != nil {
		return err
	}

	log.Println(string(key), "-- string(key)")
	log.Println(this, "-- this")

	return GetBsOriginal().BsPutItem(this.GetBsKey(), &generic.TItem{
		Key:   key,
		Value: bBlockchainAddress,
	})
}

func (this *BlockchainAddress) GetItemBytes() ([]byte, error) {
	tBlockchainAddress, err := GetBsOriginal().BsGetItem(this.GetBsKey(), generic.TItemKey(this.String()))
	if err != nil {
		return nil, err
	}

	return tBlockchainAddress.GetValue(), nil
}

func (this *BlockchainAddress) UnMarshalArrayTItem(objects []*generic.TItem) ([]BlockchainAddress, error) {
	objs := make([]BlockchainAddress, 0)

	for _, object := range objects {
		obj := BlockchainAddress{}
		err := json.Unmarshal(object.GetValue(), &obj)

		if err != nil {
			return make([]BlockchainAddress, 0), err
		}

		objs = append(objs, obj)
	}

	return objs, nil
}

func (this *BlockchainAddress) GetFromKey(key string) (*BlockchainAddress, error) {
	item, err := GetBsOriginal().BsGetItem(this.GetBsKey(), generic.TItemKey(key))
	if err != nil {
		return nil, err
	}
	addr := &BlockchainAddress{}
	err = json.Unmarshal(item.GetValue(), &addr)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (this *BlockchainAddress) UpdateTime() {
	this.UpdatedAt = time.Now().Unix()
}
