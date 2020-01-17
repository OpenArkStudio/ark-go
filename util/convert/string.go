package convert

import (
	"errors"
	"strconv"
)

func StringWithoutErr(in interface{}) string {
	ret, _ := String(in)
	return ret
}

func String(in interface{}) (string, error) {
	var ret string
	var err error

	switch in.(type) {
	case string:
		ret = in.(string)
	case []uint8:
		ret = string(in.([]uint8))
	case int64:
		ret = strconv.FormatInt(in.(int64), 10)
	case int:
		ret = strconv.Itoa(in.(int))
	case float64:
		ret = strconv.FormatFloat(in.(float64), 'E', -1, 64)
	case uint8:
		ret = strconv.FormatUint(in.(uint64), 8)
	case uint16:
		ret = strconv.FormatUint(in.(uint64), 16)
	case nil:
		ret = ""
	default:
		err = errors.New("unknown data type")
	}

	return ret, err
}
