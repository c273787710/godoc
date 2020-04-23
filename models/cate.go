package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"godoc/validation"
	"time"
)

type CateModel struct {
	Id uint `orm:"column(id)" json:"id"`
	CateName string `orm:"column(cate_name)" json:"cate_name"`
	Status uint8 `orm:"column(status)" json:"status"`
	LockTimes uint `orm:"column(lock_times)" json:"lock_times"`
	AddTime uint `orm:"column(add_time)" json:"add_time"`
	UpdateTime uint `orm:"column(update_time)" json:"update_time"`
	IsIndex uint8 `orm:"column(is_index)" json:"is_index"`
	Sort uint8 `orm:"column(sort)" json:"sort"`
}
func(c *CateModel)TableName()string{
	return GetTableName("cate")
}

func (c *CateModel)CreateCate() error{
	//判断分类是否存在
	o := orm.NewOrm()
	isExit := new(CateModel)
	err := o.QueryTable(GetTableName("cate")).Filter("cate_name",c.CateName).One(isExit)
	if err == nil {
		return errors.New("分类已存在")
	}
	_,err = o.Insert(c)
	return err
}

func GetCates()(*[]CateModel,error){
	var cateList []CateModel
	o := orm.NewOrm()
	_,err := o.QueryTable(GetTableName("cate")).All(&cateList)
	return &cateList,err
}
func UpdateCate(tpl *validation.CateValidate) (error,*CateModel){
	o := orm.NewOrm()
	exit := o.QueryTable(GetTableName("cate")).Filter("cate_name",tpl.CateName).Filter("id__nq",tpl.Id).Exist()
	if exit {
		return errors.New("分类已存在"),nil
	}
	model := new(CateModel)
	model.Id = tpl.Id
	err := o.Read(model)
	if err != nil {
		return err,nil
	}
	model.CateName = tpl.CateName
	model.Status = tpl.Status
	model.IsIndex = tpl.IsIndex
	model.UpdateTime = uint(time.Now().Unix())
	model.Sort = uint8(tpl.Sort)
	_,err = o.Update(model)
	return err,model
}
func DeleteCate(id uint) error{
	o := orm.NewOrm()
	model := new(CateModel)
	model.Id = id
	_,err := o.Delete(model)
	return err
}



