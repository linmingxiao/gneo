package mysql

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/linmingxiao/gneo/logx"
	"time"
	_ "fmt"
)

type ConnConfig struct {
	ConnStr string `json:",optional"`
	MasterName string `json:",optional"`
	MaxOpen int    `json:",default=100,range=[10:150]"`
	MaxIdle int    `json:",default=100,range=[10:150]"`
}

type MSqlX struct {
	Cli *sql.DB
	Ctx context.Context
}

func NewMysqlConn(cf *ConnConfig) *MSqlX {
	mysqlX := MSqlX{Ctx: context.Background()}

	db, err := sql.Open("mysql", cf.ConnStr)
	if err != nil {
		logx.Errorf("Conn %s err: %s", cf.MasterName, err)
	} else {
		logx.Infof("Mysql %s connect successfully.", cf.MasterName)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 60 * 24 * 30)
	db.SetMaxOpenConns(cf.MaxOpen)
	db.SetMaxIdleConns(cf.MaxIdle)

	mysqlX.Cli = db
	return &mysqlX
}



