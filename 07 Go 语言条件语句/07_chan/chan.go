package main // Go chan 通道

import (
	"fmt"
	"time"
)

func main() {
	// 构建一个通道
	ch := make(chan int)
	// 开启一个并发匿名函数
	go func() {
		// 从3循环到0
		for i := 3; i >= 0; i-- {
			// 发送3到0之间的数值
			ch <- i
			// 每次发送完时等待
			time.Sleep(time.Second)
		}
	}()
	// 遍历接收通道数据
	for data := range ch {
		// 打印通道数据
		fmt.Println(data)
		// 当遇到数据0时, 退出接收循环
		if data == 0 {
			break
		}
	}
}

/*
执行代码，输出如下：
3
2
1
0
*/
/*
第 10 行，ch := make(chan int) 			通过 make 生成一个整型元素的通道。
第 15 行，go func() { 					将匿名函数并发执行。
第 18 行，for i := 3; i >= 0; i-- {		用循环生成 3 到 0 之间的数值。
第 21 行，ch <- i 						将 3 到 0 之间的数值依次发送到通道 ch 中。
第 24 行，time.Sleep(time.Second)		每次发送后暂停 1 秒。
第 30 行，for data := range ch {		使用 for 从通道中接收数据。
第 33 行，fmt.Println(data) 			将接收到的数据打印出来。
第 36 行，if data == 0 { 				当接收到数值 0 时，停止接收。如果继续发送，由于接收 goroutine 已经退出，没有 goroutine 发送到通道，因此运行时将会触发宕机报错。
*/

/*
通道的特性
Go语言中的通道（channel）是一种特殊的类型。在任何时候，同时只能有一个 goroutine 访问通道进行发送和获取数据。goroutine 间通过通道就可以通信。

通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。
声明通道类型
通道本身需要一个类型进行修饰，就像切片类型需要标识元素类型。通道的元素类型就是在其内部传输的数据类型，声明如下：
var 通道变量 chan 通道类型

通道类型：通道内的数据类型。
通道变量：保存通道的变量。

chan 类型的空值是 nil，声明后需要配合 make 后才能使用。
创建通道
通道是引用类型，需要使用 make 进行创建，格式如下：
通道实例 := make(chan 数据类型)

数据类型：通道内传输的元素类型。
通道实例：通过make创建的通道句柄。

请看下面的例子：
ch1 := make(chan int)                 // 创建一个整型类型的通道
ch2 := make(chan interface{})         // 创建一个空接口类型的通道, 可以存放任意格式
type Equip struct{ 一些字段 }
ch2 := make(chan *Equip)             // 创建Equip指针类型的通道, 可以存放*Equip
使用通道发送数据
通道创建后，就可以使用通道进行发送和接收操作。
1) 通道发送数据的格式
通道的发送使用特殊的操作符<-，将数据通过通道发送的格式为：
通道变量 <- 值

通道变量：通过make创建好的通道实例。
值：可以是变量、常量、表达式或者函数返回值等。值的类型必须与ch通道的元素类型一致。
2) 通过通道发送数据的例子
使用 make 创建一个通道后，就可以使用<-向通道发送数据，代码如下：
// 创建一个空接口通道
ch := make(chan interface{})
// 将0放入通道中
ch <- 0
// 将hello字符串放入通道中
ch <- "hello"
3) 发送将持续阻塞直到数据被接收
把数据往通道中发送时，如果接收方一直都没有接收，那么发送操作将持续阻塞。Go 程序运行时能智能地发现一些永远无法发送成功的语句并做出提示，代码如下：
package main
func main() {
    // 创建一个整型通道
    ch := make(chan int)
    // 尝试将0通过通道发送
    ch <- 0
}
运行代码，报错：
fatal error: all goroutines are asleep - deadlock!

报错的意思是：运行时发现所有的 goroutine（包括main）都处于等待 goroutine。也就是说所有 goroutine 中的 channel 并没有形成发送和接收对应的代码。
使用通道接收数据
通道接收同样使用<-操作符，通道接收有如下特性：
① 通道的收发操作在不同的两个 goroutine 间进行。

由于通道的数据在没有接收方处理时，数据发送方会持续阻塞，因此通道的接收必定在另外一个 goroutine 中进行。

② 接收将持续阻塞直到发送方发送数据。

如果接收方接收时，通道中没有发送方发送数据，接收方也会发生阻塞，直到发送方发送数据为止。

③ 每次接收一个元素。
通道一次只能接收一个数据元素。

通道的数据接收一共有以下 4 种写法。
1) 阻塞接收数据
阻塞模式接收数据时，将接收变量作为<-操作符的左值，格式如下：
data := <-ch

执行该语句时将会阻塞，直到接收到数据并赋值给 data 变量。
2) 非阻塞接收数据
使用非阻塞方式从通道接收数据时，语句不会发生阻塞，格式如下：
data, ok := <-ch

data：表示接收到的数据。未接收到数据时，data 为通道类型的零值。
ok：表示是否接收到数据。

非阻塞的通道接收方法可能造成高的 CPU 占用，因此使用非常少。如果需要实现接收超时检测，可以配合 select 和计时器 channel 进行，可以参见后面的内容。
3) 接收任意数据，忽略接收的数据
阻塞接收数据后，忽略从通道返回的数据，格式如下：
<-ch

执行该语句时将会发生阻塞，直到接收到数据，但接收到的数据会被忽略。这个方式实际上只是通过通道在 goroutine 间阻塞收发实现并发同步。
*/
