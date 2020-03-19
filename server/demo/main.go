package main

import (
	"github.com/ArkNX/ark-go/app"
	_ "github.com/ArkNX/ark-go/plugin/httpPlugin"
	_ "github.com/ArkNX/ark-go/plugin/logPlugin"
)

func main() {
	app.Start()
}
