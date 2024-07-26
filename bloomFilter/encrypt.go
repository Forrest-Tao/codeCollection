package bloomFilter

import (
	"github.com/spaolacci/murmur3"
)

type Encrypt interface {
	Encrypt(source string, mod uint32) uint32
}

type Encryptor struct{}

func NewEncryptor() *Encryptor {
	return &Encryptor{}
}

func (e *Encryptor) Encrypt(source string, mod uint32) uint32 {
	hash := murmur3.New32()
	hash.Write([]byte(source))
	return hash.Sum32() % mod
}
