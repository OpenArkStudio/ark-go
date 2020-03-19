package kernelSrc

type (
	NameIndexList     map[string]uint32
	DataMetaList      map[uint32]*AFNodeMeta
	TableMetaList     map[uint32]*AFTableMeta
	ContainerMetaList map[uint32]*AFContainerMeta
)

type AFClassMeta struct {
	className         string                  // 类名
	resPath           string                  // 资源路径
	nameIndexList     NameIndexList           // 字段名和下表映射
	dataMetaList      DataMetaList            // 属性映射
	tableMetaList     TableMetaList           // 表映射
	containerMetaList ContainerMetaList       // 容器映射
	classMetaCallback *AFClassCallBackManager // 函数回调管理器
}

func NewAFClassMeta(name string) *AFClassMeta {
	return &AFClassMeta{
		className:         name,
		classMetaCallback: &AFClassCallBackManager{},
	}
}

func (meta *AFClassMeta) GetNodeCount() uint32 { return uint32(len(meta.dataMetaList)) }

func (meta *AFClassMeta) GetName() string { return meta.className }

func (meta *AFClassMeta) SetResPath(path string) { meta.resPath = path }

func (meta *AFClassMeta) GetResPath() string { return meta.resPath }

// create data meta
func (meta *AFClassMeta) CreateDataMeta(name string, index uint32) *AFNodeMeta {
	// check arg
	if len(name) == 0 || index <= 0 {
		return nil
	}

	if meta.dataMetaList[index] != nil {
		return nil
	}

	if !meta.nameIndexInsert(name, index) {
		return nil
	}

	pMeta := NewAFNodeMeta(name, index)
	meta.dataMetaList[index] = pMeta
	return pMeta
}

func (meta *AFClassMeta) FindDataMeta(index uint32) *AFNodeMeta { return meta.dataMetaList[index] }

// create table meta
func (meta *AFClassMeta) CreateTableMeta(name string, index uint32) *AFTableMeta {
	// check arg
	if len(name) == 0 || index <= 0 {
		return nil
	}

	if meta.tableMetaList[index] != nil {
		return nil
	}

	if !meta.nameIndexInsert(name, index) {
		return nil
	}

	pMeta := NewAFTableMeta(name, index)
	meta.tableMetaList[index] = pMeta
	return pMeta
}

func (meta *AFClassMeta) FindTableMeta(index uint32) *AFTableMeta { return meta.tableMetaList[index] }

// create container meta
func (meta *AFClassMeta) CreateContainerMeta(name string, index uint32, className string, mask ArkDataMask) *AFContainerMeta {
	// check arg
	if len(name) == 0 || index <= 0 {
		return nil
	}

	if meta.containerMetaList[index] != nil {
		return nil
	}

	if !meta.nameIndexInsert(name, index) {
		return nil
	}

	pMeta := NewAFContainerMeta(name, index, className, mask)
	meta.containerMetaList[index] = pMeta
	return pMeta
}

func (meta *AFClassMeta) FindContainerMeta(index uint32) *AFContainerMeta {
	return meta.containerMetaList[index]
}

func (meta *AFClassMeta) GetDataMetaList() DataMetaList           { return meta.dataMetaList }
func (meta *AFClassMeta) GetTableMetaList() TableMetaList         { return meta.tableMetaList }
func (meta *AFClassMeta) GetContainerMetaList() ContainerMetaList { return meta.containerMetaList }
func (meta *AFClassMeta) GetClassCallBackManager() *AFClassCallBackManager {
	return meta.classMetaCallback
}

func (meta *AFClassMeta) GetIndex(name string) uint32 {
	return meta.nameIndexList[name]
}

func (meta *AFClassMeta) IsEntityMeta() bool { return len(meta.resPath) == 0 }

func (meta *AFClassMeta) nameIndexInsert(name string, index uint32) bool {
	if _, isExist := meta.nameIndexList[name]; isExist {
		return false
	}
	meta.nameIndexList[name] = index
	return true
}
