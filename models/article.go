package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
)

type ArticleModel struct {
	Id uint `orm:"column(id)" json:"id"`
	Title string `orm:"column(title)" json:"title"`
	Desc string `orm:"column(desc)" json:"desc"`
	LookTimes uint `orm:"column(look_times)" json:"look_times"`
	Stars uint `orm:"column(stars)" json:"stars"`
	IsHot uint8 `orm:"column(is_hot)" json:"is_hot"`
	IsRecommend uint8 `orm:"column(is_recommend)" json:"is_recommend"`
	Status uint8 `orm:"column(status)" json:"status"`
	Content string `orm:"column(content)" json:"content"`
	AddTime uint `orm:"column(add_time)" json:"add_time"`
	UpdateTime uint `orm:"column(update_time)" json:"update_time"`
	Cate *CateModel `orm:"rel(one)" json:"cate"`
}

type ArticleSearch struct {
	Title string
	Cid int
	Page int
	Limit int
	IsHot int
	IsRecommend int
}

func (a *ArticleModel)TableName()string{
	return GetTableName("article")
}
func GetArticleList(option *ArticleSearch)(*[]ArticleModel,error){
	fmt.Println(option)
	var articleList []ArticleModel
	query := orm.NewOrm().QueryTable(GetTableName("article"))
	if option.Title != "" {
		query = query.Filter("title__contains",strings.Trim(option.Title,""))
	}
	if option.Cid != 0 {
		query = query.Filter("cate_id",option.Cid)
	}
	if option.IsHot != 0 {
		query = query.Filter("is_hot",option.IsHot-1)
	}
	if option.IsRecommend != 0 {
		query = query.Filter("is_recommend",option.IsRecommend-1)
	}
	start := (option.Page - 1) * option.Limit
	query = query.Limit(option.Limit,start)
	_,err := query.OrderBy("-id").RelatedSel().All(&articleList)
	return &articleList,err
}
func GetPageInfo(option *ArticleSearch)int64{
	query := orm.NewOrm().QueryTable(GetTableName("article"))
	if option.Title != "" {
		query = query.Filter("title__contains",strings.Trim(option.Title,""))
	}
	if option.Cid != 0 {
		query = query.Filter("cate_id",option.Cid)
	}
	if option.IsHot != 0 {
		query = query.Filter("is_hot",option.IsHot-1)
	}
	if option.IsRecommend != 0 {
		query = query.Filter("is_recommend",option.IsRecommend-1)
	}
	count,err := query.Count()
	if err != nil{
		return 0
	}
	return count
}

func DelArticlByID(id uint) error{
	model := new(ArticleModel)
	model.Id = id
	_,err := orm.NewOrm().Delete(model)
	return err
}

func (a *ArticleModel)CreateArticle() error{
	o := orm.NewOrm()
	_,err := o.Insert(a)
	return err
}

func GetArticleByID(id uint) *ArticleModel{
	model := new(ArticleModel)
	model.Id = id
	err := orm.NewOrm().Read(model)
	if err != nil {
		return nil
	}
	return model
}

func (a *ArticleModel)UpdateArticle() error{
	_,err := orm.NewOrm().Update(a)
	return err
}

