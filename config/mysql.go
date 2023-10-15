package config

import (
	"log"

	_ "github.com/apache/calcite-avatica-go/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func ConnectMySQL() (*gorm.DB, error) {

	connectionString := viper.GetString("MYSQL_USER") + ":" + viper.GetString("MYSQL_PASS") + "@tcp(" + viper.GetString("MYSQL_HOST") + ":" + viper.GetString("MYSQL_PORT") + ")/" + viper.GetString("MYSQL_SCHEMA") + "?parseTime=true"
	mysqlConn, err := gorm.Open("mysql", connectionString)

	if err != nil {
		log.Println("Error connect to MySQL: ", err.Error())
		return nil, err
	}

	mysqlConn.DB().SetMaxIdleConns(2)
	mysqlConn.DB().SetMaxOpenConns(1000)
	mysqlConn.LogMode(true)

	log.Println("MySQL connection success")
	return mysqlConn, nil
}
