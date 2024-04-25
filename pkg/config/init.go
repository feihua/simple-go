package config

import "fmt"

func init() {
	Server = (&server{}).Load("config/app.ini").Init()
	fmt.Println(Server)
	Mysql = (&mysql{}).Load("config/app.ini").Init()
	fmt.Println(Mysql)
	TokenInfo = (&tokenInfo{}).Load("config/app.ini").Init()
	fmt.Println(Mysql)
}
