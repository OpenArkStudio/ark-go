package common

import (
	"plugin"
)

type AFDynLib struct {
	name        string
	libInstance *plugin.Plugin
}

func (lib *AFDynLib) GetName() string { return lib.name }
func (lib *AFDynLib) Load(path string) bool {
	p, err := plugin.Open(path)
	if err != nil {
		panic(err)
	}

	lib.libInstance = p
	return true
}

func (lib *AFDynLib) Unload() bool {
	// TODO
	return true
}

func (lib *AFDynLib) GetSymbol(symbol string) plugin.Symbol {
	f, err := lib.libInstance.Lookup(symbol)
	if err != nil {
		panic(err)
	}

	//f.(func())()
	return f

}
