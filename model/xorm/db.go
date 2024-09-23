package xorm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/gotoeasy/glang/cmn"
	_ "github.com/lib/pq"
	"golangchain/pkg/settings"
)

var Db *xorm.Engine
var DbConfig = settings.AppConfig.Db

func Start() {
	var err error
	//打开数据库
	//DSN数据源字符串：用户名:密码@协议(地址:端口)/数据库?参数=参数值
	Db, err = xorm.NewEngine(DbConfig.DriverName, DbConfig.DBUrl)
	if err != nil {
		cmn.Error(err)
	}
	Db.ShowSQL(false)
	//Db.CreateTables(UserRule{})
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
}
