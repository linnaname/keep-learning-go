package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
表名默认是结构体名称的复数形式，我醉了
*/

type Team struct {
	gorm.Model
	//注意字段名一定要大写开头框架会自动转换为小写否则不行......
	Name string
	Age  int
}

func main() {

	//连接
	dsn := "root:123456@tcp(127.0.0.1:3306)/pubg_diary?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//??转换model?
	db.AutoMigrate(&Team{})

	//插入
	//db.Create(&Team{Name: "D42", Age: 20})

	var team Team
	db.First(&team, 7)
	fmt.Println(team.Name)
	db.Last(&team)
	fmt.Println(team.ID)

	var teams []Team

	db.Order("age desc, name").Offset(7).Limit(2).Find(&teams)
	fmt.Println(teams)

	var user Team
	db.First(&user, 8)
	fmt.Println(user)

	//db.Where("name = ?", "jinzhu").First(&user)

	var count int64
	db.Table("teams").Count(&count)
	fmt.Println(count)

	db.Delete(&team, 1)

}
