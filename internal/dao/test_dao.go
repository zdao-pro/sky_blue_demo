package dao

import (
	"github.com/zdao-pro/sky_blue/pkg/database/sql"
)

//NewTestModel 返回test数据库Model指针
func NewTestModel() *sql.Model {
	return sql.NewModel(testDB)
}
