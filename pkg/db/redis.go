package db

import "github.com/tangseng-vge/TangSengDaoDaoServerLib/pkg/redis"

func NewRedis(addr string, password string) *redis.Conn {
	return redis.New(addr, password)
}
