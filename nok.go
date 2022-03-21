package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// Min returns the smaller of x or y.
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// Max returns the larger of x or y.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func GetSimple() []int {
	result := make([]int, 100)

	result[0] = 2

	num := 3
	counter := 1

	for counter < 100 {
		for j := 0; j < counter; j = j + 1 {
			if num%result[j] == 0 {
				goto next
			}
		}

		result[counter] = num
		counter++

	next:
		num = num + 1
	}

	return result
}

func Divide(num int, simple []int) map[int]int {
	result := make(map[int]int)

	for _, smpl := range simple {
		for {
			if num < smpl {
				goto fin
			}

			if num%smpl == 0 {
				result[smpl] = result[smpl] + 1
				num = num / smpl
			} else {
				break
			}
		}
	}

fin:
	return result
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func nok(div1, div2 int) int {
	div_min := min(div1, div2)
	div_max := max(div1, div2)

	fmt.Println(div_min, div_max)

	simple := GetSimple()
	d_min := Divide(div_min, simple)
	d_max := Divide(div_max, simple)

	for num, cnt := range d_min {
		if d_max[num] != 0 {
			d_max[num] = d_max[num] - min(d_max[num], cnt)
		}
	}

	result := 1

	for num, cnt := range d_min {
		result = result * powInt(num, cnt)
	}

	for num, cnt := range d_max {
		result = result * powInt(num, cnt)
	}

	return result

}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Need TWO args")
		os.Exit(2)
	}

	d1, _ := strconv.Atoi(os.Args[1])
	d2, _ := strconv.Atoi(os.Args[2])

	fmt.Print(nok(d1, d2))
}
