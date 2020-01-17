package main // Go 语言指针作为函数参数
/*
Go 语言允许向函数传递指针，只需要在函数定义的参数上设置为指针类型即可。
以下实例演示了如何向函数传递指针，并在函数调用后修改函数内的值，：
*/
import "fmt"
func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
 
	fmt.Printf("交换前 a 的值 : %d\n", a )
	fmt.Printf("交换前 b 的值 : %d\n", b )
 
	/* 调用函数用于交换值
	* &a 指向 a 变量的地址
	* &b 指向 b 变量的地址
	*/
	swap(&a, &b);
 
	fmt.Printf("交换后 a 的值 : %d\n", a )
	fmt.Printf("交换后 b 的值 : %d\n", b )
 }
 
 func swap(x *int, y *int) {
	var temp int
	temp = *x    /* 保存 x 地址的值 */
	*x = *y      /* 将 y 赋值给 x */
	*y = temp    /* 将 temp 赋值给 y */
 }