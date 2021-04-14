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
	fmt.Printf("%v %v %v \n", date, month, converted)

	//mathOperations()
	//booleanTest()
	bitShifting()
}

func mathOperations() {

	first := 16
	second := 12

	println(first + second)
	println((float32(first) / float32(second)))
	// Error: print((float32(first) / second))

}
func booleanTest() {
	println("---- Booleans --")
	b1 := 8                       // 1000
	b2 := 5                       // 0101
	b3 := 3                       // 0011
	println(b1 & b2)              // add a bit in a position when both bits are on 0000
	println(b1 | b2)              // add a bit in a position where a bit is on (or just an adding operation) 1101
	println(b1^b2, "--", b2^b3)   // add a bit in a position when a bit is "on" JUST in one byte. 1101 (for b1 ^b2), 0110 (for b2^b3)
	println(b1&^b2, "--", b2&^b3) // add a bit in a position when a bit in the first byte is ON and the same position in the second byte is OFF. 1000 (for b1 ^b2), 0100 (for b2^b3)

}

func bitShifting() {
	println("---- bitShifting ---")
	a := 16         // 2^4
	println(a >> 2) // 2^4 / 2^2 -> 16/4
	println(a << 3) // 2^4 * 2^3 -> 16*8
}
