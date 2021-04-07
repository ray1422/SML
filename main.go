package main

import (
	"fmt"

	"github.com/ray1422/SML/core"
)

type Test struct {
}

func main() {
	u1, _ := core.AddUnit("px", 1, nil)
	u2, _ := core.AddUnit("em", 16, u1)
	u3, err := core.AddUnit("em", 1./16., u2)
	if err != nil {
		panic(err)
	}
	u4, err := core.AddUnit("em", 1, u3)
	u5, err := core.AddUnit("em", 16, u4)
	u6, err := core.AddUnit("em", 16, u5)
	fmt.Println(u6.Val())

}
