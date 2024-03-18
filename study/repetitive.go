package study

import "fmt"
func Repetitive() {
	fmt.Println("for문")
	repeat()
	repeat2()
	repeat3()
	forRange()
}

// go에서는 while문을 제공하지 않음 for문만 사용가능
// for 초기식; 조건식; 변화식 {
// 	// 반복할 코드
// }

func repeat() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//조건식만 사용해서 더욱 간결한 코드를 작성가능
func repeat2() {
	sum := 0
	i := 1
	for i <= 100 {
		sum += i
		i++
	}
	fmt.Println(sum)
}

//무한루프
func repeat3() {
	sum := 0
	for {
		fmt.Println("무한루프입니다. \n", sum)
		sum++
		if sum == 100 {
			break
		}
	}
	
}

func forRange() {
	var arr [6]int = [6]int{1,2,3,4,5,6}

	for index, num := range arr {
		fmt.Printf("arr[%d]의 값은 %d입니다. \n", index, num)
	}

	//index는 생략이 가능
	for _, num := range arr {
		fmt.Printf("arr의 값은 %d입니다. \n", num)
	}
	for index := range arr {
		fmt.Printf("arr[%d]의 값은 %d입니다. \n", index, arr[index])
	}

	//문자열에서 사용
	var fruits map[string]string = map[string]string{
		"apple" : "red",
		"banana" : "yellow",
		"grape" : "purple",
	}

	for fruit, color := range fruits {
		fmt.Printf("%s의 색깔은 %s입니다. \n", fruit, color)
	}
}

