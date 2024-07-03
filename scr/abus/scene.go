package abus //Array based update system ABUS

import "reflect"

type Scene struct {
	Objs []Obj
}

func NewScene() Scene {
	return Scene{}
}

func (scene *Scene) Update() {
	for i := range scene.Objs {
		scene.Objs[i].Update()
	}
}

func (scene *Scene) Create(comp ...Comp) Obj {
	scene.Objs = append(scene.Objs, Obj{ID: int32(len(scene.Objs)), Comps: comp})
	return scene.Objs[len(scene.Objs)-1]
}

func (scene *Scene) Delete(id int32) {
	for i, obj := range scene.Objs {
		if obj.ID == id {
			scene.Objs = append(scene.Objs[:i], scene.Objs[i+1:]...)
			return
		}
	}
}

func (scene *Scene) GetId(obj Obj) int32 {
	return int32(scene.Objs[obj.ID].ID)
}

func (scene *Scene) SetObject(id int32, obj Obj) {
	scene.Objs[id] = obj
}

func (scene *Scene) GetObject(id int32) Obj {
	return scene.Objs[id]
}

func (scene *Scene) SetComponent(obj Obj, comps ...Comp) {
	for _, o := range scene.Objs {
		if o.ID == obj.ID {
			o.AddComponents(comps...)
			return
		}
	}
}

func (scene *Scene) GetComponent(obj Obj, typ reflect.Type) Comp {
	for _, o := range scene.Objs {
		if o.ID == obj.ID {
			return o.GetComponentOfType(typ)
		}
	}
	return nil
}
