package main // main 패키지로 선언
import "fmt" // fmt 패키지를 가져옴

// golang에서는 fmt를 사용하지 않고도 Print, Println 등을 사용할 수 있음
// 하지만 fmt 패키지를 가져와서 사용하는 것이 좋음
// 왜냐하면 fmt 패키지는 표준 라이브러리로서 표준 출력을 제공하기 때문
func helloWorld() {
	// 변수 선언
	// var 변수명 자료형 = 값
	// 자료형을 생략하고 값을 할당하면 자료형을 추론하여 할당
	// var 변수명 = 값
	// 변수명 := 값
	// 위와 같이 선언하면 자료형을 추론하여 할
	var hello string = "hello world!"

	// 상수 선언
	// const 상수명 자료형 = 값
	// 자료형을 생략하고 값을 할당하면 자료형을 추론하여 할당
	// const 상수명 = 값
	// 상수명 := 값
	const hello2 string = "hello world2!"

	var quantity int
	var length, width float64
	var customerName string

	// 값 할당
	// 변수를 선언 한 이후 값을 할당할수 있음
	// 한번에 여러개를 설정할수 있다.


	quantity = 4
	length, width = 1.2, 2.4
	customerName = "Damon Cole"

	// 출력
	fmt.Print(hello)
	fmt.Print(hello2)	
	fmt.Print(quantity)
	fmt.Print(length, width)
	fmt.Print(customerName)

}

func zeroValue() {
	// 변수를 선언만 하고 값을 할당하지 않으면
	// zero value로 초기화 됨
	// 숫자형의 경우 0
	// bool형의 경우 false
	// 문자열의 경우 ""
	var quantity int
	var length, width float64
	var customerName string

	fmt.Print(quantity)
	fmt.Print(length, width)
	fmt.Print(customerName)
}

func shortHand() {
	// 변수를 선언할 때 선언과 동시에 값을 할당하는 경우에는 보통 단축 변수 선언을 많이 사용함
	// 변수를 선언만 하고 값을 할당하지 않으면
	// zero value로 초기화 됨
	// 숫자형의 경우 0
	// bool형의 경우 false
	// 문자열의 경우 ""
	quantity := 4
	length, width := 1.2, 2.4
	customerName := "Damon Cole"

	fmt.Print(quantity)
	fmt.Print(length, width)
	fmt.Print(customerName)

}
func main() {

	helloWorld()
	zeroValue()
	shortHand()
}