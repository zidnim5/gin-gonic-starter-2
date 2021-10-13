package postgres_connection

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var host string
var db_name string 
var username string 
var password string 
var db_port string

/** 
 * TODO: connect to db postgres
 * 
 * @return *gorm.DB
 */
func ConnectDB() *gorm.DB {
	var err error

     host = ""
     db_name = ""
     username = ""
     password = ""
     db_port = ""

	dsn := "host="+ host + " user=" + username + " password=" + password + " dbname=" + db_name + " port=" + db_port + " sslmode=disable  TimeZone=Asia/Jakarta"

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

     // to close connection
     // postgres, err := Db.DB()
     // defer postgres.Close()

     return Db

}
