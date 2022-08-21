package specialquestion

import (
	"fmt"
	"sync"
)

func PrintInfoWithTwoThread() {
	chanNum := 4
	chanQueue := make([]chan struct{}, chanNum)
	var result = 0
	exitChan := make(chan struct{})
	for i := 0; i < chanNum; i++ {
		chanQueue[i] = make(chan struct{})
		if i == chanNum-1 {
			go func(i int) {
				chanQueue[i] <- struct{}{}
			}(i)
		}
	}
	for i := 0; i < chanNum; i++ {
		var lastChan, curChan chan struct{}
		if i == 0 {
			lastChan = chanQueue[chanNum-1]
		} else {
			lastChan = chanQueue[i-1]
		}
		curChan = chanQueue[i]
		go func(i byte, lastChan, curChan chan struct{}) {
			for {
				if result > 20 {
					exitChan <- struct{}{}
				}
				<-lastChan
				fmt.Printf("%c\n", i)
				result++
				curChan <- struct{}{}
			}
		}('A'+byte(i), lastChan, curChan)
	}
	<-exitChan
	fmt.Println("done")

}

// 交替打印AB
func PrintThreadAB() {
	aChan := make(chan struct{}, 1)
	bChan := make(chan struct{}, 1)
	count := 5
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < count; i++ {
			<-aChan
			fmt.Println("A")
			bChan <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < count; i++ {
			<-bChan
			fmt.Println("B")
			// 无缓冲channel死锁的坑，最后一个 B 打印完不用给 A 发信号了，A 的循环已退出
			if i == count-1 {
				return
			}
			aChan <- struct{}{}
		}
	}()

	aChan <- struct{}{}
	wg.Wait()
}

// gorounting泄露
func GoroutineLeak() {
	chanInt := make(chan int)
	defer close(chanInt)
	go func() {
		for {
			res := <-chanInt
			//res,ok := <-chanInt
			//if !ok {
			//     break
			//}
			fmt.Println(res)
		}
	}()
	chanInt <- 1
	chanInt <- 1
}
