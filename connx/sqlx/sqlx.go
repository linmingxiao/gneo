/**
可同时处理 Mysql Oracle 相关链接操作
 */
package sqlx

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //mysql
	"github.com/jmoiron/sqlx"
	"github.com/linmingxiao/gneo/logx"
	_ "github.com/mattn/go-oci8" //oracle
	"time"
)

type ConnConfig struct {
	ConnStr string `json:",optional"`
	MasterName string `json:",optional"`
	Type string `json:",default=mysql,options=mysql|oracle"`
	MaxOpen int    `json:",default=100,range=[10:150]"`
	MaxIdle int    `json:",default=100,range=[10:150]"`
}

type SqlX struct {
	Cli *sqlx.DB
	Ctx context.Context
}

func NewSqlXConn(cf *ConnConfig) *SqlX {
	sqlX := SqlX{Ctx: context.Background()}
	var db *sqlx.DB
	var err error
	if cf.Type == "mysql"{
		db, err = sqlx.Connect("mysql", cf.ConnStr)
	} else if cf.Type == "oracle" {
		db, err = sqlx.Connect("oci8", cf.ConnStr)
	} else {
		panic(fmt.Sprintf("SqlX connect type error: %s", cf.Type))
	}
	if err != nil {
		logx.Errorf("Conn %s--%s err: %s",cf.Type, cf.MasterName, err)
	} else {
		logx.Infof("%s %s connect successfully.",cf.Type, cf.MasterName)
	}

	//SQL惰性连接池问题，这里Open成功之后需要ping一次，确保Mysql连接成功
	if errPing := db.Ping(); errPing!= nil{
		logx.Errorf("DB %s ping failed:", cf.MasterName)
		logx.Error(errPing)
		panic(errPing)
	} else {
		logx.Debugf("DB %s ping successfully.", cf.MasterName)
	}

	//See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 60 * 24 * 30)
	db.SetMaxOpenConns(cf.MaxOpen)
	db.SetMaxIdleConns(cf.MaxIdle)

	sqlX.Cli = db
	return &sqlX
}




