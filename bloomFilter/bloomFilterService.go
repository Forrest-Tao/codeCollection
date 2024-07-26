package bloomFilter

import (
	"context"
	_ "embed"
	"strconv"
)

//go:embed lua/BatchGetBits.lua
var BatchGetBitsScript string

//go:embed lua/SetBits.lua
var SetBitsScript string

type BloomFilterService struct {
	e    Encrypt
	rdb  *RedisClient
	mLen uint32
	k    int
}

func NewBloomFilterService(mLen uint32, k int, rdb *RedisClient, e Encrypt) *BloomFilterService {
	return &BloomFilterService{
		e:    e,
		rdb:  rdb,
		mLen: mLen,
		k:    k,
	}
}

func (b *BloomFilterService) Exist(ctx context.Context, key, value string) (bool, error) {
	args := make([]interface{}, b.k+1)
	//lua 中我们设置第一个arg为 hash算法的个数
	args[0] = b.k
	for _, encrypted := range b.getKEncrypted(value) {
		args = append(args, encrypted)
	}
	// 对比bitmap中相应位置的 bit值是否都为1
	res, err := b.rdb.Eval(ctx, BatchGetBitsScript, []string{key}, args...)
	if err != nil {
		return false, err
	}
	return res.(int64) == 1, nil
}

// set
func (b *BloomFilterService) set(ctx context.Context, key, value string) error {
	args := make([]interface{}, b.k+1)
	//lua 中我们设置第一个arg为 hash算法的个数
	args[0] = b.k
	for _, encrypted := range b.getKEncrypted(value) {
		args = append(args, encrypted)
	}
	eval, err := b.rdb.Eval(ctx, SetBitsScript, []string{key}, args...)
	if err != nil || eval != 1 {
		return err
	}
	return nil
}

func (b *BloomFilterService) getKEncrypted(value string) []uint32 {
	encrypteds := make([]uint32, 0, b.k)
	origin := value

	for i := 0; i < b.k; i++ {
		encrypted := b.e.Encrypt(origin, b.mLen)
		encrypteds = append(encrypteds, encrypted)
		origin = strconv.Itoa(int(encrypted))
	}
	return encrypteds
}
