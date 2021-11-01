package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "стоп" {
			return
		}
		l, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("не верно", err)
			continue
		}
		fmt.Println(l)
		intFirstChan := make(chan int)
		intFirstChan <- l
		close(intFirstChan)
		//sqrt(intFirstChan)
		//var wg sync.WaitGroup
		//wg.Add(2)
		//go sqrt(intFirstChan)
		multiplication(sqrt(intFirstChan))
		//wg.Wait()
	}
}
func sqrt(s chan int) chan int{
	y := <-s
	intSecondChan := make(chan int)
	go func() {
		fmt.Println("получаемое значение для возведения в квадрат", y)
		y = y * y
		fmt.Println("значение в квадрате равно",  y)
		intSecondChan <- y
		close(intSecondChan)
	}()
	return intSecondChan
}
func multiplication(y chan int) int{
	go func() int{
		x:= <- y
		fmt.Println("число, которое умножаем на два",x)
		x = x * 2
		fmt.Println("ПРОИЗВЕДЕНИЕ НА два",x)
		return x
	}()
	return 847
}