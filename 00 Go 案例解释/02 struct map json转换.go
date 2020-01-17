package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	//struct 定义中json``解析说明 ----------------------------
	//可以选择的控制字段有三种：
	// -：不要解析这个字段
	// omitempty：当字段为空（默认值）时，不要解析这个字段。比如 false、0、nil、长度为 0 的 array，map，slice，string
	// FieldName：当解析 json 的时候，使用这个名字
	clientsList := []struct {
		ID   int    `json:"sid"`             // 解析（encode/decode） 的时候，使用 `sname`，而不是 `Field`
		Name string `json:"class,omitempty"` // 解析的时候使用 `class`，如果struct 中这个值为空，就忽略它
		Age  int    `json:"-"`               // 解析的时候忽略该字段。默认情况下会解析这个字段，因为它是大写字母开头的
	}{
		{ID: 1, Name: "one", Age: 15},
		{ID: 1, Name: "two", Age: 16},
	}
	fmt.Println(clientsList)
	// struct转为Json
	structToJson, _ := json.Marshal(clientsList)
	fmt.Println(string(structToJson))

	// map ----------------------------
	var clientsMap = map[string]interface{}{
		"data":    nil,
		"message": "",
	}

	// Format for success {"data":{"clients":[{"id":<int>,"name":<string>}]},"message":"success"}
	var clientsListInformations = []map[string]interface{}{}
	var appendInformation = map[string]interface{}{}
	for i := 0; i < len(clientsList); i++ {
		appendInformation = map[string]interface{}{
			"id":   clientsList[i].ID,
			"name": clientsList[i].Name,
		}
		clientsListInformations = append(clientsListInformations, appendInformation)
	}
	clientsMap["message"] = "success"
	clientsMap["data"] = map[string]interface{}{
		"clients": clientsListInformations,
	}
	fmt.Println(clientsMap)
	// map转json
	mapToJson, _ := json.Marshal(clientsMap)
	fmt.Println(string(mapToJson))

	// json to struct (1) ----------------------------
	type People struct {
		Name string `json:"name_title"`
		Age  int    `json:"age_size"`
	}
	jsonStr := `{"name_title": "this is name","age_size":9}`
	people := new(People)
	err := json.Unmarshal([]byte(jsonStr), &people)
	fmt.Println(err) // 0 -> mean no error
	fmt.Println(people)
	fmt.Println(people.Name)

	// json to struct (2) ----------------------------
	// 在线http://json2struct.mervine.net/
	type StuRead struct {
		Name  interface{}
		Age   interface{}
		HIgh  interface{}
		Class json.RawMessage `json:"class"` //注意这里 解析的时候使用 `class`，如果struct 中这个值为空，就忽略它
	}

	type Class struct {
		Name  string
		Grade int
	}

	data := "{\"name\":\"张三\",\"Age\":18,\"high\":true,\"sex\":\"男\",\"CLASS\":{\"naME\":\"1班\",\"GradE\":3}}"
	str := []byte(data)
	stu := StuRead{}
	json.Unmarshal(str, &stu)

	// 注意这里：二次解析！
	cla := new(Class)
	json.Unmarshal(stu.Class, cla)

	fmt.Println("stu:", stu)
	fmt.Println("string(stu.Class):", string(stu.Class))
	fmt.Println("class:", cla)

	// 利用反射，打印变量类型
	nameType := reflect.TypeOf(stu.Name)
	ageType := reflect.TypeOf(stu.Age)
	highType := reflect.TypeOf(stu.HIgh)
	classType := reflect.TypeOf(stu.Class)

	fmt.Println("nameType:", nameType)
	fmt.Println("ageType:", ageType)
	fmt.Println("highType:", highType)
	fmt.Println("classType:", classType)

	// json转map ----------------------------
	jsonStr2 := `{"name": "ccc","age": 21}`
	var mapResult map[string]interface{}
	json.Unmarshal([]byte(jsonStr2), &mapResult)
	fmt.Println(mapResult)

	// map转struct ----------------------------
	// (1) 先将map转json, json再转struct
	// (2) 第三方库 mapstructure.Decode
	// 在命令行中运行： go get github.com/goinggo/mapstructure
	mapInstance := make(map[string]interface{})
	mapInstance["Name"] = "jqw"
	mapInstance["Age"] = 18

	//var people People
	//err := mapstructure.Decode(mapInstance, &people)
	//fmt.Println(people)

	// struct转map ----------------------------
	type Student struct {
		ID   int
		Name string
		Age  int
	}
	student := Student{20, "yyy", 12}

	obj1 := reflect.TypeOf(student)
	obj2 := reflect.ValueOf(student)

	var some = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		some[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	fmt.Println(some)
}
