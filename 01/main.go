package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Not technically a heap but also not technically NOT a heap ;)
type minheap struct {
	vals [3]int64
}

func (m *minheap) Push(val int64) bool {
	switch {
	case val < m.vals[0]:
		return false

	case val < m.vals[1]:
		m.vals[0] = val
		return true

	case val < m.vals[2]:
		m.vals[0], m.vals[1] = m.vals[1], val
		return true

	default: // case val >= m.vals[2]:
		m.vals[0], m.vals[1], m.vals[2] = m.vals[1], m.vals[2], val
		return true
	}
}

func (m *minheap) Max() int64 {
	return m.vals[2]
}

func (m *minheap) Sum() int64 {
	return m.vals[0] + m.vals[1] + m.vals[2]
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var calories int64
	var maxes minheap

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			maxes.Push(calories)
			calories = 0
			continue
		}

		n, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			panic(err)
		}

		calories += n
	}

	if calories > 0 {
		maxes.Push(calories)
	}

	fmt.Println(maxes.Max())
	fmt.Println(maxes.Sum())
}
