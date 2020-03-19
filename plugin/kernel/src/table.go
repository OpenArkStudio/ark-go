package kernelSrc

type AFCTable struct {
}

func (table *AFCTable) GetIndexByName(name string) uint32                  { return 0 }
func (table *AFCTable) GetName() string                                    { return "" }
func (table *AFCTable) GetColCount() uint32                                { return 0 }
func (table *AFCTable) GetColType(index uint32) ArkDataType                { return 0 }
func (table *AFCTable) GetMask() ArkMaskType                               { return 0 }
func (table *AFCTable) HaveMask(future ArkDataMask) bool                   { return false }
func (table *AFCTable) IsPublic() bool                                     { return false }
func (table *AFCTable) IsPrivate() bool                                    { return false }
func (table *AFCTable) IsRealTime() bool                                   { return false }
func (table *AFCTable) IsSave() bool                                       { return false }
func (table *AFCTable) FindInt32(index uint32, value int32) uint32         { return 0 }
func (table *AFCTable) FindInt64(index uint32, value int64) uint32         { return 0 }
func (table *AFCTable) FindBool(index uint32, value bool) uint32           { return 0 }
func (table *AFCTable) FindFloat32(index uint32, value float32) uint32     { return 0 }
func (table *AFCTable) FindFloat64(index uint32, value float64) uint32     { return 0 }
func (table *AFCTable) FindString(index uint32, value string) uint32       { return 0 }
func (table *AFCTable) FindGUID(index uint32, value AFGUID) uint32         { return 0 }
func (table *AFCTable) First() *AFRow                                      { return nil }
func (table *AFCTable) Next() *AFRow                                       { return nil }
func (table *AFCTable) GetIndex() uint32                                   { return 0 }
func (table *AFCTable) GetRowCount() uint32                                { return 0 }
func (table *AFCTable) AddRow(row uint32) *AFRow                           { return nil }
func (table *AFCTable) AddRowWithData(row uint32, args *AFDataList) *AFRow { return nil }
func (table *AFCTable) FindRow(row uint32) *AFRow                          { return nil }
func (table *AFCTable) RemoveRow(row uint32) error                         { return nil }
func (table *AFCTable) Clear()                                             {}
