package redis

import (
	"testing"
	"time"
)

func TestStore(t *testing.T) {
	a := Config{"127.0.0.1:6379", 0, "", ""}
	s := NewStore(&a)
	e := s.Set("hahaha", time.Second*33)
	if e != nil {
		println(e.Error())
	}
	r, _ := s.Check("hahaha")
	println(r)
	println(s.wrapperKey("hahaha"))
}
