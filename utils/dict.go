package utils

import "fmt"

type DictItem struct {
	K interface{}
	V interface{}
}

// Dict python like Dictionary
type Dict map[interface{}]interface{}

func (d *Dict) Keys() chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for k := range *d {
			ch <- k
		}

	}()
	return ch
}
func (d *Dict) Values() chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for _, v := range *d {
			ch <- v
		}
	}()
	return ch
}

func (d *Dict) Items() chan (DictItem) {
	ch := make(chan DictItem)
	go func() {
		defer close(ch)
		for k, v := range *d {
			ch <- DictItem{K: k, V: v}
		}
	}()
	return ch
}
func (d *Dict) Dump(indent int) {
	perIndStr := "    "
	indentStr := ""
	for i := 0; i < indent; i++ {
		indentStr += perIndStr
	}
	if d == nil || *d == nil {
		fmt.Println(indentStr + "<nil>")
		return
	}
	fmt.Println(indentStr + "{")

	for k, v := range *d {
		if v == nil {
			fmt.Println(indentStr+perIndStr+k.(string)+":", "<nil>")
		} else if u, ok := v.(*Dict); ok {
			u.Dump(indent + 1)
		} else if u, ok := v.(string); ok {
			fmt.Println(indentStr+perIndStr+k.(string)+":", "'"+u+"'")
		} else {
			fmt.Println(indentStr+perIndStr+k.(string)+":", v)
		}
	}
	fmt.Println(indentStr + "}")
}
