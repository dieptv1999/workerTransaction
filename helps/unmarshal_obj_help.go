package helps

import (
	"encoding/json"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"log"
)

func UnMarshalArrayTItem(objects []*generic.TItem) interface{} {
	log.Println("method @UnMarshalArrayTItem -- begin")
	objs := make([]interface{}, 0)
	var obj interface{}

	for _, object := range objects {
		err := json.Unmarshal(object.GetValue(), &obj)
		if err != nil {
			log.Println(err.Error(), "-- err.Error() helps/unmarshal_obj_help.go:18")
			return make([]interface{}, 0)
		}

		objs = append(objs, obj)
	}

	log.Println("method @UnMarshalArrayTItem -- end")
	return objs
}

func UnMarshalArrayTItemToStringKey(objects []*generic.TItem) []string {
	log.Println("method @UnMarshalArrayTItemToStringKey -- begin")
	objs := make([]string, 0)
	for _, object := range objects {
		objs = append(objs, string(object.GetKey()))
	}

	log.Println("method @UnMarshalArrayTItemToStringKey -- end")
	return objs
}

func UnMarshalArrayTItemToStringVal(objects []*generic.TItem) []string {
	log.Println("method @UnMarshalArrayTItemToStringVal -- begin")
	objs := make([]string, 0)
	for _, object := range objects {
		objs = append(objs, string(object.GetKey()))
	}

	log.Println("method @UnMarshalArrayTItemToStringVal -- end")
	return objs
}

func UnMarshalTItem(object *generic.TItem) (interface{}, error) {
	var obj interface{}

	err := json.Unmarshal(object.GetValue(), &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func UnMarshalBytes(bytes []byte) (interface{}, error) {
	var obj interface{}

	err := json.Unmarshal(bytes, &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
