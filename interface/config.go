package ark

import "encoding/xml"

type pluginConf struct {
	XMLName xml.Name `xml:"xml"`
	Plugins *plugins `xml:"plugins"`
	Res     *res     `xml:"res"`
}

type plugins struct {
	XMLName xml.Name  `xml:"plugins"`
	Path    string    `xml:"path,attr"`
	Plugin  []*plugin `xml:"plugin"`
}

type plugin struct {
	XMLName xml.Name `xml:"plugin"`
	Name    string   `xml:"name,attr"`
}

type res struct {
	XMLName xml.Name `xml:"res"`
	Path    string   `xml:"path,attr"`
}
