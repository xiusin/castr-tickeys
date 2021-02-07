package main

import (
	"github.com/xiusin/castr-tickeys/components"
	"github.com/xiusin/castr-tickeys/helper"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"github.com/xiusin/logger"
)

func init() {
	loggerName := filepath.Join(helper.AppDirPath("logger.log"))
	f, _ := os.OpenFile(loggerName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	logger.SetOutput(f)
	logger.SetLogLevel(logger.DebugLevel)
	logger.SetReportCaller(true)
	rand.Seed(time.Now().UnixNano())
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(err)
		}
	}()
	app := widgets.NewQApplication(len(os.Args), os.Args)
	app.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	app.SetStyle(widgets.QStyleFactory_Create("Funsion"))
	app.SetQuitOnLastWindowClosed(false)
	components.InitKeyboard(app)
	app.Exec()
}
