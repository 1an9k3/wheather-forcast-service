package predictor

import (
	"log"
	"path"
	"path/filepath"
	"runtime"
)

func getRootPath() string {
	var rootPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		rootPath, err := filepath.Abs(path.Join(path.Dir(filename), "..", ".."))
		if err != nil {
			log.Panicf("get path fail: %v", err)
		} else {
			log.Printf("root path: %v", rootPath)
		}
	}
	return rootPath
}

func getModelFile() {

}
