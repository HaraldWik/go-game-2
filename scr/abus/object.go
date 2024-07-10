package abus //Array based update system ABUS

import "reflect"

type Obj struct {
	ID    uint32
	Mods  []Mod         // Module
	Props []interface{} // Property
}

func NewObject(mods ...Mod) Obj {
	return Obj{Mods: mods}
}

func (obj *Obj) Update() {
	for _, mod := range obj.Mods {
		mod.Update()
	}
}

// Module
func (obj *Obj) AddModules(mods ...Mod) {
	obj.Mods = append(obj.Mods, mods...)
}

func (obj *Obj) RemoveModules(mods ...Mod) {
	for _, m := range mods {
		for i, mod := range obj.Mods {
			if mod == m {
				obj.Mods = append(obj.Mods[:i], obj.Mods[i+1:]...)
				break
			}
		}
	}
}

// Property
func (obj *Obj) AddProperty(props ...interface{}) {
	obj.Props = append(obj.Props, props...)
}

func (obj *Obj) RemoveProperty(props ...interface{}) {
	for _, p := range props {
		for i, prop := range obj.Mods {
			if prop == p {
				obj.Props = append(obj.Props[:i], obj.Props[i+1:]...)
				break
			}
		}
	}
}

func (obj *Obj) GetProperty(propType interface{}) interface{} {
	for _, p := range obj.Props {
		if p == propType {
			return p
		}
	}
	return nil
}

func (obj *Obj) ContainsProperty(target interface{}) (interface{}, bool, uint32) {
	var count uint32
	var matches []interface{}
	var hasProp bool
	targetValue := reflect.TypeOf(target)
	for _, prop := range obj.Props {
		if reflect.TypeOf(prop) == targetValue {
			count++
			matches = append(matches, prop)
		}
	}

	if count > 0 {
		hasProp = true
	}

	return matches[0], hasProp, count
}
