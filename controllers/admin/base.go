package admin

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

type rtnTpl struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func (b *BaseController)Json(r *rtnTpl){
	b.Data["json"] = r
	b.ServeJSON()
	b.StopRun()
}