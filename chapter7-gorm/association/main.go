// gorm 关联模式
// 查找关联
// 添加关联
// 替换关联
// 删除关联
// 清空关联
// 关联计数

package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Email     string
	Password  string
	Addresses []Address
}

type Address struct {
	gorm.Model
	Street  string
	City    string
	State   string
	ZipCode string
	UserID  uint
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{}, &Address{})

	// 添加一个用户
	user := User{
		Name:     "John",
		Email:    "john@example.com",
		Password: "password",
	}

	// 判断数据是否存在，不存在就添加

	db.FirstOrCreate(&user)

	// 添加一个地址到用户
	// address := Address{
	// 	Street:  "123 Main St",
	// 	City:    "Anytown",
	// 	State:   "CA",
	// 	ZipCode: "12345",
	// 	UserID:  user.ID,
	// }

	// db.Create(&address)

	// 查找一个用户及其地址
	// var user1 User
	// db.Preload("Addresses").First(&user1, 1)

	// 打印用户信息及地址
	// fmt.Println(user1.Name)
	// for _, address := range user1.Addresses {
	// 	fmt.Println(address.Street, address.City, address.State, address.ZipCode)
	// }

	// // 替换一个用户的地址
	// var address2 Address
	// db.First(&address2, 2)
	// db.Model(&user).Association("Addresses").Replace(&address2)

	// // 删除一个用户的地址
	// db.Model(&user).Association("Addresses").Delete(&address)

	// // 清空一个用户的地址
	// db.Model(&user).Association("Addresses").Clear()

	// // 获取一个用户的地址数量
	// count := db.Model(&user).Association("Addresses").Count()
	// fmt.Println(count)
}
