package user

import (
	"github.com/gin-gonic/gin"
	"iads/server/model"
	"iads/server/pkg/e"
	"iads/server/pkg/util"
	"strconv"
)

// @Summary 用户注册
// @Produce  json
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200
// @Failure 500
// @Router /api/v1/tags [post]
func register(c *gin.Context) {
	u := &model.User{}
	if err := c.ShouldBindJSON(u); err != nil {
		util.RES(c, e.INVALID_PARAMS, gin.H{
			"message": err.Error(),
		})
		return
	}

	if flag := u.GetOneByUsername(u.Username); flag == false {
		err := u.UserAdd()
		if err != nil {
			util.RES(c, e.ERROR, gin.H{
				"message": err.Error(),
			})
		} else {
			util.RES(c, e.SUCCESS, gin.H{
				"message": "注册成功",
			})
		}
	} else {
		util.RES(c, e.ERROR, gin.H{
			"message": "用户名已存在！",
		})
	}
}

//列表user数据
func UserList(c *gin.Context) {
	var users model.User
	result, err := users.UserList()
	if err != nil {
		util.RES(c, e.ERROR, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		util.RES(c, e.ERROR, gin.H{
			"data": result,
		})
	}
}

// @Summary 添加user用户
// @Description add user by username and password
// @Accept  json
// @Produce  json
// @Param  username query string true "Username"
// @Param  password query string true "Password"
// @Success 200 {string} string	"ok"
// @Router /v1.0/useradd [post]
func AddUser(c *gin.Context) {
	u := &model.User{}
	if err := c.ShouldBindJSON(u); err != nil {
		util.RES(c, e.INVALID_PARAMS, gin.H{
			"message": err.Error(),
		})
		return
	}

	if flag := u.GetOneByUsername(u.Username); flag == false {
		err := u.UserAdd()
		if err != nil {
			util.RES(c, e.ERROR, gin.H{
				"message": err.Error(),
			})
		} else {
			util.RES(c, e.SUCCESS, gin.H{
				"message": "添加成功",
			})
		}
	} else {
		util.RES(c, e.ERROR, gin.H{
			"message": "用户名已存在！",
		})
	}
}

//修改user数据
func UpdateUserByID(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	user.Password = c.Request.FormValue("password")
	_, err = user.UserUpdate(uint(id))
	if err != nil {
		util.RES(c, e.ERROR, gin.H{
			"message": err.Error(),
		})
	} else {
		util.RES(c, e.SUCCESS, gin.H{})
	}
}

//删除user数据
func DeleteUserByID(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if id == 0 {
		util.RES(c, e.INVALID_PARAMS, gin.H{
			"message": "id必须大于0",
		})
	}
	_, err = user.UserDestroy(uint(id))
	if err != nil {
		util.RES(c, e.ERROR, gin.H{
			"message": err.Error(),
		})
	} else {
		util.RES(c, e.SUCCESS, gin.H{
			"message": "删除成功",
		})
	}
}

//列表role数据
func RoleList(c *gin.Context) {
	var roles model.Role
	result, err := roles.RoleList()
	if err != nil {
		util.RES(c, e.ERROR, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		util.RES(c, e.ERROR, gin.H{
			"data": result,
		})
	}
}

// @Summary 添加role
// @Description add user by username and password
// @Accept  json
// @Produce  json
// @Param  username query string true "RoleName"
// @Param  password query string true "Password"
// @Success 200 {string} string	"ok"
// @Router /v1.0/useradd [post]
func AddRole(c *gin.Context) {
	r := &model.Role{}
	if err := c.ShouldBindJSON(r); err != nil {
		util.RES(c, e.INVALID_PARAMS, gin.H{
			"message": err.Error(),
		})
		return
	}

	if flag := r.GetOneByRolename(r.Rolename); flag == false {
		err := r.RoleAdd()
		if err != nil {
			util.RES(c, e.ERROR, gin.H{
				"message": err.Error(),
			})
		} else {
			util.RES(c, e.SUCCESS, gin.H{
				"message": "添加成功",
			})
		}
	} else {
		util.RES(c, e.ERROR, gin.H{
			"message": "role名已存在！",
		})
	}
}

//删除role数据
func DeleteRoleByID(c *gin.Context) {
	var role model.Role
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if id == 0 {
		util.RES(c, e.INVALID_PARAMS, gin.H{
			"message": "id必须大于0",
		})
	}
	_, err = role.RoleDestroy(uint(id))
	if err != nil {
		util.RES(c, e.ERROR, gin.H{
			"message": err.Error(),
		})
	} else {
		util.RES(c, e.SUCCESS, gin.H{
			"message": "删除成功",
		})
	}
}

//修改role数据
func UpdateRoleByID(c *gin.Context) {
	var role model.Role
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	role.Rolename = c.Request.FormValue("rolename")
	_, err = role.RoleUpdate(uint(id))
	if err != nil {
		util.RES(c, e.ERROR, gin.H{
			"message": err.Error(),
		})
	} else {
		util.RES(c, e.SUCCESS, gin.H{})
	}
}
