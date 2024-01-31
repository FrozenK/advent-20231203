package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type coordinate struct {
	x, y int
}
type symbol struct {
	value    string
	position coordinate
}

type number struct {
	value    string
	position coordinate
}

func parseLine(line string, y int) ([]number, []symbol) {
	symbols := []symbol{}
	numbers := []number{}

	numFound := false
	var num string
	for x, c := range line {
		_, err := strconv.Atoi(string(c))
		if err == nil {
			// we have found a partial number
			numFound = true
			num += string(c)
			continue
		}

		if numFound {
			// if we had found a number we add the final number to the number slice
			numFound = false
			n := number{
				value: num,
				position: coordinate{
					x: x - len(num),
					y: y,
				},
			}
			numbers = append(numbers, n)
			fmt.Printf("[ %s - position %d - lenght %d - x %d]\n", num, x-len(num), len(num), x)
			num = ""
		}
		if c == '.' {
			continue
		}

		// we have found a symbol
		s := symbol{
			value: string(c),
			position: coordinate{
				x: x,
				y: y,
			},
		}
		symbols = append(symbols, s)
	}

	if numFound {
		// if we had found a number we add the final number to the number slice
		numFound = false
		n := number{
			value: num,
			position: coordinate{
				x: len(line) - len(num),
				y: y,
			},
		}
		numbers = append(numbers, n)
		num = ""
	}
	return numbers, symbols
}

func (s symbol) found(c coordinate) bool {
	if c.x == s.position.x && c.y == s.position.y {
		return true
	}
	return false
}

func (n number) isAPart(s []symbol, length int, width int) bool {
	mx := []coordinate{
		{x: -1, y: -1},
		{x: -1, y: 0},
		{x: -1, y: 1},
		{x: 0, y: -1},
		{x: 1, y: -1},
		{x: 1, y: 0},
		{x: 1, y: 1},
		{x: 0, y: 1},
	}

	for i := range n.value {
		for _, c := range mx {
			px := n.position.x + i + c.x
			py := n.position.y + c.y
			if px < 0 || py < 0 {
				continue
			}
			for _, symbol := range s {
				if symbol.found(coordinate{
					x: px,
					y: py,
				}) {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	sum := 0
	f, err := os.Open("input1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	symbols := []symbol{}
	numbers := []number{}

	y := 0
	lineLen := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineLen = len(line)

		ln, ls := parseLine(line, y)
		numbers = append(numbers, ln...)
		symbols = append(symbols, ls...)
		y++
	}
	for _, n := range numbers {
		if n.isAPart(symbols, lineLen, y) {
			v, _ := strconv.Atoi(n.value)
			sum += v
		}
	}
	fmt.Printf("Sum = %d\n", sum)
}
