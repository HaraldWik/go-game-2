package scene

import (
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

type Scene2D struct {
	List []Obj2D
}

func New2D() Scene2D {
	return Scene2D{}
}

func (scene Scene2D) LoopThru() (int, Obj2D) {
	for i, obj := range scene.List {
		return i, obj
	}
	return 0, Obj2D{}
}

func (scene Scene2D) Update() {
	if len(scene.List) != 0 {
		i, _ := scene.LoopThru()
		scene.List[i].Update()
	}
}

func (scene *Scene2D) Set(id int, new Obj2D) {
	scene.List[id] = new
}

func (scene *Scene2D) Get(id int) Obj2D {
	if _, obj := scene.LoopThru(); obj.ID == id {
		return obj
	}
	return Obj2D{}
}

func (scene *Scene2D) Create(pos, size vec2.Type, rot float32, color vec3.Type, comp []func(obj *Obj2D)) Obj2D {
	obj := Obj2D{
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

func (scene *Scene2D) Delete(obj Obj2D) {
	if i, _ := scene.LoopThru(); obj.ID == i {
		scene.List = append(scene.List[:i], scene.List[i+1:]...)
		return
	}
}
