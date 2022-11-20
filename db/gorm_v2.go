package db

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
	"time"
)

func NewDbConn(dbName string) (conn *gorm.DB, sqlDb *sql.DB) {
	mysqlConf := viper.GetStringMapString("mysql")
	connection := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s`,
		mysqlConf["user"], mysqlConf["password"], mysqlConf["host"], mysqlConf["port"], mysqlConf[dbName])
	val := url.Values{}
	val.Set("charset", "utf8mb4")
	val.Add("parseTime", "True")
	val.Add("loc", "Asia/Shanghai")
	dsn := fmt.Sprintf(`%s?%s`, connection, val.Encode())

	fmt.Println(`mysql dsn:`, dsn)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println(`数据库连接异常：`, err)
		panic(err)
	}

	sqlDB, err := conn.DB()
	if err != nil {
		fmt.Println(`数据库连接池异常`, err)
		panic(err)
	}
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(2 * time.Minute)

	return conn, sqlDb
}

func NewIcemTestDbConn(dbName string) (conn *gorm.DB) {
	mysqlConf := map[string]string{
		"user":     "root",
		"password": "windows2010..",
		"host":     "test.icemdata.com",
		"port":     "3306",
	}
	connection := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s`,
		mysqlConf["user"], mysqlConf["password"], mysqlConf["host"], mysqlConf["port"], dbName)
	val := url.Values{}
	val.Set("charset", "utf8mb4")
	val.Add("parseTime", "True")
	val.Add("loc", "Asia/Shanghai")
	dsn := fmt.Sprintf(`%s?%s`, connection, val.Encode())

	fmt.Println(`mysql dsn:`, dsn)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println(`数据库连接异常：`, err)
		panic(err)
	}

	return conn
}

func NewIcemProdDbConn(dbName string) (conn *gorm.DB) {
	mysqlConf := map[string]string{
		"user":     "icem",
		"password": "7L2o66FRtGvOqmAZ",
		"host":     "sh-cynosdbmysql-grp-kehpn5lk.sql.tencentcdb.com",
		"port":     "20668",
	}
	connection := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s`,
		mysqlConf["user"], mysqlConf["password"], mysqlConf["host"], mysqlConf["port"], dbName)
	val := url.Values{}
	val.Set("charset", "utf8mb4")
	val.Add("parseTime", "True")
	val.Add("loc", "Asia/Shanghai")
	dsn := fmt.Sprintf(`%s?%s`, connection, val.Encode())

	fmt.Println(`mysql dsn:`, dsn)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println(`数据库连接异常：`, err)
		panic(err)
	}

	return conn
}
