package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @Description: 互斥锁   不加锁的话，几个线程回同时访问到一个set，导致读取的数据不准确   在num还未呗设置true的时候被其他协程读取到，造成多次打印。
 */
var m sync.Mutex

var set = make(map[int]bool, 0)

/**
 * @Description: 打印函数
 */
func printFunc(num int) {
	m.Lock()
	defer m.Unlock()
	if _, ok := set[num]; !ok {
		fmt.Println(num)
	}
	set[num] = true
}

func main() {
	for i := 0; i < 10; i++ {
		go printFunc(100)
		//time.Sleep(time.Second)
	}
	time.Sleep(time.Second)
}
