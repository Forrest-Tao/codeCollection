package rand_test

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/rand"
	"testing"
)

// 测试 rand 包的功能
func TestRand(t *testing.T) {
	t.Run("TestPerm", func(t *testing.T) {
		result := rand.Perm(10)
		fmt.Println(result)
	})
}
