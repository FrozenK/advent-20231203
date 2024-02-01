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

type matrix struct {
	value    string
	position coordinate
}

func parseLine(line string, y int) ([]number, []symbol, []matrix) {
	symbols := []symbol{}
	numbers := []number{}
	matrixs := []matrix{}

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

			for i := range num {
				m := matrix{
					value: num,
					position: coordinate{
						x: x - len(num) + i,
						y: y,
					},
				}
				matrixs = append(matrixs, m)
			}
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
		n := number{
			value: num,
			position: coordinate{
				x: len(line) - len(num),
				y: y,
			},
		}
		numbers = append(numbers, n)

		for i := range num {
			m := matrix{
				value: num,
				position: coordinate{
					x: len(line) - len(num) + i,
					y: y,
				},
			}
			matrixs = append(matrixs, m)
		}
		num = ""
	}
	return numbers, symbols, matrixs
}

func (s symbol) getRatio(m []matrix) int {
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

	var matrixs []matrix
	for _, c := range mx {
		px := s.position.x + c.x
		py := s.position.y + c.y
		if px < 0 || py < 0 {
			continue
		}
		for _, matrix := range m {
			if matrix.found(coordinate{
				x: px,
				y: py,
			}) {
				if notContains(matrixs, matrix) {
					fmt.Print(" Found : ")
					fmt.Println(matrix)
					matrixs = append(matrixs, matrix)
				}
			}
		}
	}

	if len(matrixs) != 2 {
		return 0
	}

	fmt.Println(matrixs)
	v1, _ := strconv.Atoi(matrixs[0].value)
	v2, _ := strconv.Atoi(matrixs[1].value)
	return v1 * v2
}

func (s symbol) found(c coordinate) bool {
	if c.x == s.position.x && c.y == s.position.y {
		return true
	}
	return false
}

func notContains(m []matrix, i matrix) bool {
	for _, a := range m {
		if a.value == i.value {
			return false
		}
	}
	return true
}

func (m matrix) found(c coordinate) bool {
	if c.x == m.position.x && c.y == m.position.y {
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
	ratio := 0
	f, err := os.Open("input1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	symbols := []symbol{}
	numbers := []number{}
	matrixs := []matrix{}

	y := 0
	lineLen := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineLen = len(line)

		ln, ls, lm := parseLine(line, y)
		numbers = append(numbers, ln...)
		matrixs = append(matrixs, lm...)
		symbols = append(symbols, ls...)
		y++
	}
	for _, n := range numbers {
		if n.isAPart(symbols, lineLen, y) {
			v, _ := strconv.Atoi(n.value)
			sum += v
		}
	}
	for _, s := range symbols {
		fmt.Println(s)
		if s.value != "*" {
			continue
		}
		ratio += s.getRatio(matrixs)
	}
	fmt.Printf("Sum = %d\n", sum)
	fmt.Printf("Ratio = %d\n", ratio)
}
