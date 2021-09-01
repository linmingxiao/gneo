package mysql

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type ConnConfig struct {
	ConnStr string `json:",optional"`
	MaxOpen int    `json:",default=100,range=[10:1000]"`
	MaxIdle int    `json:",optional"`
}

type MSqlX struct {
	Cli *sql.DB
	Ctx context.Context
}

func NewMysqlConn(cf *ConnConfig) *MSqlX {
	mysqlX := MSqlX{Ctx: context.Background()}

	db, err := sql.Open("mysql", cf.ConnStr)
	if err != nil {
		log.Fatalf("Conn %s err: %s", cf.ConnStr, err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(cf.MaxOpen)
	db.SetMaxIdleConns(cf.MaxIdle)

	mysqlX.Cli = db
	return &mysqlX
}
