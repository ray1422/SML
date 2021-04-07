package unit

import (
	"math"
	"testing"
)

func TestDeleteUnit(t *testing.T) {
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
}

func Test_AbsBehavior(t *testing.T) {
	rel, _ := AddUnit("rel", 1, nil, REL)
	_, err4 := AddUnit("abs", 1, rel, ABS)
	if err4 == nil {
		t.Error("an absolute unit shouldn't reference an non-absolute unit")
	}
	if u, _ := AddUnit("abs2", 1, nil, ABS); u.UpdateScale(2) == nil {
		t.Error("scale of an absolute unit should be constant")
	}
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

func Test_RelativeScaleChain(t *testing.T) {
	px, err1 := AddUnit("px", 1, nil, REL)
	em, err2 := AddUnit("em", 16, px, REL)
	kem, err3 := AddUnit("kem", 1000, em, REL)
	if err1 != nil && err2 != nil && err3 != nil {
		t.Error("add unit fail")
		return
	}
	if val, _ := kem.getVal(map[*Unit]bool{}); math.Abs(val-16000) > 1e-6 {
		t.Error("Val Error")
	}
	if val, _ := kem.Val(); math.Abs(val-16000.) > 1e-4 {
		t.Error("wrong scale")
	}
	px.UpdateScale(2)
	if val, _ := kem.Val(); math.Abs(val-16000.*2) > 1e-4 {
		t.Error("wrong scale (after update scale)")
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
