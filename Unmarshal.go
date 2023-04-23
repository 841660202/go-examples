package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ProductId int64 `json:"product_id"`
	Name      string
	ImgUrl    string `json:"img_url"`
	Number    int
	Price     float64
	IsOnSale  bool
	Color     []string
}

type Book struct {
	Name     string
	Author   string
	pageSize int
}

// 使用Marshal把Product转换为json格式 , 参数为对象的方式
func marshalToJson1(product Product) string {
	b, err := json.Marshal(product)
	if err == nil {
		return string(b)
	}
	return "错误"
}

// 使用Marshal把Product转换为json格式 , 参数为指针的方式
func marshalToJson2(book *Book) string {
	b, err := json.Marshal(book)
	if err == nil {
		return string(b)
	}
	return "错误"
}

func main() {
	// ----------------- 测试json.Marshal -----------------
	colorArr := []string{"红色", "蓝色", "绿色"}
	product := Product{1, "Mac Pro电脑", "http://mac-pro.com", 12000, 12, true, colorArr}
	fmt.Println("\n--------------------------------------------")
	fmt.Println(marshalToJson1(product))
	// 结果：{"product_id":1,"Name":"Mac Pro电脑","img_url":"http://mac-pro.com","Number":12000,"Price":12,"IsOnSale":true,"Color":["红色","蓝色","绿色"]}
	// 说明：如果变量打上了json标签，则导出的json数据格式按照定义的json标签导出（product_id, img_url）

	book := Book{}
	book.Name = "GoLang"
	book.Author = "哈哈哈"
	book.pageSize = 500
	fmt.Println(marshalToJson2(&book))
	// 结果：{"Name":"GoLang","Author":"哈哈哈"}
	// 说明：pageSize没有导出来，只有数据结构中变量首字母大写才可以导出。

	var booklist []Book // 定义一个书籍的切片
	book1 := Book{"java", "小小", 801}
	book2 := Book{"python", "天天", 621}
	book3 := Book{"零值", "测试", 0}
	booklist = append(booklist, book1)
	booklist = append(booklist, book2)
	booklist = append(booklist, book3)
	bookJson, err := json.Marshal(booklist)
	if err == nil {
		fmt.Println("===> 切片转json：" + string(bookJson))
	}
	// 结果：[{"Name":"java","Author":"小小"},{"Name":"python","Author":"天天"},{"Name":"零值","Author":"测试"}]
	// 说明：零值的没有导出

	// ----------------- 测试json.MarshalIndent -----------------
	jsIndent, _ := json.MarshalIndent(&book, "", "\t")
	fmt.Println("===> 测试MarshalIndent结果: ")
	fmt.Println(string(jsIndent))
	// 结果如下：
	//{
	// "Name": "GoLang",
	// "Author": "哈哈哈"
	//}
	jsIndent1, _ := json.MarshalIndent(&book, "哈哈", "\t")
	fmt.Println("===> 测试MarshalIndent结果: ")
	fmt.Println(string(jsIndent1))
	//{
	//哈哈    "Name": "GoLang",
	//哈哈    "Author": "哈哈哈"
	//哈哈}
	// 说明：MarshalIndent对读的结果做了一些处理，简单说就是json多了一些格式处理。
	fmt.Println("\n------------------测试json.Unmarshal--------------------------")

	product1 := Product{1, "Mac Pro电脑", "http://mac-pro.com", 12000, 12, true, colorArr}
	// ----------------- 测试json.Unmarshal -----------------
	productstr := marshalToJson1(product1) // {"product_id":1,"Name":"Mac Pro电脑","img_url":"http://mac-pro.com","Number":12000,"Price":12,"IsOnSale":true}
	fmt.Println(productstr)
	fmt.Println("\n--------------------------------------------")

	// productstr := marshalToJson1({"product_id":1,"Name":"Mac Pro电脑"}) // {"product_id":1,"Name":"Mac Pro电脑","img_url":"http://mac-pro.com","Number":12000,"Price":12,"IsOnSale":true}
	productresult := Product{}
	err1 := json.Unmarshal([]byte(productstr), &productresult)
	fmt.Println("===> 测试json.Unmarshal结果: ")
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(productresult)
	}
	// 结果：{1 Mac Pro电脑 http://mac-pro.com 12000 12 true [红色 蓝色 绿色]}
	for _, val := range productresult.Color {
		fmt.Println("===>color: ", val)
	}
	//===>color:  红色
	//===>color:  蓝色
	//===>color:  绿色
	fmt.Println("\n--------------------------------------------")

	fmt.Println(bookJson)
	fmt.Println("\n--------------------------------------------")

	var booklistdata []Book                                 // 定义一个切片
	err2 := json.Unmarshal([]byte(bookJson), &booklistdata) // bookJson为：[{"Name":"java","Author":"小小"},{"Name":"python","Author":"天天"},{"Name":"零值","Author":"测试"}]
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(booklistdata)
	}
	// 结果： [{java 小小 0} {python 天天 0} {零值 测试 0}]

	// 说明：Json Unmarshal是将json字符串解码到相应的数据结构中。 json字符串解析时，需要一个"接收体"接收解析后的数据，
	// 并且这个接收体必须是传递指针（例如：&productresult），否则解析虽然不报错，但数据无法赋值到接收体中。
}

// Gin框架中处理请求参数的零值问题
// https://juejin.cn/post/7016514018487566343

// 解决go gin框架 binding:"required"`无法接收零值的问题
// https://www.cnblogs.com/rainbow-tan/p/15457818.html
