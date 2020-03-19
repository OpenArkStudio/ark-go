package kernelSrc

type ArkDataType uint32

const (
	DT_EMPTY     ArkDataType = iota // unknown
	DT_BOOLEAN                      // bool
	DT_INT32                        // int32
	DT_UINT32                       // uint32
	DT_INT64                        // int64
	DT_UINT64                       // uint64
	DT_FLOAT32                      // float32
	DT_FLOAT64                      // float64
	DT_STRING                       // string
	DT_INTERFACE                    // interface
	DT_VECTOR3D                     // vector3d
	DT_ARRAY                        // array
	DT_TABLE                        // DataTable
	DT_GUID                         // guid
)

var ArkDataTypeMap = map[string]ArkDataType{
	"":          DT_EMPTY,
	"bool":      DT_BOOLEAN,
	"int32":     DT_INT32,
	"uint32":    DT_UINT32,
	"int64":     DT_INT64,
	"uint64":    DT_UINT64,
	"float32":   DT_FLOAT32,
	"float64":   DT_FLOAT64,
	"string":    DT_STRING,
	"interface": DT_INTERFACE,
	"vector3d":  DT_VECTOR3D,
	"array":     DT_ARRAY,
	"table":     DT_TABLE,
	"guid":      DT_GUID,
}
