package kernelSrc

type AFDelaySyncRow struct {
	needClear bool
	row       uint32
	// TODO: replace with set
	nodeList map[*AFNode]interface{}
}

func NewAFDelaySyncRow(row uint32) *AFDelaySyncRow {
	return &AFDelaySyncRow{row: row}
}

type AFDelaySyncTable struct {
	needClear bool
	rowList   map[uint32]AFDelaySyncRow
}

type AFDelaySyncContainer struct {
	indexList   map[uint32]interface{}
	destroyList map[uint32]interface{}
}

type AFDelaySyncData struct {
	nodeList      map[*AFNode]interface{}
	tableList     map[uint32]AFDelaySyncTable
	containerList map[uint32]AFDelaySyncContainer
}

//////////////////////////////////////////////////////
type DelaySyncMaskData map[ArkMaskType]AFDelaySyncData
type DelaySyncDataList map[AFGUID]DelaySyncMaskData
type NODE_SYNC_FUNCTOR func(AFGUID, uint32, ArkDataType, AFData) int
type TABLE_SYNC_FUNCTOR func(AFGUID, uint32, ArkDataType, AFData) int

type AFClassCallBackManager struct {
}
