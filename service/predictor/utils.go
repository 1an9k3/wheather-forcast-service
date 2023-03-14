package predictor

import (
	"github.com/fyk1234/wheather-forcast-service/service/config"
	"path"
	"path/filepath"
	"runtime"
)

func getCurrentAbsDir() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

func getRootPath() string {
	var rootPath string
	rootPath = filepath.Join(getCurrentAbsDir(), "..", "..")
	return rootPath
}

func getModelPath() string {
	tomlCfg := config.InitConfig()
	var modelPath string
	modelPath = filepath.Join(getRootPath(), tomlCfg.Predictor.ModelPath)
	return modelPath
}
