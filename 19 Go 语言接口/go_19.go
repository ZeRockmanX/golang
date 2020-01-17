package main // Go 语言接口
/*
Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，
任何其他类型只要实现了这些方法就是实现了这个接口。
*/
import (
	"fmt"
)

/* 定义接口 */
type Phone interface {
	//  method_name1 [return_type]
	call()
}

/* 定义结构体 */
type NokiaPhone struct {
	one string
}

/* 实现接口方法 */
func (nokiaPhone NokiaPhone) call() {
	/* 方法实现 */
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone
	// 接口调用
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	//结构体调用
	var test NokiaPhone
	test.one = "testuse"
	fmt.Println(test.one)
	fmt.Println(test)

}

/*
在上面的例子中，我们定义了一个接口Phone，接口里面有一个方法call()。
然后我们在main函数里面定义了一个Phone类型变量，并分别为之赋值为NokiaPhone和IPhone。
然后调用call()方法，输出结果如下：
I am Nokia, I can call you!
I am iPhone, I can call you!
*/

/*
给接口增加参数：

package main

import (
    "fmt"
)
type Man interface {
    name() string;
    age() int;
}

type Woman struct {
}

func (woman Woman) name() string {
   return "Jin Yawei"
}
func (woman Woman) age() int {
   return 23;
}

type Men struct {
}

func ( men Men) name() string {
   return "liweibin";
}
func ( men Men) age() int {
    return 27;
}

func main(){
    var man Man;

    man = new(Woman);
    fmt.Println( man.name());
    fmt.Println( man.age());
    man = new(Men);
    fmt.Println( man.name());
    fmt.Println( man.age());
}
// -------------------------------------------------
func (name string) imp() string{
    print("这是实现方法的写法")
}

func sum(x int,y int) int{
    print("这是正常写法")
}
// -------------------------------------------------
package main

import "fmt"

type Phone interface {
    call(param int) string
    takephoto()
}

type Huawei struct {
}

func (huawei Huawei) call(param int) string{
    fmt.Println("i am Huawei, i can call you!", param)
    return "damon"
}

func (huawei Huawei) takephoto() {
    fmt.Println("i can take a photo for you")
}

func main(){
    var phone Phone
    phone = new(Huawei)
    phone.takephoto()
    r := phone.call(50)
    fmt.Println(r)
}
// -------------------------------------------------
接口案例:

package main
import (
    "fmt"
)


//定义接口
type Phone interface {
    call()
    call2()
}


//一直都搞不懂这是干啥的
//原来是用来定义结构体内的数据类型的

type Phone1 struct {
    id            int
    name          string
    category_id   int
    category_name string
}

//第一个类的第一个回调函数
func (test Phone1) call() {
    fmt.Println("这是第一个类的第一个接口回调函数 结构体数据：", Phone1{id: 1, name: "浅笑"})
}

//第一个类的第二个回调函数
func (test Phone1) call2() {
    fmt.Println("这是一个类的第二个接口回调函数call2", Phone1{id: 1, name: "浅笑", category_id: 4, category_name: "分类名称"})
}



//第二个结构体的数据类型
type Phone2 struct {
    member_id       int
    member_balance  float32
    member_sex      bool
    member_nickname string
}

//第二个类的第一个回调函数
func (test2 Phone2) call() {
    fmt.Println("这是第二个类的第一个接口回调函数call", Phone2{member_id: 22, member_balance: 15.23, member_sex: false, member_nickname: "浅笑18"})
}


//第二个类的第二个回调函数
func (test2 Phone2) call2() {
    fmt.Println("这是第二个类的第二个接口回调函数call2", Phone2{member_id: 44, member_balance: 100, member_sex: true, member_nickname: "陈超"})
}

//开始运行
func main() {
    var phone Phone

    //先实例化第一个接口
    phone = new(Phone1)
    phone.call()
    phone.call2()

    //实例化第二个接口
    phone = new(Phone2)
    phone.call()
    phone.call2()
}
// -------------------------------------------------
将接口做为参数

package main

import (
    "fmt"
)

type Phone interface {
    call() string
}

type Android struct {
    brand string
}

type IPhone struct {
    version string
}

func (android Android) call() string {
    return "I am Android " + android.brand
}

func (iPhone IPhone) call() string {
    return "I am iPhone " + iPhone.version
}

func printCall(p Phone) {
    fmt.Println(p.call() + ", I can call you!")
}

func main() {
    var vivo = Android{brand:"Vivo"}
    var hw = Android{"HuaWei"}

    i7 := IPhone{"7 Plus"}
    ix := IPhone{"X"}

    printCall(vivo)
    printCall(hw)
    printCall(i7)
    printCall(ix)
}
输出结果：

I am Android Vivo, I can call you!
I am Android HuaWei, I can call you!
I am iPhone 7 Plus, I can call you!
I am iPhone X, I can call you!
// -------------------------------------------------
如果想要通过接口方法修改属性，需要在传入指针的结构体才行，具体代码入下的 [1][2] 处：

type fruit interface{
    getName() string
    setName(name string)
}
type apple struct{
    name string
}
//[1]
func (a *apple) getName() string{
    return a.name
}
//[2]
func (a *apple) setName(name string) {
   a.name = name
}
func testInterface(){
    a:=apple{"红富士"}
    fmt.Print(a.getName())
    a.setName("树顶红")
    fmt.Print(a.getName())
}
*/
