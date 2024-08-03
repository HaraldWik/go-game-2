package s2d

import (
	d2d "github.com/HaraldWik/go-game-2/scr/2d/data"
	"github.com/HaraldWik/go-game-2/scr/ups"
)

type AABB struct{}

func (a AABB) Start(obj *ups.Object) {
	obj.Tags.Add("AABB")
}

func (a AABB) Update(obj *ups.Object, deltaTime float32) {
	transform := obj.Data.Get("Transform").(d2d.Transform2D)
	aabbs := obj.Scene.GetByTag("AABB")

	for _, aabb := range aabbs {
		if aabb.Name != obj.Name {
			aabbTransform := aabb.Data.Get("Transform").(d2d.Transform2D)

			// Check for collision in the X direction
			if transform.Position.X < aabbTransform.Position.X+aabbTransform.Scale.X-transform.Scale.X/2 &&
				transform.Position.X+transform.Scale.X-transform.Scale.X/2 > aabbTransform.Position.X {
				// Collision detected on X axis, resolve
				if transform.Position.X < aabbTransform.Position.X {
					transform.Position.X -= 5.0 * deltaTime
				} else {
					transform.Position.X += 5.0 * deltaTime
				}
			}

			// Check for collision in the Y direction
			if transform.Position.Y < aabbTransform.Position.Y+aabbTransform.Scale.Y-transform.Scale.Y/2 &&
				transform.Position.Y+transform.Scale.Y-transform.Scale.Y/2 > aabbTransform.Position.Y {
				// Collision detected on Y axis, resolve
				if transform.Position.Y < aabbTransform.Position.Y {
					transform.Position.Y -= 3.0 * deltaTime
				} else {
					transform.Position.Y += 3.0 * deltaTime
				}
			}
		}
	}
	obj.Data.Set("Transform", transform)
}

type StaticAABB struct{}

func (a StaticAABB) Start(obj *ups.Object) {
	obj.Data.Set("Start", obj.Data.Get("Transform"))
}

func (a StaticAABB) Update(obj *ups.Object, deltaTime float32) {
	transform := obj.Data.Get("Start").(d2d.Transform2D)
	obj.Data.Set("Transform", transform)
}
