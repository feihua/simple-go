package dao

import "github.com/feihua/simple-go/internal/model"

type DaoImpl struct {
	UserDao     *UserDao
	UserRoleDao *UserRoleDao
	RoleDao     *RoleDao
	RoleMenuDao *RoleMenuDao
	MenuDao     *MenuDao
	LogDao      *LogDao
	DictDao     *DictDao
	DeptDao     *DeptDao
}

func NewDaoImpl() *DaoImpl {

	db := model.DB
	return &DaoImpl{
		UserDao:     NewUserDao(db),
		UserRoleDao: NewUserRoleDao(db),
		RoleDao:     NewRoleDao(db),
		RoleMenuDao: NewRoleMenuDao(db),
		MenuDao:     NewMenuDao(db),
		LogDao:      NewLogDao(db),
		DictDao:     NewDictDao(db),
		DeptDao:     NewDeptDao(db),
	}
}

var Dao *DaoImpl

func InitDao() {
	model.Init()
	Dao = NewDaoImpl()

}
