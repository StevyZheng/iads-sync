package redis

import (
	"testing"
	"time"
)

func TestStore(t *testing.T) {
	a := Config{"localhost:6379", 0, "", ""}
	s := NewStore(&a)
	e := s.Set("hahaha1", time.Second*33)
	if e != nil {
		println(e.Error())
	}
	r, _ := s.Check("hahaha")
	println(r)
	println(s.wrapperKey("hahaha"))
}
