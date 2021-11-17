package oracle

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/linmingxiao/gneo/logx"
	"log"
	"os"
	"time"
	_ "github.com/mattn/go-oci8"
)



type ConnConfig struct {
	ConnStr string `json:",optional"`
	MasterName string `json:",optional"`
	MaxOpen int    `json:",default=100,range=[10:150]"`
	MaxIdle int    `json:",default=100,range=[10:150]"`
}

type OracleX struct {
	Cli *sql.DB
	Ctx context.Context

}

func NewOracleConn(cf *ConnConfig) *OracleX {
	oracleX := OracleX{Ctx: context.Background()}

	db, err := sql.Open("oci8", cf.ConnStr)
	if err != nil {
		logx.Errorf("Conn oracle %s err: %s", cf.MasterName, err)
	} else {
		logx.Infof("Oracle %s connect successfully.", cf.MasterName)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 60 * 24 * 30)
	db.SetMaxOpenConns(cf.MaxOpen)
	db.SetMaxIdleConns(cf.MaxIdle)
	oracleX.Cli = db
	return &oracleX
}



//示例方法，请勿调用
func Examplequery() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8")
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	db, err := sql.Open("oci8", "hssale/gmsale@192.168.19.21:1521/dspdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from v$version")
	if err != nil {
		log.Fatal(err)
	}
	cols, _ := rows.Columns()
	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))
	dest := make([]interface{}, len(cols))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}
	for rows.Next() {
		err = rows.Scan(dest...)
		for i, raw := range rawResult {
			if raw == nil {
				result[i] = ""
			} else {
				result[i] = string(raw)
			}
		}
		fmt.Printf("%s\n", result[0])
	}
	rows.Close()
}
