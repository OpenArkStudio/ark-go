package kernelSrc

import "errors"

type AFDataList struct {
}

func (list *AFDataList) Concat(src *AFDataList) bool  { return false }
func (list *AFDataList) Split(src, split string) bool { return false }

func (list *AFDataList) Append(data *AFData) bool                              { return false }
func (list *AFDataList) AppendDataList(src *AFDataList, start, count int) bool { return false }

func (list *AFDataList) Clear()      {}
func (list *AFDataList) Empty() bool { return false }

func (list *AFDataList) GetCount() int                 { return 0 }
func (list *AFDataList) GetType(index int) ArkDataType { return 0 }

// add data
func (list *AFDataList) AddBool(value bool)             {}
func (list *AFDataList) AddInt32(value int32)           {}
func (list *AFDataList) AddInt64(value int64)           {}
func (list *AFDataList) AddUInt32(value uint32)         {}
func (list *AFDataList) AddUInt64(value uint64)         {}
func (list *AFDataList) AddFloat32(value float32)       {}
func (list *AFDataList) AddFloat64(value float64)       {}
func (list *AFDataList) AddString(value string)         {}
func (list *AFDataList) AddInterface(value interface{}) {}

// get data
func (list *AFDataList) GetBool(value int) bool             { return false }
func (list *AFDataList) GetInt32(value int) int32           { return 0 }
func (list *AFDataList) GetInt64(value int) int64           { return 0 }
func (list *AFDataList) GetUInt32(value int) uint32         { return 0 }
func (list *AFDataList) GetUInt64(value int) uint64         { return 0 }
func (list *AFDataList) GetFloat32(value int) float32       { return 0 }
func (list *AFDataList) GetFloat64(value int) float64       { return 0 }
func (list *AFDataList) GetString(value int) string         { return "" }
func (list *AFDataList) GetInterface(value int) interface{} { return nil }

func (list *AFDataList) GetMemUsage() int64 { return 0 }

func (list *AFDataList) TypeEx(args ...ArkDataType) bool {
	for k, v := range args {
		if v == DT_EMPTY || list.GetType(k) != v {
			return false
		}
	}
	return true
}

func (list *AFDataList) ToAFIData(index int, data *AFData) error {
	if list.GetType(index) != data.GetType() {
		return errors.New("type mismatch")
	}

	switch data.GetType() {
	case DT_BOOLEAN:
		data.SetBool(list.GetBool(index))
	case DT_INT32:
		data.SetInt32(list.GetInt32(index))
	case DT_INT64:
		data.SetInt64(list.GetInt64(index))
	case DT_FLOAT32:
		data.SetFloat32(list.GetFloat32(index))
	case DT_FLOAT64:
		data.SetFloat64(list.GetFloat64(index))
	case DT_STRING:
		data.SetString(list.GetString(index))
	default:
		return errors.New("invalid data type")
	}

	return nil
}

func (list *AFDataList) Equal(index int, data *AFData) bool {
	if list.GetType(index) != data.GetType() {
		return false
	}

	flag := false
	switch data.GetType() {
	case DT_BOOLEAN:
		flag = data.GetBool() == list.GetBool(index)
	case DT_INT32:
		flag = data.GetInt32() == list.GetInt32(index)
	case DT_UINT32:
		flag = data.GetUInt32() == list.GetUInt32(index)
	case DT_INT64:
		flag = data.GetInt64() == list.GetInt64(index)
	case DT_UINT64:
		flag = data.GetUInt64() == list.GetUInt64(index)
	case DT_FLOAT32:
		flag = data.GetFloat32() == list.GetFloat32(index)
	case DT_FLOAT64:
		flag = data.GetFloat64() == list.GetFloat64(index)
	case DT_STRING:
		flag = data.GetString() == list.GetString(index)
	}

	return flag
}
