package admin

import (
	"godoc/models"
	"encoding/json"
	"strconv"
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

//@Title Login
//@Description 通过token和uid获取账号信息
//@Param token header string true "token"
//@Param uid header string true "用户uid"
//@Success 200
//@router /login [get]
func (l *LoginController)GetInfo(){
	rtn := new(rtnTpl)
	uidString := l.Ctx.Request.Header.Get("uid")
	token := l.Ctx.Request.Header.Get("token")
	if uidString == "" || token == "" {
		rtn.Code = "1004"
		rtn.Msg = "用户登陆失效"
		l.Json(rtn)
	}
	uid,err := strconv.Atoi(uidString)
	if err != nil {
		rtn.Code = "1004"
		rtn.Msg = "登陆信息错误"
		l.Json(rtn)
	}
	model := models.GetTokenModelByToken(token)
	if model == nil {
		rtn.Code = "1005"
		rtn.Msg = "用户登陆失效"
		l.Json(rtn)
	}
	if uint(uid) != model.Uid {
		rtn.Code = "1006"
		rtn.Msg = "TOKEN无效"
		l.Json(rtn)
	}
	admin := models.GetAdminInfoById(model.Uid)
	if admin == nil {
		rtn.Code = "1007"
		rtn.Msg = "TOKEN无效"
		l.Json(rtn)
	}
	admin.Token = token
	rtn.Code = "0"
	rtn.Msg = "登录成功"
	rtn.Data = admin
	l.Json(rtn)
}
