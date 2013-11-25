package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	// 必须加载
	s1, _ := template.ParseFiles("header.tpl", "content.tpl", "footer.tpl")
	// s1.ExecuteTemplate(os.Stdout, "header", nil)
	// fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", nil)

	fmt.Println()
	// s1.ExecuteTemplate(os.Stdout, "footer", nil)
	// fmt.Println()
	s1.Execute(os.Stdout, nil)
}
