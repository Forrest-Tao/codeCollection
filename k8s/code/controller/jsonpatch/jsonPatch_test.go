package jsonpatch

import (
	"fmt"
	jsonpatch "github.com/evanphx/json-patch"
	"k8s.io/apimachinery/pkg/util/json"
	"log"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr"`
}

func (u *User) Empty() bool {
	return u.Age == 0 && u.Name == "" && u.Addr == ""
}

func (u *User) ToJson() []byte {
	res, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func TestJsonPath(t *testing.T) {
	user1 := &User{
		Name: "jiang",
		Age:  19,
		Addr: "tutu",
	}
	user2 := &User{
		Name: "jiang",
		Age:  20,       //changed
		Addr: "London", //changed
	}
	// original obj

	user3 := &User{
		Name: "jiang",
	}
	patch, err := jsonpatch.CreateMergePatch(user1.ToJson(), user2.ToJson())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("patch: ", string(patch))
	mergePatch, err := jsonpatch.MergePatch(user3.ToJson(), patch)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("final res", string(mergePatch))
}
