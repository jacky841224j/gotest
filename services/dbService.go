package services

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	UserName     string = "root"
	Password     string = "test123"
	Addr         string = "127.0.0.1"
	Port         int    = 3306
	Database     string = "DevDb"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

var DB *sql.DB

func Open() {
	//組合sql連線字串
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", UserName, Password, Addr, Port, Database)
	//連接MySQL
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}
	DB.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	DB.SetMaxOpenConns(MaxOpenConns)
	DB.SetMaxIdleConns(MaxIdleConns)
}

func Create(sqlcommnd string, data interface{}) (int64, error) {
	//組合sql連線字串
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", UserName, Password, Addr, Port, Database)
	//連接MySQL
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return 0, err
	}
	DB.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	DB.SetMaxOpenConns(MaxOpenConns)
	DB.SetMaxIdleConns(MaxIdleConns)

	row, err := DB.Exec(sqlcommnd)

	if err != nil {
		return 0, err
	}

	result, err := row.RowsAffected()
	if err != nil {
		return 0, err
	}

	return result, nil
}
