package bigsetservice

import (
	"fmt"
	"log"

	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

var (
	etcdEndpoint = []string{"127.0.0.1:2379"}
)

func init() {
	log.Println("method @ services/bigsetservice/big_set_service.go:19")
}


func GetBigSet(sid, host, port string) StringBigsetService.StringBigsetServiceIf {
	key := fmt.Sprintf("%s:%s", host, port)
	log.Println("--------------", sid, host, port, etcdEndpoint)
	log.Println(key, "-- key")

	return StringBigsetService.NewStringBigsetServiceModel(sid,
		etcdEndpoint,
		GoEndpointBackendManager.EndPoint{
			Host:      host,
			Port:      port,
			ServiceID: sid,
		})
}
