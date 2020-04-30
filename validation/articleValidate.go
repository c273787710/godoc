package validation

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/validation"
)

type ArticleValidate struct {
	Id uint `json:"id"`
	Title string `json:"title" valid:"Required"`
	Desc string `json:"desc" valid:"Required"`
	Cid uint `json:"cid" valid:"Required"`
	IsHot uint8 `json:"is_hot" valid:"ZeroOrOne"`
	IsRecommend uint8 `json:"is_recommend" valid:"ZeroOrOne"`
	Status uint8 `json:"status" valid:"ZeroOrOne"`
	Content string `json:"content" valid:"Required"`
}

func (c *ArticleValidate)ValidArticle()error{
	valid := validation.Validation{}
	b,_ := valid.Valid(c)
	if !b {
		fmt.Println(valid.Errors[0].Error())
		//验证不通过
		return errors.New("传递参数错误")
	}
	return nil
}
