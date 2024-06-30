package ecs

// Scene
type Scene struct {
	List []Obj
}

func NewScene() Scene {
	return Scene{}
}

func (scene *Scene) Create(comp ...Commponent) Obj {
	scene.List = append(scene.List, Obj{ID: len(scene.List), Components: comp})
	return scene.List[len(scene.List)-1]
}

func (scene *Scene) Delete(obj Obj) {
	for i, obj := range scene.List {
		if obj.ID == i {
			scene.List = append(scene.List[:i], scene.List[i+1:]...)
			return
		}
	}
}

func (scene *Scene) Update() {
	for i := range scene.List {
		scene.List[i].Update()
	}
}

// Object
type Obj struct {
	ID         int
	Components []Commponent
}

func (obj *Obj) Update() {
	for _, comp := range obj.Components {
		comp.Update()
	}
}

func (obj *Obj) AddComponents(comps ...Commponent) {
	obj.Components = append(obj.Components, comps...)
}

// RemoveComponents method for Object to remove specified components
func (obj *Obj) RemoveComponents(comps ...Commponent) {
	for _, compToRemove := range comps {
		for i, comp := range obj.Components {
			if comp == compToRemove {
				obj.Components = append(obj.Components[:i], obj.Components[i+1:]...)
				break // Continue removing if there are duplicates of the same component
			}
		}
	}
}

// Component
type Commponent interface {
	Update()
}
