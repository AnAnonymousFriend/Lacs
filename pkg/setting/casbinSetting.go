package setting

import "github.com/casbin/casbin/v2"

func CasbinSetup()  {
	ef, err := casbin.NewSyncedEnforcer("../conf/acl_simple_model.conf", "../conf/acl_simple_policy.csv")
	if err!=nil {
		println(err)
	}
	CabinEnforcer = ef
}