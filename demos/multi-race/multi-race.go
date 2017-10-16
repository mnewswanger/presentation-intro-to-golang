package main

import (
	"bufio"
	"bytes"
	"strconv"
	"sync"

	"github.com/fatih/color"
)

func main() {
	fizzBuzz()
}

func fizzBuzz() {
	buffer := bytes.Buffer{}
	scanner := bufio.NewScanner(&buffer)
	writeWaitGroup := sync.WaitGroup{}
	colors := []func(string, ...interface{}) string{color.RedString, color.BlueString, color.YellowString, color.GreenString}
	writeWaitGroup.Add(len(colors))
	for i, color := range colors {
		go func(wg *sync.WaitGroup, start uint16, outputColor func(string, ...interface{}) string) {
			defer wg.Done()
			printFizzBuzz(start, 100, uint16(len(colors)), &buffer, outputColor)
		}(&writeWaitGroup, uint16(i+1), color)
	}
	writeWaitGroup.Wait()
	for scanner.Scan() {
		println(scanner.Text())
	}
}

func printFizzBuzz(start uint16, end uint16, increment uint16, b *bytes.Buffer, outputFunc func(string, ...interface{}) string) {
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
		b.WriteString(outputFunc(output) + "\n")
	}
}
