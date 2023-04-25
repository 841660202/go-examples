package main

import "fmt"

func main() {
	var arr = [5]int{1, 23, 34, 66, 43}
	//定义一个切片，让切片引用一个已创建好的数组
	//表示 slice 引用到 arr 这个数组
	//引用范围数组下标[1,3)
	slice := arr[1:4] //[:]代表全部，[start:]代表开始到结束
	fmt.Println("输出1:------------------------")
	fmt.Println(slice)
	fmt.Println("容量 =", cap(slice)) //一般为切片元素个数的2倍，可动态变化

	fmt.Println("输出2:------------------------")
	//第二种方式使用make
	var slice1 = make([]float64, 5, 10) //必须为[],参数为数据类型，大小，容量
	fmt.Println(slice1)

	fmt.Println("输出3:------------------------")
	//第三种方式，原理类似make
	var slice2 = []int{1, 2, 3}
	fmt.Println(slice2)

	fmt.Println("输出4:------------------------")

	//遍历切片
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	fmt.Println("输出5:------------------------")
	for i, v := range slice {
		fmt.Printf("i = %v,v = %v\n", i, v)
	}

	fmt.Println("输出6:------------------------")
	//切片的动态追加,底层原理是新建数组，然后切片指向新的数组
	slice = append(slice, 88, 99)
	fmt.Println(slice)

	fmt.Println("输出7:------------------------")
	//切片拷贝
	slice3 := make([]int, 10)
	copy(slice3, slice) //把slice拷贝给slice3，必须是切片类型才可以
	fmt.Println(slice3)

	fmt.Println("输出8:------------------------")
	//字符串与切片之间的关系
	str := "hello@go中"
	//string底层是一个数组，因此string可以进行切片
	//string是不可变的，因此不能用string[0]修改字符串
	//如果修改string，将string转成[]byte或[]run（字符处理兼容汉字）切片
	slice4 := str[6:]
	fmt.Println(slice4)
	fmt.Println("输出 rune:------------------------")
	arr1 := []rune(str)
	fmt.Println(arr1)

	arr1[8] = '国'
	str = string(arr1)
	fmt.Println(str)
}

// Go语言之数组与切片 https://blog.csdn.net/Dreaming_My_Dream/article/details/125388538
