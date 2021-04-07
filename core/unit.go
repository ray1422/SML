package core

import (
	"errors"
	"math"
)

var (
	units map[string]*Unit = map[string]*Unit{}
)

type Unit struct {
	name  string
	scale float64
	ref   *Unit
	val   float64
}

func (u *Unit) String() string {
	return u.name
}
func (u *Unit) Val() (float64, error) {
	if math.IsNaN(u.val) {
		return u.getVal(map[*Unit]bool{})
	} else {
		return u.val, nil
	}
}

func (u *Unit) getVal(vis map[*Unit]bool) (float64, error) {
	if _, exists := vis[u]; exists {
		return math.NaN(), errors.New("unit" + u.String() + "already visited 😬😬😬")
	}

	if u.ref == nil {
		// abs unit
		return u.scale, nil
	}
	vis[u] = true
	ref_val, err := u.ref.getVal(vis)
	if err != nil {
		return math.NaN(), errors.New("reference tracing failed 😬😬😬")
	}
	return ref_val * u.scale, nil

}
func AddUnit(name string, scale float64, ref *Unit) (*Unit, error) {
	if _, exists := units[name]; exists {
		return nil, errors.New("unit " + name + " already defined 🙄🙄🙄")
	}
	u := &Unit{name: name, scale: scale, ref: ref, val: math.NaN()}
	units[name] = u
	return u, nil
}
