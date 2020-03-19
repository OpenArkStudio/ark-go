package kernelSrc

type AFContainerMeta struct {
	containerName  string // 容器名
	containerIndex uint32 // 下标
	className      string // 容器容纳的对象名
	mask           ArkDataMask
}

func NewAFContainerMeta(name string, index uint32, className string, mask ArkDataMask) *AFContainerMeta {
	return &AFContainerMeta{
		containerName:  name,
		containerIndex: index,
		className:      className,
		mask:           mask,
	}
}

func (meta *AFContainerMeta) GetName() string {
	return meta.containerName
}

func (meta *AFContainerMeta) GetClassName() string {
	return meta.className
}

func (meta *AFContainerMeta) GetIndex() uint32 {
	return meta.containerIndex
}

func (meta *AFContainerMeta) GetMask() ArkDataMask {
	return meta.mask
}
