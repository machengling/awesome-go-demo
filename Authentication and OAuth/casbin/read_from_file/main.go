package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func main() {

	e, _ := casbin.NewEnforcer("conf/model.conf", "conf/policy.csv")

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

	results, err := e.BatchEnforce([][]interface{}{{"alice", "data1", "read"}, {"bob", "data2", "write"}, {"jack", "data3", "read"}})
	fmt.Println(results, err)


}
