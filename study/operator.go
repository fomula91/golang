package study

import "fmt"

func Operator() {
	operator()
	boolOperator()
	chanelOperator()
	pointerOperator()
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

func chanelOperator() {
	ch := make(chan int)

	go func ()  {
		ch <- 10
	}()

	result := <- ch
	fmt.Println(result)
}

func pointerOperator() {
	var num int = 2
	var pnum = &num

	fmt.Println("num : ", num)   //num 값
	fmt.Println("pnum :", pnum)  //num의 메모리 주소
	fmt.Println("pnum :", *pnum) //num의 주소로 메모리에 할당돼있는 값 접근

	*pnum++
	fmt.Println("num : ", num)
	fmt.Println("pnum :", *pnum)
	//포인터 연산자를 이용한 값 변경
}

// 연산자 우선순위
// 1. (), [], ->, ++, -- ::: 결합방향 ->
// 2. +, -, !, ~, (type), *, &, sizeof ::: 결합방향 <-
// 3. *,/,% ::: 곱셈, 나눗셈 관련 연산	결합방향 →
// 4 +,- ::: 덧셈, 뺄셈 결합방향 →
// 5 <<, >> ::: 비트 이동 결합방향 →
// 6 <, <=, >, => ::: 대소 비교  결합방향 →
// 7 ==, != ::: 동등 비교 결합방향 →
// 8 & ::: 비트 AND 결합방향 →
// 9 ^ ::: 비트 XOR 결합방향 →
// 10 | ::: 비트 OR 결합방향 →
// 11 && ::: 논리 AND 결합방향 →
// 12 || ::: 논리 OR 결합방향 →
// 13 ? : ::: 조건 연산 결합방향 ←
// 14 =, +=, -=, *=, /=, %=, <<=, >>=, &=, ^=, |= ::: 대입 연산 결합방향 ←
// 15 , ::: 콤마 연산 결합방향 →