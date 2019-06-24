package models

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Rolename string `json:"rolename" gorm:"unique"`
}

func (role *Role) GetOneByRolename(rolename string) bool {
	var r Role
	DBConnect.Eloquent.Select("id").Where("rolename = ?", rolename).First(&r)
	if r.ID > 0 {
		return true
	}
	return false
}

func (role Role) RoleAdd() (err error) {
	ret := DBConnect.Eloquent.Create(&role)
	if ret.Error != nil {
		err = ret.Error
		return
	}
	return
}

func (role *Role) RoleList() (roles []Role, err error) {
	if err = DBConnect.Eloquent.Find(&roles).Error; err != nil {
		return
	}
	return
}

//修改role
func (role *Role) RoleUpdate(id uint) (updateRole Role, err error) {
	if err = DBConnect.Eloquent.Select([]string{"id", "rolename"}).First(&updateRole, id).Error; err != nil {
		return
	}
	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = DBConnect.Eloquent.Model(&updateRole).Update(&role).Error; err != nil {
		return
	}
	return
}

//删除role数据
func (role *Role) RoleDestroy(id uint) (Result Role, err error) {
	if err = DBConnect.Eloquent.Select([]string{"id"}).First(&role, id).Error; err != nil {
		return
	}
	if err = DBConnect.Eloquent.Delete(&role).Error; err != nil {
		return
	}
	Result = *role
	return
}
