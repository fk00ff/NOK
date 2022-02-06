package main

import (
	"fmt"
	"os"
	"strconv"
)

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// Max returns the larger of x or y.
func Max(x, y int) int {
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
			if num/result[j] == 0 {
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
	return nil
}

func nok(div1, div2 int) int {
	div_min := Min(div1, div2)
	div_max := Max(div1, div2)

	fmt.Println(div_min, div_max)

	simple := GetSimple()
	d_min := Divide(div_min, simple)
	// d_max:=Divide(div_max, simple)

	for num, cnt := range d_min {
		if d_min[num] != 0 {
			d_min[num] = cnt + 1
		}
	}

	return 1

}

func main() {
	fmt.Print("test")

	if len(os.Args) < 3 {
		fmt.Println("Need TWO args")
		os.Exit(2)
	}

	d1, _ := strconv.Atoi(os.Args[1])
	d2, _ := strconv.Atoi(os.Args[2])

	fmt.Print(nok(d1, d2))
}