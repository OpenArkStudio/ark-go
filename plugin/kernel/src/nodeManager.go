package kernelSrc

type (
	NODE_MANAGER_FUNCTOR func(*AFNode, *AFData, *AFData) int
	DataList             map[uint32]AFNode
)

type AFNodeManager struct {
	classMeta *AFClassMeta
	func_     NODE_MANAGER_FUNCTOR
	dataList  *DataList
}

func NewAFNodeManagerFromAFClassMeta(pClassMeta *AFClassMeta) *AFNodeManager {
	return &AFNodeManager{
		classMeta: pClassMeta,
	}
}

func NewAFNodeManager(pClassMeta *AFClassMeta, dataList *AFDataList, func_ NODE_MANAGER_FUNCTOR) *AFNodeManager {
	ret := &AFNodeManager{
		classMeta: pClassMeta,
		func_:     func_,
	}
	ret.InitData(dataList)

	return ret
}

func (manager *AFNodeManager) InitData(args *AFDataList) {

}

func (manager *AFNodeManager) IsEmpty() bool { return false }

// data operation
func (manager *AFNodeManager) CreateData(pDataMeta *AFNodeMeta) *AFNode { return nil }
func (manager *AFNodeManager) CreateDataByName(name string) *AFNode     { return nil }
func (manager *AFNodeManager) CreateDataByAFNode(pData *AFNode) bool    { return true }

// get node
func (manager *AFNodeManager) GetNodeByName(name string) *AFNode   { return nil }
func (manager *AFNodeManager) GetNodeByIndex(index uint32) *AFNode { return nil }

// query data
func (manager *AFNodeManager) GetBoolByName(name string) bool       { return false }
func (manager *AFNodeManager) GetInt32ByName(name string) int32     { return 0 }
func (manager *AFNodeManager) GetUInt32ByName(name string) uint32   { return 0 }
func (manager *AFNodeManager) GetInt64ByName(name string) int64     { return 0 }
func (manager *AFNodeManager) GetUInt64ByName(name string) uint64   { return 0 }
func (manager *AFNodeManager) GetFloat32ByName(name string) float32 { return 0 }
func (manager *AFNodeManager) GetFloat64ByName(name string) float64 { return 0 }
func (manager *AFNodeManager) GetStringByName(name string) string   { return "" }
func (manager *AFNodeManager) GetGuidByName(name string) AFGUID     { return 0 }

func (manager *AFNodeManager) GetBoolByIndex(index uint32) bool       { return false }
func (manager *AFNodeManager) GetInt32ByIndex(index uint32) int32     { return 0 }
func (manager *AFNodeManager) GetUInt32ByIndex(index uint32) uint32   { return 0 }
func (manager *AFNodeManager) GetInt64ByIndex(index uint32) int64     { return 0 }
func (manager *AFNodeManager) GetUInt64ByIndex(index uint32) uint64   { return 0 }
func (manager *AFNodeManager) GetFloat32ByIndex(index uint32) float32 { return 0 }
func (manager *AFNodeManager) GetFloat64ByIndex(index uint32) float64 { return 0 }
func (manager *AFNodeManager) GetStringByIndex(index uint32) string   { return "" }
func (manager *AFNodeManager) GetGuidByIndex(index uint32) AFGUID     { return 0 }

// set data
func (manager *AFNodeManager) SetBoolByName(name string, value error) error      { return nil }
func (manager *AFNodeManager) SetInt32ByName(name string, value int32) error     { return nil }
func (manager *AFNodeManager) SetUInt32ByName(name string, value uint32) error   { return nil }
func (manager *AFNodeManager) SetInt64ByName(name string, value int64) error     { return nil }
func (manager *AFNodeManager) SetUInt64ByName(name string, value uint64) error   { return nil }
func (manager *AFNodeManager) SetFloat32ByName(name string, value float32) error { return nil }
func (manager *AFNodeManager) SetFloat64ByName(name string, value float64) error { return nil }
func (manager *AFNodeManager) SetStringByName(name string, value string) error   { return nil }
func (manager *AFNodeManager) SetGuidByName(name string, value AFGUID) error     { return nil }

func (manager *AFNodeManager) SetBoolByIndex(index uint32, value bool) error       { return nil }
func (manager *AFNodeManager) SetInt32ByIndex(index uint32, value int32) error     { return nil }
func (manager *AFNodeManager) SetUInt32ByIndex(index uint32, value uint32) error   { return nil }
func (manager *AFNodeManager) SetInt64ByIndex(index uint32, value int64) error     { return nil }
func (manager *AFNodeManager) SetUInt64ByIndex(index uint32, value uint64) error   { return nil }
func (manager *AFNodeManager) SetFloat32ByIndex(index uint32, value float32) error { return nil }
func (manager *AFNodeManager) SetFloat64ByIndex(index uint32, value float64) error { return nil }
func (manager *AFNodeManager) SetStringByIndex(index uint32, value string) error   { return nil }
func (manager *AFNodeManager) SetGuidByIndex(index uint32, value AFGUID) error     { return nil }

// other query
func (manager *AFNodeManager) GetDataList() *DataList { return nil }

func (manager *AFNodeManager) getIndex(name string) uint32                { return 0 }
func (manager *AFNodeManager) findData(index uint32) *AFNode              { return nil }
func (manager *AFNodeManager) findDataWithoutChange(index uint32) *AFNode { return nil }
