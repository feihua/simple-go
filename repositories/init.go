package repositories

import (
	"github.com/feihua/simple-go/models"
)

func init() {
	//config.Mysql.Load("../config/app.ini").Init()
	models.Init()
}
