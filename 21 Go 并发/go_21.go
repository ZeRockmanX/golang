package main //Go 并发
/*
Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。
goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。
goroutine 语法格式：
go 函数名( 参数列表 )
例如：
go f(x, y, z)
开启一个新的 goroutine:
f(x, y, z)
Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。 同一个程序中的所有 goroutine 共享同一个地址空间。
*/
import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

//我们单独写一个 say2 函数来跑 goroutine，并且 Sleep 时间设置长一点，150 毫秒
/*
package main
import (
    "fmt"
    "time"
)
func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s, (i+1)*100)
    }
}
func say2(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(150 * time.Millisecond)
        fmt.Println(s, (i+1)*150)
    }
}
func main() {
    go say2("world")
    say("hello")
}
*/
/*
问题来了，say2 只执行了 3 次，而不是设想的 5 次，为什么呢？

原来，在 goroutine 还没来得及跑完 5 次的时候，主函数已经退出了。

我们要想办法阻止主函数的结束，要等待 goroutine 执行完成之后，再退出主函数：
*/

/*
package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s, (i+1)*100)
    }
}
func say2(s string, ch chan int) {
    for i := 0; i < 5; i++ {
        time.Sleep(150 * time.Millisecond)
        fmt.Println(s, (i+1)*150)
    }
    ch <- 0
    close(ch)
}

func main() {
    ch := make(chan int)
    go say2("world", ch)
    say("hello")
    fmt.Println(<-ch)
}
// 我们引入一个信道，默认的，信道的存消息和取消息都是阻塞的，
// 在 goroutine 中执行完成后给信道一个值 0，
// 则主函数会一直等待信道中的值，一旦信道有值，主函数才会结束。
*/
