package main

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"log"
)

func main() {
	text := `
    [request_definition]
    r = sub, obj, act

    [policy_definition]
    p = sub, obj, act

    [role_definition]
    g = _, _

    [policy_effect]
    e = some(where (p.eft == allow))

    [matchers]
    m = r.sub == p.sub && keyMatch2(r.obj, p.obj) && r.act == p.act
    `
	m, err := model.NewModelFromString(text)
	if err != nil {
		log.Fatalf("model: %s", err)
	}

	e, err := casbin.NewEnforcer(m)
	if err != nil {
		log.Fatalf("enforcer: %s", err)
	}

	// Add policy rules
	e.AddPolicy("alice", "/foo/:id", "GET")

	// Test permissions
	ok, err := e.Enforce("alice", "/foo/123", "GET")
	if err != nil {
		log.Fatalf("enforce: %s", err)
	}
	if ok {
		log.Println("Access granted")
	} else {
		log.Println("Access denied")
	}
}
