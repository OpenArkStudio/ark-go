package logSrc

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/ArkNX/ark-go/interface"
	logInterface "github.com/ArkNX/ark-go/plugin/log/interface"
	"github.com/ArkNX/ark-go/util"
)

var (
	logModuleType   = util.GetType((*AFCLogModule)(nil))
	logModuleName   = util.GetName((*AFCLogModule)(nil))
	logModuleUpdate = fmt.Sprintf("%p", (&AFCLogModule{}).Update) != fmt.Sprintf("%p", (&ark.AFCModule{}).Update)
)

func init() {
	logInterface.AFILogModuleName = util.GetName((*AFCLogModule)(nil))
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
