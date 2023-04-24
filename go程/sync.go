package

import {

"sync"
}

var l sync.Mutex
var a string

func f() {
	a = "hello, world"
	l.Unlock()
}

func main() {
	l.Lock()
	go f()
	l.Lock()
	print(a)
}


// https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-sync-primitives/