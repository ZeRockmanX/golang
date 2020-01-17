package main // Go 语言结构体
/*
Go 语言结构体
Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。
结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
结构体表示一项记录，比如保存图书馆的书籍记录，每本书有以下属性：
Title ：标题
Author ： 作者
Subject：学科
ID：书籍ID
定义结构体
结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体有中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}
一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：
variable_name := structure_variable_type {value1, value2...valuen}
或
variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
*/
import "fmt"

// 定义结构体
type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	// -----------------------------------------------------------------
	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})
	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})
	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})

	// -----------------------------------------------------------------
	// 访问结构体成员
	// 如果要访问结构体成员，需要使用点号 . 操作符
	var Book1 Books        /* 声明 Book1 为 Books 类型 */
	var Book2 Books        /* 声明 Book2 为 Books 类型 */

   	/* book 1 描述 */
   	Book1.title = "Go 语言"
   	Book1.author = "www.runoob.com"
  	Book1.subject = "Go 语言教程"
  	Book1.book_id = 6495407

   	/* book 2 描述 */
   	Book2.title = "Python 教程"
   	Book2.author = "www.runoob.com"
   	Book2.subject = "Python 语言教程"
   	Book2.book_id = 6495700

   	/* 打印 Book1 信息 */
   	fmt.Printf( "Book 1 title : %s\n", Book1.title)
   	fmt.Printf( "Book 1 author : %s\n", Book1.author)
   	fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
   	fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

   	/* 打印 Book2 信息 */
   	fmt.Printf( "Book 2 title : %s\n", Book2.title)
   	fmt.Printf( "Book 2 author : %s\n", Book2.author)
   	fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
	fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
	   
	// -----------------------------------------------------------------
	// 结构体作为函数参数
	// 你可以像其他数据类型一样将结构体类型作为参数传递给函数
	/* 打印 Book1 信息 */
	printBook(Book1)

	/* 打印 Book2 信息 */
	printBook(Book2)

	// -----------------------------------------------------------------
	// 结构体指针
	// 你可以像其他数据类型一样将结构体类型作为参数传递给函数
	/*
	你可以定义指向结构体的指针类似于其他指针变量，格式如下：
	var struct_pointer *Books
	以上定义的指针变量可以存储结构体变量的地址。查看结构体变量地址，可以将 & 符号放置于结构体变量前：
	struct_pointer = &Book1
	使用结构体指针访问结构体成员，使用 "." 操作符：
	struct_pointer.title
	*/
	/* 打印 Book1 信息 */
	ptrPrintBook(&Book1)
	/* 打印 Book2 信息 */
	ptrPrintBook(&Book2)
}
// 结构体作为函数参数
func printBook( book Books ) {
	fmt.Printf( "Book title : %s\n", book.title)
	fmt.Printf( "Book author : %s\n", book.author)
	fmt.Printf( "Book subject : %s\n", book.subject)
	fmt.Printf( "Book book_id : %d\n", book.book_id)
 }
// 结构体指针
 func ptrPrintBook( book *Books ) {
	fmt.Printf( "Book title : %s\n", book.title)
	fmt.Printf( "Book author : %s\n", book.author)
	fmt.Printf( "Book subject : %s\n", book.subject)
	fmt.Printf( "Book book_id : %d\n", book.book_id)
 }
 // 结构体是作为参数的值传递
 // 如果想在函数里面改变结果体数据内容，需要传入指针

/*
// struct 类似于 java 中的类，可以在 struct 中定义成员变量。
// 要访问成员变量，可以有两种方式：
// 1.通过 struct 变量.成员 变量来访问。
// 2.通过s truct 指针.成员 变量来访问。
// 不需要通过 getter, setter 来设置访问权限
type Rect struct{   		//定义矩形类
    x,y float64       		//类型只包含属性，并没有方法
    width,height float64
}
func (r *Rect) Area() float64{    //为Rect类型绑定Area的方法，*Rect为指针引用可以修改传入参数的值
    return r.width*r.height       //方法归属于类型，不归属于具体的对象，声明该类型的对象即可调用该类型的方法
}
*/