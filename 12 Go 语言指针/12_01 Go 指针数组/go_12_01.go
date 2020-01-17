package main // Go 语言指针数组
/*
Go 语言指针数组

*/
import "fmt"

const MAX int = 3

func main() {

	a := []int{10, 100, 200}
	var i int

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}

	// 有一种情况，我们可能需要保存数组，这样我们就需要使用到指针。
	b := []int{10, 100, 200}
	var j int
	var ptr [MAX]*int

	for j = 0; j < MAX; j++ {
		ptr[j] = &b[j] /* 整数地址赋值给指针数组 */
	}

	for j = 0; j < MAX; j++ {
		fmt.Printf("b[%d] = %d\n", j, *ptr[j])
	}
}
