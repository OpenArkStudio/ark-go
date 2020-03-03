package main

import (
	"github.com/ArkNX/ark-go/app"
	_ "github.com/ArkNX/ark-go/plugin/http"
	_ "github.com/ArkNX/ark-go/plugin/log"
)

func main() {
	app.Start()
}
