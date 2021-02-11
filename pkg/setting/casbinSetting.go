package setting

import "github.com/casbin/casbin/v2"

var CasbinEnforcer *casbin.SyncedEnforcer
func CasbinSetting()  {
	ef, err := casbin.NewSyncedEnforcer("../conf/acl_simple_model.conf", "../conf/acl_simple_policy.csv")
	if err!=nil {
		println(err)
	}
	CasbinEnforcer = ef
}
