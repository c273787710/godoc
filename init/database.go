package init

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func DBinit(){
	dbType := beego.AppConfig.String("db_type")
	dbAlias := beego.AppConfig.String("db_alias")
	dbHost := beego.AppConfig.String("db_host")
	dbPort := beego.AppConfig.String("db_port")
	dbUser := beego.AppConfig.String("db_user")
	dbPwd := beego.AppConfig.String("db_pwd")
	dbName := beego.AppConfig.String("db_name")
	dbCharset := beego.AppConfig.String("db_charset")
	dataSource := dbUser + ":" + dbPwd + "@tcp("+ dbHost +":"+ dbPort +")/"+ dbName + "?charset=" + dbCharset
	err := orm.RegisterDataBase(dbAlias,dbType,dataSource)
	if err != nil{
		panic("数据库连接失败："+err.Error())
	}
}
