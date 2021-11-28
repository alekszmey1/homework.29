package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	c0 := make(chan int)
	c1 := make(chan int)
	go readOut(c0)
	go sqrt(c0, c1)
	printGopher(c1)
}

func readOut(downstream chan int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "стоп" {
			close(downstream)
			return
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

func sqrt(upstream, downstream chan int){
	for  item := range upstream {
		fmt.Println("получаемое значение для возведения в квадрат", item)
		x := item * item
		fmt.Println("значение в квадрате равно", x)
		downstream <- x
	}
	close(downstream)
}

func printGopher(upstream chan int) {
	x:= <- upstream
	fmt.Println("число, которое умножаем на два",x)
	x = x * 2
	fmt.Println("ПРОИЗВЕДЕНИЕ НА два",x)
}