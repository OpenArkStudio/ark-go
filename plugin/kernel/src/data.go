package kernelSrc

import (
	"fmt"
	"github.com/spf13/cast"
	"unsafe"
)

type AFData struct {
	dataType ArkDataType
	value    interface{}
}

func NewAFData(value interface{}) *AFData {
	ret := new(AFData)

	if value == nil {
		// return as soon as possible if empty
		ret.dataType = DT_EMPTY
		return ret
	}

	switch value.(type) {
	case AFData:
		temp := value.(AFData)
		ret.dataType = temp.dataType
		ret.value = temp.value
		return ret
	case *AFData:
		temp := value.(*AFData)
		ret.dataType = temp.dataType
		ret.value = temp.value
		return ret
	case bool:
		ret.dataType = DT_BOOLEAN
	case int32:
		ret.dataType = DT_INT32
	case uint32:
		ret.dataType = DT_UINT32
	case int64:
		ret.dataType = DT_INT64
	case uint64:
		ret.dataType = DT_UINT64
	case float32:
		ret.dataType = DT_FLOAT32
	case float64:
		ret.dataType = DT_FLOAT64
	case string:
		ret.dataType = DT_STRING
	case interface{}:
		ret.dataType = DT_INTERFACE
	default:
		ret.dataType = DT_EMPTY
		return ret
	}

	ret.value = value
	return ret
}

/////////////////////////////////////////////////////////////////
func (data *AFData) Release() {
	data.dataType = DT_EMPTY
	data.value = nil
}

func (data *AFData) GetMemUsage() int64 { return int64(unsafe.Sizeof(data.value)) }

func (data *AFData) ToString() string { return fmt.Sprintf("%v", data.value) }

func (data *AFData) IsNilValue() bool { return data.value == nil }

func (data *AFData) SetDefaultValue(t ArkDataType) {
	switch t {
	case DT_BOOLEAN:
		data.SetBool(false)
	case DT_INT32:
		data.SetInt32(0)
	case DT_UINT32:
		data.SetUInt32(0)
	case DT_INT64:
		data.SetInt64(0)
	case DT_UINT64:
		data.SetUInt64(0)
	case DT_FLOAT32:
		data.SetFloat32(0)
	case DT_FLOAT64:
		data.SetFloat64(0)
	case DT_STRING:
		data.SetString("")
	default:
		data.value = nil
	}
}

/////////////////////////////////////////////////////////////////
func (data *AFData) GetType() ArkDataType { return data.dataType }

// Set Data
func (data *AFData) SetUnknown()                    { data.setValue(DT_EMPTY, nil) }
func (data *AFData) SetBool(value bool)             { data.setValue(DT_BOOLEAN, value) }
func (data *AFData) SetInt32(value int32)           { data.setValue(DT_INT32, value) }
func (data *AFData) SetUInt32(value uint32)         { data.setValue(DT_UINT32, value) }
func (data *AFData) SetInt64(value int64)           { data.setValue(DT_INT64, value) }
func (data *AFData) SetUInt64(value uint64)         { data.setValue(DT_UINT64, value) }
func (data *AFData) SetFloat32(value float32)       { data.setValue(DT_FLOAT32, value) }
func (data *AFData) SetFloat64(value float64)       { data.setValue(DT_FLOAT64, value) }
func (data *AFData) SetString(value string)         { data.setValue(DT_STRING, value) }
func (data *AFData) SetInterface(value interface{}) { data.setValue(DT_INTERFACE, value) }
func (data *AFData) setValue(t ArkDataType, v interface{}) {
	data.Release()
	data.dataType = t
	data.value = v
}

/////////////////////////////////////////////////////////////////
// Get Data
func (data *AFData) GetBool() bool {
	if data.dataType != DT_BOOLEAN {
		return false
	}
	return cast.ToBool(data.value)
}

func (data *AFData) GetInt32() int32 {
	if data.dataType != DT_INT32 {
		return 0
	}
	return cast.ToInt32(data.value)
}

func (data *AFData) GetUInt32() uint32 {
	if data.dataType != DT_UINT32 {
		return 0
	}
	return cast.ToUint32(data.value)
}

func (data *AFData) GetInt64() int64 {
	if data.dataType != DT_INT64 {
		return 0
	}
	return cast.ToInt64(data.value)
}

func (data *AFData) GetUInt64() uint64 {
	if data.dataType != DT_UINT64 {
		return 0
	}
	return cast.ToUint64(data.value)
}

func (data *AFData) GetFloat32() float32 {
	if data.dataType != DT_FLOAT32 {
		return 0
	}
	return cast.ToFloat32(data.value)
}

func (data *AFData) GetFloat64() float64 {
	if data.dataType != DT_FLOAT64 {
		return 0
	}
	return cast.ToFloat64(data.value)
}

func (data *AFData) GetString() string {
	if data.dataType != DT_STRING {
		return ""
	}
	return cast.ToString(data.value)
}

func (data *AFData) GetInterface() interface{} {
	if data.dataType != DT_INTERFACE {
		return nil
	}
	return data.value
}
