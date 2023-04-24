package main

import (
	"fmt"
	"sort"
)

func main() {
	//声明map并不会分配内存
	var a map[string]string
	//使用map前需要make，make的作用就是给map分配内存空间
	a = make(map[string]string, 10) //10：创建长度
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no3"] = "卢俊义"
	a["no3"] = "武松"
	fmt.Println("输出1:", a)

	//声明并初始化
	var heroes = map[string]string{"hero1": "刘备",
		"hero2": "关羽",
		"hero3": "张飞"}
	fmt.Println("输出2:", heroes)

	//value是map
	student := make(map[string]map[string]string) //空间大小不写也可以
	student["01"] = make(map[string]string)
	student["01"]["name"] = "tom"
	student["01"]["sex"] = "男"

	student["02"] = make(map[string]string)
	student["02"]["name"] = "jack"
	student["02"]["sex"] = "女"
	fmt.Println("输出3:", student)

	//增删改查
	heroes["hero4"] = "吕布"  //如果key存在则是增加，否则为修改
	delete(heroes, "hero4") //如果key不存在，删除不会操作，也不会报错
	//一次性删除所有的key，可遍历key删除，可创建新的空间
	heroes = make(map[string]string) //推荐
	val, findRes := a["no2"]         //val是key对应的值，findRes 为是否存在key
	if findRes {
		fmt.Println("输出4:", "key的值为%v\n", val)
	}

	fmt.Println("输出5:------------------------")
	//遍历map，只能用for-range,无序输出
	for s, s2 := range a {
		fmt.Printf("key=%v,value=%v\n", s, s2)
	}
	fmt.Println("输出6:------------------------")
	for s, m := range student {
		fmt.Println("key=", s)
		for s2, s3 := range m {
			fmt.Printf("\tkey=%v,value=%v\n", s2, s3)
		}
	}
	fmt.Println("输出7:------------------------")
	//map的长度,有多少对key-value
	fmt.Println("map 的长度", len(a))

	//map切片，可以理解为数组元素为map,可以动态增加map
	monster := make([]map[string]string, 2)

	if monster[0] == nil {
		monster[0] = make(map[string]string)
		monster[0]["name"] = "牛魔王"
		monster[0]["age"] = "400"
	}
	if monster[1] == nil {
		monster[1] = make(map[string]string)
		monster[1]["name"] = "狮子精"
		monster[1]["age"] = "300"
	}
	newMonster := map[string]string{"name": "白骨精", "age": "500"}
	monster = append(monster, newMonster)
	fmt.Println("输出8:------------------------")
	fmt.Println(monster[0])
	fmt.Println(monster[1])
	fmt.Println(monster[2])
	fmt.Println(monster)

	fmt.Println("输出9:------------------------")
	//按照map的key进行排序输出
	map1 := map[int]int{15: 14, 67: 23, 34: 45, 20: 11}
	keys := []int{}
	for key, _ := range map1 { //创建切片将key存入切片
		keys = append(keys, key)
	}
	sort.Ints(keys) //递增顺序排列
	for _, key := range keys {
		fmt.Printf("map[%v] = %v\n", key, map1[key])
	}
}

// Go 基础篇之 Map: https://zhuanlan.zhihu.com/p/608079943
