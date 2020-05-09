package main

import (
    "fmt"

    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

func main() {
    _, err := sqlConnect()
    if err != nil {
        panic(err.Error())
    } else {
        fmt.Println("DB接続成功")
    }
}

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
    DBMS := "mysql"
    USER := "go_example"
    PASS := "12345!"
    PROTOCOL := "tcp(localhost:3306)"
    DBNAME := "go_example"

    CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
    return gorm.Open(DBMS, CONNECT)
}