package ark

import (
	"path/filepath"
	"reflect"
)

func GetType(i interface{}) reflect.Type {
	return reflect.TypeOf(i).Elem()
}

func GetName(i interface{}) string {
	t := reflect.TypeOf(i).Elem()
	return filepath.Join(t.PkgPath(), t.Name())
}
