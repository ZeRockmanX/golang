package main // Go 函数
/*
Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。语法格式如下：

func (variable_name variable_data_type) function_name() [return_type]{
   // 函数体
}
下面定义一个结构体类型和该类型的一个方法：
*/

import (
	"fmt"
)

/* 定义结构体 */
// 就像是一个类
type Circle struct {
	radius float64
}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

/*
以上代码执行结果为：
的面积 =  314
*/
/*
Go 没有面向对象，而我们知道常见的 Java。
C++ 等语言中，实现类的方法做法都是编译器隐式的给函数加一个 this 指针，
而在 Go 里，这个 this 指针需要明确的申明出来，其实和其它 OO 语言并没有很大的区别。
*/

/*
// 貌似类的写法
package main

import "fmt"

type Cat struct {
	name string
}
func (Cat) getName1() {
	fmt.Println("get name is ")
}
func (c Cat) getName2() {
	fmt.Println(c.name)
}
func (c Cat) getName3(n string) {
	fmt.Println(n + c.name)
}
func (c *Cat) setName(n string) {
	c.name = n
}

func main() {
	var newcat Cat
	newcat.name = "mini"
	newcat.getName1()
	newcat.getName2()
	newcat.getName3("I think its name is ")

	newcat.setName("Super cat")
	newcat.getName1()
	newcat.getName2()
	newcat.getName3("I think its name is ")
}
*/
