package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

func printNumbers(num int, c0 chan int, c1 chan int, cerr chan error) {
	for i := 1; i <= num; i++ {
		if i > 10 {
			cerr <- errors.New(fmt.Sprintf("Warning: i=%d is larger than 10", i))
			continue
		}
		if i%2 == 1 {
			c1 <- i
		} else {
			c0 <- i
		}
	}
}

func printLetters() {
	var letter rune = 'a'
	for i := 0; i < 10; i++ {
		fmt.Printf("%c\n", letter+rune(i))
	}
}

func main() {
	var wg sync.WaitGroup
	c0 := make(chan int)
	c1 := make(chan int)
	cerr := make(chan error)
	num := 20
	wg.Add(3)
	go func() {
		for i := 0; i < num; i++ {
			select {
			case x := <-c0:
				fmt.Println("Even number:", x)
			case x := <-c1:
				fmt.Println("Odd number:", x)
			case x := <-cerr:
				log.Println(x)
			}
		}
		wg.Done()
	}()
	go func() {
		printNumbers(num, c0, c1, cerr)
		wg.Done()
	}()
	go func() {
		printLetters()
		wg.Done()
	}()
	wg.Wait()
}
