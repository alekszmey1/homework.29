package main

import (
	"fmt"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	firstChan := make(chan int)
	go readOut(firstChan, &wg)
	go sqrt(firstChan,&wg)

	wg.Wait()
	fmt.Println("the end")

}

func sqrt(f chan int, wg *sync.WaitGroup){
	wg.Add(1)
	defer func() {
		wg.Done()
	}()
	 for x := range f {
	 	y := x*x
		fmt.Println("квадрат равен", y)
	}
}
func readOut(downstream chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	defer func() {
		fmt.Println("считыватель завершает работу")
		wg.Done()
	}()
	defer func() {
		fmt.Println("считыватель закрывает канал")
		close(downstream)
	}()
	for i:= 0; i< 10 ; i++ {
				fmt.Println("вы ввели число", i)
		downstream <- i
	}
}