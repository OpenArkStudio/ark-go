package common

import (
	"ark-go/util/convert"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
	"time"
)

// bus address, like IP address, 8.8.8.8
type AFBusAddr struct {
	ChannelId uint8
	ZoneId    uint8
	AppType   uint8
	InstId    uint8
	BudId     int
}

func NewAFBusAddrFromInt(id int) *AFBusAddr {
	return &AFBusAddr{
		ChannelId: uint8(id >> 24),
		ZoneId:    uint8(id >> 16),
		AppType:   uint8(id >> 8),
		InstId:    uint8(id),
		BudId:     id,
	}
}

func NewAFBusAddr(cId uint8, zId uint8, pId uint8, iId uint8) *AFBusAddr {
	return &AFBusAddr{
		ChannelId: cId,
		ZoneId:    zId,
		AppType:   pId,
		InstId:    iId,
		BudId:     int(binary.BigEndian.Uint32([]uint8{cId, zId, pId, iId})),
	}
}

func (a *AFBusAddr) ToString() string {
	return fmt.Sprintf("%d.%d.%d.%d", a.ChannelId, a.ZoneId, a.AppType, a.InstId)
}

func (a *AFBusAddr) FromString(busName string) error {
	if busName == "" {
		return errors.New("bus name is empty")
	}

	strs := strings.Split(busName, ".")
	if len(strs) != 4 {
		return errors.New("Bus id ` " + busName + " ` is invalid, it likes 8.8.8.8")
	}

	var uint8Arr []uint8
	for _, str := range []string{strs[0], strs[1], strs[2], strs[3]} {
		i, err := convert.Uint8(str)
		if err != nil {
			return err
		}
		uint8Arr = append(uint8Arr, i)
	}

	a.ChannelId = uint8Arr[0]
	a.ZoneId = uint8Arr[1]
	a.AppType = uint8Arr[2]
	a.InstId = uint8Arr[3]
	a.BudId = int(binary.BigEndian.Uint32([]uint8{uint8Arr[0], uint8Arr[1], uint8Arr[2], uint8Arr[3]}))

	return nil
}

// bus relation, app connect other app with direct way or waiting sync message
type AFBusRelation struct {
	AppType        uint8
	TargetAppType  uint8
	ConnectionType bool
}

type AFProcConfig struct {
	BusId         int
	MaxConnection uint32
	ThreadNum     uint8
	IntranetEp    AFEndpoint
	ServerEp      AFEndpoint
	// to add other fieldss
}

type AFRegCenter struct {
	Ip            string
	Port          uint16
	ServiceName   string
	CheckInterval time.Duration
	CheckTimeout  time.Duration
}

type AFAppConfig struct {
	RegCenter           AFRegCenter
	Name2types          map[string]ARKAppType
	Type2names          map[ARKAppType]string
	ConnectionRelations map[ARKAppType][]ARKAppType
	SelfProc            AFProcConfig
}
