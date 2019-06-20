package user

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"iads/server/pkg/util"
)

var dbConnect *util.Connection

func Init() {
	dbConnect = util.NewConnection()
}

//用户类
type User struct {
	gorm.Model
	Username string `json:"username" form:"username" validate:"required" gorm:"unique"`
	Password string `json:"password" form:"password" validate:"required"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	RoleID   uint   `json:"role_id"`
	Role     Role   `json:"role"`
}

type Role struct {
	gorm.Model
	Rolename string `json:"rolename" gorm:"unique"`
}

type Permissions struct {
	gorm.Model
	PermissionsEntry string `json:"permissions_entry" gorm:"unique"`
}

func CreateTable() {
	dbConnect.Eloquent.AutoMigrate(&Permissions{}, &Role{}, &User{})
	dbConnect.Eloquent.Model(&User{}).AddForeignKey("role_id", "roles(id)", "no action", "no action")
}

type Login struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}

// Validator .
func (login *Login) Validator() (*User, string, bool) {
	user := &User{Username: login.Username}
	err := dbConnect.Eloquent.Where("username = ?", login.Username).First(&user).Error
	fmt.Println(user)
	var msg string
	if err != nil {
		msg = "没有该账户！"
		return nil, msg, false
	}

	if user.Password != login.Password {
		msg = "密码错误！"
		return nil, msg, false
	}
	msg = "登录成功！"
	return user, msg, true
}

func (user *User) GetOneByUsername(username string) bool {
	var u User
	dbConnect.Eloquent.Select("id").Where("username = ?", username).First(&u)
	if u.ID > 0 {
		return true
	}
	return false
}

//添加user用户
func (user User) UserAdd() (err error) {
	ret := dbConnect.Eloquent.Create(&user)
	if ret.Error != nil {
		err = ret.Error
		return
	}
	return
}

//用户user列表
func (user *User) UserList() (users []User, err error) {
	if err = dbConnect.Eloquent.Preload("Role").Find(&users).Error; err != nil {
		return
	}
	return
}

//修改user
func (user *User) UserUpdate(id uint) (updateUser User, err error) {
	if err = dbConnect.Eloquent.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
		return
	}
	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = dbConnect.Eloquent.Model(&updateUser).Update(&user).Error; err != nil {
		return
	}
	return
}

//删除user数据
func (user *User) UserDestroy(id uint) (Result User, err error) {
	if err = dbConnect.Eloquent.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}
	if err = dbConnect.Eloquent.Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}

//****************************************

func (role *Role) GetOneByRolename(rolename string) bool {
	var r Role
	dbConnect.Eloquent.Select("id").Where("rolename = ?", rolename).First(&r)
	if r.ID > 0 {
		return true
	}
	return false
}

func (role Role) RoleAdd() (err error) {
	ret := dbConnect.Eloquent.Create(&role)
	if ret.Error != nil {
		err = ret.Error
		return
	}
	return
}

func (role *Role) RoleList() (roles []Role, err error) {
	if err = dbConnect.Eloquent.Find(&roles).Error; err != nil {
		return
	}
	return
}

//修改role
func (role *Role) RoleUpdate(id uint) (updateRole Role, err error) {
	if err = dbConnect.Eloquent.Select([]string{"id", "rolename"}).First(&updateRole, id).Error; err != nil {
		return
	}
	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = dbConnect.Eloquent.Model(&updateRole).Update(&role).Error; err != nil {
		return
	}
	return
}

//删除role数据
func (role *Role) RoleDestroy(id uint) (Result Role, err error) {
	if err = dbConnect.Eloquent.Select([]string{"id"}).First(&role, id).Error; err != nil {
		return
	}
	if err = dbConnect.Eloquent.Delete(&role).Error; err != nil {
		return
	}
	Result = *role
	return
}
