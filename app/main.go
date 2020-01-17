package main

import (
	"ark-go/model"
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"

	"ark-go/common"
	"ark-go/util/convert"
)

const logo = `
*************************************************
       _         _    
      / \   _ __| | __
     / _ \ | '__| |/ /
    / ___ \| |  |   < 
   /_/   \_\_|  |_|\_\

Copyright 2019 (c) ArkNX. All Rights Reserved.
Website: https://arknx.com
Github:  https://github.com/ArkNX
*************************************************
`

var (
	busId      = ""
	serverName = ""
	plugin     = ""
	logPath    = ""
)

func isFlagParsed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func parseFlags() error {
	flag.StringVar(&busId, "busid", "", "Set application id(like IP address: 8.8.8.8)")
	flag.StringVar(&serverName, "name", "", "Set application name")
	flag.StringVar(&plugin, "plugin", "", "plugin config path")
	flag.StringVar(&logPath, "logpath", "", "Set application log output path")
	flag.Parse()

	// check the required flags
	for _, name := range []string{"busid", "name", "plugin", "logpath"} {
		if !isFlagParsed(name) {
			return errors.New("flag ` " + name + " ` is absent")
		}
	}

	// parse bus id
	strArr := strings.Split(busId, ".")
	if len(strArr) != 4 {
		return errors.New("Bus id ` " + busId + " ` is invalid, it likes 8.8.8.8")
	}

	var uint8Arr []uint8
	for _, str := range strArr {
		i, err := convert.Uint8(str)
		if err != nil {
			return err
		}
		uint8Arr = append(uint8Arr, i)
	}

	model.GetAFPluginManagerInstance().SetBusID(common.NewAFBusAddr(uint8Arr[0], uint8Arr[1], uint8Arr[2], uint8Arr[3]).BudId)

	// set app name
	model.GetAFPluginManagerInstance().SetAppName(serverName)

	// set plugin config path
	model.GetAFPluginManagerInstance().SetAppName(plugin)

	// set log path
	model.GetAFPluginManagerInstance().SetAppName(logPath)

	return nil
}

func printLogo() {
	fmt.Println(logo)
}
func main() {
	if err := parseFlags(); err != nil {
		log.Fatal(err)
	}

	printLogo()

	// TODO: delete
	temp := common.NewAFBusAddrFromInt(16909060)
	fmt.Println(temp.ToString())
	fmt.Println(temp.BudId)

	temp = common.NewAFBusAddr(1, 2, 3, 4)
	fmt.Println(temp.ToString())
	fmt.Println(temp.BudId)

	temp = &common.AFBusAddr{}
	temp.FromString("9.9.9.9")
	fmt.Println(temp.ToString())
	fmt.Println(temp.BudId)
}
