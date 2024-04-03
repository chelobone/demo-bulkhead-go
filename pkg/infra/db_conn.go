package infra

import (
	"fmt"
	"log"

	"github.com/chelobone/demo_bulkhead_go/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConnector struct {
	Conn *gorm.DB
}

func NewConfig() *config.ClientConfig {
	return config.LoadConfig().ClientInfo
}

func NewMySQLConnector() *MySQLConnector {
	conf := config.LoadConfig()

	dsn := mysqlConnInfo(*conf.MySQLInfo)
	log.Println(dsn)
	//conn, err := sql.Open(driverName, dsn)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &MySQLConnector{
		Conn: conn,
	}
}

func mysqlConnInfo(mysqlInfo config.MySqlInfo) string {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		mysqlInfo.MySQLUser,
		mysqlInfo.MySQLPassword,
		mysqlInfo.MySQLAddr,
		mysqlInfo.MySQLDBName)

	return dataSourceName
}
