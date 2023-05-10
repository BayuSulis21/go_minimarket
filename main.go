package main

import (
	"go-minimarket/config"
	"go-minimarket/database"
	"go-minimarket/services"
	"go-minimarket/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var DbConn *gorm.DB

func main() {
	// koneksi db
	conf := config.Config
	DbConn = database.InitDB(conf.Db)
	//export var public ke subfolder services
	services.DbConn = DbConn

	router := gin.Default()
	router.Use(utils.LoggerMiddleware(utils.Logger))

	router.GET("/generatedSignature", services.GeneratedSignatureHandler)
	router.GET("/barang", services.ListBarangHandler)
	router.GET("/search_barang", services.SearchBarangHandler)
	router.POST("/register", services.RegisterHandler)
	router.POST("/sendWA", services.SendWAHandler)
	router.Run(":8888")
}
