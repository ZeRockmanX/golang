package main // Go 语言变量作用域
/*
作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围。
Go 语言中变量可以在三个地方声明：

函数内定义的变量称为局部变量
函数外定义的变量称为全局变量
函数定义中的变量称为形式参数
接下来让我们具体了解局部变量、全局变量和形式参数。
*/
import (
	"fmt"
)

/* 声明全局变量 */
// 全局变量可以在整个包甚至外部包（被导出后）使用。
var a int = 20

func main() {
	/* main 函数中声明局部变量 */
	// Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。实例如下：
	var a int = 10
	var b int = 20
	var c int = 0

	fmt.Printf("main()函数中 a = %d\n", a)
	c = sum(a, b)
	fmt.Printf("main()函数中 a = %d\n", a)
	fmt.Printf("main()函数中 c = %d\n", c)
}

/* 函数定义-两数相加 a b为形式参数*/
func sum(a, b int) int {
	// 只对本函数内的形式参数a有影响
	a = a + 1
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)
	return a + b
}

/*
初始化局部和全局变量
不同类型的局部和全局变量默认值为：

数据类型	初始化默认值
int			0
float32		0
pointer		nil
*/

/*
// 语句中的也是局部变量，并且可以使用相同的名字
// 例如在 for 循环的 initialize（a:=0） 中，此时 initialize 中的 a 与外层的 a 不是同一个变量，
// initialize 中的 a 为 for 循环中的局部变量，因此在执行完 for 循环后，输出 a 的值仍然为 0。
func main(){
  var a int = 0
  fmt.Println("for start")
  for a:=0; a < 10; a++ {
    fmt.Println(a)
  }
  fmt.Println("for end")

  fmt.Println(a)
}
*/
