package rand

import (
	"math/rand"
	"time"
)

import "sync"

var rng = struct {
	sync.Mutex
	rand *rand.Rand
}{
	rand: rand.New(rand.NewSource(time.Now().UnixNano())),
}

func Int() int {
	rng.Lock()
	defer rng.Unlock()
	return rng.rand.Int()
}

func Intn(max int) int {
	rng.Lock()
	defer rng.Unlock()
	return rng.rand.Intn(max)
}

func IntRange(minn, maxn int) int {
	rng.Lock()
	defer rng.Unlock()
	return rng.rand.Intn(maxn-minn) + minn
}

func Perm(n int) []int {
	rng.Lock()
	defer rng.Unlock()
	return rng.rand.Perm(n)
}

func Seed(seed int64) {
	rng.Lock()
	defer rng.Unlock()
	rng.rand = rand.New(rand.NewSource(seed))
}
