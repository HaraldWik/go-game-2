package scene

import vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"

type Obj3D struct {
	ID    int
	Pos   vec3.Type
	Size  vec3.Type
	Rot   vec3.Type
	Color vec3.Type // X=Red, Y=Green, Z=Blue
	Comp  []func(obj *Obj3D)
}

func (obj *Obj3D) Update() {
	for _, comp := range obj.Comp {
		comp(obj)
	}
}
