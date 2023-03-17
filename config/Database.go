package config

import (
	"Product/app/Models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type config struct {
	db_connection string
	db_name       string
	db_host       string
	db_port       string
	db_username   string
	db_password   string
}

func Init() *gorm.DB {

	InitialDB, _ := gorm.Open(connectionMap().db_connection, assembleConfig())

	InitialDB.AutoMigrate(&Models.Products{})

	return InitialDB
}

func connectionMap() config {

	conf := config{
		db_connection: "root",
		db_name:       "railway",
		db_host:       "containers-us-west-17.railway.app",
		db_port:       "7059",
		db_username:   "root",
		db_password:   "xUKdpIqo3UZ7NiTy7lao",
	}

	return conf
}

func assembleConfig() string {

	conf := connectionMap().db_username + ":" +
		connectionMap().db_password + "@(" +
		connectionMap().db_host + ":" +
		connectionMap().db_port + ")/" +
		connectionMap().db_name + "?" +
		"parseTime=true"

	return conf
}
