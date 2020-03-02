package logInterface

import (
	"io"

	"github.com/sirupsen/logrus"

	ark "ark-go/interface"
)

var AFILogModuleName string

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
