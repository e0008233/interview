package threading

import (
	"fmt"
	"sync"
)

func Print1() {
	var chOdd = make(chan int, 1)
	var chEven = make(chan int, 1)

	wg := sync.WaitGroup{}
	chEven <- 1 // 偶数chan中放一个元素
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i <= 100; i += 2 {
			<-chEven
			fmt.Printf("g2||%d\n", i)
			chOdd <- 1
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i += 2 {
			// 如果chOdd 中有数据
			<-chOdd
			fmt.Printf("g1||%d\n", i)
			chEven <- 1
		}
	}()

	wg.Wait()
}

func Print2() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	isEven := true

	wg.Add(2)
	go printNumbers(&wg, &mutex, isEven)
	go printNumbers(&wg, &mutex, !isEven)

	wg.Wait()

}

func printNumbers(wg *sync.WaitGroup, mutex *sync.Mutex, isEven bool) {
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		mutex.Lock()
		if (isEven && i%2 == 0) || (!isEven && i%2 != 0) {
			fmt.Printf("%s: %d\n", map[bool]string{true: "Even", false: "Odd"}[isEven], i)
			isEven = !isEven // toggle isEven flag
		}
		mutex.Unlock()
	}
}
