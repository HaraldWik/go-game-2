package scene

import (
	"fmt"

	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

type Obj2D struct {
	ID    int
	Pos   vec2.Type
	Size  vec2.Type
	Rot   float32
	Color vec3.Type // X=Red, Y=Green, Z=Blue
	Comp  []func(obj *Obj2D)
}

func (obj *Obj2D) Update() {
	for _, comp := range obj.Comp {
		comp(obj)
	}
}

func (obj *Obj2D) AddCommponents(comp []func(obj *Obj2D)) {
	for i := 0; i < len(comp); i++ {
		obj.Comp = append(obj.Comp, comp[i])
	}
}

func Physics2D(obj *Obj2D) {
	obj.Pos.Y -= 0.1
}

func Hello2D(obj Obj2D) {
	fmt.Println(obj)
}
