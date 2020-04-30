package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

func init(){
	orm.RegisterModel(new(AdminModel))
	orm.RegisterModel(new(TokenModel))
	orm.RegisterModel(new(CateModel))
	orm.RegisterModel(new(ArticleModel))
}

func GetTableName(name string)string{
	return beego.AppConfig.String("db_prefix") + name
}

type BaseTpl struct {
	validation.Validation
}
