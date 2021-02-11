package middleware

import (
	"Lacs/pkg/setting"
	"fmt"
	_ "Lacs/pkg/setting"
)



func CasbinVerification(sub string,dom string,obj string,act string)  {
	if passed, _ := setting.CabinEnforcer.Enforce(sub, dom, obj, act); passed {
		// permit clark to read data1
		fmt.Println("Enforce policy passed.")
	} else {
		// deny the request, show an error
		fmt.Println("Enforce policy denied.")
	}

}