package kernelSrc

type ROW_CALLBACK_FUNCTOR func(uint32, *AFNode, AFData, AFData) int

type AFRow struct {
	row          uint32
	func_        ROW_CALLBACK_FUNCTOR
	pNodeManager *AFNodeManager
}

func NewAFRow(pClassMeta *AFClassMeta, row uint32, args *AFDataList, func_ ROW_CALLBACK_FUNCTOR) *AFRow {
	ret := &AFRow{
		row:   row,
		func_: func_,
	}
	ret.pNodeManager = NewAFNodeManager(pClassMeta, args, ret.OnDataCallBack)

	return ret
}

// query row
func (row *AFRow) GetRowCount() uint32 { return row.row }

// get row value
// by index
func (row *AFRow) GetBoolByIndex(index uint32) bool       { return false }
func (row *AFRow) GetInt32ByIndex(index uint32) int32     { return 0 }
func (row *AFRow) GetInt64ByIndex(index uint32) int64     { return 0 }
func (row *AFRow) GetUInt32ByIndex(index uint32) uint32   { return 0 }
func (row *AFRow) GetUInt64ByIndex(index uint32) uint64   { return 0 }
func (row *AFRow) GetFloat32ByIndex(index uint32) float32 { return 0 }
func (row *AFRow) GetFloat64ByIndex(index uint32) float64 { return 0 }
func (row *AFRow) GetStringByIndex(index uint32) string   { return "" }
func (row *AFRow) GetGUIDByIndex(index uint32) AFGUID     { return 0 }

// by name
func (row *AFRow) GetBoolByName(name uint32) bool       { return false }
func (row *AFRow) GetInt32ByName(name uint32) int32     { return 0 }
func (row *AFRow) GetInt64ByName(name uint32) int64     { return 0 }
func (row *AFRow) GetUInt32ByName(nam uint32) uint32    { return 0 }
func (row *AFRow) GetUInt64ByName(nam uint32) uint64    { return 0 }
func (row *AFRow) GetFloat32ByName(name uint32) float32 { return 0 }
func (row *AFRow) GetFloat64ByName(name uint32) float64 { return 0 }
func (row *AFRow) GetStringByName(name uint32) string   { return "" }
func (row *AFRow) GetGUIDByName(name uint32) AFGUID     { return 0 }

// set row value
// by index
func (row *AFRow) SetBoolByIndex(index uint32, value bool) error       { return nil }
func (row *AFRow) SetInt32ByIndex(index uint32, value int32) error     { return nil }
func (row *AFRow) SetInt64ByIndex(index uint32, value int64) error     { return nil }
func (row *AFRow) SetUInt32ByIndex(index uint32, value uint32) error   { return nil }
func (row *AFRow) SetUInt64ByIndex(index uint32, value uint64) error   { return nil }
func (row *AFRow) SetFloat32ByIndex(index uint32, value float32) error { return nil }
func (row *AFRow) SetFloat64ByIndex(index uint32, value float64) error { return nil }
func (row *AFRow) SetStringByIndex(index uint32, value string) error   { return nil }
func (row *AFRow) SetGUIDByIndex(index uint32, value AFGUID) error     { return nil }

// by name
func (row *AFRow) SetBoolByName(index string, value bool) error       { return nil }
func (row *AFRow) SetInt32ByName(index string, value int32) error     { return nil }
func (row *AFRow) SetInt64ByName(index string, value int64) error     { return nil }
func (row *AFRow) SetUInt32ByName(index string, value uint32) error   { return nil }
func (row *AFRow) SetUInt64ByName(index string, value uint64) error   { return nil }
func (row *AFRow) SetFloat32ByName(index string, value float32) error { return nil }
func (row *AFRow) SetFloat64ByName(index string, value float64) error { return nil }
func (row *AFRow) SetStringByName(index string, value string) error   { return nil }
func (row *AFRow) SetGUIDByName(index string, value AFGUID) error     { return nil }

func (row *AFRow) initData(args *AFDataList) {}

func (row *AFRow) getNodeManager() *AFNodeManager { return nil }

func (row *AFRow) OnDataCallBack(pNode *AFNode, oldData *AFData, newData *AFData) int {
	return 0
}
