package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	//------------------------------------------------------------------------------------------
	// 单个定义 type xxx struct{}
	//---------------------------------------------
	type T1 struct {
		C3 string
		C4 string
	}
	//--实例化-------------------------------------------
	var R0 T1
	R0.C3 = "Real0 C3"
	R0.C4 = "Real0 C4"
	fmt.Println(R0)
	R0 = T1{C3: "str1", C4: "str2"}
	fmt.Println(R0)
	//--嵌套-------------------------------------------
	zeroEX := struct {
		C1 bool
		C2 T1
	}{
		C1: true,
		C2: T1{C3: "T0 C3"},
	}
	fmt.Println(zeroEX)
	//--匿名嵌套匿名写法----------------------------------------
	oneEX := struct {
		C1 bool
		C2 struct {
			C3 bool
			C4 string
		}
	}{
		C1: true,
		C2: struct {
			C3 bool
			C4 string
		}{C3: true},
	}
	fmt.Println(oneEX)
	//------------------------------------------------------------------------------------------
	// 数组定义 type xxx []struct{}
	//---------------------------------------------
	type T2 []struct {
		C3 string
		C4 string
	}
	//--实例化-------------------------------------------
	var R2 T2
	R2 = T2{{C3: "str1", C4: "str2"}}
	fmt.Println(R2)
	//--嵌套-------------------------------------------
	towEX := []struct {
		C1 bool
		C2 T2
	}{
		{C1: true},
		{C2: T2{{C3: "T0 C3"}}},
	}
	fmt.Println(towEX)
	//---匿名数组嵌套匿名数组------------------------------------------
	conditions := []struct {
		C1 bool
		C2 []struct {
			C3 bool
			C4 string
		}
	}{
		{C1: true},
		{C2: []struct {
			C3 bool
			C4 string
		}{{C3: true}}},
	}
	fmt.Println(conditions)
	fmt.Println(conditions[1].C2[0].C3)
	//------------------------------------------------------------------------------------------
	// Define a template.
	const temp = `
{{if .C1}}
        mutiple conditions passed
{{else}}
        mutiple conditions not passed
{{end}}
`
	t := template.Must(template.New("content").Parse(temp))
	for i, r := range conditions {
		fmt.Printf("Iteration: %d:", i)
		err := t.Execute(os.Stdout, r)
		fmt.Println(err)
	}
}
