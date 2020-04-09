package admin

import (
	"doc/models"
	"encoding/json"
)

type LoginController struct {
	BaseController
}

//@Title Login
//@Description 超级管理员后台登录
//@Param username body string true "登录名"
//@Param password body string true "密码"
//@Success 200
//@router /login [post]
func (l *LoginController)Login(){
	rtn := new(rtnTpl)
	paramTpl := new(models.AdminTpl)
	err := json.Unmarshal(l.Ctx.Input.RequestBody,paramTpl)
	if err != nil {
		rtn.Code = "1000"
		rtn.Msg = "用户名或密码缺失"
		l.Json(rtn)
	}
	//验证参数
	if err := paramTpl.Check();err != nil {
		rtn.Code = "1001"
		rtn.Msg = err.Error()
		l.Json(rtn)
	}
	//验证用户名密码是否正确
	model := new(models.AdminModel)
	model.Username = paramTpl.Username
	if err := model.CheckLoginByUsername(paramTpl.Password);err!=nil{
		rtn.Code = "1002"
		rtn.Msg = err.Error()
		l.Json(rtn)
	}
	//登录成功
	err = model.CreateToken()
	if err != nil {
		rtn.Code = "1003"
		rtn.Msg = err.Error()
		l.Json(rtn)
	}
	rtn.Code = "0"
	rtn.Msg = "登录成功"
	rtn.Data = model
	l.Json(rtn)
}

