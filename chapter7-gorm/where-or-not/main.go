package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	IsActive bool
}

func main() {
	// 创建 SQLite 数据库连接
	dsn := "test.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移 User 模型到数据库中
	db.AutoMigrate(&User{})

	// 插入一些示例数据
	db.Create(&User{Name: "Alice", Age: 25, IsActive: true})
	db.Create(&User{Name: "Bob", Age: 30, IsActive: true})
	db.Create(&User{Name: "Charlie", Age: 35, IsActive: false})

	// 使用 Where 查询 Name 为 Alice 的用户
	var user User
	db.Where("name = ?", "Alice").First(&user)
	fmt.Println("User:", user)

	// 使用 Or 查询 Age 为 25 或 30 的用户
	var users []User
	db.Where("age = ?", 25).Or("age = ?", 30).Find(&users)
	fmt.Println("Users:", users)

	// 使用 Not 查询 IsActive 不为 true 的用户
	db.Not("is_active = ?", true).Find(&users)
	fmt.Println("Users:", users)
}
