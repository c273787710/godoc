package init

import "godoc/validation"

func init(){
	DBinit()
	validation.ValidateInit()
}
