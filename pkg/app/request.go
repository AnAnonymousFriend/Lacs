package app

import "github.com/astaxie/beego/validation"

func MarkErrors(errors []*validation.Error)  {
	for _,err:= range errors{
		println(err.Key + err.Message)
	}
	return
}
