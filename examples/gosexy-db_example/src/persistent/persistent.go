package persistent

import (
	"github.com/gosexy/db"
	"fmt"
	_ "github.com/gosexy/db/mongo"
	"github.com/astrata/tango/datasource"
)

var sess = make(map[string]db.Database)

func Database(name string) db.Database {
	if _, ok := sess[name]; ok == false {
		driver, settings := datasource.Config(name)
		sess[name] = db.Open(driver, settings)
		if sess[name] == nil {
			panic(fmt.Sprintf("resource: Cannot open resource %s.", name))
		}
	}
	return sess[name]
}
