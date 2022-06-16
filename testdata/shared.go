package testdata

import (
	"path"
	"runtime"
)

var HtmlExt = ".html"
var JsonExt = ".json"
var NewExt = ".new"

var TestdataDir = currentPath() + "/"

func currentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
