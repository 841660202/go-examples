package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var mutex sync.Mutex
	wait := sync.WaitGroup{}

	fmt.Println("Locked")
	mutex.Lock()
	// for i := 1; i <= 3; i++ {
	// 	fmt.Println("普通打印：", i)
	// }
	// fmt.Println("-------------------------")
	for i := 1; i <= 3; i++ {
		// fmt.Println("普通打印：", i)
		wait.Add(1)

		go func(i int) {
			fmt.Println("Not lock:", i)

			mutex.Lock()
			fmt.Println("Lock:", i)

			time.Sleep(time.Second)

			fmt.Println("Unlock:", i)
			mutex.Unlock()

			defer wait.Done()
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Unlocked")
	// 锁永远得不到释放就是死锁
	// fatal error: all goroutines are asleep - deadlock!
	// time.Sleep(10 * time.Second)
	mutex.Unlock()
	wait.Wait()

}

// https://shockerli.net/post/golang-pkg-mutex/
// https://blog.csdn.net/K346K346/article/details/95641101
// 多次执行输出结果中的先后顺序是不定的，因为多协程并发执行的顺序是随机的。

// [Running] go run "/Users/haotian/Desktop/go-example/go程/sync2.go"
// Locked
// Not lock: 1
// Not lock: 3
// Not lock: 2
// Unlocked
// Lock: 1
// Unlock: 1
// Lock: 3
// Unlock: 3
// Lock: 2
// Unlock: 2

// [Done] exited with code=0 in 4.645 seconds

// [Running] go run "/Users/haotian/Desktop/go-example/go程/sync2.go"
// Locked
// Not lock: 3
// Not lock: 1
// Not lock: 2
// Unlocked
// Lock: 3
// Unlock: 3
// Lock: 1
// Unlock: 1
// Lock: 2
// Unlock: 2

// [Done] exited with code=0 in 4.184 seconds

// [Running] go run "/Users/haotian/Desktop/go-example/go程/sync2.go"
// Locked
// Not lock: 1
// Not lock: 2
// Not lock: 3
// Unlocked
// Lock: 1
// Unlock: 1
// Lock: 2
// Unlock: 2
// Lock: 3
// Unlock: 3

// [Done] exited with code=0 in 4.179 seconds
