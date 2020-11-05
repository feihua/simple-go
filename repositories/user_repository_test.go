package repositories

import (
	"fmt"
	"simple-go/dto"
	"simple-go/models"
	"simple-go/pkg/config"
	"testing"
)

const (
	username        = "liufeihua"
	password        = "123456"
	update_username = "liufeihua_update"
)

func init() {
	config.Mysql.Load("../config/app.ini").Init()
	models.Init()
}

func TestCreateUser(t *testing.T) {
	userDto := dto.UserDto{
		Username: username,
		Password: password,
	}
	err := CreateUser(userDto)
	if err != nil {
		t.Fail()
	}
}

func TestGetUserByUsername(t *testing.T) {
	//user := GetUserByUsername(username)
	//if user.Name == "" {
	//	t.Fail()
	//}
	//fmt.Println(user)
}

func TestGetUserList(t *testing.T) {
	user := GetUserByUsername(username)

	fmt.Println(user)
}

func TestUpdateUser(t *testing.T) {
	//user := GetUserByUsername(username)
	//userDto := dto.UserDto{
	//	ID:       user,
	//	Username: update_username,
	//	Password: user.Password,
	//}
	//err := UpdateUser(userDto)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	t.Fail()
	//}
}

func TestDeleteUserById(t *testing.T) {
	//user := GetUserList(username)
	//err := DeleteUserById(user.ID)
	//if err != nil {
	//	t.Fail()
	//}
}
