package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Initialize a Xorm adapter with MySQL database.
	a, err := xormadapter.NewAdapter("mysql", "mysql:123456@tcp(127.0.0.1:3306)/casbin")
	if err != nil {
		log.Fatalf("error: adapter: %s", err)
	}

	m, err := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`)
	if err != nil {
		log.Fatalf("error: model: %s", err)
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}

	sub := "alice"  // the user that wants to access a resource.
	obj := "data12" // the resource that is going to be accessed.
	act := "read"   // the operation that the user performs on the resource.

	if res, err := e.Enforce(sub, obj, act); res {
		// permit alice to read data1
		fmt.Println(res)
	} else {
		fmt.Println(err)
		// deny the request, show an error
	}
}
