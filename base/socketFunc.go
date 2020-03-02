package base

import (
	"errors"
	"regexp"

	"github.com/spf13/cast"
)

type ProtoType string

const (
	protoTypeUnknown ProtoType = "unknown"
	protoTypeTcp     ProtoType = "tcp"
	protoTypeUdp     ProtoType = "udp"
	protoTypeHttp    ProtoType = "http"
	protoTypeHttps   ProtoType = "https"
	protoTypeWs      ProtoType = "ws"
	protoTypeWss     ProtoType = "wss"
)

func ProtoTypeToStr(t ProtoType) string {
	return cast.ToString(t)
}

func StrToProtoType(t string) ProtoType {
	switch t {
	case "tcp":
		return protoTypeTcp
	case "udp":
		return protoTypeUdp
	case "http":
		return protoTypeHttp
	case "https":
		return protoTypeHttps
	case "ws":
		return protoTypeWs
	case "wss":
		return protoTypeWss
	default:
		return protoTypeUnknown
	}
}

// ***** AFEndpoint *****
type AFSocket struct{}

// ***** AFEndpoint *****
type AFEndpoint struct {
	isIpv6 bool
	ext    struct {
		Proto ProtoType
		Ip    string
		Port  uint16
		Path  string
	}
}

func NewFromString(url string) (*AFEndpoint, error) {
	if url == "" {
		return nil, errors.New("AFEndpoint url is empty")
	}

	r, err := regexp.Compile("((.*)://)?([^:/]+)(:(\\d+))?(/.*)?")
	if err != nil {
		return nil, err
	}

	if r.MatchString(url) {
		return nil, errors.New("unmatched url ` " + url + " `")
	}

	strArr := r.FindStringSubmatch(url)

	port, err := cast.ToUint16E(strArr[5])
	if err != nil {
		return nil, err
	}

	ep := &AFEndpoint{
		isIpv6: false,
		ext: struct {
			Proto ProtoType
			Ip    string
			Port  uint16
			Path  string
		}{
			Proto: StrToProtoType(strArr[2]),
			Ip:    strArr[3],
			Port:  port,
			Path:  strArr[6],
		},
	}
	return ep, nil
}

func (a *AFEndpoint) ToString() string {
	var url string
	if a.ext.Proto != protoTypeUnknown {
		url += string(a.ext.Proto)
	}

	url += a.GetIP() + ":" + cast.ToString(a.GetPort()) + a.GetPath()

	return url
}

//******* GET & SET ********
func (a *AFEndpoint) Proto() ProtoType {
	return a.ext.Proto
}

func (a *AFEndpoint) SetProto(proto ProtoType) {
	a.ext.Proto = proto
}

func (a *AFEndpoint) GetIP() string {
	return a.ext.Ip
}

func (a *AFEndpoint) SetIP(ip string) {
	a.ext.Ip = ip
}

func (a *AFEndpoint) GetPath() string {
	return a.ext.Path
}

func (a *AFEndpoint) SetPath(path string) {
	a.ext.Path = path
}

func (a *AFEndpoint) GetPort() uint16 {
	return a.ext.Port
}

func (a *AFEndpoint) SetPort(port uint16) {
	a.ext.Port = port
}

func (a *AFEndpoint) IsV6() bool {
	return a.isIpv6
}

func (a *AFEndpoint) SetIsV6(v6 bool) {
	a.isIpv6 = v6
}
