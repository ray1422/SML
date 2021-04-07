package core

import (
	"math"
	"testing"
)

func Test_AddUnit(t *testing.T) {
	name := "test_unit"
	u, err := UnitDefine(name, 1, nil)
	if u == nil || err != nil {
		t.Error("init failed")
		return
	}
	if u.name != name {
		t.Error("wrong name")
	}
	_, err2 := UnitDefine(name, 1, nil)
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
		t.Error("wrong scale (cached Val")
	}
}

func Test_ScaleChain(t *testing.T) {
	px, err1 := UnitDefine("px", 1, nil)
	em, err2 := UnitDefine("em", 16, px)
	kem, err3 := UnitDefine("kem", 1000, em)
	if err1 != nil && err2 != nil && err3 != nil {
		t.Error("add unit fail")
		return
	}
	if val, _ := kem.getVal(map[*Unit]bool{}); math.Abs(val-16000) > 1e-6 {
		t.Error("Val Error")
	}
}

func Test_UnitCircle(t *testing.T) {
	a, _ := UnitDefine("a", 1, nil)
	b, _ := UnitDefine("b", 16, a)
	c, _ := UnitDefine("c", 1000, b)
	a.ref = c
	if v, _ := a.Val(); !math.IsNaN(v) {
		t.Error("error")
	}
}
