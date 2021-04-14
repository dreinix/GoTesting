package main

import (
	"fmt"
	"strconv"
)

var date = 16

func main() {
	/*var date int
	date = 16*/
	var month int = 12
	year := 99
	converted := strconv.Itoa(year)
	fmt.Printf("%v %v %v", date, month, converted)
}
