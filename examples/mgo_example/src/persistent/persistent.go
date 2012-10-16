package persistent

import (
	"github.com/astrata/tango/datasource"
	"labix.org/v2/mgo"
	"fmt"
)

var sess *mgo.Session

func Database(name string) *mgo.Database {
	if sess == nil {
		var err error
		var url string
		_, settings := datasource.Config(name)
		if settings.Port == 0 {
			settings.Port = 27017
		}
		if settings.Database == "" {
			panic("mgo: Missing database name.")
		}
		if settings.User != "" && settings.Password != "" {
			url = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", settings.User, settings.Password, settings.Host, settings.Port, settings.Database)
		} else {
			url = fmt.Sprintf("mongodb://%s:%d/%s", settings.Host, settings.Port, settings.Database)
		}
		fmt.Printf("url: %s\n", url)
		sess, err = mgo.Dial(url)
		if err != nil {
			panic(fmt.Sprintf("mgo: %s", err.Error()))
		}
	}
	return sess.DB("")
}
