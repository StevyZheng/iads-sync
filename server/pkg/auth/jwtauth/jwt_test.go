package jwtauth

import (
	"iads/server/pkg/auth/jwtauth/store/redis"
	"testing"
)

func TestJWT(t *testing.T) {
	a := redis.Config{"127.0.0.1:6379", 0, "", ""}
	t1 := New(redis.NewStore(&a))
	token, err := t1.GenerateToken("1")
	if err != nil || token != nil {
		println(err.Error())
	}
}
