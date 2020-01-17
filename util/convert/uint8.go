package convert

import (
	"errors"
	"strconv"
)

func Uint8(in interface{}) (uint8, error) {
	var ret uint8
	var err error

	switch in.(type) {
	case string:
		inp := in.(string)
		if inp == "" {
			err = errors.New("in is empty string")
			break
		}
		temp, err := strconv.ParseUint(inp, 10, 8)
		if err != nil {
			break
		}
		ret = uint8(temp)
	case int:
		ret = uint8(in.(int))
	case int32:
		ret = uint8(in.(int32))
	case int64:
		ret = in.(uint8)
	case uint8:
		ret = in.(uint8)
	case uint16:
		ret = uint8(in.(uint16))
	case uint32:
		ret = uint8(in.(uint32))
	case uint64:
		ret = uint8(in.(uint64))
	case float64:
		ret = uint8(in.(float64))
	default:
		err = errors.New("unknown data type")
	}

	return ret, err
}

func Uint8WithoutErr(str interface{}) uint8 {
	i, _ := Uint8(str)
	return i
}
