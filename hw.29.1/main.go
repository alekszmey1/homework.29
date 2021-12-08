package main

import (
	"fmt"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("запускается функция main")
	firstChan := make(chan int)
	/*c := make(chan os.Signal, 1)
	signal.Notify(c,syscall.SIGINT,syscall.SIGTERM)*/
	go readOut(firstChan, &wg)
	go sqrt(firstChan,/*c*/ &wg)
	wg.Wait()
	fmt.Println("the end")
}

func sqrt(f chan int,/*c chan os.Signal,*/ wg *sync.WaitGroup){
	//wg.Add(1)
	fmt.Println("запускается функция по возведению в квадрат")
	defer func() {
		wg.Done()
	}()
	 for x := range f {
	 	fmt.Println("запускается цикл по возведению в квадрат")
	 	y := x*x
		fmt.Println("квадрат равен", y)
		 /*select {
		 case <-f:
		 	x := <-f
		 	y:=x*x
		 	fmt.Printf("квадрат %v равен %v\n",x,y)
		 case <-c:
		 	break
		 default:
		 	fmt.Println("выхожу из программы")*/
		 }
	}
//}
func readOut(downstream chan int, wg *sync.WaitGroup) {
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
	for i := 0; i < 1000000; i++ {
		fmt.Println("запускается цикл по сканированию", i)
			fmt.Printf("число %v положен в канал\n", i)
		downstream <- i
	}
}

