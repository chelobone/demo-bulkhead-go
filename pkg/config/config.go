package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type appConfig struct {
	HTTPInfo   *HTTPInfo
	MySQLInfo  *MySqlInfo
	ClientInfo *ClientConfig
}

type ClientConfig struct {
	GetById    string
	GetAll     string
	PostClient string
}
type HTTPInfo struct {
	Addr string
}

type MySqlInfo struct {
	MySQLUser     string
	MySQLPassword string
	MySQLAddr     string
	MySQLDBName   string
}

func LoadConfig() *appConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	addr := ":" + os.Getenv("PORT")

	httpInfo := &HTTPInfo{
		Addr: addr,
	}

	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlAddr := os.Getenv("MYSQL_ADDR")
	mysqlDBName := os.Getenv("MYSQL_DATABASE")

	dbInfo := &MySqlInfo{
		MySQLUser:     mysqlUser,
		MySQLPassword: mysqlPassword,
		MySQLAddr:     mysqlAddr,
		MySQLDBName:   mysqlDBName,
	}

	clientInfo := &ClientConfig{
		GetById:    os.Getenv("MYSQL_SP_GET_CLIENTBYID"),
		GetAll:     os.Getenv("MYSQL_SP_GET_CLIENTS"),
		PostClient: os.Getenv("MYSQL_SP_POST_CLIENT"),
	}

	conf := appConfig{
		MySQLInfo:  dbInfo,
		HTTPInfo:   httpInfo,
		ClientInfo: clientInfo,
	}

	return &conf
}
