package main

import (
	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"sync"
	"wokerTrans/bigsetservice"
)

var (
	onceBS     sync.Once
	mapBigset  map[string]*MyStringBSService
	originSID  = "/openstars/cryptocurrency/cryptoapi/services/bigset-skv"
	originHost = "127.0.0.1"
	originPort = "18407"
	txSID      = "/openstars/cryptocurrency/crypto-transaction/services/thrift"
	txHost     = "127.0.0.1"
	txPort     = "18407"
	txProtocol = "binary"
)

type MyStringBSService struct {
	StringBigsetService.StringBigsetServiceIf
}

func InitBigSetIf() {
	onceBS.Do(func() {
		mapBigset = map[string]*MyStringBSService{}
		mapBigset[BS_ORIGIN_SERVICE] = &MyStringBSService{
			bigsetservice.GetBigSet(originSID, originHost, originPort)}
		mapBigset[BS_TRANSACTION_SERVICE] = &MyStringBSService{
			bigsetservice.GetBigSet(txSID, txHost, txPort)}
	})
}

func GetBsOriginal() *MyStringBSService {
	return mapBigset[BS_ORIGIN_SERVICE]
}

func GetBsTransService() *MyStringBSService {
	return mapBigset[BS_TRANSACTION_SERVICE]
}
