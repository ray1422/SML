package container

import (
	"fmt"
	"reflect"
	"strings"
)

type directionT uint16
type inlineBlockT uint16

const (
	// VERTICAL id
	VERTICAL directionT = iota
	// HORIZONTAL id
	HORIZONTAL directionT = iota
)

// RenderGeneralAttr general attr
type RenderGeneralAttr struct {
	Height  int
	Width   int
	Margin  [4]float64 // top, right, borrom, left
	Padding [4]float64 // top, right, borrom, left
}

func (r *RenderGeneralAttr) RenderAttr() *RenderGeneralAttr {
	return r
}

// VarBlock interface
type VarBlock interface {
	Block
	Append(u Block)
}

// Block interface
type Block interface {
	Children() []Block
	RenderAttr() *RenderGeneralAttr
}

// BaseBlock basic container of content
type BaseBlock struct {
	RenderGeneralAttr
	children []Block
}
type CodeBlock struct {
	RenderGeneralAttr
	Inline  bool
	Lang    string
	Content string
}

func (blk *CodeBlock) Children() []Block {
	return []Block{blk}
}

// Children get children
func (blk *BaseBlock) Children() []Block {
	return []Block(blk.children)
}

// Append append a child
func (blk *BaseBlock) Append(u Block) {
	if u == nil {
		return
	}
	blk.children = append(blk.children, u)
}

// Clear clear all children
func (blk *BaseBlock) Clear() {
	blk.children = []Block{}
}

func Dump(blk Block, indent int) {
	if reflect.ValueOf(blk).Kind() == reflect.Ptr && reflect.ValueOf(blk).IsNil() {
		fmt.Println("<nil>")
		return
	}
	indentStr := ""
	indentStrSingle := "    "
	for i := 0; i < indent; i++ {
		indentStr += indentStrSingle
	}
	// indentStr = indentStr[:indent*2]
	switch u := blk.(type) {
	case *TextBlock:
		if u != nil {
			fmt.Println(indentStr + "text: '" + esc(u.Content) + "'")
		} else {
			fmt.Println(indentStr + "text: " + "<nil>")
		}
	case *CodeBlock:
		fmt.Println(indentStr + indentStrSingle + "Lang: " + u.Lang)
		fmt.Println(indentStr + indentStrSingle + "Content:\n" + u.Content)
	default:
		fmt.Println(indentStr + getType(blk) + ":")
		switch v := blk.(type) {
		case *InlineBlock:
			fmt.Println(indentStr+indentStrSingle+"InlineBlockType:", v.InlineBlockType)
		case *ListBlock:
			fmt.Println(indentStr+indentStrSingle+"Ordered:", v.Ordered)
		case *ImageBlock:
			fmt.Println(indentStr+indentStrSingle+"Alt:", v.Alt)
			fmt.Println(indentStr+indentStrSingle+"Src:", v.Src)
			fmt.Println(indentStr + indentStrSingle + "Attr:")
			v.Attr.Dump(indent + 1)
		}
		for _, v := range blk.Children() {
			Dump(v, indent+1)

		}
		// fmt.Println(indentStr + "}")
	}

}

func getType(v interface{}) string {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func (t inlineBlockT) String() string {
	return inlineBlockT2Str[t]
}

func esc(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\n", "\\n")
	return s
}
