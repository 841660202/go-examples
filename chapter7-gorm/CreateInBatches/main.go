package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Language struct {
	gorm.Model
	Name string
}

type User struct {
	gorm.Model
	Name      string
	Languages []*Language `gorm:"many2many:user_languages"`
}

func main() {
	// 连接 SQLite 数据库
	dsn := "test.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移表结构
	db.AutoMigrate(&User{}, &Language{})

	// 创建语言记录
	langs := []*Language{
		{Name: "Go"},
		{Name: "Python"},
		{Name: "Java"},
	}
	db.CreateInBatches(langs, len(langs))

	// 创建用户并分配语言
	users := []*User{
		{
			Name: "Alice",
			Languages: []*Language{
				langs[0],
				langs[1],
			},
		},
		{
			Name: "Bob",
			Languages: []*Language{
				langs[1],
				langs[2],
			},
		},
	}
	db.CreateInBatches(users, len(users))
}
