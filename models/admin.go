package models

import (
	"doc/util"
	"errors"
	"github.com/astaxie/beego/orm"
)

type AdminModel struct {
	Id uint `orm:"column(id)" json:"id"`
	Username string  `orm:"column(username)" json:"username"`
	Password string `orm:"column(password)" json:"-"`
	LastLoginIp string `orm:"column(last_login_ip)" json:"-"`
	LastLoginTime uint `orm:"column(last_login_time)" json:"-"`
	Nickname string `orm:"column(nickname)" json:"nickname"`
	Img string `orm:"column(img)" json:"img"`
	Content string `orm:"column(content)" json:"content"`
	Desc string `orm:"column(desc)" json:"desc"`
	Token string `orm:"-" json:"token"`
}
func(a *AdminModel)TableName()string{
	return GetTableName("admin")
}

func (a *AdminModel)CheckLoginByUsername(pass string)error{
	o := orm.NewOrm()
	err := o.QueryTable(GetTableName("admin")).Filter("username",a.Username).One(a)
	if err == orm.ErrNoRows{
		//不存在该用户
		return errors.New("用户错误")
	}
	if err == orm.ErrMultiRows {
		return errors.New("不用不存在")
	}
	if err != nil {
		return errors.New("登录失败")
	}
	if _pass := util.MD5Encry(pass);_pass!= a.Password {
		return errors.New("密码错误")
	}
	return nil
}
func (a *AdminModel)CreateToken()error{
	token,err := CreateToken(a.Id)
	if err != nil {
		return errors.New("登录失败")
	}
	a.Token = token
	return nil
}


//登录用户接受的参数
type AdminTpl struct {
	BaseTpl
	Username string `valid:"Required;MinSize(5);MaxSize(20)"`
	Password string `valid:"Required;MinSize(6);MaxSize(20)"`
}
func (a *AdminTpl)Check()error{
	result,err := a.Valid(a)
	if err != nil || !result {
		return errors.New("参数传递错误")
	}
	return nil
}
