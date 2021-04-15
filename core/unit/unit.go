package unit

import (
	"errors"
	"math"
)

type unitTypeID int

const (
	// ABS Type
	ABS unitTypeID = iota
	// REL Type
	REL unitTypeID = iota
)

var (
	units        map[string]*Unit = make(map[string]*Unit)
	unitType2Str                  = map[unitTypeID]string{ABS: "absolute", REL: "relative"}
)

// Unit struct
type Unit struct {
	name     string
	scale    float64
	ref      *Unit
	val      float64
	unitType unitTypeID
}

// // prevent unit being updated while getting value.
// func Lock() {
// 	updateLock.Lock()
// }

// // unlock the updating lock
// func Unlock() {
// 	updateLock.Unlock()
// }

// human readable unit name with type
func (u *Unit) String() string {
	return u.name + "(" + unitType2Str[u.unitType] + " unit)"
}

// Name of unit
func (u *Unit) Name() string {
	return u.name
}

// Val calculated value of a unit
func (u *Unit) Val() (float64, error) {

	if u.unitType == REL || math.IsNaN(u.val) { // rel unit can't be cache
		return u.getVal(map[*Unit]bool{})
	}
	return u.val, nil

}

// RelVal set base rel unit value and then return u's value
func (u *Unit) RelVal(ref *Unit, val float64, visS ...map[*Unit]bool) (float64, error) {
	if len(visS) == 1 {
		if _, ok := visS[0][u]; ok {
			return math.NaN(), errors.New("unit" + u.String() + "already visited 😬😬😬")
		}
		visS[0][u] = true
	} else {
		visS = []map[*Unit]bool{{u: true}}
	}
	if u == ref {
		if u.ref != nil {
			return math.NaN(), errors.New("only root unit can set reference value")
		}
		return val, nil
	}
	if u.ref == nil {
		return math.NaN(), errors.New("parent can't be nil finding fail")
	}
	v, err := u.ref.RelVal(ref, val, visS[0])
	return v * u.scale, err
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
	refVal, err := u.ref.getVal(vis)
	if err != nil {
		return math.NaN(), errors.New("reference tracing failed 😬😬😬")
	}
	return refVal * u.scale, nil

}

// AddUnit define new unit
func AddUnit(name string, scale float64, ref *Unit, unitType unitTypeID) (*Unit, error) {
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

// RemoveUnit remove a unit
func RemoveUnit(name string) {
	delete(units, name)
}

// GetUnit get unit by it's name
func GetUnit(name string) *Unit {
	if v, ok := units[name]; ok {
		return v
	}
	return nil
}

// ClearUnits delete all units
func ClearUnits() {
	units = make(map[string]*Unit)
}
