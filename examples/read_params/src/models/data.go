package models

import (
	"github.com/astrata/tango"
	"github.com/astrata/tango/app"
)

type Data struct {
	Params tango.Value
	Files tango.Files
}

func init() {
	app.Register("Data", &Data{})
	app.Route("/data", app.App("Data"))
}

func (self *Data) StartUp() {

}

func (self *Data) Debug(url ...string) map[string] interface{} {

	response := map[string] interface{}{}

	response["test_url"] = url

	response["test_text"] = self.Params.GetString("test_text")

	response["test_text_get"] = self.Params.GetString("test_text_get")

	file := self.Files.GetFile("test_file")

	if file != nil {
		response["test_file(name)"] = file.Name
	}

	return response
}
