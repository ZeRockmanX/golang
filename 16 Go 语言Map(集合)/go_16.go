package main // Go 语言Map(集合)
/*
Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。
定义 Map
可以使用内建函数 make 也可以使用 map 关键字来定义 Map:
// 声明变量，默认 map 是 nil
var map_variable map[key_data_type]value_data_type
// 使用 make 函数
map_variable := make(map[key_data_type]value_data_type)
如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
*/
import "fmt"

func main() {
	// -----------------------------------------------------------------
	// 定义 Map
	var countryCapitalMap1 map[string]string     // 声明变量，默认 map 是 nil
	countryCapitalMap1 = make(map[string]string) // 使用 make 函数 初始化

	//或者直接创建直接countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap1["France"] = "巴黎"
	countryCapitalMap1["Italy"] = "罗马"
	countryCapitalMap1["Japan"] = "东京"
	countryCapitalMap1["India"] = "新德里"
	countryCapitalMap1["American"] = "华盛顿"

	//或者直接创建不用定义
	//countryCapitalMap1 := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	/*使用键输出地图值 */
	for country := range countryCapitalMap1 {
		fmt.Println(country, "首都是", countryCapitalMap1[country])
	}

	/*查看元素在集合中是否存在 */
	// 给key , value赋值 countryCapitalMap[键]的值
	capital, ok := countryCapitalMap1["American"] /*如果确定是真实的,则存在,否则不存在 */
	fmt.Println(capital)                          // key
	fmt.Println(ok)                               // vaule
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	// -----------------------------------------------------------------
	// delete() 函数
	/* 创建map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	/* 打印地图 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*删除元素*/
	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
}
