package main // Go 变量函数
/*
Go 函数

Go 语言可以很灵活的创建函数，并作为另外一个函数的实参。
以下实例中我们在定义的函数中初始化一个变量，该函数仅仅是为了使用内置函数 math.sqrt()，实例为：
*/

import (
	"fmt"
	"math"
)

func main() {
	/* 声明函数变量 */
	// 求平方根
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* 使用函数 */
	// 变量内置function方法
	fmt.Println(getSquareRoot(9))
	// 独立方法
	fmt.Println(getSquartRootSinglefuntion(16))
}

func getSquartRootSinglefuntion(x float64) float64 {
	return math.Sqrt(x)
}
