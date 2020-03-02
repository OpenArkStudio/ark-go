package ark

import "encoding/xml"

type pluginConf struct {
	XMLName xml.Name `xml:"xml"`
	Res     *res     `xml:"res"`
}

type res struct {
	XMLName xml.Name `xml:"res"`
	Path    string   `xml:"path,attr"`
}
