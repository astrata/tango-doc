package persistent

import (
	"github.com/astrata/tango/datasource"
	"database/sql"
	_ "github.com/ziutek/mymysql/godrv"
	"fmt"
)

var sess *sql.DB

func Database(name string) *sql.DB {
	if sess == nil {
		var err error
		var url string

		_, settings := datasource.Config(name)

		if settings.Port == 0 {
			settings.Port = 3306
		}

		if settings.Database == "" {
			panic("sql: Missing database name.")
		}

		if settings.User != "" && settings.Password != "" {
			url = fmt.Sprintf("tcp:%s:%d*%s/%s/%s", settings.Host, settings.Port, settings.Database, settings.User, settings.Password)
		} else {
			url = fmt.Sprintf("tcp:%s:%d*%s", settings.Host, settings.Port, settings.Database)
		}

		sess, err = sql.Open("mymysql", url)

		if err != nil {
			panic(fmt.Sprintf("sql: %s", err.Error()))
		}
	}
	return sess
}
