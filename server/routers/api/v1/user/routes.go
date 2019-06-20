package user

import (
	"github.com/gin-gonic/gin"
)

// 用户RegisterRouter 注册路由
func RegisterUserRouter(r *gin.RouterGroup) {

	// 注册
	r.POST("/register", register)
	// 登录
	//r.POST("/Login", Auth.LoginHandler)

	auth := r.Group("")
	auth.Use(Auth.MiddlewareFunc())
	{
		// 用户列表
		auth.GET("/list", UserList)
		//添加用户
		auth.POST("/add", AddUser)
		// 删除用户
		auth.DELETE("/:id", DeleteUserByID)
		// 更新用户信息
		auth.PUT("/:id", UpdateUserByID)
	}
}

// 角色RegisterRouter 注册路由
func RegisterRoleRouter(r *gin.RouterGroup) {
	auth := r.Group("")
	auth.Use(Auth.MiddlewareFunc())
	{
		// 角色列表
		auth.GET("/list", RoleList)
		//添加角色
		auth.POST("/add", AddRole)
		// 删除角色
		auth.DELETE("/:id", DeleteRoleByID)
		// 更新角色信息
		auth.PUT("/:id", UpdateRoleByID)
	}
}
