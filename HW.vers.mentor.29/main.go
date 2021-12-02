package main

import (

	"fmt"
	"log"
	"strconv"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	in := sender(&wg)
	middle := summator(in,&wg)
	out := multiplier(middle,&wg)
	receiver(out,&wg)
	wg.Wait()
	fmt.Println("the end!!!")

}

func sender (wg *sync.WaitGroup)chan int {
	out := make(chan int)
	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("отправитель завершает работу")
			wg.Done()
		}()
		defer func() {
			fmt.Println("отправитель закрывает канал")
			close(out)
		}()
		var scan string
		var digit int
		for  {
			_,err := fmt.Scan(&scan)
			if err != nil {
				log.Println(err)
				continue
			}
			digit, err = strconv.Atoi(scan)
			if err != nil {
				if scan == "stop"{
					break
				}
				log.Println(err)
				continue
			}
			fmt.Printf("отправитель отправил %v\n", digit)
			out <-digit
		}
	}()
	return out
}

func summator (in chan int, wg *sync.WaitGroup) chan int {
	out :=make(chan int)
	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("сумматор завершает работу")
			wg.Done()
		}()
		defer func() {
			fmt.Println("сумматор закрывает канал")
			close(out)
		}()
		for value := range in {
			result  := value + value
			fmt.Printf("сумматор принял %v, Отправил %v\n",value,result)
			out <- result
		}
	}()
	return out
}

func multiplier(in chan int, wg *sync.WaitGroup) chan int {
	out := make(chan int)
	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("мультиплеер завершает работу")
			wg.Done()
		}()
		defer func() {
			fmt.Println("мультиплеер закрывает канал")
			close(out)
		}()
		for value := range in {
			result  := value * value
			fmt.Printf("мультиплеер принял %v, мультиплеер отправил %v\n",value,result)
			out <- result
		}
	}()
return out
}

func receiver (in chan int,wg *sync.WaitGroup)  {
	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("получатель завершает работу")
			wg.Done()
		}()
		for value := range in{
			fmt.Printf("получатель принял %v\n", value)
		}
	}()
}