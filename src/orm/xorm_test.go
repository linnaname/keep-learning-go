package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)
import "xorm.io/xorm"

const (
	DRIVER_NAME = "mysql"
	DATESOURCE  = "root:123456@/pubg_diary?charset=utf8"
)

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func TestORM(t *testing.T) {
	engine, err := xorm.NewEngine(DRIVER_NAME, DATESOURCE)
	assert.NoError(t, err)
	err = engine.Sync2(new(User))
	assert.NoError(t, err)
}

func TestInsert(t *testing.T) {
	engine, err := xorm.NewEngine(DRIVER_NAME, DATESOURCE)
	assert.NoError(t, err)

	//单个插入
	user := &User{Name: "lzy", Age: 50}
	affected, _ := engine.Insert(user)
	assert.Equal(t, affected, int64(1))
	fmt.Printf("%d records inserted, user.id:%d\n", affected, user.Id)

	//批量插入
	users := make([]*User, 2)
	users[0] = &User{Name: "xhq", Age: 41}
	users[1] = &User{Name: "lhy", Age: 12}
	affected, _ = engine.Insert(&users)
	fmt.Printf("%d records inserted, id1:%d, id2:%d", affected, users[0].Id, users[1].Id)
}

func TestGet(t *testing.T) {
	engine, err := xorm.NewEngine(DRIVER_NAME, DATESOURCE)
	assert.NoError(t, err)
	engine.ShowSQL(true)

	user1 := &User{}
	has, _ := engine.ID(1).Get(user1)
	if has {
		fmt.Printf("user1:%v\n", user1)
	}

	user2 := &User{}
	has, _ = engine.Where("name=?", "dj").Get(user2)
	if has {
		fmt.Printf("user2:%v\n", user2)
	}

	user3 := &User{Id: 5}
	has, _ = engine.Get(user3)
	if has {
		fmt.Printf("user3:%v\n", user3)
	}

	user4 := &User{Name: "pipi"}
	has, _ = engine.Get(user4)
	if has {
		fmt.Printf("user4:%v\n", user4)
	}
}

func TestFind(t *testing.T) {
	engine, err := xorm.NewEngine(DRIVER_NAME, DATESOURCE)
	assert.NoError(t, err)
	engine.ShowSQL(true)

	slcUsers := make([]User, 1)
	engine.Where("age > ? and age < ?", 12, 30).Find(&slcUsers)
	fmt.Println("users whose age between [12,30]:", slcUsers)

	mapUsers := make(map[int64]User)
	engine.Where("length(name) = ?", 3).Find(&mapUsers)
	fmt.Println("users whose has name of length 3:", mapUsers)
}

func TestUpdate(t *testing.T) {
	engine, err := xorm.NewEngine(DRIVER_NAME, DATESOURCE)
	assert.NoError(t, err)
	engine.ShowSQL(true)

	engine.ID(1).Update(&User{Name: "ldj"})
	engine.ID(1).Cols("name", "age").Update(&User{Name: "dj"})
	engine.Table(&User{}).ID(1).Update(map[string]interface{}{"age": 18})
}

func TestDelete(t *testing.T) {
	engine, err := xorm.NewEngine(DRIVER_NAME, DATESOURCE)
	assert.NoError(t, err)
	engine.ShowSQL(true)
	affected, _ := engine.Where("name = ?", "lzy").Delete(&User{})
	fmt.Printf("%d records deleted", affected)
}

func TestOrgin(t *testing.T) {
	engine, err := xorm.NewEngine(DRIVER_NAME, DATESOURCE)
	assert.NoError(t, err)

	querySql := "select * from user limit 1"
	reuslts, _ := engine.Query(querySql)
	for _, record := range reuslts {
		for key, val := range record {
			fmt.Println(key, string(val))
		}
	}

	updateSql := "update `user` set name=? where id=?"
	res, _ := engine.Exec(updateSql, "ldj", 1)
	fmt.Println(res.RowsAffected())
}
