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
}

var conf = &Conf{}

func init() {
	byts, err := ioutil.ReadFile(AppDirPath("package.json"))
	if err != nil {
		_ = json.Unmarshal(byts, conf)
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
