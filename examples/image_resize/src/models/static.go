package models

import (
	"fmt"
	"github.com/astrata/tango"
	"github.com/astrata/tango/app"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

// staticRoot directory for static files.
var staticRoot = "static"

type Static struct {
}

func init() {
	app.Register("Static", &Static{})
	app.Fallback("/", app.App("Static"))
}

// Checking root directory on start.
func (self *Static) StartUp() {
	info, err := os.Stat(staticRoot)
	if err == nil {
		if info.IsDir() == false {
			panic(fmt.Sprintf("%s is not a directory.\n", staticRoot))
		}
	} else {
		panic(err.Error())
	}
}

// Catches all requests and serves files.
func (self *Static) CatchAll(path ...string) []byte {

	route := staticRoot + tango.PS + strings.Trim(strings.Join(path, tango.PS), tango.PS)

	info, err := os.Stat(route)

	if err == nil {

		if info.IsDir() == true {

			route = strings.Trim(route, tango.PS) + tango.PS + "index.html"

			info, err = os.Stat(route)
		}

		file, err := os.Open(route)

		if err == nil {

			defer file.Close()

			content := make([]byte, info.Size())

			file.Read(content)

			fileType := mime.TypeByExtension(filepath.Ext(route))

			app.Server.Context.SetHeader("Content-Type", fileType)

			return content

		}

	}

	app.Server.Context.HttpError(404)

	return []byte{}
}
