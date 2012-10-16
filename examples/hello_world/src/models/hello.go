package models

import (
	"github.com/astrata/tango/app"
)

type Hello struct {
}

func init() {
	app.Register("Hello", &Hello{})
	app.Route("/", app.App("Hello"))
}

func (self *Hello) StartUp() {

}

func (self *Hello) Index() string {
	return "Hello, world!\n"
}

