package kernelSrc

import (
	"fmt"
	"github.com/spf13/cast"
	"unsafe"
)

type AFNode struct {
	dataMeta *AFNodeMeta
	value    interface{}
}

func NewAFNodeByMeta(pDataMeta *AFNodeMeta) *AFNode {
	switch pDataMeta.GetType() {
	case DT_BOOLEAN, DT_INT32, DT_UINT32, DT_INT64, DT_UINT64, DT_FLOAT32, DT_FLOAT64, DT_STRING, DT_GUID:
		return &AFNode{dataMeta: pDataMeta}
	}

	return nil
}

func (node *AFNode) SetMeta(pDataMeta *AFNodeMeta) {
	node.dataMeta = pDataMeta
}

func (node *AFNode) GetMeta() *AFNodeMeta {
	return node.dataMeta
}

func (node *AFNode) GetName() string {
	if node.dataMeta != nil {
		return node.dataMeta.GetName()
	}
	return ""
}

func (node *AFNode) GetIndex() uint32 {
	if node.dataMeta != nil {
		return node.dataMeta.GetIndex()
	}
	return 0
}

func (node *AFNode) GetType() ArkDataType {
	if node.dataMeta != nil {
		return node.dataMeta.GetType()
	}
	return DT_EMPTY
}

func (node *AFNode) HaveMask(maskType ArkMaskType) bool {
	if node.dataMeta != nil {
		return node.dataMeta.HaveMask(maskType)
	}
	return false
}

func (node *AFNode) GetMask() ArkDataMask {
	if node.dataMeta != nil {
		return node.dataMeta.GetMask()
	}
	return ArkDataMask{}
}

func (node *AFNode) Reset() {
	node.dataMeta = nil
}

func (node *AFNode) IsNil() bool {
	return node.value == nil
}

func (node *AFNode) CopyFrom(other *AFNode) {
	if node.dataMeta == nil || other == nil {
		return
	}
	if node.dataMeta.GetType() != other.GetType() {
		return
	}

	node.value = other.value
}

func (node *AFNode) SaveTo(other *AFNode) {
	if node.dataMeta == nil || other == nil {
		return
	}
	if node.dataMeta.GetType() != other.GetType() {
		return
	}

	other.value = node.value
}

func (node *AFNode) ToString() string {
	return fmt.Sprintf("%v", node.value)
}

func (node *AFNode) GetValue() ID_TYPE {
	return ID_TYPE(cast.ToUint32(node.value))
}

func (node *AFNode) GetMemUsage() int64 { return int64(unsafe.Sizeof(node.value)) }

func (node *AFNode) FromString(str string) {
	switch node.dataMeta.GetType() {
	case DT_BOOLEAN:
		node.value = cast.ToBool(str)
	case DT_INT32:
		node.value = cast.ToInt32(str)
	case DT_UINT32:
		node.value = cast.ToUint32(str)
	case DT_INT64:
		node.value = cast.ToInt64(str)
	case DT_UINT64:
		node.value = cast.ToUint64(str)
	case DT_FLOAT32:
		node.value = cast.ToFloat32(str)
	case DT_FLOAT64:
		node.value = cast.ToFloat64(str)
	case DT_STRING:
		node.value = cast.ToString(str)
	case DT_GUID:
		node.value = cast.ToBool(str)
	}
}

/////////////////////////////////////////////////////////////////
// get
func (node *AFNode) GetBool() bool {
	if node.dataMeta.type_ != DT_UINT64 {
		return false
	}
	return cast.ToBool(node.value)
}

func (node *AFNode) GetInt32() int32 {
	if node.dataMeta.type_ != DT_INT32 {
		return 0
	}
	return cast.ToInt32(node.value)
}

func (node *AFNode) GetUInt32() uint32 {
	if node.dataMeta.type_ != DT_UINT32 {
		return 0
	}
	return cast.ToUint32(node.value)
}

func (node *AFNode) GetInt64() int64 {
	if node.dataMeta.type_ != DT_INT64 {
		return 0
	}
	return cast.ToInt64(node.value)
}

func (node *AFNode) GetUInt64() uint64 {
	if node.dataMeta.type_ != DT_UINT64 {
		return 0
	}
	return cast.ToUint64(node.value)
}

func (node *AFNode) GetFloat32() float32 {
	if node.dataMeta.type_ != DT_FLOAT32 {
		return 0
	}
	return cast.ToFloat32(node.value)
}

func (node *AFNode) GetFloat64() float64 {
	if node.dataMeta.type_ != DT_FLOAT64 {
		return 0
	}
	return cast.ToFloat64(node.value)
}

func (node *AFNode) GetString() string {
	if node.dataMeta.type_ != DT_STRING {
		return ""
	}
	return cast.ToString(node.value)
}

func (node *AFNode) GetObject() AFGUID {
	if node.dataMeta.type_ != DT_GUID {
		return 0
	}
	return AFGUID(cast.ToInt64(node.value))
}

func (node *AFNode) GetInterface() interface{} {
	if node.dataMeta.type_ != DT_INT32 {
		return 0
	}
	return node.value
}

// set
func (node *AFNode) SetUnknown()                    {}
func (node *AFNode) SetBool(value bool)             { node.setValueWithCheck(DT_BOOLEAN, value) }
func (node *AFNode) SetInt64(value int64)           { node.setValueWithCheck(DT_INT64, value) }
func (node *AFNode) SetUInt64(value uint64)         { node.setValueWithCheck(DT_UINT64, value) }
func (node *AFNode) SetFloat32(value float32)       { node.setValueWithCheck(DT_FLOAT32, value) }
func (node *AFNode) SetFloat64(value float64)       { node.setValueWithCheck(DT_FLOAT64, value) }
func (node *AFNode) SetString(value string)         { node.setValueWithCheck(DT_STRING, value) }
func (node *AFNode) SetInterface(value interface{}) { node.setValueWithCheck(DT_INTERFACE, value) }
func (node *AFNode) SetInt32(value int32)           { node.setValueWithCheck(DT_INT32, value) }
func (node *AFNode) SetUInt32(value uint32)         { node.setValueWithCheck(DT_UINT32, value) }
func (node *AFNode) SetObject(value AFGUID)         { node.setValueWithCheck(DT_GUID, value) }
func (node *AFNode) setValueWithCheck(t ArkDataType, value interface{}) {
	if node.dataMeta.type_ == t {
		node.value = value
	}
}
