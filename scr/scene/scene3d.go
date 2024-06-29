package scene

import (
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

type Scene3D struct {
	List []Obj3D
}

func New3D() Scene3D {
	return Scene3D{}
}

func (scene Scene3D) LoopThru() (int, Obj3D) {
	for i, obj := range scene.List {
		return i, obj
	}
	return 0, Obj3D{}
}

func (scene Scene3D) Update() {
	if len(scene.List) != 0 {
		i, _ := scene.LoopThru()
		scene.List[i].Update()
	}
}

func (scene *Scene3D) Set(obj Obj3D, new Obj3D) {
	scene.List[obj.ID] = new
}

func (scene *Scene3D) Get(id int) Obj3D {
	if _, obj := scene.LoopThru(); obj.ID == id {
		return obj
	}
	return Obj3D{}
}

func (scene *Scene3D) Create(pos, size, rot, color vec3.Type, comp []func(obj *Obj3D)) Obj3D {
	obj := Obj3D{
		ID:    len(scene.List),
		Pos:   pos,
		Size:  size,
		Rot:   rot,
		Color: color,
		Comp:  comp,
	}
	scene.List = append(scene.List, obj)
	return obj
}

func (scene *Scene3D) Delete(obj Obj3D) {
	if i, _ := scene.LoopThru(); obj.ID == i {
		scene.List = append(scene.List[:i], scene.List[i+1:]...)
		return
	}
}
