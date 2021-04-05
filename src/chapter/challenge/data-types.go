package main

import (
	"fmt"
)

// 编写一个程序来计算斐波纳契数列
func fbnq()  {
	var count int = 7
	numbers := []int{1, 1}
	for i := 2; i < count; i++ {
		val := numbers[i-1] + numbers[i-2]
        numbers = append(numbers, val)
    }
	fmt.Printf("%v", numbers)
}

func fibonacci(n int) []int {
	if n < 2 {
		return make([]int, 0)
	}

	nums := make([]int, n)
	nums[0], nums[1] = 1, 1
	for i :=2; i < n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}
	return nums
}

// 创建罗马数字转换器
func romanToArabic(numeral string) int {
	romanMap := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	arabicVals := make([]int, len(numeral)+1)

    for index, digit := range numeral {
        if val, present := romanMap[digit]; present {
            arabicVals[index] = val
        } else {
            fmt.Printf("Error: The roman numeral %s has a bad digit: %c\n", numeral, digit)
            return 0
        }
	}
	fmt.Print(arabicVals)

    total := 0

    for index := 0; index < len(numeral); index++ {
        if arabicVals[index] < arabicVals[index+1] {
            arabicVals[index] = -arabicVals[index]
        }
        total += arabicVals[index]
    }

    return total
}

func main() {
	fibonacci(6)

	//fmt.Println("MCLX is: ", romanToArabic("MCLX"))
    // fmt.Println("MCMXCIX is: ", romanToArabic("MCMXCIX"))
	fmt.Println("MCM is: ", romanToArabic("MCM"))
	// fmt.Println("MCMZ is: ", romanToArabic("MCMZ"))
}