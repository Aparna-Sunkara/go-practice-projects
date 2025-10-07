package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("c")

	var a int = 123
	var b uint = 10
	var c byte = 's'
	var d float32 = 333.5
	var f rune = 555

	fmt.Println(ConvertStringToInt("123"))
	fmt.Println(ConvertStringToUint("10"))
	fmt.Println(ConvertstringTobyte("s"))
	fmt.Println(ConvertStringTofloat32("333.5"))
	fmt.Println(ConvertStringToRune("555"))
	fmt.Println(a, b, c, d, f)

	fmt.Println(ConvertInttofloat32(10))
	fmt.Println(convertInttostring(15))
	fmt.Println(convertInttobyte(400))
	fmt.Println(convertintTorune(700))

}

func ConvertStringToInt(s string) (int, error) {
	num, error := strconv.Atoi(s)
	return num, error
}

func ConvertStringToUint(s string) (uint, error) {
	num, error := strconv.ParseUint(s, 10, 32)
	return uint(num), error
}
func ConvertstringTobyte(s string) (byte, error) {
	num, error := strconv.Atoi(s)
	return byte(num), error
}
func ConvertStringToRune(s string) (rune, error) {
	num, error := strconv.Atoi(s)
	return rune(num), error

}
func ConvertStringTofloat32(s string) (float32, error) {
	num, error := strconv.ParseFloat(s, 1)
	return float32(num), error
}

func ConvertInttofloat32(i int) (float32, error) {
	return float32(i), nil
}
func convertInttostring(i int) (string, error) {
	return string(i), nil
}
func convertInttobyte(i int) (byte, error) {
	return byte(i), nil
}
func convertintTorune(i int) (rune, error) {
	return rune(i), nil
}
