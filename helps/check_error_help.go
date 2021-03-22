package helps

import (
	"strings"

	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
)

func IsError(err error) bool {
	return err != nil && !strings.Contains(err.Error(), generic.TErrorCode_EGood.String())
}

func IsNotExisted(err error) bool {
	return err != nil && strings.Contains(err.Error(), generic.TErrorCode_EItemNotExisted.String())
}

func IsUnknown(err error) bool {
	return err != nil && strings.Contains(err.Error(), generic.TErrorCode_EUnknownException.String())
}
func IsNotFoundKey(err error) bool {
	return strings.Contains(err.Error(), "Can not found key")
}

//func GetSoftError(err error) (int, string) {
//	log.Printf("%s: Error: %v", time.Now().Format(consts.DATE_DEFAULT_FORMAT), err)
//	return consts.REST_ERROR_CODE_CLIENT, fmt.Sprintf("%s: Error: %v\n", time.Now().Format(consts.DATE_DEFAULT_FORMAT), err)
//}
//
//func GetSoftError501(err error) (int, string) {
//	log.Printf("%s: Error: %v\n", time.Now().Format(consts.DATE_DEFAULT_FORMAT), err)
//	return consts.REST_ERROR_CODE_SERVER, fmt.Sprintf("%s: Error: %v", time.Now().Format(consts.DATE_DEFAULT_FORMAT), err)
//}
