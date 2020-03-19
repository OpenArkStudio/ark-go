package logModule

import (
	ark "github.com/ArkNX/ark-go/interface"
	"github.com/sirupsen/logrus"
	"io"
	"reflect"
)

var (
	ModuleName   string
	ModuleType   reflect.Type
	ModuleUpdate string
)

type AFILogModule interface {
	ark.AFIModule
	// logrus setting
	SetFormatter(formatter logrus.Formatter)
	SetOutput(out io.Writer)
	SetReportCaller(include bool)
	SetLevel(level logrus.Level)
	AddHook(hook logrus.Hook)
	// logrus logger
	GetLogger() *logrus.Logger
}
