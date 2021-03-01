package main

import (

	"fmt"
	"github.com/casbin/casbin/v2"
)


func main()  {



	e,err:= casbin.NewDistributedEnforcer("../conf/acl_simple_model.conf","../conf/acl_simple_policy.csv")
	if err !=nil {
		panic(err)
		return
	}
	//e, _ := casbin.NewSyncedEnforcer("../conf/acl_simple_model.conf",stringSoure)


	//e, _ := casbin.NewSyncedEnforcer("../conf/acl_simple_model.conf", "../conf/acl_simple_policy.csv")
	sub := "p"   // the user that wants to access a resource.
	dom := "admin" // the user who belongs to.
	obj := "/api/v1/role/all"   // the resource that is going to be accessed.
	act := "get"    // the operation that the user performs on the resource.
	if passed, _ := e.Enforce(sub, dom, obj, act); passed {
		// permit clark to read data1
		fmt.Println("Enforce policy passed.")
	} else {
		// deny the request, show an error
		fmt.Println("Enforce policy denied.")
	}
}