package testdata

import (
	"path"
	"runtime"
)

var TestdataDir = currentPath() + "/"

func currentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
