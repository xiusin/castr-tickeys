package helper

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

type Conf struct {
	SoundType string `json:"sound_type"`
	Style     string `json:"style"`
	Pos       [2]int `json:"pos"`
	Delay     int    `json:"delay"` // 延迟刷新时间
}

var internalConf = Conf{
	SoundType: "bubble",
	Style:     "color: #fff; border: 2px solid #fff; border-radius: 0px; background: rgba(0,0,0,0.4); font-size: 30px;",
	Pos:       [2]int{15, -35},
	Delay:     2000,
}

var conf = &Conf{}

func init() {
	byts, err := ioutil.ReadFile(AppDirPath("package.json"))
	if err == nil {
		err = json.Unmarshal(byts, conf)
		if err != nil {
			conf = &internalConf
		}
	} else {
		conf = &internalConf
	}
}

func getCurrentDir() string {
	workingDir, _ := os.Getwd()
	if runtime.GOOS == "darwin" {
		workingDir, _ = os.Executable()
		workingDir = filepath.Dir(workingDir)
	}
	return workingDir
}

func AppDirPath(name string) string {
	return filepath.Join(getCurrentDir(), name)
}

func GetConf() *Conf {
	return conf
}
