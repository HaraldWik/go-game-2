package mod

import (
	"github.com/HaraldWik/go-game-2/scr/abus"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

type Transform3D struct {
	Pos, Size, Rot vec3.Type
	hasStarted     bool
}

func (transform Transform3D) Start() {

}

func (transform *Transform3D) Update() {
	if !transform.hasStarted {
		curObj := abus.SceneManager.GetCurrentScene().CurObj
		curObj.AddProperty(transform)
		transform.hasStarted = true
	}
}

func GetTransform3D(obj abus.Obj) Transform3D {
	if obj.GetProperty(Transform3D{}) != nil {
		return obj.GetProperty(Transform3D{}).(Transform3D)
	}
	return Transform3D{}
}
