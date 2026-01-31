package casbin

import (
	"github.com/casbin/casbin/v3"
	"github.com/casbin/casbin/v3/model"
	entAdapter "github.com/casbin/ent-adapter"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"system/conf"
)

var CasbinEnforcer *casbin.Enforcer

func InitCasbin() {
	var err error
	CasbinEnforcer, err = newCasbin()
	if err != nil {
		hlog.Fatal(err)
	}

}

func newCasbin() (enforcer *casbin.Enforcer, err error) {
	//adapter, err := entAdapter.NewAdapter("mysql", conf.GetConf().MySQL.DSN)

	adapter, err := entAdapter.NewAdapter("postgres", conf.GetConf().PostgreSQL.DSN)
	//adapter, err := entAdapter.NewAdapter("postgres", "user=root password=kcer913639 host=101.126.9.226 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		hlog.Error(err)
		return
	}

	var text string
	if conf.GetConf().Casbin.ModelText == "" {
		text = `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
	} else {
		text = conf.GetConf().Casbin.ModelText
	}

	m, err := model.NewModelFromString(text)
	if err != nil {
		hlog.Error(err)
		return
	}

	enforcer, err = casbin.NewEnforcer(m, adapter)
	if err != nil {
		hlog.Error(err)
		return
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		hlog.Error(err)
		return
	}

	return
}
