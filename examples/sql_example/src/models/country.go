package models

import (
	"github.com/gosexy/sugar"
	"database/sql"
	"github.com/astrata/tango/app"
	"github.com/astrata/tango/body"
	"persistent"
)

type Country struct {
	Database *sql.DB
}

func init() {
	app.Register("Country", &Country{})
	app.Route("/country", app.App("Country"))
}

func (self *Country) StartUp() {
	self.Database		= persistent.Database("default")
}

type CountryValue struct {
	Code string
	Name string
}

func (self *Country) Find(code string) body.Body {

	response := body.Json()

	country := CountryValue{}

	row := self.Database.QueryRow(
		"SELECT code, name FROM countries WHERE code = ?",
		code,
	)

	err := row.Scan(&country.Code, &country.Name)

	if err != nil {
		response.Set(
			sugar.Tuple {
				"error": "Not found",
			},
		)
	} else {
		response.Set(
			sugar.Tuple {
				"success": "Found",
				"data": sugar.Tuple{
					"code": country.Code,
					"name": country.Name,
				},
			},
		)
	}

	return response
}

