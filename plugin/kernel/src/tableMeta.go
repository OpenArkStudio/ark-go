package kernelSrc

// 表结构定义
type AFTableMeta struct {
	tableName string       // 表名
	index     uint32       // 下标
	typeName  string       // 表对应的AFClassMeta 名称 string
	classMeta *AFClassMeta // 对应的AFClassMeta
	mask      ArkDataMask  // 是否存储同步等特性
}

func NewAFTableMeta(name string, index uint32) *AFTableMeta {
	return &AFTableMeta{
		tableName: name,
		index:     index,
	}
}

func (meta *AFTableMeta) GetName() string {
	return meta.tableName
}

func (meta *AFTableMeta) GetColCount() uint32 {
	if meta.classMeta == nil {
		return 0
	}
	return meta.classMeta.GetNodeCount()
}

func (meta *AFTableMeta) SetClassMeta(pClassMeta *AFClassMeta) {
	meta.classMeta = pClassMeta
}

func (meta *AFTableMeta) FindNodeMeta(index uint32) *AFNodeMeta {
	if meta.classMeta == nil {
		return nil
	}

	return meta.classMeta.FindDataMeta(index)
}

func (meta *AFTableMeta) GetClassMeta() *AFClassMeta {
	return meta.classMeta
}

func (meta *AFTableMeta) GetColType(index uint32) ArkDataType {
	if meta.classMeta == nil {
		return DT_EMPTY
	}

	pMeta := meta.classMeta.FindDataMeta(index)
	if pMeta == nil {
		return DT_EMPTY
	}
	return pMeta.GetType()
}

func (meta *AFTableMeta) GetMask() ArkDataMask { return meta.mask }

func (meta *AFTableMeta) HaveMask(feature ArkMaskType) bool { return meta.mask.HaveMask(feature) }

func (meta *AFTableMeta) SetMask(feature ArkDataMask) { meta.mask = feature }

func (meta *AFTableMeta) IsPublic() bool { return meta.mask.HaveMask(PF_SYNC_VIEW) }

func (meta *AFTableMeta) IsPrivate() bool { return meta.mask.HaveMask(PF_SYNC_SELF) }

func (meta *AFTableMeta) IsRealTime() bool { return meta.mask.HaveMask(PF_REAL_TIME) }

func (meta *AFTableMeta) IsSave() bool { return meta.mask.HaveMask(PF_SAVE) }

func (meta *AFTableMeta) SetTypeName(value string) { meta.typeName = value }

func (meta *AFTableMeta) GetTypeName() string { return meta.typeName }

func (meta *AFTableMeta) GetIndex() uint32 { return meta.index }

func (meta *AFTableMeta) GetIndexByName(name string) uint32 {
	if meta.classMeta == nil {
		return 0
	}
	return meta.classMeta.GetIndex(name)
}
