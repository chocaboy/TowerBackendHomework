package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var x int64
	var num int
	var bit string
	fmt.Print("Number: ")
	fmt.Scan(&x)
	fmt.Print("Number of bit: ")
	fmt.Scan(&num)
	fmt.Print("Bit to change to: ")
	fmt.Scan(&bit)

	bin_x := strconv.FormatInt(x, 2)
	runes := []rune(bin_x)
	runes[len(bin_x)-num] = rune(bit[0]) // через срезы
	fmt.Println(string(runes))

	if bit == "1" {
		mask, _ := strconv.ParseInt("1"+strings.Repeat("0", num-1), 2, 64) // через маски
		fmt.Println(strconv.FormatInt(x|mask, 2))
	} else {
		mask, _ := strconv.ParseInt(strings.Repeat("1", len(strconv.FormatInt(x, 2))-num)+"0"+strings.Repeat("1", num-1), 2, 64)
		fmt.Println(strconv.FormatInt(x&mask, 2))
	}
}
