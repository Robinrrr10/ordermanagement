package utils

import (
	"fmt"
	"os"

	"github.com/magiconair/properties"
)

var ServerPort, DBhost, DBport, DBname, DBuser, DBpass string

func ReadProperty() {
	//fmt.Println("Read config values")
	ServerPort = giveValue("server.port", "8080")
	DBhost = giveValue("db.mysql.host", "localhost")
	DBport = giveValue("db.mysql.port", "3306")
	DBname = giveValue("db.mysql.dbname", "business")
	DBuser = giveValue("db.mysql.user", "root")
	DBpass = giveValue("db.mysql.password", "root")
	//fmt.Println(ServerPort, DBhost, DBport, DBname, DBuser, DBpass)
}

func giveValue(key string, def string) (value string) {
	filePath := "config/main.properties"
	p := properties.MustLoadFile(filePath, properties.UTF8)
	value = os.Getenv(key)
	if value == "" {
		value = p.GetString(key, def)
	}
	fmt.Println(key + ":" + value)
	return value
}
