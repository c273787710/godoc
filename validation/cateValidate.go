package validation

import (
	"github.com/astaxie/beego/validation"
	"errors"
)

type CateValidate struct {
	Id uint `json:"id"`
	CateName string `json:"cate_name" valid:"Required"`
	Status uint8 `json:"status" valid:"ZeroOrOne"`
	IsIndex uint8 `json:"is_index" valid:"ZeroOrOne"`
	Sort int `json:"sort" valid:"Max(200);Min(1)"`
}

func (c *CateValidate)ValidCate()error{
	valid := validation.Validation{}
	b,_ := valid.Valid(c)
	if !b {
		//验证不通过
		return errors.New("传递参数错误")
	}
	return nil
}


