package main // Go 语言常量
import (
	"fmt"
	"unsafe"
)

func main() {
	// ----------------------------------------------------------
	/*
	   常量是一个简单值的标识符，在程序运行时，不会被修改的量。
	   常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
	   常量的定义格式：
	   const identifier [type] = value
	   你可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。
	*/
	// 显式类型定义： const b string = "abc"
	const LENGTH int = 10
	// 隐式类型定义： const b = "abc"
	const WIDTH = 5
	// 多个相同类型的声明可以简写为：
	const a1, b1, c1 = 1, false, "str" //多重赋值
	var area int
	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a1, b1, c1)
	/*
	   以上实例运行结果为：
	   面积为 : 50
	   1 false str
	*/
	// ----------------------------------------------------------
	// 常量还可以用作枚举：
	const (
		a2 = "abc"
		b2 = len(a2)
		c2 = unsafe.Sizeof(a2)
	)
	println(a2, b2, c2)
	// 常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过：
	// 以上实例运行结果为：
	// abc 3 16
	// ----------------------------------------------------------
	// iota，特殊常量，可以认为是一个可以被编译器修改的常量。
	/*
		iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
		iota 可以被用作枚举值：
		const (
		    a = iota
		    b = iota
		    c = iota
		)
		第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：
		const (
		    a = iota
		    b
		    c
		)
	*/
	const (
		a3 = iota //0
		b3        //1
		c3        //2
		d3 = "ha" //独立值，iota += 1
		e3        //"ha"   iota += 1
		f3 = 100  //iota +=1
		g3        //100  iota +=1
		h3 = iota //7,恢复计数
		i3        //8
	)
	fmt.Println(a3, b3, c3, d3, e3, f3, g3, h3, i3)
	// 以上实例运行结果为：
	// 0 1 2 ha ha 100 100 7 8
	// ----------------------------------------------------------
	// iota 表示从 0 开始自动加 1，所以 i=1<<0, j=3<<1（<< 表示左移的意思），即：i=1, j=6，这没问题，关键在 k 和 l，从输出结果看 k=3<<2，l=3<<3。
	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
	/*
		以上实例运行结果为：
		i= 1
		j= 6
		k= 12
		l= 24
		简单表述:
		i=1：1的二进制为  1 左移 0 位,不变仍为 1;
		j=3：3的二进制为 11 左移 1 位,变为二进制 110, 即 6;
		k=3：3的二进制为 11 左移 2 位,变为二进制 1100, 即 12;
		l=3：3的二进制为 11 左移 3 位,变为二进制 11000,即 24。
	*/
}
