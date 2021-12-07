package main

import (
	"fmt"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	mu := new(sync.Mutex)
	wg.Add(1)
	fmt.Println("запускается функция main")
	firstChan := make(chan int)
	for i := 0; i <10 ; i++ {
		go readOut(firstChan, &wg,i, mu)
	}
	for i := 0; i < 10; i++ {
		go sqrt(firstChan,&wg, mu)
	}
	wg.Wait()

	fmt.Println("the end")
}

func sqrt(f chan int, wg *sync.WaitGroup, mu *sync.Mutex){
	wg.Add(1)
	fmt.Println("запускается функция по возведению в квадрат")
	defer func() {
		wg.Done()
	}()
	 for x := range f {
	 	mu.Unlock()
	 	fmt.Println("запускается цикл по возведению в квадрат")
	 	y := x*x
		fmt.Println("квадрат равен", y)
	}
}
func readOut(downstream chan int, wg *sync.WaitGroup, i int, mu *sync.Mutex) {
	wg.Add(1)
	fmt.Println("запускается функция чтение")
	defer func() {
		fmt.Println("считыватель завершает работу")
		wg.Done()
	}()
	defer func() {
		fmt.Println("считыватель закрывает канал")
		close(downstream)
	}()
	/*for i:= 0; i< 10000; i++ {
		fmt.Println("запускается цикл по сканированию",i)*/
		fmt.Printf("число %v положен в канал\n",i )
		downstream <- i
		mu.Lock()
	}

