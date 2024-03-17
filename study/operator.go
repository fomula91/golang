package study

import "fmt"

func Operator() {
	operator()
	boolOperator()
}

func operator() {
	// 수식 연산자는 기본적인 사칙연산과 나머지 연산자도 사용할수 있다.
	// + 연산자는 문자열도 연산가능하다

	num1, num2 := 11, 9
	str1, str2 := "Hello", "World"

	fmt.Println("num1 + num2 ==> ",num1 + num2)
	fmt.Println("str1 + str2 ==> ",str1 + str2)

	// 증감연산자도 사용할수 있다.
	// 증감연산자는 대입과 동시에 사용할수 없다.
	// 증감연산자는 전위연산을 사용할수 없다.

	count1, count2 := 10, 14.1

	count1++
	count2--

	fmt.Println("count1 ==> ",count1)
	fmt.Println("count2 ==> ",count2)

	// 할당연산자
	a := 2
	var num int

	num = a
	fmt.Println("num = a :", num)
	num += 4
	fmt.Println("num += 4 :", num)
	num -= 2
	fmt.Println("num -= 2 :", num)
	num *= 5
	fmt.Println("num *= 5 :", num)
	num /= 2
	fmt.Println("num /= 2 :", num)
	num %= 3
	fmt.Println("num %= 3 :", num)

	num = 3  //00000011
	num &= 2 //00000010
	fmt.Printf("num &= 2 : %08b, %d\n", num, num)
	num |= 5 //00000101
	fmt.Printf("num |= 5 : %08b, %d\n", num, num)
	num ^= 4 //00000100
	fmt.Printf("num ^= 4 : %08b, %d\n", num, num)
	num &^= 2 //00000010
	fmt.Printf("num &^= 2 : %08b, %d\n", num, num)
	num <<= 9 //00001001
	fmt.Printf("num <<= 9 : %08b, %d\n", num, num)
	num >>= 8 //00001000
	fmt.Printf("num >>= 8 : %08b, %d\n", num, num)
}

func boolOperator() {
	var a bool = true
	b := false

	fmt.Println("0 && 0 : ", b && b)
	fmt.Println("0 && 1 : ", b && a)
	fmt.Println("1 && 1 : ", a && a)
	fmt.Println("0 || 0 : ", b || b)
	fmt.Println("0 || 1 : ", b || a)
	fmt.Println("1 || 1 : ", a || a)

	fmt.Println("!1 ", !true)
	fmt.Println("!0 ", !false)
}