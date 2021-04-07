package unit

import (
	"errors"
	"math"
)

type UNIT_TYPE int

const (
	ABS UNIT_TYPE = iota
	REL UNIT_TYPE = iota
)

var (
	units        map[string]*Unit = make(map[string]*Unit)
	UnitType2Str                  = map[UNIT_TYPE]string{ABS: "absolute", REL: "relative"}
)

type Unit struct {
	name     string
	scale    float64
	ref      *Unit
	val      float64
	unitType UNIT_TYPE
}

func (u *Unit) String() string {
	return u.name + "(" + UnitType2Str[u.unitType] + " unit)"
}
func (u *Unit) Name() string {
	return u.name
}
func (u *Unit) Val() (float64, error) {

	if u.unitType == REL || math.IsNaN(u.val) { // rel unit can't be cache
		return u.getVal(map[*Unit]bool{})
	} else {
		return u.val, nil
	}
}
func (u *Unit) UpdateScale(scale float64) error {
	if u.unitType == ABS {
		return errors.New("an absolute unit can't update it's value")
	}
	u.scale = scale
	return nil

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

func AddUnit(name string, scale float64, ref *Unit, unitType UNIT_TYPE) (*Unit, error) {
	if unitType == ABS && ref != nil {
		if ref.unitType != ABS {
			return nil, errors.New("an absolute unit can only reference an absolute unit")
		}
	}
	if _, exists := units[name]; exists {
		return nil, errors.New("unit " + name + " has already been defined 🙄🙄🙄")
	}
	u := &Unit{name: name, scale: scale, ref: ref, val: math.NaN(), unitType: unitType}

	units[name] = u
	return u, nil
}

func RemoveUnit(name string) {
	delete(units, name)
}
func GetUnit(name string) *Unit {
	if v, ok := units[name]; ok {
		return v
	}
	return nil
}
func ClearUnits() {
	units = make(map[string]*Unit)
}
