package main // Go 语言函数闭包
/*
Go 语言支持匿名函数，可作为闭包。
匿名函数是一个"内联"语句或表达式。
匿名函数的优越性在于可以直接使用函数内的变量，不必申明。

以下实例中，我们创建了函数 getSequence() ，返回另外一个函数。
该函数的目的是在闭包中递增 i 变量，代码如下：
*/

import "fmt"

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}

/*
GO语言的匿名函数就是闭包，以下是《GO语言编程》中对闭包的解释

 基本概念
闭包是可以包含自由（未绑定到特定对象）变量的代码块，这些变量不在这个代码块内或者
任何全局上下文中定义，而是在定义代码块的环境中定义。要执行的代码块（由于自由变量包含
在代码块中，所以这些自由变量以及它们引用的对象没有被释放）为自由变量提供绑定的计算环
境（作用域）。
 闭包的价值
闭包的价值在于可以作为函数对象或者匿名函数，对于类型系统而言，这意味着不仅要表示
数据还要表示代码。支持闭包的多数语言都将函数作为第一级对象，就是说这些函数可以存储到
变量中作为参数传递给其他函数，最重要的是能够被函数动态创建和返回。
*/
// 闭包带参数
/*
func main() {
    add_func := add(1,2)
    fmt.Println(add_func(1,1))
    fmt.Println(add_func(0,0))
    fmt.Println(add_func(2,2))
}
// 闭包使用方法
func add(x1, x2 int) func(x3 int,x4 int)(int,int,int)  {
    i := 0
    return func(x3 int,x4 int) (int,int,int){
       i++
       return i,x1+x2,x3+x4
    }
}
*/
