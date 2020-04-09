package models

import (
	"doc/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type TokenModel struct {
	Id uint `orm:"column(id)"`
	Value string `orm:"column(value)"`
	Uid uint `orm:"column(uid)"`
	ExpireTime uint `orm:"column(expire_time)"`
	AddTime uint `orm:"column(add_time)"`
	UpdateTime uint `orm:"column(update_time)"`
}
var syncToken = make(map[string]*TokenModel)

func (t *TokenModel)TableName()string{
	return GetTableName("token")
}

func CreateToken(uid uint)(string,error){
	model := new(TokenModel)
	o := orm.NewOrm()
	//判断是否存在未过期的token
	err := o.QueryTable(GetTableName("token")).Filter("uid",uid).Filter("expire_time__gte",time.Now().Unix()).One(model)
	if err == nil {
		//不存在
		model.ExpireTime = uint(time.Now().Unix()) + 86400
		model.UpdateTime = uint(time.Now().Unix())
		_,_ = o.Update(model)
		syncToken[model.Value] = model
		return model.Value,nil
	}
	s := fmt.Sprintf("solt=%s;time=%d;uid=%d",beego.AppConfig.String("admin_solt"),time.Now().Unix(),uid)
	model.Id = 0
	model.Value = util.MD5Encry(util.MD5Encry(s))
	model.Uid = uid
	model.ExpireTime = uint(time.Now().Unix()) + 86400
	model.AddTime = uint(time.Now().Unix())
	model.UpdateTime = uint(time.Now().Unix())
	_,err = o.Insert(model)
	if err != nil {
		return "",err
	}
	syncToken[model.Value] = model
	return model.Value,nil
}