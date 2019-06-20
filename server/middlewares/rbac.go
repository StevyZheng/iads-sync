package middlewares

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RbacInitReturnEnforcer() *casbin.Enforcer {
	a := gormadapter.NewAdapter("mysql", "root:000000@tcp(127.0.0.1:3306)/iads?charset=utf8&parseTime=True&loc=Local&timeout=10ms", true)
	casbinModel := `
	[request_definition]
	r = sub, obj, act
	[policy_definition]
	p = sub, obj, act
	[role_definition]
	g = _, _
	[policy_effect]
	e = some(where (p.eft == allow))
	[matchers]
	m = ((r.sub == p.sub) || g(r.sub,p.sub))  && ((r.obj == p.obj) || g(r.obj,p.obj)) && ((r.act == p.act) || g(r.act,p.act)) || r.sub == "admin"
	`
	e, _ := casbin.NewEnforcerSafe(casbin.NewModel(casbinModel), a)
	//从DB加载策略
	_ = e.LoadPolicy()
	return e
}

// NewAuthorizer returns the authorizer, uses a Casbin enforcer as input
func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	a := &BasicAuthorizer{enforcer: e}

	return func(c *gin.Context) {
		if !a.CheckPermission(c) {
			a.RequirePermission(c)
		}
	}
}

// BasicAuthorizer stores the casbin handler
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

// GetRoleName gets the user name from the request.
// Currently, only HTTP basic authentication is supported
func (a *BasicAuthorizer) GetRoleName(c *gin.Context) interface{} {
	fmt.Println("GetRoleName")
	session := sessions.Default(c)
	fmt.Println(session)
	role := session.Get("role")
	fmt.Println(role)
	if role != nil && role != "" {
		fmt.Println("session role --->", role)
		return role
	}
	fmt.Println("defaulte role ---> common")
	return "common"
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(c *gin.Context) bool {
	fmt.Println("CheckPermission")
	role := a.GetRoleName(c)
	method := c.Request.Method
	path := c.Request.URL.Path
	return a.enforcer.Enforce(role, path, method)
}

// RequirePermission returns the 403 Forbidden to the client
func (a *BasicAuthorizer) RequirePermission(c *gin.Context) {
	c.AbortWithStatus(403)
}

//拦截器
/*
func RbacHandler(e *casbin.Enforcer) gin.HandlerFunc {

	return func(c *gin.Context) {

		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		//获取请求方法
		act := c.Request.Method
		//获取用户的角色
		sub := "admin"

		//判断策略中是否存在
		if e.Enforce(sub, obj, act) {
			fmt.Println("通过权限")
			c.Next()
		} else {
			fmt.Println("权限没有通过")
			c.Abort()
		}
	}
}
func RbacHandler(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		authz.NewAuthorizer(e)
	}
}*/
