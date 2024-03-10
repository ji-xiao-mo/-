package models

import (
	"GOwork/project/gin-ranking/dao"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (User) TableName() string {
	return "user"
}

// 根据指定的ID来查询一行数据
func GetUserTest(id int) (User, error) {
	var user User
	err := dao.Db.Where("id = ?", id).First(&user).Error
	return user, err
}

func GetUserListTest() ([]User, error) {
	var users []User
	err := dao.Db.Where("id < ?", 3).Find(&users).Error
	return users, err
}

// 用来在表格中增加一行数据
func AddUser(username string) (int, error) {
	user := User{Username: username}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}

// 用来指定更新数据
func UpdataUser(id int, username string) {
	dao.Db.Model(&User{}).Where("id = ?", id).Update("username", username)
}

// 用来删除指定id的数据
func DeleteUser(id int) error {
	err := dao.Db.Delete(&User{}, id).Error
	return err
}
