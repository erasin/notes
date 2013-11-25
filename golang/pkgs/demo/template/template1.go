package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name    string
	Age     int
	Emails  []string
	Friends []*Friend
}

type Friend struct {
	Name  string
	Email string
}

func main() {
	f1 := Friend{Name: "f1.xxx"}
	f2 := Friend{Name: "f2.yyy"}

	p := Person{
		Name:    "w.er",
		Age:     20,
		Emails:  []string{"xxx@xxx.com", "yyy@yyy.com"},
		Friends: []*Friend{&f1, &f2}}

	temp1 := `你好，已经{{.Age}}的{{.Name}}!
{{range .Emails}}
Email: {{.}}
{{end}}
{{with .Friends}}
	{{range .}}
你朋友： {{.Name}}
	{{end}}
{{end}}
`

	// 创建一个 template
	t := template.New("Person Info")
	// 解析模板
	t, err := t.Parse(temp1)
	checkError(err)
	// 输出到 io.Writer 接口
	err = t.Execute(os.Stdout, p)
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
