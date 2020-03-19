package kernelSrc

type AFNodeMeta struct {
	name  string
	type_ ArkDataType
	index uint32
	mask  ArkDataMask
}

func NewAFNodeMeta(name string, index uint32) *AFNodeMeta {
	return &AFNodeMeta{
		name:  name,
		index: index,
	}
}

func (meta *AFNodeMeta) GetName() string {
	return meta.name
}

func (meta *AFNodeMeta) SetType(t ArkDataType) {
	meta.type_ = t
}

func (meta *AFNodeMeta) GetType() ArkDataType {
	return meta.type_
}

func (meta *AFNodeMeta) SetMask(mask ArkDataMask) {
	meta.mask = mask
}

func (meta *AFNodeMeta) HaveMask(maskType ArkMaskType) bool {
	return meta.mask.HaveMask(maskType)
}

func (meta *AFNodeMeta) GetIndex() uint32 {
	return meta.index
}

func (meta *AFNodeMeta) GetMask() ArkDataMask {
	return meta.mask
}
