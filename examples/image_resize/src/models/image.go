package models

import (
	"crypto"
	"fmt"
	"github.com/astrata/tango"
	"github.com/astrata/tango/app"
	"github.com/gosexy/canvas"
	"github.com/gosexy/checksum"
	"github.com/gosexy/resource"
	"github.com/gosexy/to"
	"os"
	"path"
	"strconv"
	"strings"
)

var ImageRoot = "thumbnails"

type Image struct {
	Params tango.Value
}

func init() {
	app.Register("Image", &Image{})
	app.Route("/image", app.App("Image"))
}

func (self *Image) StartUp() {
	resource.Root = "temp" + tango.PS + "downloads"
}

func (self *Image) Resize(size string) {
	var err error

	url := to.String(self.Params.Get("url"))

	relPath := ImageRoot + tango.PS + size + tango.PS + checksum.String(fmt.Sprintf("%s/%s", size, url), crypto.SHA1) + ".png"
	fullPath := Root + tango.PS + relPath

	_, err = os.Stat(fullPath)

	if err == nil {

		app.Server.Context.Redirect("/" + relPath)

		app.Server.Context.HttpError(200)

		return

	} else {

		filePath, err := resource.Download(url)

		if err == nil {

			thumb := canvas.New()

			opened := thumb.Open(filePath)

			if opened == true {

				resize := strings.Split(size, "x")

				width, _ := strconv.Atoi(resize[0])
				height, _ := strconv.Atoi(resize[1])

				thumb.AutoOrientate()
				thumb.Thumbnail(uint(width), uint(height))

				os.MkdirAll(path.Dir(fullPath), os.ModeDir|0755)

				written := thumb.Write(fullPath)

				if written {
					app.Server.Context.Redirect("/" + relPath)
				} else {
					app.Server.Context.HttpError(500)
				}

				return
			}

		}

	}

	app.Server.Context.HttpError(404)
}
