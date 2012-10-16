package models

import (
	"github.com/astrata/tango"
	"github.com/astrata/tango/app"
	"github.com/astrata/tango/validation"
	"github.com/astrata/tango/body"
	"github.com/gosexy/sugar"
)

type Data struct {
	Params tango.Value
	Rules *validation.Rules
}

func init() {
	app.Register("Data", &Data{})
	app.Route("/data", app.App("Data"))
}

func (self *Data) StartUp() {
	self.Rules = validation.New()

	self.Rules.Add("email", validation.Email, "Please enter a valid e-mail address.")
	self.Rules.Add("phone", validation.Numeric, "Please enter some numbers.")
}

func (self *Data) Validate() body.Body {

	response := body.Json()

	content := sugar.Tuple{}

	params := self.Params

	isValid, messages := self.Rules.Validate(params)

	if isValid == true {
		content["success"] = "Thank you."
	} else {
		content["error"] = "Some errors were found."
		content["data"] = messages
	}

	response.Set(content)

	return response
}

func (self *Data) Required() body.Body {

	response := body.Json()

	content := sugar.Tuple{}

	params := self.Params

	var isValid bool
	var messages map[string][]string

	isValid, messages = params.Require("email", "name", "last_name")

	if isValid == false {
		content["error"] = "Missing required data."
		content["data"] = messages
	}

	isValid, messages = self.Rules.Validate(params)

	if isValid == true {
		content["success"] = "Thank you."
	} else {
		content["error"] = "Some errors were found."
		content["data"] = messages
	}

	response.Set(content)

	return response
}

