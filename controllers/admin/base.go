package admin

import (
	"github.com/astaxie/beego"
	"strconv"
	"godoc/models"
)

type BaseController struct {
	beego.Controller
}
var notAuthController = []string{"LoginController"}
func (b *BaseController)Prepare(){
	controllerName,_ := b.GetControllerAndAction()
	isAuth := isAuth(controllerName)
	if isAuth {
		token := b.Ctx.Input.Header("token")
		uid := b.Ctx.Input.Header("uid")
		if token == "" || uid == "" {
			rtnTpl := new(rtnTpl)
			rtnTpl.Code = "5001"
			rtnTpl.Msg = "登陆失效"
			b.Json(rtnTpl)
		}
		id,err := strconv.Atoi(uid)
		if err != nil {
			rtnTpl := new(rtnTpl)
			rtnTpl.Code = "5001"
			rtnTpl.Msg = "登陆失效"
			b.Json(rtnTpl)
		}
		tokenModel := models.GetTokenModelByToken(token)
		if tokenModel == nil {
			rtnTpl := new(rtnTpl)
			rtnTpl.Code = "5001"
			rtnTpl.Msg = "登陆失效"
			b.Json(rtnTpl)
		}
		if tokenModel.Uid != uint(id) {
			rtnTpl := new(rtnTpl)
			rtnTpl.Code = "5001"
			rtnTpl.Msg = "登陆失效"
			b.Json(rtnTpl)
		}
	}
}
func isAuth(controller string) bool{
	for _,v := range notAuthController {
		if v == controller {
			return false
		}
	}
	return true
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