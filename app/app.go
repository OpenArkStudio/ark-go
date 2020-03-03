package app

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/spf13/cast"

	"github.com/ArkNX/ark-go/base"
	ark "github.com/ArkNX/ark-go/interface"
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
Github:  https://github.com/ArkNX/ark-go

Version : %s
Branch : %s
CommitID : %s
*************************************************
`

var (
	// version args
	commit  string
	branch  string
	version = "no-version"
	v       bool
	// command line args
	busId      string
	serverName string
	plugin     string
	logPath    string
)

func parseFlags() error {
	flag.StringVar(&busId, "busid", "", "Set application id(like IP address: 8.8.8.8)")
	flag.StringVar(&serverName, "name", "", "Set application name")
	flag.StringVar(&plugin, "plugin", "", "plugin config path")
	flag.StringVar(&logPath, "logpath", "", "Set application log output path")
	flag.BoolVar(&v, "v", false, "show the version")
	flag.Parse()

	// show the version
	if v {
		return nil
	}

	// check the required flags
	for _, name := range []string{"busid", "name", "plugin", "logpath"} {
		found := false
		flag.Visit(func(f *flag.Flag) {
			if f.Name == name {
				found = true
			}
		})

		if !found {
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
		i, err := cast.ToUint8E(str)
		if err != nil {
			return err
		}
		uint8Arr = append(uint8Arr, i)
	}

	ark.GetAFPluginManagerInstance().SetBusID(base.NewAFBusAddr(uint8Arr[0], uint8Arr[1], uint8Arr[2], uint8Arr[3]).BudId)

	// set app name
	ark.GetAFPluginManagerInstance().SetAppName(serverName)

	// set plugin config path
	ark.GetAFPluginManagerInstance().SetPluginConf(plugin)

	// set log path
	ark.GetAFPluginManagerInstance().SetLogPath(logPath)

	return nil
}

func printLogo() {
	fmt.Printf(logo, version, branch, commit)
}

func Start() {
	printLogo()

	if err := parseFlags(); err != nil {
		log.Fatal(err)
	}

	if v {
		return
	}

	if err := ark.GetAFPluginManagerInstance().Start(); err != nil {
		log.Fatal(err)
	}

	defer ark.GetAFPluginManagerInstance().Stop()

	// start server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(time.Millisecond)
	for {
		select {
		case <-sigChan:
			return
		case <-ticker.C:
			ark.GetAFPluginManagerInstance().Update()
		}
	}
}
