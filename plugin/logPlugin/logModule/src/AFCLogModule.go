package src

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"reflect"
	"runtime"

	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/logPlugin/logModule"
)

func init() {
	logModule.ModuleName = ark.GetName((*AFCLogModule)(nil))
	logModule.ModuleType = ark.GetType((*AFCLogModule)(nil))
	logModule.ModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&AFCLogModule{}).Update).Pointer()).Name()
}

type AFCLogModule struct {
	ark.AFCModule
	// other data
	logger *logrus.Logger
}

func (logModule *AFCLogModule) Init() error {
	logModule.logger = &logrus.Logger{
		Out:          os.Stdout,
		Formatter:    &logrus.JSONFormatter{},
		ReportCaller: true,
		Level:        logrus.WarnLevel,
	}

	return nil
}

// ------------------- logrus options -------------------
func (logModule *AFCLogModule) SetFormatter(formatter logrus.Formatter) {
	logModule.logger.SetFormatter(formatter)
}

func (logModule *AFCLogModule) SetOutput(out io.Writer) {
	logModule.logger.SetOutput(out)
}

func (logModule *AFCLogModule) SetReportCaller(include bool) {
	logModule.logger.SetReportCaller(include)
}

func (logModule *AFCLogModule) SetLevel(level logrus.Level) {
	logModule.logger.SetLevel(level)
}

func (logModule *AFCLogModule) AddHook(hook logrus.Hook) {
	logModule.logger.AddHook(hook)
}

func (logModule *AFCLogModule) GetLogger() *logrus.Logger {
	return logModule.logger
}
