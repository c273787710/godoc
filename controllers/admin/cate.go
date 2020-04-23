package admin

import (
	"godoc/models"
	"encoding/json"
	"time"
	"godoc/validation"
)

type CateController struct {
	BaseController
}

//@Title Add
//@Description 添加文章分类
//@Success 200
//@router /add [post]
func (c *CateController)Add(){
	rtnTpl := new(rtnTpl)
	paramTpl := new(validation.CateValidate)
	err := json.Unmarshal(c.Ctx.Input.RequestBody,paramTpl)
	if err != nil {
		rtnTpl.Code = "2000"
		rtnTpl.Msg = err.Error()
		c.Json(rtnTpl)
	}
	if err := paramTpl.ValidCate();err != nil {
		rtnTpl.Code = "2000"
		rtnTpl.Msg = err.Error()
		c.Json(rtnTpl)
	}
	model := new(models.CateModel)
	model.CateName = paramTpl.CateName
	model.AddTime = uint(time.Now().Unix())
	model.UpdateTime = uint(time.Now().Unix())
	model.IsIndex = paramTpl.IsIndex
	model.LockTimes = 0
	model.Sort = uint8(paramTpl.Sort)
	model.Status = paramTpl.Status
	err = model.CreateCate()
	if err != nil {
		rtnTpl.Code = "2004"
		rtnTpl.Msg = err.Error()
		c.Json(rtnTpl)
	}
	rtnTpl.Code = "0"
	rtnTpl.Msg = "添加成功"
	rtnTpl.Data = model
	c.Json(rtnTpl)
}

//@Title GetCate
//@Description 获取分类列表
//@Success 200
//@router /list [get]
func (c *CateController)List(){
	rtnTpl := new(rtnTpl)
	list,err := models.GetCates()
	if err != nil {
		rtnTpl.Code = "2005"
		rtnTpl.Msg = err.Error()
		c.Json(rtnTpl)
	}
	rtnTpl.Code = "0"
	rtnTpl.Msg = "获取成功"
	rtnTpl.Data = list
	c.Json(rtnTpl)
}

//@Title GetCate
//@Description 修改分类列表
//@Success 200
//@router /edit [post]
func (c *CateController)Edit(){
	rtnTpl := new(rtnTpl)
	cateValidate := new(validation.CateValidate)
	err := json.Unmarshal(c.Ctx.Input.RequestBody,cateValidate)
	if err != nil {
		rtnTpl.Code = "2006"
		rtnTpl.Msg = err.Error()
		c.Json(rtnTpl)
	}
	err = cateValidate.ValidCate()
	if err != nil {
		rtnTpl.Code = "2006"
		rtnTpl.Msg = err.Error()
		c.Json(rtnTpl)
	}
	if cateValidate.Id == 0 {
		rtnTpl.Code = "2006"
		rtnTpl.Msg = "分类id错误"
		c.Json(rtnTpl)
	}
	err,data := models.UpdateCate(cateValidate)
	rtnTpl.Code = "0"
	rtnTpl.Msg = "更新成功"
	rtnTpl.Data = data
	c.Json(rtnTpl)
}


//@Title DelCate
//@Description 删除分类列表
//@Success 200
//@router /del [delete]
func (c *CateController)DelCate(){
	cateValidate := new(validation.CateValidate)
	err := json.Unmarshal(c.Ctx.Input.RequestBody,cateValidate)
	rtnTpl := new(rtnTpl)
	if err != nil {
		rtnTpl.Code = "2007"
		rtnTpl.Msg = err.Error()
		c.Json(rtnTpl)
	}
	if cateValidate.Id == 0 {
		rtnTpl.Code = "2008"
		rtnTpl.Msg = "分类ID错误"
		c.Json(rtnTpl)
	}
	if err := models.DeleteCate(cateValidate.Id);err != nil {
		rtnTpl.Code = "2008"
		rtnTpl.Msg = err.Error()
		c.Json(rtnTpl)
	}
	rtnTpl.Code = "0"
	rtnTpl.Msg = "删除成功"
	c.Json(rtnTpl)
}