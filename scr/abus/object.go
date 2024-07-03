package abus //Array based update system ABUS

import "reflect"

type Obj struct {
	ID    int32
	Comps []Comp
}

func (obj *Obj) Update() {
	for _, comp := range obj.Comps {
		comp.Update()
	}
}

func (obj *Obj) AddComponents(comps ...Comp) {
	obj.Comps = append(obj.Comps, comps...)
}

// RemoveComponents method for Object to remove specified components
func (obj *Obj) RemoveComponents(comps ...Comp) {
	for _, compToRemove := range comps {
		for i, comp := range obj.Comps {
			if comp == compToRemove {
				obj.Comps = append(obj.Comps[:i], obj.Comps[i+1:]...)
				break // Continue removing if there are duplicates of the same component
			}
		}
	}
}

func (obj *Obj) GetComponentOfType(typ reflect.Type) Comp {
	for _, comp := range obj.Comps {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}
	return nil
}
