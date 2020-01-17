package main // Go 语言循环语句
/*
Go 语言提供了以下几种类型循环处理语句：

循环类型	描述
for 循环	重复执行语句块
循环嵌套	在 for 循环中嵌套一个或多个 for 循环

循环控制语句
循环控制语句可以控制循环体内语句的执行过程。
GO 语言支持以下几种循环控制语句：

控制语句		描述
break 语句		经常用于中断当前 for 循环或跳出 switch 语句
continue 语句	跳过当前循环的剩余语句，然后继续进行下一轮循环。
goto 语句		将控制转移到被标记的语句。
*/
import "fmt"

func main() {
	// ！！！ 在条件表达式前面声明的的变量只能在条件语句块中使用
	/*
		if a, b := 21, 31; a > b {
		    fmt.Println("a>b ? true")
		}else {
		    fmt.Println(a,b) //Ok
		}
		fmt.Println(a,b)    //error: undefined a ,undefined b
	*/
	// ----------------------------------------------------------
	/*
		init： 一般为赋值表达式，给控制变量赋初值；
		condition： 关系表达式或逻辑表达式，循环控制条件；
		post： 一般为赋值表达式，给控制变量增量或减量。
		for语句执行过程如下：
		①先对表达式1赋初值；
		②判别赋值表达式 init 是否满足给定条件，若其值为真，满足循环条件，则执行循环体内语句，
		然后执行 post，进入第二次循环，再判别 condition；否则判断 condition 的值为假，
		不满足条件，就终止for循环，执行循环体外语句。
	*/
	// for 循环
	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	/* 和 C 语言的 for 一样： for init; condition; post { }*/
	// ！！！ a := 0 a重新被定义，则循环体内的a为for定义的a，而不是var定义的a
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	/* 和 C 的 while 一样： for condition { }*/
	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	/* for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环*/
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
	// ----------------------------------------------------------
	// Go 语言循环嵌套
	/* 定义局部变量 */
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
	// ----------------------------------------------------------
	// Go 语言 break 语句
	/*
		Go 语言中 break 语句用于以下两方面：

		用于循环语句中跳出循环，并开始执行循环之后的语句。
		break 在 switch（开关语句）中在执行一条case后跳出语句的作用。
	*/
	/* 定义局部变量 */
	var a3 int = 10

	/* for 循环 */
	for a3 < 20 {
		fmt.Printf("a3 的值为 : %d\n", a3)
		a3++
		if a3 > 15 {
			/* 使用 break 语句跳出循环 */
			break
		}
	}
	// ----------------------------------------------------------
	// Go 语言 continue 语句
	/*
		Go 语言的 continue 语句 有点像 break 语句。但是 continue 不是跳出循环，而是跳过当前循环执行下一次循环语句。
		for 循环中，执行 continue 语句会触发for增量语句的执行
	*/
	/* 定义局部变量 */
	var a4 int = 10

	/* for 循环 */
	for a4 < 20 {
		if a4 == 15 {
			/* 跳过此次循环 */
			a4 = a4 + 1
			continue
		}
		fmt.Printf("a4 的值为 : %d\n", a4)
		a4++
	}
	// ----------------------------------------------------------
	// Go 语言 goto 语句
	/*
		Go 语言的 goto 语句可以无条件地转移到过程中指定的行。

		goto 语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。
		但是，在结构化程序设计中一般不主张使用 goto 语句， 以免造成程序流程的混乱，使理解和调试程序都产生困难。
	*/
	/* 定义局部变量 */
	var a5 int = 10

	/* 循环 */
LOOP1:
	for a5 < 20 {
		if a5 == 15 {
			/* 跳过迭代 */
			a5 = a5 + 1
			goto LOOP1
		}
		fmt.Printf("a5 的值为 : %d\n", a5)
		a5++
	}
	//for循环配合goto打印九九乘法表
	for m := 1; m < 10; m++ {
		n := 1
	LOOP2:
		if n <= m {
			fmt.Printf("%dx%d=%d ", n, m, m*n)
			n++
			goto LOOP2
		} else {
			fmt.Println("")
		}
		n++
	}
}
