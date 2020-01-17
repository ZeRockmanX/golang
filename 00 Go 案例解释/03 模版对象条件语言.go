package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	var clients = []map[string]interface{}{
		{"id": "01", "name": "one"},
		{"id": "02", "name": "two"},
		{"id": "03", "name": "three"},
	}
	fmt.Println(clients)

	const temp = `
			{{if (.id) and (eq .name "one")}}
				mutiple conditions passed
			{{else}}
				mutiple conditions not passed
			{{end}}
			`
	t := template.Must(template.New("content").Parse(temp))
	for i, r := range clients {
		fmt.Printf("Iteration: %d:", i)
		err := t.Execute(os.Stdout, r)
		fmt.Println(err)
	}

}
