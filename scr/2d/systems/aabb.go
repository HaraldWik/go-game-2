package s2d

import (
	"fmt"

	d2d "github.com/HaraldWik/go-game-2/scr/2d/data"
	"github.com/HaraldWik/go-game-2/scr/ups"
)

type AABB struct{}

func (a AABB) Start(obj *ups.Object) {
	obj.Tags.Add("AABB")
}

func (a AABB) Update(obj *ups.Object, deltaTime float32) {
	var (
		transform = obj.Data.Get("Transform").(d2d.Transform2D)
		aabbs     = obj.Scene.FindByTag("AABB")
	)

	fmt.Println(aabbs)
	fmt.Println(obj.Tags)

	for _, aabb := range aabbs {
		if obj.Name != aabb.Name {
			aabbTransform := aabb.Data.Get("Transform").(d2d.Transform2D)

			if transform.Position.X < aabbTransform.Position.X+aabbTransform.Scale.X {
				aabbTransform.Position.Y += 5 * deltaTime
			}
			if transform.Position.X+transform.Scale.X > aabbTransform.Position.X {
				aabbTransform.Position.Y -= 5 * deltaTime
			}
			if transform.Position.Y < aabbTransform.Position.Y+aabbTransform.Scale.Y {
				aabbTransform.Position.Y -= 5 * deltaTime
			}
			if transform.Position.Y+transform.Scale.Y > aabbTransform.Position.Y {
				aabbTransform.Position.Y += 5 * deltaTime
			}
		}
	}

	obj.Data.Set("Transform", transform)
}
