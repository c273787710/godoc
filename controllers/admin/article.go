package admin

import (
	"encoding/json"
	"godoc/models"
	"godoc/validation"
	"math"
	"time"
)

type ArticleController struct {
	BaseController
}
//APIVersion 1.0.0
//@Title Add
//@Description 添加文章
//@Success 200
//@router /add [post]
func (a *ArticleController)Add(){
	params := new(validation.ArticleValidate)
	err := json.Unmarshal(a.Ctx.Input.RequestBody,params)
	rtnTpl := new(rtnTpl)
	if err != nil {
		rtnTpl.Code = "3001"
		rtnTpl.Msg = "传递参数错误"
		a.Json(rtnTpl)
	}
	err = params.ValidArticle()
	if err != nil {
		rtnTpl.Code = "3002"
		rtnTpl.Msg = err.Error()
		a.Json(rtnTpl)
	}
	//根据分类ID获取分类信息
	cate := models.GetCateByID(params.Cid)
	if cate == nil{
		rtnTpl.Code = "3003"
		rtnTpl.Msg = "文章分类不存在"
		a.Json(rtnTpl)
	}
	model := new(models.ArticleModel)
	model.Title = params.Title
	model.Desc = params.Desc
	model.LookTimes = 0
	model.Stars = 0
	model.Cate = cate
	model.IsHot = params.IsHot
	model.IsRecommend = params.IsRecommend
	model.Status = params.Status
	model.Content = params.Content
	model.AddTime = uint(time.Now().Unix())
	model.UpdateTime = uint(time.Now().Unix())
	err = model.CreateArticle()
	if err != nil {
		rtnTpl.Code = "3003"
		rtnTpl.Msg = err.Error()
		a.Json(rtnTpl)
	}
	rtnTpl.Code = "0"
	rtnTpl.Msg = "添加文章成功"
	rtnTpl.Data = model
	a.Json(rtnTpl)
}
//APIVersion 1.0.0
//@Title Edit
//@Description 修改文章
//@Success 200
//@router /edit [post]
func (a *ArticleController)Edit(){
	params := new(validation.ArticleValidate)
	err := json.Unmarshal(a.Ctx.Input.RequestBody,params)
	rtnTpl := new(rtnTpl)
	if err != nil {
		rtnTpl.Code = "3001"
		rtnTpl.Msg = "传递参数错误"
		a.Json(rtnTpl)
	}
	err = params.ValidArticle()
	if err != nil {
		rtnTpl.Code = "3002"
		rtnTpl.Msg = err.Error()
		a.Json(rtnTpl)
	}
	if params.Id == 0 {
		rtnTpl.Code = "3003"
		rtnTpl.Msg = "文章ID错误"
		a.Json(rtnTpl)
	}
	model := models.GetArticleByID(params.Id)
	if model == nil {
		rtnTpl.Code = "3004"
		rtnTpl.Msg = "文章不存在"
		a.Json(rtnTpl)
	}
	cateModel := models.GetCateByID(params.Cid)
	if cateModel == nil {
		rtnTpl.Code = "3005"
		rtnTpl.Msg = "分类不存在"
		a.Json(rtnTpl)
	}
	model.Cate = cateModel
	model.UpdateTime = uint(time.Now().Unix())
	model.Title = params.Title
	model.Content = params.Content
	model.Status = params.Status
	model.IsRecommend = params.IsRecommend
	model.IsHot = params.IsHot
	model.Desc = params.Desc
	err = model.UpdateArticle()
	if err != nil {
		rtnTpl.Code = "3006"
		rtnTpl.Msg = err.Error()
		a.Json(rtnTpl)
	}
	rtnTpl.Code = "0"
	rtnTpl.Msg = "修改成功"
	rtnTpl.Data = model
	a.Json(rtnTpl)
}
//APIVersion 1.0.0
//@Title List
//@Description 添加文章
//@Success 200
//@router /list [get]
func (a *ArticleController)List(){
	options := new(models.ArticleSearch)
	options.Title = a.GetString("title")
	options.Cid,_ = a.GetInt("cid")
	options.IsHot,_ = a.GetInt("is_hot")
	options.IsRecommend,_ = a.GetInt("is_recommend")
	page,_ := a.GetInt("page")
	if page == 0 {
		options.Page = 1
	}else {
		options.Page = page
	}
	limit,_ := a.GetInt("limit")
	if limit == 0 {
		options.Limit = 10
	}else{
		options.Limit = limit
	}
	rtnTpl := new(rtnTpl)
	data,err := models.GetArticleList(options)
	if err != nil{
		rtnTpl.Code = "3003"
		rtnTpl.Msg = err.Error()
		a.Json(rtnTpl)
	}
	count := models.GetPageInfo(options)
	//pages := int64(math.Ceil(float64(count / int64(options.Limit))))
	rtn := struct {
		List *[]models.ArticleModel `json:"list"`
		Count int64 `json:"count"`
		Page int64 `json:"page"`
	}{
		List:data,
		Count:count,
		Page:int64(math.Ceil(float64(count) / float64(options.Limit))),
	}
	rtnTpl.Code = "0"
	rtnTpl.Msg = "获取成功"
	rtnTpl.Data = &rtn
	a.Json(rtnTpl)
}

//APIVersion 1.0.0
//@Title Del
//@Description 删除文章
//@Success 200
//@router /del [post]
func (a *ArticleController)Del(){
	var aid = struct {
		Id uint `json:"id"`
	}{}
	rtnTpl := new(rtnTpl)
	err := json.Unmarshal(a.Ctx.Input.RequestBody,&aid)
	if err != nil {
		rtnTpl.Code = "3003"
		rtnTpl.Msg = err.Error()
		a.Json(rtnTpl)
	}
	err = models.DelArticlByID(aid.Id)
	if err != nil {
		rtnTpl.Code = "3003"
		rtnTpl.Msg = err.Error()
		a.Json(rtnTpl)
	}
	rtnTpl.Code = "0"
	rtnTpl.Msg = "删除成功"
	a.Json(rtnTpl)
}
