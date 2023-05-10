package database

import (
	"fmt"
	Config "go-minimarket/config"
	"net/url"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB(configDb Config.Db) *gorm.DB {

	dbString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		configDb.User,
		configDb.Password,
		configDb.Host,
		configDb.Port,
		configDb.Name,
		url.QueryEscape(configDb.Location))

	Dbconn, err := gorm.Open("mysql", dbString)

	//Dbconn.LogMode(configDb.Debug)
	if err != nil {
		//Logger.LogError("INITIALIZE", err.Error())
		fmt.Print(err.Error())
		return nil
	}

	return Dbconn
}
