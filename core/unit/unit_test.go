package unit

import (
	"fmt"
	"math"
	"testing"
)

func Test_RelValChain(t *testing.T) {
	defer ClearUnits()
	p, _ := AddUnit("p", 1, nil, REL)
	q, _ := AddUnit("q", 1, p, REL)
	r, _ := AddUnit("r", 1, q, REL)
	s, _ := AddUnit("s", 1, q, REL)
	s2, _ := AddUnit("s2", 1, nil, REL)
	p.ref = r
	_, err := r.RelVal(s, 2)
	if err == nil {
		t.Error("chain detect failed!")
	}
	_, err = s2.RelVal(s, 2)
	if err == nil {
		t.Error("parent detect failed!")
	}
	_, err = s2.RelVal(nil, 2)
	if err == nil {
		t.Error("nil parent detect failed!")
	}
	ClearUnits()
	p, _ = AddUnit("p", 1, nil, REL)
	q, _ = AddUnit("q", 1, p, REL)
	r, _ = AddUnit("r", 1, q, REL)
	_, err = r.RelVal(q, 1)
	if err == nil {
		t.Error("non-root should be unable to set reference value")
	}
}
func Test_RelVal(t *testing.T) {
	defer ClearUnits()
	p, _ := AddUnit("p", 1, nil, REL)
	q, _ := AddUnit("q", 1, p, REL)
	r, _ := AddUnit("r", 1, q, REL)
	v, _ := r.RelVal(p, 10)
	if v != 10 {
		t.Error(v)
	}
	fmt.Println(v)
}
func Test_DeleteUnit(t *testing.T) {
	defer ClearUnits()
	AddUnit("test", 1, nil, ABS)
	if RemoveUnit("test"); GetUnit("test") != nil {
		t.Error("delete")
	}
}
func Test_AddAbsUnit(t *testing.T) {
	defer ClearUnits()
	name := "test_unit"
	u, err := AddUnit(name, 1, nil, ABS)
	if err != nil {
		t.Fatal("unit " + u.String() + " init failed")
	}
	if GetUnit(name) == nil {
		t.Error("func UnitGet failed")
	}
	if u.Name() != name {
		t.Error("name failed")
	}
	if u == nil || err != nil {
		t.Error("init failed")
		return
	}
	if u.name != name {
		t.Error("wrong name")
	}
	_, err2 := AddUnit(name, 1, nil, ABS)
	if err2 == nil {
		t.Error("unit with the same name can be re-defined")
	}
	if val, _ := u.getVal(map[*Unit]bool{}); math.Abs(val-1) > 1e-6 {
		t.Error("wrong scale (getVal)")
	}
	if val, _ := u.Val(); math.Abs(val-1) > 1e-6 {
		t.Error("wrong scale (Val)")
	}
	if val, err := u.Val(); err == nil {
		u.val = val
	} else {
		t.Error(err)
		return
	}
	if val, _ := u.Val(); math.Abs(val-1) > 1e-6 {
		t.Error("wrong scale (cached Val)")
	}

	ClearUnits()
	r, _ := AddUnit("r", 1, nil, REL)
	a, err := AddUnit("a", 1., r, ABS)
	if a != nil || err == nil {
		t.Error("abs shouldn't be able to ref a rel unit")
	}
	ClearUnits()
}

func Test_ScaleChain(t *testing.T) {
	defer ClearUnits()
	px, err1 := AddUnit("px", 1, nil, ABS)
	em, err2 := AddUnit("em", 16, px, ABS)
	kem, err3 := AddUnit("kem", 1000, em, ABS)
	if err1 != nil && err2 != nil && err3 != nil {
		t.Error("add unit fail")
		return
	}
	if val, _ := kem.getVal(map[*Unit]bool{}); math.Abs(val-16000) > 1e-6 {
		t.Error("Val Error")
	}
}

func Test_UnitCircle(t *testing.T) {
	defer ClearUnits()
	a, _ := AddUnit("a", 1, nil, ABS)
	b, _ := AddUnit("b", 16, a, ABS)
	c, _ := AddUnit("c", 1000, b, ABS)
	a.ref = c
	if v, _ := a.Val(); !math.IsNaN(v) {
		t.Error("error")
	}
}
