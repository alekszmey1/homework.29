package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	fmt.Println("запускается функция main")
	var wg sync.WaitGroup
	c0 := make(chan int)
	c1 := make(chan int)
	go readOut (c0, &wg)
	go sqrt (c0, c1, &wg)
	go printGopher (c1, &wg)
	wg.Wait()
	fmt.Println("the end")
}

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
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println("запускается цикл по сканированию")
		line := scanner.Text()
		if line == "стоп" {
			break
		}
		l, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("не верно", err)
			continue
		}
		fmt.Println("вы ввели число", l)
		downstream <- l
	}
}
func sqrt(upstream, downstream chan int, wg *sync.WaitGroup){
	wg.Add(1)
	fmt.Println("запускается функция по возведению квадрата")
	defer func() {
		fmt.Println("множитель завершает работу")
		wg.Done()
	}()
	defer func() {
		fmt.Println("множитель закрывает канал")
		close(downstream)
	}()
	for  item := range upstream {
		fmt.Println("запустился цикл в функции по возведению в квадрат")
		fmt.Println("получаемое значение для возведения в квадрат", item)
		x := item * item
		fmt.Println("значение в квадрате равно", x)
		downstream <- x
	}
}
func printGopher(upstream chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	fmt.Println("запускается функция по умножению на два")
	defer func() {
		fmt.Println("множитель на два завершает работу")
		wg.Done()
	}()
	for  x := range upstream {
		fmt.Println("запускается цикл по умножению на два")
		fmt.Println("число, которое умножаем на два", x)
		y := x * 2
		fmt.Println("ПРОИЗВЕДЕНИЕ НА два", y)
	}
}