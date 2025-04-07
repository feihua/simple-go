package dao

import (
	"github.com/feihua/simple-go/internal/models"
	"github.com/feihua/simple-go/pkg/config"
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

// func TestCreateUser(t *testing.T) {
//	userDto := dto.UserDto{
//		Mobile:   "",
//		UserName: "",
//		Password: "",
//		StatusId: 0,
//		Sort:     0,
//		Remark:   "",
//	}
//	err := CreateUser(userDto)
//	if err != nil {
//		t.Fail()
//	}
// }
//
// func TestGetUserByUsername(t *testing.T) {
//	//user := QueryUserByUsername(username)
//	//if user.Name == "" {
//	//	t.Fail()
//	//}
//	//fmt.Println(user)
// }
//
// func TestGetUserList(t *testing.T) {
//	list, _ := QueryUserList(1, 2)
//
//	fmt.Println(list)
// }
//
// func TestUpdateUser(t *testing.T) {
//	//user := QueryUserByUsername(username)
//	//userDto := dto.UserDto{
//	//	ID:       user,
//	//	Username: update_username,
//	//	Password: user.Password,
//	//}
//	//err := UpdateUser(userDto)
//	//if err != nil {
//	//	fmt.Println(err.Error())
//	//	t.Fail()
//	//}
// }
//
// func TestDeleteUserByIds(t *testing.T) {
//	//user := QueryUserList(username)
//	//err := DeleteUserByIds(user.ID)
//	//if err != nil {
//	//	t.Fail()
//	//}
// }
