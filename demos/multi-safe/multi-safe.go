package main

import (
	"strconv"
	"sync"

	"github.com/fatih/color"
)

func main() {
	fizzBuzz()
}

func fizzBuzz() {
	c := make(chan string, 1000)
	writeWaitGroup := sync.WaitGroup{}
	colors := []func(string, ...interface{}) string{color.RedString, color.BlueString, color.YellowString, color.GreenString}
	writeWaitGroup.Add(len(colors))
	for i, color := range colors {
		go func(wg *sync.WaitGroup, start uint16, outputColor func(string, ...interface{}) string) {
			defer wg.Done()
			printFizzBuzz(start, 100, uint16(len(colors)), c, outputColor)
		}(&writeWaitGroup, uint16(i+1), color)
	}
	printWaitGroup := sync.WaitGroup{}
	printWaitGroup.Add(1)
	go func(c chan string, wg *sync.WaitGroup) {
		defer wg.Done()
		for s := range c {
			println(s)
		}
	}(c, &printWaitGroup)
	writeWaitGroup.Wait()
	close(c)
	printWaitGroup.Wait()
}

func printFizzBuzz(start uint16, end uint16, increment uint16, c chan string, outputFunc func(string, ...interface{}) string) {
	output := ""
	for i := start; i <= end; i += increment {
		if i%3 == 0 || i%5 == 0 {
			output = ""
			if i%3 == 0 {
				output += "Fizz"
			}
			if i%5 == 0 {
				output += "Buzz"
			}
		} else {
			output = strconv.Itoa(int(i))
		}
		c <- outputFunc(output)
	}
}
