package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once
var onceBody = func() {
	fmt.Println("Only once")
}

func main() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			once.Do(onceBody)
			fmt.Println("i=", i)
		}(i)
	}
	time.Sleep(time.Second) // 睡眠 1s 等待 go 程执行完，注意睡眠时间不能太短。
}

// 程序运行输出：

// Only once
// i= 3
// i= 4
// i= 0
// i= 1
// i= 2
// 从输出结果可以看出，尽管 for 循环每次都会调用 once.Do() 方法，但是函数 onceBody() 却只会被执行一次。
