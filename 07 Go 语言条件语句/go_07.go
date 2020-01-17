package main // Go 语言条件语句
/*
条件语句需要开发者通过指定一个或多个条件，并通过测试条件是否为 true 来决定是否执行指定语句，并在条件为 false 的情况在执行另外的语句。
Go 语言条件语句
Go 语言提供了以下几种条件判断语句：
语句			 描述
if 	语句		 if 语句 由一个布尔表达式后紧跟一个或多个语句组成。
if...else		语句	if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。
if 嵌套语句		 你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。
switch 语句		 switch 语句用于基于不同条件执行不同动作。
select 语句		 select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
*/
import "fmt"

func main() {
	// ----------------------------------------------------------
	// if 嵌套语句
	/* 定义局部变量 */
	var a1 int = 100
	var b1 int = 200

	/* 判断条件*/
	if a1 == 100 {
		/* if 条件语句为 true 执行 */
		if b1 == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Println("a1 的值为 100 ， b1 的值为 200")
		}
	} else {
		/* if 条件语句为 非true 执行 */
		fmt.Println("未触发语句")
	}
	fmt.Printf("a1 值为 : %d\n", a1)
	fmt.Printf("b2 值为 : %d\n", b1)
	// ----------------------------------------------------------
	/*
	   switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。
	   switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。
	   switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 fallthrough 。
	*/
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	// 支持多条件匹配
	case 50:
	case 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
	/*
	   以上代码执行结果为：
	   优秀!
	   你的等级是 A
	*/
	// ----------------------------------------------------------
	/*
	   Type Switch
	   switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
	*/
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T\n", i)
	case int:
		fmt.Printf("x 是 int 型\n")
	case float64:
		fmt.Printf("x 是 float64 型\n")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型\n")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型\n")
	default:
		fmt.Printf("未知型\n")
	}
	/* x 的类型 :<nil> */
	// ----------------------------------------------------------
	/*
	   fallthrough
	   使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
	*/
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
	/*
	   以上代码执行结果为：
	   2、case 条件语句为 true
	   3、case 条件语句为 false
	   4、case 条件语句为 true
	*/
	// ----------------------------------------------------------
	/*
		select 语句
		select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。
		select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。

		以下描述了 select 语句的语法：
		每个 case 都必须是一个通信
		所有 channel 表达式都会被求值
		所有被发送的表达式都会被求值
		如果任意某个通信可以进行，它就执行，其他被忽略。
		如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
		否则：
		如果有 default 子句，则执行该语句。
		如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。
	*/
	// chan是golang中非常重要的一个东西，用来做goroutine的通信，因为golang程序必然会有多个goroutine，如何同步这些goroutine就很重要了。
	// 例子见 chan
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Println("received ", i1, " from c1")
	case c2 <- i2:
		fmt.Println("sent ", i2, " to c2")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Println("received ", i3, " from c3")
		} else {
			fmt.Println("c3 is closed")
		}
	default:
		fmt.Println("no communication")
	}
	/*
		以上代码执行结果为：
		no communication
	*/
}
