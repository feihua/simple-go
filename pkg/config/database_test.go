package config

import (
	"fmt"
	"testing"
)

func TestMysql(t *testing.T) {

	Mysql = (&mysql{}).Load("../../config/app.ini").Init()
	fmt.Println(Mysql)
	if Mysql == nil {
		t.Fail()
	}
}
