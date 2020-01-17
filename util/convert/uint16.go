package convert

import (
	"errors"
	"strconv"
)

func Uint16(in interface{}) (uint16, error) {
	var ret uint16
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
		ret = uint16(temp)
	case int:
		ret = uint16(in.(int))
	case int32:
		ret = uint16(in.(int32))
	case int64:
		ret = uint16(in.(int64))
	case uint8:
		ret = uint16(in.(uint8))
	case uint16:
		ret = in.(uint16)
	case uint32:
		ret = uint16(in.(uint32))
	case uint64:
		ret = uint16(in.(uint64))
	case float64:
		ret = uint16(in.(float64))
	default:
		err = errors.New("unknown data type")
	}

	return ret, err
}

func Uint16WithoutErr(in interface{}) uint16 {
	i, _ := Uint16(in)
	return i
}
