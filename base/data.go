package base

import (
	"fmt"
)

type AFData struct {
	value interface{}
}

func (data *AFData) Set(value interface{}) {
	data.value = value
}

func (data *AFData) GetInt64() (value int64, err error) {
	value, ret := data.value.(int64)
	if !ret {
		return 0, fmt.Errorf("GetInt64 failed, %v", value)
	}

	return value, nil
}

func (data *AFData) GetFloat32() (value float32, err error) {
	value, ret := data.value.(float32)
	if !ret {
		return 0, fmt.Errorf("GetInt64 failed, %v", value)
	}

	return value, nil
}
