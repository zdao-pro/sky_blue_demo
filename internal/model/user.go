package model

import (
	"context"
	"fmt"
	"time"

	"sky_blue_demo/internal/dao"

	"github.com/zdao-pro/sky_blue/pkg/database/sql"
)

/*
user_info表的结构
+-------------+---------------+------+-----+-------------------+-------------------+
| Field       | Type          | Null | Key | Default           | Extra             |
+-------------+---------------+------+-----+-------------------+-------------------+
| id          | bigint        | NO   | PRI | NULL              | auto_increment    |
| name        | varchar(50)   | NO   | MUL |                   |                   |
| age         | int           | YES  |     | NULL              |                   |
| regist_time | timestamp     | NO   |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED |
| status      | tinyint       | NO   |     | 0                 |                   |
| create_time | bigint        | NO   |     | 0                 |                   |
| misc        | varchar(1024) | NO   |     |                   |                   |
+-------------+---------------+------+-----+-------------------+-------------------+
*/

// 使用时需在结构体tag指明orm
// 当字段存储的是int类型时间戳，需指明time_format:  unix(秒级时间戳) unixmilli(毫秒级时间戳) unixnano(纳秒级时间戳)，才可以自动解析为time.Time
// 当字段为json字符串,需要在解析结构体tag指明json, 才可以自动解析
// 结构体需要声明*sql.Model,并赋值
// 使用事务时,可以直接使用Begin(context.Background()),Commit(),Rollback(),不需要另外声明
type UserInfo struct {
	*sql.Model
	ID         int16     `orm:"id"`
	Name       string    `orm:"name"`
	Age        int       `orm:"age"`
	Status     uint8     `orm:"status"`
	RegistTime time.Time `orm:"regist_time"`
	CreateTime time.Time `orm:"create_time" time_format:"unixmilli"`
	MiscInfo   Misc      `orm:"misc"`
}

//Misc misc字段解析的json字符串映射的结构体
type Misc struct {
	Email string `json:"email"`
	Phone int    `json:"phone"`
}

//NewUserInfo ..
func NewUserInfo() *UserInfo {
	return &UserInfo{
		Model: dao.NewTestModel(),
	}
}

//Insert 插入信息
func (u *UserInfo) Insert(name string, age int, createTime int) (err error) {
	_, err = u.Exec(context.Background(), "insert into user_info set name = ?,age = ?,create_time = ?", name, age, createTime)
	return
}

//QueryUserByID ..
func (u *UserInfo) QueryUserByID(id int) error {
	err := u.Select(context.Background(), u, "select id,name,age,regist_time,status,create_time,misc from user_info where id = ?", id)
	if nil != err {
		return err
	}
	fmt.Println(u)
	return nil
}

//QueryUserByName ..
func (u *UserInfo) QueryUserByName(name string) (*[]UserInfo, error) {
	list := make([]UserInfo, 0)
	err := u.Select(context.Background(), &list, "select id,name,age,regist_time,status,create_time,misc from user_info where name = ?", name)
	if nil != err {
		return nil, err
	}
	// fmt.Println(list)
	return &list, nil
}
