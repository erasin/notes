// package html/template

package template

type Template struct{}

func Must(t *Template, err error) *Template

func (t *Template) Execute(wr io.Writer, data interface{}) error
