package model

import . "ark-go/common"

type DynEntryPluginFunc func(*AFPluginManager)
type DynExitPluginFunc func(*AFPluginManager)

type AFPluginManager struct {
	busId          int    // bus id
	timestamp      int64  // loop timestamp
	pluginPath     string // the plugin.so filepath
	resPath        string // the resource filepath
	pluginConfPath string // plugin configuration filepath

	pluginNameList            map[string]bool      // plugin_name -> bool
	orderedPluginNameList     []string             // ordered plugin names
	pluginLibList             map[string]*AFDynLib // dynamic libraries
	pluginInstanceList        map[string]*AFPlugin // plugin instances
	moduleInstanceList        map[string]*AFModule // module instances
	orderedModuleInstanceList []*AFModule          // ordered module instances

	moduleWithUpdateFuncList map[string]*AFModule // the list of modules who have the `update` function
}

// AFPluginManager functions
